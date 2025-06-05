import {
  DeletePersistentVolumeByName,
  GetPersistentVolumeByName,
  GetPersistentVolumes,
  UpdatePersistentVolumeByName,
} from '../../wailsjs/go/middleware/StorageMiddleware'
export const fetchGetPersistentVolumes = async () => GetPersistentVolumes()
export const fetchGetPersistentVolumeByName = async (name) => GetPersistentVolumeByName(name)
export const fetchUpdatePersistentVolumeByName = async (name, dto) =>
  UpdatePersistentVolumeByName(name, dto)
export const fetchDeletePersistentVolumeByName = async (name) => DeletePersistentVolumeByName(name)
//# sourceMappingURL=storage.service.js.map
