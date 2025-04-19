import {
    DeleteNamespaceByName,
    GetNamespace,
    GetNamespaceByName,
    GetNodeByName,
    GetNodes,
    UpdateNamespaceByName
} from "../../wailsjs/go/middleware/GeneralMiddleware";
import {objects} from "../../wailsjs/go/models";
import NamespaceDto = objects.NamespaceDto;

export const fetchGetNodes = async() => GetNodes();
export const fetchGetNodeByName = async(name: string) => GetNodeByName(name)

export const fetchGetNamespace= async() : Promise<NamespaceDto[]> => GetNamespace();
export const fetchGetNamespaceByName= async(name: string) => GetNamespaceByName(name);
export const fetchUpdateNamespaceByName = async(name: string) => UpdateNamespaceByName(name);
export const fetchDeleteNamespaceByName = async(name: string) => DeleteNamespaceByName(name);