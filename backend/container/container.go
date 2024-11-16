package container

import "fmt"

type Container struct {
	services map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]interface{}),
	}
}

func (c *Container) Register(name string, service interface{}) {
	c.services[name] = service
}

func (c *Container) Resolve(name string) (interface{}, error) {
	service, exist := c.services[name]
	if !exist {
		return nil, fmt.Errorf("service %s not found", name)
	}

	return service, nil
}

func (c *Container) MustResolve(name string) interface{} {
	service, err := c.Resolve(name)
	if err != nil {
		panic(err)
	}

	return service
}
