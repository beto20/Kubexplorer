package objects

import (
	"Kubessistant/backend/config"
	"Kubessistant/backend/knowledge"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
	"math/rand"
	"strconv"
)

type IPodObject interface {
	GetPodsMock() []PodDto
	GetPods() []PodDto
	GetPodByName(name string) PodDto
	UpdatePodByName(name string) bool
	DeletePodByName(name string) bool
	GetPodMetric(namespace string) []PodMetricDto
}

type podImpl struct{}

func NewPodObject() IPodObject {
	return &podImpl{}
}

type PodDto struct {
	Name      string
	Namespace string
	Replicas  int32
	Container Container
	Age       string
	Status    string
}

type PodMetricDto struct {
	Name      string
	Namespace string
	Consume   Resource
}

type Container struct {
	Limit   Resource
	Request Resource
}

type Resource struct {
	Cpu              string
	Memory           string
	Storage          string
	StorageEphemeral string
}

var client = config.GetKubeInstance()

func (p *podImpl) GetPods() []PodDto {
	podsClient := client.CoreV1().Pods("TODO")
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get pods")
	}
	var podArray []PodDto

	for _, pod := range pods.Items {
		p := PodDto{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Replicas:  1,
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

func (p *podImpl) GetPodsMock() []PodDto {
	var pods []PodDto
	for i := 0; i < 10; i++ {

		p := PodDto{
			Name:      fmt.Sprintf("pod %d", i),
			Namespace: "TODO",
			Replicas:  1,
			Container: Container{
				Limit: Resource{
					Cpu:    strconv.Itoa(rand.Intn(1000)) + "mi",
					Memory: strconv.Itoa(rand.Intn(1000)),
				},
				Request: Resource{
					Cpu:    strconv.Itoa(rand.Intn(1000)),
					Memory: strconv.Itoa(rand.Intn(1000)),
				},
			},
			Status: "Alive",
			Age:    strconv.Itoa(rand.Intn(1000)),
		}

		pods = append(pods, p)
	}

	return pods
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

func (p *podImpl) GetPodByName(name string) PodDto {
	var pods PodDto

	return pods
}

func (p *podImpl) UpdatePodByName(name string) bool {
	return true
}

func (p *podImpl) DeletePodByName(name string) bool {
	return true
}

func (p *podImpl) GetPodMetric(namespace string) []PodMetricDto {
	inClusterConfig, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error creating in-cluster config: %v", err)
	}

	metricsClient, err := metricsv.NewForConfig(inClusterConfig)
	if err != nil {
		fmt.Print(err)
	}

	isWatchEnable := validateWatchFeature(inClusterConfig)
	if isWatchEnable && false {
		watchMetrics()
	}

	return pollMetrics(metricsClient, namespace)
}

func pollMetrics(metricsClient *metricsv.Clientset, namespace string) []PodMetricDto {
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error to get pod metrics")
	}

	var metrics []PodMetricDto

	for _, podMetric := range podMetrics.Items {
		for _, container := range podMetric.Containers {
			metrics = append(metrics, PodMetricDto{
				Name:      podMetric.Name,
				Namespace: podMetric.Namespace,
				Consume: Resource{
					Cpu:              container.Usage.Cpu().String(),
					Memory:           container.Usage.Memory().String(),
					Storage:          container.Usage.Storage().String(),
					StorageEphemeral: container.Usage.StorageEphemeral().String(),
				},
			})
		}
	}

	return metrics
}

func watchMetrics() {
	// TODO: next feature
}

func validateWatchFeature(inClusterConfig *rest.Config) bool {

	discoveryClient := discovery.NewDiscoveryClientForConfigOrDie(inClusterConfig)
	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion("metrics.k8s.io/v1beta1")
	if err != nil {
		fmt.Printf("Error getting API resources: %v", err)
	}

	for _, resources := range apiResourceList.APIResources {
		if resources.Name == "pods" {
			for _, verb := range resources.Verbs {
				if verb == "watch" {
					return true
				}
			}
		}
	}

	return false
}
