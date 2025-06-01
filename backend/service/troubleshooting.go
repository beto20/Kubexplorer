package service

import (
	"context"
	"fmt"
	app "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SourceObject string

const (
	Pod        SourceObject = "POD"
	Service    SourceObject = "SERVICE"
	Deployment SourceObject = "DEPLOYMENT"
	Secret     SourceObject = "SECRET"
)

type WellKnowPodError string
type WellKnowDeploymentError string
type WellKnowServiceError string
type WellKnowSecretError string

const (
	CrashLoopBackOff    WellKnowPodError        = "CrashLoopBackOff"
	OomKilled           WellKnowPodError        = "OOMKilled"
	ImagePullBackOff    WellKnowPodError        = "ImagePullBackOff"
	ErrImagePull        WellKnowPodError        = "ErrImagePull"
	Evicted             WellKnowPodError        = "Evicted"
	Pending             WellKnowPodError        = "Pending"
	UnavailableReplicas WellKnowDeploymentError = "UnavailableReplicas"
	DeploymentPending   WellKnowDeploymentError = "DeploymentPending"
	ConnectionRefused   WellKnowServiceError    = "ConnectionRefused"
	IpPending           WellKnowServiceError    = "IPPending"
	Unauthorized        WellKnowSecretError     = "Unauthorized"
	InvalidSecretData   WellKnowSecretError     = "InvalidSecretData"
)

type Troubleshooting struct {
	client kubernetes.Interface
}

func NewTroubleshooting(client kubernetes.Interface) *Troubleshooting {
	return &Troubleshooting{client: client}
}

func (t *Troubleshooting) Analyse(name string, namespace string, object SourceObject) {

	switch object {
	case Pod:
		pod, err := t.client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			panic("Pod not found")
		}
		identifyPodPotentialError(pod)
	case Service:
		service, err := t.client.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			panic("Service not found")
		}
		identifyServicePotentialError(service)
	case Deployment:
		deployment, err := t.client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			panic("Deployment not found")
		}
		identifyDeploymentPotentialError(deployment)
	default:
		fmt.Println("Not a valid object")
	}

}

func identifyPodPotentialError(pod *core.Pod) {

	for _, container := range pod.Status.ContainerStatuses {
		if container.State.Waiting != nil {
			reason := container.State.Waiting.Reason
			switch reason {
			case "CrashLoopBackOff":
				fmt.Println("Container is crashing. Suggest checking logs.")
			case "OOMKilled":
				fmt.Println("Container is crashing. Suggest increase memory usage.")
			case "ImagePullBackOff":
				fmt.Println("Container is crashing. Suggest check out the image policy.")
			case "ErrImagePull":
				fmt.Println("Container is crashing. Suggest check out the image repository.")
			case "Evicted":
				fmt.Println("Container is crashing. Probably the node doesn't have enough resources.")
			case "Pending":
				fmt.Println("Container is crashing. Probably there are other pods in schedule queue")
			}
		}
	}

}

func identifyServicePotentialError(service *core.Service) {
	for _, condition := range service.Status.Conditions {
		fmt.Println(condition.Message)
		fmt.Println(condition.Reason)
		fmt.Println(condition.Status)
	}
}

func identifyDeploymentPotentialError(deployment *app.Deployment) {
	for _, condition := range deployment.Status.Conditions {
		fmt.Println(condition.Type)
		fmt.Println(condition.Status)
		fmt.Println(condition.Reason)
		fmt.Println(condition.Message)
	}
}
