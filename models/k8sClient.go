package models

import (
	"flag"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset
var once sync.Once

func GetClientset() *kubernetes.Clientset {
	once.Do(func() {
		kubeconfig := flag.String("kubeconfig", "/go/src/.kube/config", "absolute path to the kubeconfig file")
		flag.Parse()
		// uses the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags(clientcmd.ClusterDefaults.Server, *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err = kubernetes.NewForConfig(config)
	})
	return clientset
}
