package objects

type IStorageObject interface {
	GetPersistentVolumes() []PersistentVolumeDto
	GetPersistentVolumeByName(name string) PersistentVolumeDto
	UpdatePersistentVolumeByName(name string) bool
	DeletePersistentVolumeByName(name string) bool
}

type storageImpl struct {
}

type PersistentVolumeDto struct {
}

func NewStorageObject() IStorageObject {
	return &storageImpl{}
}

func (s *storageImpl) GetPersistentVolumes() []PersistentVolumeDto {
	return []PersistentVolumeDto{}
}

func (s *storageImpl) GetPersistentVolumeByName(name string) PersistentVolumeDto {
	return PersistentVolumeDto{}
}

func (s *storageImpl) UpdatePersistentVolumeByName(name string) bool {
	return true
}

func (s *storageImpl) DeletePersistentVolumeByName(name string) bool {
	return true
}
