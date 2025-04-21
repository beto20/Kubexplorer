package kubeclient

//
//import (
//	"fmt"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/rest"
//	"k8s.io/client-go/tools/clientcmd"
//	"sync"
//)
//
//type ClusterManager struct {
//	mu      sync.RWMutex
//	clients map[string]*kubernetes.Clientset
//}
//
//func NewClusterManager() *ClusterManager {
//	return &ClusterManager{
//		clients: make(map[string]*kubernetes.Clientset),
//	}
//}
//
//func (cm *ClusterManager) GetClient(clusterId string, kubeConfigPath string) (*kubernetes.Clientset, error) {
//	cm.mu.RLock()
//	if client, exists := cm.clients[clusterId]; exists {
//		cm.mu.RUnlock()
//		return client, nil
//	}
//	cm.mu.Unlock()
//
//	cm.mu.Lock()
//	defer cm.mu.Unlock()
//
//	if client, exists := cm.clients[clusterId]; exists {
//		return client, nil
//	}
//
//	config, err := buildRestConfig(kubeConfigPath, clusterId)
//	if err != nil {
//		return nil, fmt.Errorf("error building config: %s", err)
//	}
//
//	clientSet, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		return nil, fmt.Errorf("error building clientSet: %s", err)
//	}
//
//	cm.clients[clusterId] = clientSet
//
//	return clientSet, nil
//}
//
//func buildRestConfig(kubeConfigPath, context string) (*rest.Config, error) {
//	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
//	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
//	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
//
//	return clientConfig.ClientConfig()
//}
