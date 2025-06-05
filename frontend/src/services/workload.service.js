import {
  DeleteDeployment,
  GetDeployment,
  GetDeployments,
  GetPod,
  GetPods,
  RestartPod,
  UpdateDeployment,
  UpdatePod,
} from '../../wailsjs/go/middleware/WorkloadMiddleware'
export const fetchGetPods = async () => GetPods()
export const fetchGetPod = async (name, namespace) => GetPod(name, namespace)
export const fetchUpdatePod = async (name, namespace, dto) => UpdatePod(name, namespace, dto)
export const fetchRestartPod = async (name, namespace) => RestartPod(name, namespace)
export const fetchGetDeployments = async () => GetDeployments()
export const fetchGetDeployment = async (name, namespace) => GetDeployment(name, namespace)
export const fetchUpdateDeployment = async (name, namespace, dto) =>
  UpdateDeployment(name, namespace, dto)
export const fetchDeleteDeployment = async (name, namespace) => DeleteDeployment(name, namespace)
//# sourceMappingURL=workload.service.js.map
