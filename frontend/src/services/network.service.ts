import {GetIngresses, GetServices} from "../../wailsjs/go/middleware/NetworkMiddleware";

export const fetchGetServices = async () => GetServices();
export const fetchGetIngresses = async() => GetIngresses();