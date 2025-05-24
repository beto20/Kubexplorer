package model

type ServiceDto struct {
	Name              string
	Namespace         string
	Labels            map[string]string
	Status            string
	CreationTimestamp string
	Spec              string
}
