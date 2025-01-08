package objects

type INamespaceObject interface {
	GetNamespace() []NamespaceDto
	GetNamespaceByName(name string) NamespaceDto
	UpdateNamespaceByName(name string) bool
	DeleteNamespaceByName(name string) bool
}

type namespaceImpl struct {
}

func NewNamespaceObject() INamespaceObject {
	return &namespaceImpl{}
}

type NamespaceDto struct {
	Name     string
	Resource Resource
	Roles    []string
	Version  string
	Age      string
	Status   bool
}

func (n *namespaceImpl) GetNamespace() []NamespaceDto {
	return []NamespaceDto{
		{
			Name: "node-1",
			Resource: Resource{
				Cpu:    "6Gi",
				Memory: "20Gi",
			},
			Roles: []string{
				"ADMIN",
				"GENERAL",
			},
			Version: "1.24.0",
			Age:     "20 day",
			Status:  true,
		},
		{
			Name: "node-2",
			Resource: Resource{
				Cpu:    "6Gi",
				Memory: "20Gi",
			},
			Roles: []string{
				"ADMIN",
				"GENERAL",
			},
			Version: "1.28.0",
			Age:     "100 day",
			Status:  true,
		},
	}
}
func (n *namespaceImpl) GetNamespaceByName(name string) NamespaceDto {
	return NamespaceDto{
		Name: "node-1",
		Resource: Resource{
			Cpu:    "6Gi",
			Memory: "20Gi",
		},
		Roles: []string{
			"ADMIN",
			"GENERAL",
		},
		Version: "1.24.0",
		Age:     "20 day",
		Status:  true,
	}
}

func (n *namespaceImpl) UpdateNamespaceByName(name string) bool {
	return true
}

func (n *namespaceImpl) DeleteNamespaceByName(name string) bool {
	return true
}
