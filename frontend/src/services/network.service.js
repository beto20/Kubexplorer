import {
  DeleteIngressByName,
  DeleteServiceByName,
  GetIngressByName,
  GetIngresses,
  GetServiceByName,
  GetServices,
  UpdateIngressByName,
  UpdateServiceByName,
} from '../../wailsjs/go/middleware/NetworkMiddleware'
export const fetchGetServices = async (namespace) => GetServices(namespace)
export const fetchGetServiceByName = async (name, namespace) => GetServiceByName(name, namespace)
export const fetchUpdateServiceByName = async (name, namespace, dto) =>
  UpdateServiceByName(name, namespace, dto)
export const fetchDeleteServiceByName = async (name, namespace) =>
  DeleteServiceByName(name, namespace)
export const fetchGetIngresses = async (namespace) => GetIngresses(namespace)
export const fetchGetIngressByName = async (name, namespace) => GetIngressByName(name, namespace)
export const fetchUpdateIngressByName = async (name, namespace, dto) =>
  UpdateIngressByName(name, namespace, dto)
export const fetchDeleteIngressByName = async (name, namespace) =>
  DeleteIngressByName(name, namespace)
//# sourceMappingURL=network.service.js.map
