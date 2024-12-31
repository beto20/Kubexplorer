import { ref, onMounted } from 'vue';
import { fetchCommonParameters, fetchKubernetesParameters, fetchAllEnvironments } from '../services/layout.service';
import {database, objects} from "../../wailsjs/go/models";
import EnvironmentDto = objects.EnvironmentDto;
import CommonParameterDto = database.CommonParameterDto;

export function useLayoutComposable() {
    const commonParameters = ref<CommonParameterDto[]>([]);
    const kubernetesParameters = ref<CommonParameterDto[]>([]);
    const environments = ref<EnvironmentDto[]>([]);

    const fetchData = async () => {
        try {
            commonParameters.value = await fetchCommonParameters();
            kubernetesParameters.value = await fetchKubernetesParameters();
            environments.value = await fetchAllEnvironments();
        } catch (error) {
            console.error('Error fetching environment data:', error);
            throw error;
        }
    };

    // TODO: No es necesario agregar async, porque luego se tiene un async
    // onMounted es parte del lifecycle de vuejs
    onMounted(async () => {
        console.log('onMounted triggered');
        await fetchData()
    });

    return {
        commonParameters,
        kubernetesParameters,
        environments,
        fetchData
    };
}
