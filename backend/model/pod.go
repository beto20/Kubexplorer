package model

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
