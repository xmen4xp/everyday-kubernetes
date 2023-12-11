package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

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
	home := homedir.HomeDir()

	flag.StringVar(&kubeconfig, "k", filepath.Join(home, ".kube", "config"), "Absolute path to the kubeconfig file. Defaults to ~/.kube/config.")
	flag.IntVar(&num_watchers, "num-watchers", 1, "Number of watchers to start. Defaults to 1.")

	flag.Parse()

	ctxt, cancelFn := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for iter := 0; iter < num_watchers; iter++ {
		go func() {
			defer wg.Done()
			start_namespace_watcher(&kubeconfig, ctxt, iter)
		}()
	}

	fmt.Println("Blocking, press ctrl+c to terminatett...")
	<-done
	cancelFn()
	wg.Wait()
	fmt.Println("Program terminated successfully. Goodbye !")
	fmt.Println()
}
