package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var lock = &sync.Mutex{}

var kubeClient *kubernetes.Clientset

func GetKubeInstance() *kubernetes.Clientset {
	if kubeClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if kubeClient == nil {
			fmt.Println("Creating a single instance")
			kubeClient = buildKubeClient()
		} else {
			fmt.Println("Single instance already exists")
		}
	} else {
		fmt.Println("Single instance already exists")
	}

	return kubeClient
}

func buildKubeClient() *kubernetes.Clientset {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	if err != nil {
		fmt.Println("Error when get Kubeconfig")
	}

	client := kubernetes.NewForConfigOrDie(config)

	return client
}
