package model

type NamespaceDto struct {
	Name     string
	Resource Resource
	Roles    []string
	Version  string
	Age      string
	Status   bool
}
