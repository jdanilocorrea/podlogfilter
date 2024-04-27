package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient(kubeconfigPath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic("Falha ao construir kubeconfig: " + err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("Falha ao criar cliente Kubernetes: " + err.Error())
	}
	return clientset
}
