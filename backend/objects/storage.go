package objects

type IStorageObject interface {
	GetPersistentVolumes() []PersistentVolumeDto
	GetPersistentVolumeByName() PersistentVolumeDto
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

func (s *storageImpl) GetPersistentVolumeByName() PersistentVolumeDto {
	return PersistentVolumeDto{}
}
