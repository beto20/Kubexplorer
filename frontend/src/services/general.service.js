import {
  DeleteNamespaceByName,
  GetNamespace,
  GetNamespaceByName,
  GetNodeByName,
  GetNodes,
  UpdateNamespaceByName,
} from '../../wailsjs/go/middleware/GeneralMiddleware'
export const fetchGetNodes = async () => GetNodes()
export const fetchGetNodeByName = async (name) => GetNodeByName(name)
export const fetchGetNamespace = async () => GetNamespace()
export const fetchGetNamespaceByName = async (name) => GetNamespaceByName(name)
export const fetchUpdateNamespaceByName = async (name, dto) => UpdateNamespaceByName(name, dto)
export const fetchDeleteNamespaceByName = async (name) => DeleteNamespaceByName(name)
//# sourceMappingURL=general.service.js.map
