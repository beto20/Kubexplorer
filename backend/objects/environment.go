package objects

type IEnvironmentObject interface {
	GetCurrentEnvironment(envName string) EnvironmentDto
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

func (e *environmentImpl) GetCurrentEnvironment(envName string) EnvironmentDto {
	dto := EnvironmentDto{
		Name:        "test",
		Description: "test description",
		Env:         "Dev",
		Status:      true,
	}

	return dto
}

func (e *environmentImpl) GetAllEnvironment() []EnvironmentDto {
	dto := EnvironmentDto{
		Name:        "test",
		Description: "test description",
		Env:         "Dev",
		Status:      true,
	}

	var res []EnvironmentDto
	res = append(res, dto)

	return res
}
