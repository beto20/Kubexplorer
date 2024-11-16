package objects

type IServiceObject interface {
	GetServices(namespace string) []ServiceDto
	GetServicesMock(namespace string) []ServiceDto
}

type serviceImpl struct{}

type ServiceDto struct {
}

func NewServiceObject() IServiceObject {
	return serviceImpl{}
}

func (s serviceImpl) GetServices(namespace string) []ServiceDto {
	var services []ServiceDto

	return services
}

func (s serviceImpl) GetServicesMock(namespace string) []ServiceDto {
	var services []ServiceDto

	return services
}
