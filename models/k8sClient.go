package models

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

var clientset *kubernetes.Clientset
var once sync.Once

func GetClientset() *kubernetes.Clientset {
	once.Do(func() {
		kubeconfig := flag.String("kubeconfig", "C://Users//hy352//.kube//config", "absolute path to the kubeconfig file")
		flag.Parse()
		// uses the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("192.168.34.41:8080", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err = kubernetes.NewForConfig(config)
	})
	return clientset
}
