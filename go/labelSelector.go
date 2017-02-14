package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/labels"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
	fmt.Printf("starting")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("kubernetes:8080", "")
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	labelSelector := labels.Set(map[string]string{"release": "zw-test-11"}).AsSelector()
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector.String(),
	}

	for {

		pods, err := clientset.CoreV1().Pods("").List(listOptions)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(2 * time.Second)
	}
}
