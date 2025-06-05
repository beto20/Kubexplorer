import { ref, onMounted } from 'vue'
import {
  fetchCommonParameters,
  fetchKubernetesParameters,
  fetchAllEnvironments,
} from '../services/layout.service'
export function sidebarComposable() {
  const commonParameters = ref([])
  const kubernetesParameters = ref([])
  const environments = ref([])
  const fetchData = async () => {
    try {
      commonParameters.value = await fetchCommonParameters()
      kubernetesParameters.value = await fetchKubernetesParameters()
      environments.value = await fetchAllEnvironments()
      console.log('environments.value: ', environments.value)
    } catch (error) {
      console.error('Error fetching environment data:', error)
      throw error
    }
  }
  // TODO: No es necesario agregar async, porque luego se tiene un async
  // onMounted es parte del lifecycle de vuejs
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  return {
    commonParameters,
    kubernetesParameters,
    environments,
    fetchData,
  }
}
//# sourceMappingURL=SidebarComposable.js.map
