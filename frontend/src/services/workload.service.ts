import {
    DeleteDeployment,
    GetDeployment,
    GetDeployments,
    GetPod,
    GetPods,
    RestartPod, UpdateDeployment,
    UpdatePod
} from "../../wailsjs/go/middleware/WorkloadMiddleware";

export const fetchGetPods = async() => GetPods();
export const fetchGetPod = async(name: string) => GetPod(name);
export const fetchUpdatePod = async(name: string) => UpdatePod(name);
export const fetchRestartPod = async(name: string) => RestartPod(name);

export const fetchGetDeployments = async() => GetDeployments();
export const fetchGetDeployment = async(name: string) => GetDeployment(name);
export const fetchUpdateDeployment = async(name: string) => UpdateDeployment(name);
export const fetchDeleteDeployment = async(name: string) => DeleteDeployment(name);