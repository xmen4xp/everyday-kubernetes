package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	sockshopv1 "perfdm/build/apis/root.sockshop.com/v1"
	raw_clientset "perfdm/build/client/clientset/versioned"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func start_write_crd(kubeconfig *string, ctxt context.Context, writerID int) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := raw_clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	writerPrefix := fmt.Sprintf("writer-%d", writerID)
	sockshopObj := sockshopv1.SockShop{
		ObjectMeta: v1.ObjectMeta{Name: writerPrefix},
		Spec: sockshopv1.SockShopSpec{
			OrgName: writerPrefix,
		},
	}
	obj, err := clientset.RootSockshopV1().SockShops().Create(ctxt, &sockshopObj, v1.CreateOptions{})
	if err != nil {
		if k8serrors.IsAlreadyExists(err) {
			obj, err = clientset.RootSockshopV1().SockShops().Get(ctxt, writerPrefix, v1.GetOptions{})
		}
		if err != nil {
			fmt.Printf("Stopping writer %s due to create err: %+v\n", writerPrefix, err)
			return
		}
	}

	counter := 1
	for {
		//fmt.Printf("%+v\n", *obj)
		obj.Spec.OrgName = fmt.Sprintf("%s-%d", writerPrefix, counter)
		obj, err = clientset.RootSockshopV1().SockShops().Update(ctxt, obj, v1.UpdateOptions{})
		if k8serrors.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(100 * time.Millisecond)
		} else if errors.Is(err, context.Canceled) {
			return
		} else if err != nil {
			fmt.Printf("Stopping writer %s due to update err: %+v\n,", writerPrefix, err)
			return
		}
		counter += 1
	}
}

func start_namespace_watcher(kubeconfig *string, ctxt context.Context, watcherID int) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	timeOut := int64(60)
	watcher, _ := clientset.CoreV1().Namespaces().Watch(ctxt, metav1.ListOptions{TimeoutSeconds: &timeOut})

	for event := range watcher.ResultChan() {
		item, ok := event.Object.(*corev1.Namespace)
		if !ok {
			status := event.Object.(*v1.Status)
			fmt.Println("Status", *status)
			return
		}
		fmt.Println("Watcher: ", watcherID, "Received event", event.Type, "for namespace", item.Name)
	}
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	var kubeconfig string
	var num_watchers int
	var num_raw_crd_writers int

	home := homedir.HomeDir()

	flag.StringVar(&kubeconfig, "k", filepath.Join(home, ".kube", "config"), "Absolute path to the kubeconfig file. Defaults to ~/.kube/config.")
	flag.IntVar(&num_watchers, "num-watchers", 0, "Number of watchers to start. Defaults to 0.")
	flag.IntVar(&num_raw_crd_writers, "num-raw-crd-writers", 0, "Number of k8s clientset based writes to a singe CRD to start. Defaults to 0.")

	flag.Parse()

	ctxt, cancelFn := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for iter := 0; iter < num_watchers; iter++ {
		go func() {
			defer wg.Done()
			start_namespace_watcher(&kubeconfig, ctxt, iter)
			wg.Add(1)
		}()
	}

	for iter := 0; iter < num_raw_crd_writers; iter++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Starting writer %d ...\n", id)
			start_write_crd(&kubeconfig, ctxt, id)
			fmt.Printf("Stopped writer %d ...\n", id)
			wg.Add(1)
		}(iter)
	}

	fmt.Println("Blocking, press ctrl+c to terminatett...")
	<-done
	cancelFn()
	wg.Wait()
	fmt.Println("Program terminated successfully. Goodbye !")
	fmt.Println()
}
