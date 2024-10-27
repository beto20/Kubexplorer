package objects

import (
	"context"
	"fmt"

	"github.com/beto20/kubessitant/config"
	"github.com/beto20/kubessitant/knowledge"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	Name      string
	Namespace string
	Container Container
	Age       string
	Status    string
}

type Container struct {
	Limit   Resource
	Request Resource
}

type Resource struct {
	Cpu    string
	Memory string
}

var client = config.GetKubeInstance()

func GetPods(namespace string) []Pod {
	podsClient := client.CoreV1().Pods(namespace)
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get pods")
	}
	var podArray []Pod

	for _, pod := range pods.Items {
		p := Pod{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Container: Container{
				Limit: Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Limits.Cpu().String(),
					Memory: pod.Spec.Containers[0].Resources.Limits.Memory().String(),
				},
				Request: Resource{
					Cpu:    pod.Spec.Containers[0].Resources.Requests.Cpu().String(),
					Memory: pod.Spec.Containers[0].Resources.Requests.Memory().String(),
				},
			},
			Status: string(pod.Status.Phase),
			Age:    pod.Status.StartTime.String(),
		}

		podArray = append(podArray, p)

		// fmt.Printf("pod name: %s namespace: %s requestCPU: %s limitsCPU: %s requestMemory: %s limitsMemory: %s storage: %s startTime: %s status: %s\n",
		// 	pod.Name,
		// 	pod.Namespace,
		// 	pod.Spec.Containers[0].Resources.Requests.Cpu(),
		// 	pod.Spec.Containers[0].Resources.Limits.Cpu(),
		// 	pod.Spec.Containers[0].Resources.Requests.Memory(),
		// 	pod.Spec.Containers[0].Resources.Limits.Memory(),
		// 	pod.Spec.Containers[0].Resources.Limits.Storage(),
		// 	pod.Status.StartTime,
		// 	pod.Status.Phase,
		// )
	}

	return podArray
}

func status(namespace string) {
	podsClient := client.CoreV1().Pods(namespace)
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error to get pods")
	}

	for _, pod := range pods.Items {

		if pod.Status.ContainerStatuses[0].State.Waiting != nil && pod.Status.ContainerStatuses[0].State.Waiting.Reason == string(knowledge.CRASH_LOOP_BACK_OFF) {
			fmt.Printf("Pod %s is in CrashLoopBackOff: %s\n", pod.Name, pod.Status.ContainerStatuses[0].State.Waiting.Reason)
		}
		fmt.Printf("pod: %s\n", pod.Status.ContainerStatuses[0].State.Waiting)
	}

	knowledge.ErrorSource(knowledge.PODS, string(knowledge.CRASH_LOOP_BACK_OFF))
}
