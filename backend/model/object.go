package model

type ObjectMapDto struct {
	Cluster                ClusterInfo
	Node                   []NodeDto
	Namespace              []NamespaceDto
	Pods                   []PodDto
	Deployments            []DeploymentDto
	Services               []ServiceDto
	Ingresses              []IngressDto
	PersistentVolumeClaims []PersistentVolumeClaimDto
	PersistentVolumes      []PersistentVolumeDto
	ConfigMaps             []ConfigMapDto
	Secrets                []Secret
	Jobs                   []PodDto
	CronJobs               []PodDto
}

type ConfigMapDto struct {
	Name      string
	Namespace string
	Values    []string
}

type Secret struct {
	Name      string
	Namespace string
	Values    []string
}
