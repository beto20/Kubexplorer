import {onMounted, ref} from "vue";
import {fetchRestartPod} from "../services/workload.service";


interface DeleteResponse {
    fetchData: () => Promise<any>,
    isRestarted?: any
}

export function gridGeneralComposable(podName: string): DeleteResponse {
    return restartPod(podName)
}


export function restartPod(podName: string) {
    const isRestarted = ref<any>();

    const fetchData = async () => {
        try {
            isRestarted.value = await fetchRestartPod(podName, "east") // TODO: replace with a dynamic value
            console.log("VAL", isRestarted.value)
        } catch (error) {
            console.log("Error fetching pod data: ", error);
            throw error;
        }
    }

    onMounted(async () => {
        console.log("onMounted triggered");
        await fetchData()
    })

    const response: DeleteResponse = {
        fetchData: async () => {},
        isRestarted: isRestarted
    }

    return response
}
