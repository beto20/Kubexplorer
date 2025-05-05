package model

type ClusterInfo struct {
	Name      string
	Server    string
	Current   bool
	User      string
	Namespace string
}

type ClusterProfile struct {
	ContextName    string
	KubeConfigPath string
}

type EnvironmentDto struct {
	Name        string
	Description string
	Env         string
	Status      bool
}

type NodeDto struct {
	Name     string
	Resource Resource
	Roles    []string
	Version  string
	Age      string
	Status   bool
}
