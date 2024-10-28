package main

import (
	"github.com/beto20/kubessitant/objects"
)

func main() {
	// home, _ := os.UserHomeDir()
	// kubeConfigPath := filepath.Join(home, ".kube/config")

	// config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// client := kubernetes.NewForConfigOrDie(config)

	// getPods(client)
	// createPod(client)
	// updatePod(client)
	// deletePod(client)
	// getServices(client)
	// deleteDeployment((client))

	// watchLogs(client)
	// metrics(config)
	// podMetrics(config)

	/////////

	// x := objects.GetPods("")
	objects.GetDeployments("")

	// for _, pod := range x {

	// 	fmt.Printf("pod name: %s namespace: %s requestCPU: %s limitCPU: %s requestMemory: %s limitMemory: %s status: %s startTime: %s\n",
	// 		pod.Name,
	// 		pod.Namespace,
	// 		pod.Container.Request.Cpu,
	// 		pod.Container.Limit.Cpu,
	// 		pod.Container.Request.Memory,
	// 		pod.Container.Limit.Memory,
	// 		pod.Status,
	// 		pod.Age,
	// 	)
	// }

	// for i := 0; i < 30; i++ {
	// 	go config.GetKubeInstance()
	// }
	// fmt.Scanln()
}

// func podMetrics(config *rest.Config) {
// 	// Create the metrics client
// 	metricsClient, err := metricsv1beta1.NewForConfig(config)
// 	if err != nil {
// 		log.Fatalf("Error creating metrics client: %v", err)
// 	}

// 	// Fetch Pod metrics
// 	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses("mock").Get(context.TODO(), "mock-56964cb474-psck2", metav1.GetOptions{})
// 	if err != nil {
// 		log.Fatalf("Error getting pod metrics: %v", err)
// 	}

// 	// Variables to accumulate pod-level CPU and memory usage
// 	var totalCPUUsage resource.Quantity
// 	var totalMemoryUsage resource.Quantity

// 	// Loop through each container's metrics and accumulate the resource usage
// 	fmt.Printf("Aggregating metrics for Pod: %s/%s\n", podMetrics.Namespace, podMetrics.Name)
// 	for _, container := range podMetrics.Containers {
// 		cpuUsage := container.Usage["cpu"]
// 		memoryUsage := container.Usage["memory"]

// 		fmt.Printf("Container: %s | CPU Usage: %s | Memory Usage: %s\n", container.Name, cpuUsage.String(), memoryUsage.String())

// 		// Add the container's CPU and memory usage to the pod-level total
// 		totalCPUUsage.Add(cpuUsage)
// 		totalMemoryUsage.Add(memoryUsage)
// 	}

// 	// Print the total pod-level CPU and memory usage
// 	fmt.Printf("Total CPU Usage (Pod-level): %s\n", totalCPUUsage.String())
// 	fmt.Printf("Total Memory Usage (Pod-level): %s\n", totalMemoryUsage.String())

// }

// func metrics(config *rest.Config) {
// 	// Create a Metrics client
// 	metricsClient, err := metricsv1beta1.NewForConfig(config)
// 	if err != nil {
// 		log.Fatalf("Error creating metrics client: %v", err)
// 	}

// 	// Get Pod metrics
// 	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses("mock").Get(context.TODO(), "mock-56964cb474-psck2", metav1.GetOptions{})
// 	if err != nil {
// 		log.Fatalf("Error getting pod metrics: %v", err)
// 	}

// 	// Print the resource usage of each container in the pod
// 	fmt.Printf("Metrics for Pod: %s/%s\n", podMetrics.Namespace, podMetrics.Name)
// 	for _, container := range podMetrics.Containers {
// 		fmt.Printf("Container: %s\n", container.Name)
// 		cpuUsage := container.Usage["cpu"]
// 		memoryUsage := container.Usage["memory"]
// 		fmt.Printf("CPU Usage: %s\n", cpuUsage.String())
// 		fmt.Printf("Memory Usage: %s\n", memoryUsage.String())
// 	}
// }

// func watchLogs(client *kubernetes.Clientset) {
// 	informer := informers.NewSharedInformerFactoryWithOptions(client, time.Minute, informers.WithNamespace("mock"))

// 	podInformer := informer.Core().V1().Pods().Informer()

