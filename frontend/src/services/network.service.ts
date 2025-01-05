import {GetIngresses, GetServices} from "../../wailsjs/go/middleware/ServiceMiddleware";

export const fetchGetServices = async () => GetServices();
export const fetchGetIngresses = async() => GetIngresses();