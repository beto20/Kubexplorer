import {onMounted, ref} from "vue";
import {database, objects} from "../../wailsjs/go/models";
import PodDto = objects.PodDto;
import {fetchGetPods} from "../services/workload.service";
import DeploymentDto = objects.DeploymentDto;
import {fetchHeaderParams} from "../services/layout.service";
import HeadParamsDto = database.HeadParamsDto;


export function gridBodyPodsComposable(namespace: string) {
    const pods = ref<PodDto[]>([]);
    const headers = ref<HeadParamsDto[]>([]);

    const fetchData = async (): Promise<PodDto[]> => {
        try {
            pods.value = await fetchGetPods(namespace);
            headers.value = await fetchHeaderParams(namespace);

            return pods.value
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
        pods,
        fetchData
    }
}

export function gridBodyDeploymentsComposable(namespace: string) {
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

export function gridBodyNamespacesComposable(namespace: string) {
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