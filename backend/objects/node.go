package objects

type INodeObject interface {
	GetNodes() []NodeDto
	GetNodeByName(name string) NodeDto
}

type nodeImpl struct {
}

func NewNodeObject() INodeObject {
	return &nodeImpl{}
}

type NodeDto struct {
	Name     string
	Resource Resource
	Roles    []string
	Version  string
	Age      string
	Status   bool
}

func (n *nodeImpl) GetNodes() []NodeDto {
	return []NodeDto{
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

func (n *nodeImpl) GetNodeByName(name string) NodeDto {
	return NodeDto{
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
