import {
    DeleteIngressByName,
    DeleteServiceByName, GetIngressByName,
    GetIngresses,
    GetServiceByName,
    GetServices, UpdateIngressByName,
    UpdateServiceByName
} from "../../wailsjs/go/middleware/NetworkMiddleware";

export const fetchGetServices = async () => GetServices();
export const fetchGetServiceByName = async (name: string) => GetServiceByName(name);
export const fetchUpdateServiceByName = async (name: string) => UpdateServiceByName(name);
export const fetchDeleteServiceByName = async (name: string) => DeleteServiceByName(name);

export const fetchGetIngresses = async() => GetIngresses();
export const fetchGetIngressByName = async(name: string) => GetIngressByName(name);
export const fetchUpdateIngressByName = async(name: string) => UpdateIngressByName(name);
export const fetchDeleteIngressByName = async(name: string) => DeleteIngressByName(name);
