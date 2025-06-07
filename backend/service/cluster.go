package service

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

type ClusterService interface {
	BuildPodRelationships(ctx context.Context) (model.ObjectMapDto, error)
}

type clusterService struct {
	client kubernetes.Interface
}

func NewClusterService(client kubernetes.Interface) ClusterService {
	return &clusterService{client: client}
}

func (c *clusterService) BuildPodRelationships(ctx context.Context) (model.ObjectMapDto, error) {
	pods, err := c.client.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing pods: %w", err)
	}

	rs, err := c.client.AppsV1().ReplicaSets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing replicasets: %w", err)
	}

	dps, err := c.client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing deployments: %w", err)
	}

	svcs, err := c.client.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing services: %w", err)
	}

	nodes, err := c.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing nodes: %w", err)
	}

	pvcs, err := c.client.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing PVCs: %w", err)
	}

	cms, err := c.client.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing configmaps: %w", err)
	}

	secrets, err := c.client.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing secrets: %w", err)
	}

	nss, err := c.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("listing namespaces: %w", err)
	}

	objMap := &model.ObjectMapDto{
		Pods:                   model.MapPodsToDto(pods),
		ReplicaSets:            model.MapReplicaSetsToDto(rs),
		Deployments:            model.MapDeploymentsToDto(dps),
		Services:               model.MapServicesToDto(svcs),
		Nodes:                  model.MapNodesToDto(nodes),
		PersistentVolumeClaims: model.MapPVCsToDto(pvcs),
		ConfigMaps:             model.MapConfigMapsToDto(cms),
		Secrets:                model.MapSecretsToDto(secrets),
		Namespaces:             model.MapNamespacesToDto(nss),
	}

	BuildObjectMap(objMap)

	return *objMap, nil
}

// Relationship resource logic
type objectMapIndex struct {
	Pods        map[string]*model.Pod
	Deployments map[string]*model.Deployment
	ReplicaSets map[string]*model.ReplicaSet
	Nodes       map[string]*model.Node
	Namespaces  map[string]*model.Namespace
	PVCs        map[string]*model.PersistentVolumeClaim
	ConfigMaps  map[string]*model.ConfigMap
	Secrets     map[string]*model.Secret
	Jobs        map[string]*model.Job
}

func BuildObjectMap(obj *model.ObjectMapDto) *model.ObjectMapDto {
	idx := buildIndex(obj)
	resolvePodRelationships(obj, idx)
	resolveServiceRelationships(obj)
	// TODO: add other resources resolvers
	return obj
}

// Move to an index struct with maps to improve search O(1)
func buildIndex(obj *model.ObjectMapDto) *objectMapIndex {
	key := func(ns, name string) string { return ns + "/" + name }
	idx := &objectMapIndex{
		Pods:        map[string]*model.Pod{},
		Deployments: map[string]*model.Deployment{},
		ReplicaSets: map[string]*model.ReplicaSet{},
		Nodes:       map[string]*model.Node{},
		Namespaces:  map[string]*model.Namespace{},
		PVCs:        map[string]*model.PersistentVolumeClaim{},
		ConfigMaps:  map[string]*model.ConfigMap{},
		Secrets:     map[string]*model.Secret{},
		Jobs:        map[string]*model.Job{},
	}

	for i := range obj.Pods {
		p := &obj.Pods[i]
		idx.Pods[key(p.Namespace, p.Name)] = p
	}

	for i := range obj.Deployments {
		d := &obj.Deployments[i]
		idx.Deployments[key(d.Namespace, d.Name)] = d
	}

	for i := range obj.ReplicaSets {
		rs := &obj.ReplicaSets[i]
		idx.ReplicaSets[key(rs.Namespace, rs.Name)] = rs
	}

	for i := range obj.Nodes {
		n := &obj.Nodes[i]
		idx.Nodes[n.Name] = n
	}

	for i := range obj.Namespaces {
		ns := &obj.Namespaces[i]
		idx.Namespaces[ns.Name] = ns
	}

	for i := range obj.PersistentVolumeClaims {
		pvc := &obj.PersistentVolumeClaims[i]
		idx.PVCs[key(pvc.Namespace, pvc.Name)] = pvc
	}

	for i := range obj.ConfigMaps {
		cm := &obj.ConfigMaps[i]
		idx.ConfigMaps[key(cm.Namespace, cm.Name)] = cm
	}

	for i := range obj.Secrets {
		s := &obj.Secrets[i]
		idx.Secrets[key(s.Namespace, s.Name)] = s
	}

	for i := range obj.Jobs {
		j := &obj.Jobs[i]
		idx.Jobs[key(j.Namespace, j.Name)] = j
	}

	return idx
}

func resolvePodRelationships(obj *model.ObjectMapDto, idx *objectMapIndex) {
	for i := range obj.Pods {
		pod := &obj.Pods[i]

		if node := idx.Nodes[pod.NodeName]; node != nil {
			node.PodNames = append(node.PodNames, pod.Name)
		}

		if ns := idx.Namespaces[pod.Namespace]; ns != nil {
			ns.PodNames = append(ns.PodNames, pod.Name)
		}

		switch pod.OwnerKind {
		case "ReplicaSet":
			if rs := idx.ReplicaSets[pod.Namespace+"/"+pod.OwnerName]; rs != nil {
				rs.PodNames = append(rs.PodNames, pod.Name)
				pod.Deployment = rs.Deployment
				if dep := idx.Deployments[pod.Namespace+"/"+rs.Deployment]; dep != nil {
					dep.PodNames = append(dep.PodNames, pod.Name)
				}
			}

		case "Job":
			if job := idx.Jobs[pod.Namespace+"/"+pod.OwnerName]; job != nil {
				job.PodNames = append(job.PodNames, pod.Name)
			}
		}

		for _, pvcName := range pod.PVCNames {
			if pvc := idx.PVCs[pod.Namespace+"/"+pvcName]; pvc != nil {
				pvc.UsedByPods = append(pvc.UsedByPods, pod.Name)
			}
		}

		for _, cmName := range pod.ConfigMapRefs {
			if cm := idx.ConfigMaps[pod.Namespace+"/"+cmName]; cm != nil {
				cm.UsedByPods = append(cm.UsedByPods, pod.Name)
			}
		}

		for _, secName := range pod.SecretRefs {
			if sec := idx.Secrets[pod.Namespace+"/"+secName]; sec != nil {
				sec.UsedByPods = append(sec.UsedByPods, pod.Name)
			}
		}
	}
}

func resolveServiceRelationships(obj *model.ObjectMapDto) {
	for i := range obj.Services {
		svc := &obj.Services[i]

		for j := range obj.Pods {
			pod := &obj.Pods[j]

			if pod.Namespace != svc.Namespace {
				continue
			}

			if selectorLabelsMatcher(svc.Selector, pod.Labels) {
				svc.PodNames = append(svc.PodNames, pod.Name)
				pod.ServiceRefs = append(pod.ServiceRefs, svc.Name)
			}
		}
	}
}

func selectorLabelsMatcher(selector map[string]string, labels map[string]string) bool {
	if len(selector) == 0 {
		return false
	}

	for k, v := range selector {
		if labels[k] != v {
			return false
		}
	}

	return true
}
