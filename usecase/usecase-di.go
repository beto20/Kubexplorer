package usecase

type K8sObject interface {
	GetAll()
	GetDetailsById()
	DeleteOneById()
	EditOneById()
}

type Sql interface {
	GetAll()
}
