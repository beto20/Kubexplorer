package model

type NamespaceDto struct {
	Name         string
	Version      string
	CreationTime string
	Labels       map[string]string
	Status       string
}
