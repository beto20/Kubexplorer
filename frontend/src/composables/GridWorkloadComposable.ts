import {onMounted, ref} from "vue";
import {database, objects} from "../../wailsjs/go/models";
import PodDto = objects.PodDto;
import {fetchGetDeployments, fetchGetPods} from "../services/workload.service";
import DeploymentDto = objects.DeploymentDto;
import {fetchHeaderParams} from "../services/layout.service";
import HeadParamsDto = database.HeadParamsDto;
import {fetchGetNamespace} from "../services/general.service";
import NamespaceDto = objects.NamespaceDto;


interface GridResponse {
    body: any,
    header: any,
    fetchData: () => Promise<any>,
}

export function gridBodyComposable(namespace: string, k8sObject: string): GridResponse {

    switch (k8sObject) {
        case "namespace": {
            console.log("Get Namespaces")
            return gridBodyNamespacesComposable(k8sObject)
        }
        case "pod": {
            console.log("Get Pod")
            return gridBodyPodsComposable(namespace, k8sObject)
        }
        case "deployment": {
            console.log("Get Deployment")
            return gridBodyDeploymentsComposable(namespace, k8sObject)
        }
        default: {
            console.log("K8sObject not found")
            return {
                body: null,
                header: null,
                fetchData: async () => {}
            }
        }
    }
}

export function gridBodyPodsComposable(namespace: string, k8sObject: string) {
    const body = ref<PodDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<any> => {
        try {
            body.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(k8sObject);
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    const res: GridResponse = {
        body: body,
        header: headers,
        fetchData: fetchData
    }

    return res
}

export function gridBodyDeploymentsComposable(namespace: string, k8sObject: string) {
    const body = ref<DeploymentDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<any> => {
        try {
            body.value = await fetchGetDeployments(namespace);
            headers.value = await fetchHeaderParams(k8sObject);
        } catch (error) {
            console.log("Error fetching deployment data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    const res: GridResponse = {
        body: body,
        header: headers,
        fetchData: fetchData
    }

    return res
}

export function gridBodyServicesComposable(namespace: string) {
    const deployments = ref<DeploymentDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<DeploymentDto[]> => {
        try {
            deployments.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(namespace);

            return deployments.value
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    return {
        deployments,
        fetchData
    }
}

export function gridBodyIngressesComposable(namespace: string) {
    const deployments = ref<DeploymentDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<DeploymentDto[]> => {
        try {
            deployments.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(namespace);

            return deployments.value
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    return {
        deployments,
        fetchData
    }
}

export function gridBodyPersistantVolumesComposable(namespace: string) {
    const deployments = ref<DeploymentDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<DeploymentDto[]> => {
        try {
            deployments.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(namespace);

            return deployments.value
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    return {
        deployments,
        fetchData
    }
}

export function gridBodyNamespacesComposable(k8sObject: string) {
    const body = ref<NamespaceDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<any> => {
        try {
            body.value = await fetchGetNamespace();
            headers.value = await fetchHeaderParams(k8sObject);

        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    const res: GridResponse = {
        body: body,
        header: headers,
        fetchData: fetchData
    }

    return res
}

export function gridBodyNodesComposable(namespace: string) {
    const deployments = ref<DeploymentDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<DeploymentDto[]> => {
        try {
            deployments.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(namespace);

            return deployments.value
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    return {
        deployments,
        fetchData
    }
}