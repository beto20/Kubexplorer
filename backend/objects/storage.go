package objects

type IStorageObject interface {
	GetPersistentVolumes() []PersistentVolumeDto
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
