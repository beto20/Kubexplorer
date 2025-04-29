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
