package objects

type IServiceObject interface {
	GetServices(namespace string) []ServiceDto
	GetServicesMock() []ServiceDto
	GetServiceByName(name string) ServiceDto
	UpdateServiceByName(name string) bool
	DeleteServiceByName(name string) bool
}

type serviceImpl struct{}

type ServiceDto struct {
}

func NewServiceObject() IServiceObject {
	return &serviceImpl{}
}

func (s *serviceImpl) GetServices(namespace string) []ServiceDto {
	var services []ServiceDto

	return services
}

func (s *serviceImpl) GetServicesMock() []ServiceDto {
	var services []ServiceDto

	return services
}

func (s *serviceImpl) GetServiceByName(name string) ServiceDto {
	var services ServiceDto

	return services
}

func (s *serviceImpl) UpdateServiceByName(name string) bool {
	return true
}

func (s *serviceImpl) DeleteServiceByName(name string) bool {
	return true
}
