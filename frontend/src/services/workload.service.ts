import {
    DeleteDeploymentByName,
    GetDeploymentByName,
    GetDeployments,
    GetPodByName,
    GetPods,
    RestartPodByName, UpdateDeploymentByName,
    UpdatePodByName
} from "../../wailsjs/go/middleware/WorkloadMiddleware";

export const fetchGetPods = async() => GetPods();
export const fetchGetPodByName = async(name: string) => GetPodByName(name);
export const fetchUpdatePodByName = async(name: string) => UpdatePodByName(name);
export const fetchRestartPodByName = async(name: string) => RestartPodByName(name);

export const fetchGetDeployments = async() => GetDeployments();
export const fetchGetDeploymentByName = async(name: string) => GetDeploymentByName(name);
export const fetchUpdateDeploymentByName = async(name: string) => UpdateDeploymentByName(name);
export const fetchDeleteDeploymentByName = async(name: string) => DeleteDeploymentByName(name);