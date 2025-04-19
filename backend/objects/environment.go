package objects

type IEnvironmentObject interface {
	GetCurrentEnvironment(env string, name string) EnvironmentDto
	GetAllEnvironment() []EnvironmentDto
}

type environmentImpl struct{}

func NewEnvironmentObject() IEnvironmentObject {
	return &environmentImpl{}
}

type EnvironmentDto struct {
	Name        string
	Description string
	Env         string
	Status      bool
}

func (e *environmentImpl) GetCurrentEnvironment(env string, name string) EnvironmentDto {
	dto := EnvironmentDto{
		Name:        "minikube",
		Description: "minikube description",
		Env:         "Dev",
		Status:      true,
	}

	return dto
}

func (e *environmentImpl) GetAllEnvironment() []EnvironmentDto {
	return []EnvironmentDto{
		{
			Name:        "minikube",
			Description: "minikube description",
			Env:         "Dev",
			Status:      true,
		},
		{
			Name:        "minikube-2",
			Description: "minikube description",
			Env:         "UAT",
			Status:      true,
		},
		{
			Name:        "minikube-3",
			Description: "minikube description",
			Env:         "PRD",
			Status:      false,
		},
		{
			Name:        "minikube-4",
			Description: "minikube description",
			Env:         "PRD-2",
			Status:      true,
		},
		{
			Name:        "minikube-5",
			Description: "minikube description",
			Env:         "Sandbox",
			Status:      false,
		},
	}
}
