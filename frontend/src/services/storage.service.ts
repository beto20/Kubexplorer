import {
  DeletePersistentVolumeByName,
  GetPersistentVolumeByName,
  GetPersistentVolumes,
  UpdatePersistentVolumeByName,
} from '../../wailsjs/go/middleware/StorageMiddleware'

export const fetchGetPersistentVolumes = async () => GetPersistentVolumes()
export const fetchGetPersistentVolumeByName = async (name: string) =>
  GetPersistentVolumeByName(name)
export const fetchUpdatePersistentVolumeByName = async (name: string, dto: any) =>
  UpdatePersistentVolumeByName(name, dto)
export const fetchDeletePersistentVolumeByName = async (name: string) =>
  DeletePersistentVolumeByName(name)