// 	// Add event handlers for Pod events
// 	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc: func(obj interface{}) {
// 			pod := obj.(*v1.Pod)
// 			fmt.Printf("Pod Added: %s Image: %s \n", pod.Name, pod.Spec.Containers[0].Image)
// 			fmt.Printf("Pod Added: %s Image: %s \n", pod.Name, pod.Spec.Containers[0].Image)
// 		},
// 		UpdateFunc: func(oldObj, newObj interface{}) {
// 			oldPod := oldObj.(*v1.Pod)
// 			newPod := newObj.(*v1.Pod)
// 			fmt.Printf("Pod Updated: %s -> %s\n", oldPod.Name, newPod.Name)
// 		},
// 		DeleteFunc: func(obj interface{}) {
// 			pod := obj.(*v1.Pod)
// 			fmt.Printf("Pod Deleted: %s\n", pod.Name)
// 		},
// 	})

// 	// Start the informer
// 	stopCh := make(chan struct{})
// 	defer close(stopCh)
// 	informer.Start(stopCh)

// 	// Wait for the informer cache to sync
// 	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
// 		log.Fatalf("Error syncing informer cache")
// 	}

// 	// Block forever
// 	<-stopCh
// }

// func deleteDeployment(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "mock"

// 	deploymentClient := client.AppsV1().Deployments(namespace)

// 	ds, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for _, d := range ds.Items {

// 		if strings.HasPrefix(d.Name, "mock-") {
// 			fmt.Printf("deployment: %s - namespace: %s\n", d.Name, d.Namespace)

// 			// deploymentClient.Delete(context.TODO(), d.Name, metav1.DeleteOptions{}) NO USAR
// 		}
// 	}

// }

// func getPods(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "eu2-mock-ns"

// 	// get the Pod interface (easy for later use)
// 	podsClient := client.CoreV1().Pods(namespace)

// 	// read all pods
// 	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

// 	// loop through pod list to get names
// 	for i, pod := range pods.Items {
// 		fmt.Printf("Name of %dth pod: %s\n", i, pod.Name)
// 	}
// }

// func createPod(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "eu2-mock-ns"

// 	// get the Pod interface (easy for later use)
// 	podsClient := client.CoreV1().Pods(namespace)

// 	podDefintion := &v1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			GenerateName: "demo-k8s-",
// 			Namespace:    namespace,
// 		},
// 		Spec: v1.PodSpec{
// 			Containers: []v1.Container{
// 				{
// 					Name:  "nginx-container",
// 					Image: "nginx:latest",
// 					Resources: v1.ResourceRequirements{
// 						Requests: v1.ResourceList{},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	newPod, err := podsClient.Create(context.TODO(), podDefintion, metav1.CreateOptions{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("Pod '%s' is created!", newPod.Name)

// }

// func updatePod(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "eu2-mock-ns"

// 	// get the Pod interface (easy for later use)
// 	podsClient := client.CoreV1().Pods(namespace)

// 	fmt.Println("Updating pod...")
// 	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() objects
// 		// retrive the latest pod
// 		currentPod, updateErr := podsClient.Get(context.TODO(), "mock-56b64bc885-5wx6b", metav1.GetOptions{})
// 		if updateErr != nil {
// 			panic(updateErr.Error())
// 		}

// 		// change container image
// 		// registry/image:v1.16.1-3
// 		currentPod.Spec.Containers[0].Image = "registry/image:v1.19"

// 		// update pod
// 		updatedPod, updateErr := podsClient.Update(context.TODO(), currentPod, metav1.UpdateOptions{})
// 		fmt.Printf("Updated pod: %s", updatedPod.Name)
// 		return updateErr
// 	})
// 	if retryErr != nil {
// 		panic(retryErr.Error())
// 	}

// }

// func deletePod(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "eu2-mock-ns"

// 	// get the Pod interface (easy for later use)
// 	podsClient := client.CoreV1().Pods(namespace)

// 	deleteErr := podsClient.Delete(context.TODO(), "mock-56b64bc885-5wx6b", metav1.DeleteOptions{})
// 	if deleteErr != nil {
// 		panic(deleteErr.Error())
// 	}

// }

// func getServices(client *kubernetes.Clientset) {
// 	// define the namespace
// 	namespace := "eu2-mock-ns"

// 	dep, err := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for _, d := range dep.Items {
// 		fmt.Printf("Service: %s\n", d.Name)
// 	}

// }
