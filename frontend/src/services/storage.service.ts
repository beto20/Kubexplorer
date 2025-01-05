import {GetPersistentVolumes} from "../../wailsjs/go/middleware/StorageMiddleware";

export const fetchGetPersistentVolumes = async() => GetPersistentVolumes();