package pod

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ErrorsWellKnow string

const (
	CrashLoopBackOff ErrorsWellKnow = "CrashLoopBackOff"
	OOMKilled        ErrorsWellKnow = "OOMKilled"
	ImagePullBackOff ErrorsWellKnow = "ImagePullBackOff"
	Evicted          ErrorsWellKnow = "Evicted"
	Pending          ErrorsWellKnow = "Pending"
)

func setClient() *kubernetes.Clientset {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	if err != nil {
		fmt.Println("Error when get Kubeconfig")
	}

	client := kubernetes.NewForConfigOrDie(config)

	return client
}

func GetPods(namespace string) {
	client := setClient()
	podsClient := client.CoreV1().Pods(namespace)
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error to get pods")
	}

	for _, pod := range pods.Items {
		if pod.Status.ContainerStatuses[0].State.Waiting != nil && pod.Status.ContainerStatuses[0].State.Waiting.Reason == string(CrashLoopBackOff) {
			fmt.Printf("Pod %s is in CrashLoopBackOff: %s\n", pod.Name, pod.Status.ContainerStatuses[0].State.Waiting.Reason)
		}
		fmt.Printf("pod: %s\n", pod.Status.ContainerStatuses[0].State.Waiting)
	}

	applySolution(CrashLoopBackOff)
}

func rollbackImage() {
	fmt.Print("rollback")
}

func applySolution(errorType ErrorsWellKnow) {
	fmt.Print("applySolution")
	switch errorType {
	case ErrorsWellKnow(OOMKilled):
		calibrateResources(errorType)
	case ErrorsWellKnow(Evicted):
		calibrateResources(errorType)
	case ErrorsWellKnow(ImagePullBackOff):
		rollbackImage()
	case ErrorsWellKnow(CrashLoopBackOff):
		rollbackImage()
	case ErrorsWellKnow(Pending):
		calibrateResources(errorType)
	}
}

func calibrateResources(errorType ErrorsWellKnow) {

	if errorType == ErrorsWellKnow(OOMKilled) {
		increaseMemory()
	} else if errorType == ErrorsWellKnow(Pending) {
		increaseCPU()
	} else {
		increaseMemory()
		increaseCPU()
	}

}

func increaseMemory() {
	fmt.Print("memory increase")
}

func increaseCPU() {
	fmt.Print("cpu increase")
}
