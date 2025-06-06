package kubeclient

import (
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

type ClusterManager struct {
	mu             sync.RWMutex
	clients        map[string]*kubernetes.Clientset
	dynamicClients map[string]*dynamic.DynamicClient
}

func NewClusterManager() *ClusterManager {
	return &ClusterManager{
		clients: make(map[string]*kubernetes.Clientset),
	}
}

func (cm *ClusterManager) GetClient(clusterName string, kubeConfigPath string) (*kubernetes.Clientset, error) {
	cm.mu.RLock()
	if client, exists := cm.clients[clusterName]; exists {
		cm.mu.RUnlock()
		return client, nil
	}
	cm.mu.RUnlock()

	cm.mu.Lock()
	defer cm.mu.Unlock()

	if client, exists := cm.clients[clusterName]; exists {
		return client, nil
	}

	config, err := buildRestConfig(kubeConfigPath, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error building config: %s", err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building clientSet: %s", err)
	}

	cm.clients[clusterName] = clientSet

	return clientSet, nil
}

func (cm *ClusterManager) GetDynamicClient(clusterName string, kubeConfigPath string) (*dynamic.DynamicClient, error) {
	cm.mu.RLock()
	if dynamicClient, exists := cm.dynamicClients[clusterName]; exists {
		cm.mu.RUnlock()
		return dynamicClient, nil
	}
	cm.mu.RUnlock()

	cm.mu.Lock()
	defer cm.mu.Unlock()

	if dynamicClient, exists := cm.dynamicClients[clusterName]; exists {
		return dynamicClient, nil
	}

	config, err := buildRestConfig(kubeConfigPath, clusterName)
	if err != nil {
		return nil, fmt.Errorf("error building config: %s", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building dynamicClient: %s", err)
	}

	cm.dynamicClients[clusterName] = dynamicClient

	return dynamicClient, nil
}

func buildRestConfig(kubeConfigPath, context string) (*rest.Config, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	return clientConfig.ClientConfig()
}
