import {
    DeleteNamespaceByName,
    GetNamespace,
    GetNamespaceByName,
    GetNodeByName,
    GetNodes,
    UpdateNamespaceByName
} from "../../wailsjs/go/middleware/GeneralMiddleware";

export const fetchGetNodes = async() => GetNodes();
export const fetchGetNodeByName = async(name: string) => GetNodeByName(name)

export const fetchGetNamespace= async() => GetNamespace();
export const fetchGetNamespaceByName= async(name: string) => GetNamespaceByName(name);
export const fetchUpdateNamespaceByName = async(name: string) => UpdateNamespaceByName(name);
export const fetchDeleteNamespaceByName = async(name: string) => DeleteNamespaceByName(name);