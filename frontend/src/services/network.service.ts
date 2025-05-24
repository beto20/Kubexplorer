import {
    DeleteIngressByName,
    DeleteServiceByName, GetIngressByName,
    GetIngresses,
    GetServiceByName,
    GetServices, UpdateIngressByName,
    UpdateServiceByName
} from "../../wailsjs/go/middleware/NetworkMiddleware";

export const fetchGetServices = async (namespace: string) => GetServices(namespace);
export const fetchGetServiceByName = async (name: string, namespace: string) => GetServiceByName(name, namespace);
export const fetchUpdateServiceByName = async (name: string, namespace: string, dto: any) => UpdateServiceByName(name, namespace, dto);
export const fetchDeleteServiceByName = async (name: string, namespace: string) => DeleteServiceByName(name, namespace);

export const fetchGetIngresses = async(namespace: string) => GetIngresses(namespace);
export const fetchGetIngressByName = async(name: string, namespace: string) => GetIngressByName(name, namespace);
export const fetchUpdateIngressByName = async(name: string, namespace: string, dto: any) => UpdateIngressByName(name, namespace, dto);
export const fetchDeleteIngressByName = async(name: string, namespace: string) => DeleteIngressByName(name, namespace);
