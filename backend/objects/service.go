package objects

type IServiceObject interface {
	GetServices(namespace string) []ServiceDto
	GetServicesMock(namespace string) []ServiceDto
	GetIngresses(namespace string) []IngressDto
}

type serviceImpl struct{}

type ServiceDto struct {
}

type IngressDto struct {
}

func NewServiceObject() IServiceObject {
	return &serviceImpl{}
}

func (s *serviceImpl) GetServices(namespace string) []ServiceDto {
	var services []ServiceDto

	return services
}

func (s *serviceImpl) GetServicesMock(namespace string) []ServiceDto {
	var services []ServiceDto

	return services
}

func (s *serviceImpl) GetIngresses(namespace string) []IngressDto {
	var ingress []IngressDto

	return ingress
}
