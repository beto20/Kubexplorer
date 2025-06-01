package service

import (
	"context"
	"fmt"
	app "k8s.io/api/apps/v1"
	batch "k8s.io/api/batch/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SourceObject string

const (
	Pod        SourceObject = "POD"
	Job        SourceObject = "JOB"
	Deployment SourceObject = "DEPLOYMENT"
)

type WellKnownPodError string
type WellKnownDeploymentError string

type WellKnownJobError string
type WellKnowServiceError string
type WellKnowSecretError string

const (
	// Pod-level
	CrashLoopBackOff           WellKnownPodError = "CrashLoopBackOff"
	OOMKilled                  WellKnownPodError = "OOMKilled"
	ImagePullBackOff           WellKnownPodError = "ImagePullBackOff"
	ErrImagePull               WellKnownPodError = "ErrImagePull"
	CreateContainerConfigError WellKnownPodError = "CreateContainerConfigError"
	ContainerCannotRun         WellKnownPodError = "ContainerCannotRun"
	Unschedulable              WellKnownPodError = "Unschedulable"
	Evicted                    WellKnownPodError = "Evicted"
	NodeLost                   WellKnownPodError = "NodeLost"

	// Deployment-level
	UnavailableReplicas        WellKnownDeploymentError = "UnavailableReplicas"
	MinimumReplicasUnavailable WellKnownDeploymentError = "MinimumReplicasUnavailable"

	// Job-level
	DeadlineExceeded     WellKnownJobError = "DeadlineExceeded"
	BackoffLimitExceeded WellKnownJobError = "BackoffLimitExceeded"
)

type Diagnostic struct {
	client kubernetes.Interface
}

func NewDiagnosticService(client kubernetes.Interface) *Diagnostic {
	return &Diagnostic{client: client}
}

func (d *Diagnostic) Analyse(name string, namespace string, object SourceObject) {

	switch object {
	case Pod:
		pod, err := d.client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error retrieving Pod: %v\n", err)
			return
		}
		identifyPodPotentialError(pod)
	case Job:
		job, err := d.client.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error retrieving Job: %v\n", err)
			return
		}
		identifyJobPotentialError(job)
	case Deployment:
		deployment, err := d.client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error retrieving Deployment: %v\n", err)
			return
		}
		identifyDeploymentPotentialError(deployment)
	default:
		fmt.Println("Not a valid object")
	}

}

func identifyPodPotentialError(pod *core.Pod) {
	fmt.Println("Pod:", pod.Name)
	fmt.Println("Phase:", pod.Status.Phase)
	fmt.Println("PodIP:", pod.Status.PodIP)
	fmt.Println("Message:", pod.Status.Message)
	fmt.Println("Reason:", pod.Status.Reason)

	switch pod.Status.Reason {
	case string(Evicted):
		fmt.Println("Evicted")
	case string(NodeLost):
		fmt.Println("NodeLost")
	}

	for _, container := range pod.Status.ContainerStatuses {
		fmt.Println("ContainerID:", container.ContainerID)
		fmt.Println("State:", container.State)
		fmt.Println("Name:", container.Name)
		fmt.Println("Image:", container.Image)
		fmt.Println("ImageID:", container.ImageID)

		if container.State.Terminated != nil {
			if container.LastTerminationState.Terminated.Reason == string(OOMKilled) {
				fmt.Println("Container is crashing. Suggest increase memory usage.")
			}
		}

		if container.State.Waiting != nil {
			fmt.Println("Waiting.Message:", container.State.Waiting.Message)
			reason := container.State.Waiting.Reason
			fmt.Println("Waiting.Reason:", reason)

			switch reason {
			case string(CrashLoopBackOff):
				fmt.Println("Container is crashing. Suggest checking logs.")
			case string(ImagePullBackOff):
				fmt.Println("Container is crashing. Suggest check out the image policy.")
			case string(ErrImagePull):
				fmt.Println("Container is crashing. Suggest check out the image repository.")
			case string(ContainerCannotRun):
				fmt.Println("Container is crashing. Probably the Container Cannot Run")
			case string(CreateContainerConfigError):
				fmt.Println("Container is crashing. Probably the Create Container Config Error")
			}
		}
	}

	for _, condition := range pod.Status.Conditions {
		if condition.Reason == string(Unschedulable) {
			fmt.Println("Container is crashing. Unschedulable state")
		}
	}

}

func identifyJobPotentialError(job *batch.Job) {

	for _, condition := range job.Status.Conditions {
		if condition.Reason == string(BackoffLimitExceeded) {
			fmt.Println("Container is crashing. BackoffLimitExceeded.")
		}
		if condition.Reason == string(DeadlineExceeded) {
			fmt.Println("Container is crashing. DeadlineExceeded.")
		}
	}
}

func identifyDeploymentPotentialError(deployment *app.Deployment) {
	if deployment.Status.UnavailableReplicas > 0 {
		fmt.Println("Deployment is unavailable")
	}

	for _, condition := range deployment.Status.Conditions {
		if condition.Reason == string(MinimumReplicasUnavailable) {
			fmt.Println("Deployment error. MinimumReplicasUnavailable.")
		}
	}
}
