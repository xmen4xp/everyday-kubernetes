package helper

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/elliotchance/orderedmap"

	datamodel "perfdm/build/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const DEFAULT_KEY = "default"
const DISPLAY_NAME_LABEL = "nexus/display_name"
const IS_NAME_HASHED_LABEL = "nexus/is_name_hashed"

func GetCRDParentsMap() map[string][]string {
	return map[string][]string{
		"orderses.root.sockshop.com":  {"sockshops.root.sockshop.com"},
		"shippings.root.sockshop.com": {"sockshops.root.sockshop.com"},
		"sockses.root.sockshop.com":   {"sockshops.root.sockshop.com"},
		"sockshops.root.sockshop.com": {},
	}
}

func GetObjectByCRDName(dmClient *datamodel.Clientset, crdName string, name string) interface{} {
	if crdName == "orderses.root.sockshop.com" {
		obj, err := dmClient.RootSockshopV1().Orderses().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "shippings.root.sockshop.com" {
		obj, err := dmClient.RootSockshopV1().Shippings().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "sockses.root.sockshop.com" {
		obj, err := dmClient.RootSockshopV1().Sockses().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "sockshops.root.sockshop.com" {
		obj, err := dmClient.RootSockshopV1().SockShops().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}

	return nil
}

func ParseCRDLabels(crdName string, labels map[string]string) *orderedmap.OrderedMap {
	parents := GetCRDParentsMap()[crdName]

	m := orderedmap.NewOrderedMap()
	for _, parent := range parents {
		if label, ok := labels[parent]; ok {
			m.Set(parent, label)
		} else {
			m.Set(parent, DEFAULT_KEY)
		}
	}

	return m
}

func GetHashedName(crdName string, labels map[string]string, name string) string {
	orderedLabels := ParseCRDLabels(crdName, labels)

	var output string
	for i, key := range orderedLabels.Keys() {
		value, _ := orderedLabels.Get(key)

		output += fmt.Sprintf("%s:%s", key, value)
		if i < orderedLabels.Len()-1 {
			output += "/"
		}
	}

	output += fmt.Sprintf("%s:%s", crdName, name)

	h := sha1.New()
	_, _ = h.Write([]byte(output))
	return hex.EncodeToString(h.Sum(nil))
}
