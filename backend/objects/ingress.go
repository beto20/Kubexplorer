package objects

type IIngressObject interface {
	GetIngresses(namespace string) []IngressDto
	GetIngressByName(name string) IngressDto
	UpdateIngressByName(name string) bool
	DeleteIngressByName(name string) bool
}

type ingressImpl struct{}

type IngressDto struct {
}

func NewIngressObject() IIngressObject {
	return &ingressImpl{}
}

func (i *ingressImpl) GetIngresses(namespace string) []IngressDto {
	var ingress []IngressDto

	return ingress
}

func (i *ingressImpl) GetIngressByName(name string) IngressDto {
	var ingress IngressDto

	return ingress
}

func (i *ingressImpl) UpdateIngressByName(name string) bool {
	return true
}

func (i *ingressImpl) DeleteIngressByName(name string) bool {
	return true
}
