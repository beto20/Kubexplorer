import { onMounted, ref } from 'vue'
import { fetchGetDeployments, fetchGetPods } from '../services/workload.service'
import { fetchHeaderParams } from '../services/layout.service'
import { fetchGetNamespace, fetchGetNodes } from '../services/general.service'
import { fetchGetIngresses, fetchGetServices } from '../services/network.service'
import { fetchGetPersistentVolumes } from '../services/storage.service'
export function gridComposable(k8sObject) {
  switch (k8sObject) {
    case 'node': {
      console.log('Get Nodes')
      return gridBodyNodes(k8sObject)
    }
    case 'namespace': {
      console.log('Get Namespaces')
      return gridBodyNamespaces(k8sObject)
    }
    case 'service': {
      console.log('Get Services')
      return gridBodyServices(k8sObject)
    }
    case 'ingress': {
      console.log('Get Ingresses')
      return gridBodyIngresses(k8sObject)
    }
    case 'persistentVolumen': {
      console.log('Get Persistent Volumen')
      return gridBodyPersistentVolumes(k8sObject)
    }
    case 'deployment': {
      console.log('Get Deployments')
      return gridBodyDeployments(k8sObject)
    }
    case 'pod': {
      console.log('Get Pods')
      return gridBodyPods(k8sObject)
    }
    default: {
      console.log('K8s object not found')
      return {
        fetchData: async () => {},
        content: {
          head: null,
          body: null,
        },
      }
    }
  }
}
export function gridBodyPods(k8sObject) {
  const head = ref([])
  const body = ref([])
  const fetchData = async () => {
    try {
      const pods = ref([])
      const header = ref([])
      pods.value = await fetchGetPods()
      header.value = await fetchHeaderParams(k8sObject)
      console.log('PODS: ', pods.value)
      body.value = pods.value.map((p) => ({
        name: p.Name,
        namespace: p.Namespace,
        replicas: p.Replicas,
        cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
        memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
        age: p.Age,
        status: p.Status,
      }))
      head.value = header.value.map((h) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    // await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: head,
      body: body,
    },
  }
  return response
}
export function gridBodyDeployments(k8sObject) {
  const head = ref([])
  const body = ref([])
  const fetchData = async () => {
    try {
      const deployment = ref([])
      const header = ref([])
      deployment.value = await fetchGetDeployments()
      header.value = await fetchHeaderParams(k8sObject)
      console.log('DEPLOYMENTS', body.value)
      body.value = deployment.value.map((d) => ({
        name: d.Name,
        namespace: d.Namespace,
        status: d.Status,
        age: d.Age,
      }))
      head.value = header.value.map((h) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching deployment data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: head,
      body: body,
    },
  }
  return response
}
export function gridBodyServices(k8sObject) {
  const body = ref([])
  const header = ref([])
  const fetchData = async () => {
    try {
      body.value = await fetchGetServices('east') // TODO: replace with a dynamic value
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }
  return response
}
export function gridBodyIngresses(k8sObject) {
  const body = ref([])
  const header = ref([])
  const fetchData = async () => {
    try {
      body.value = await fetchGetIngresses('east') // TODO: replace with a dynamic value
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }
  return response
}
export function gridBodyPersistentVolumes(k8sObject) {
  const body = ref([])
  const header = ref([])
  const fetchData = async () => {
    try {
      body.value = await fetchGetPersistentVolumes()
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }
  return response
}
export function gridBodyNamespaces(k8sObject) {
  const head = ref([])
  const body = ref([])
  const fetchData = async () => {
    try {
      const namespaces = ref([])
      const header = ref([])
      namespaces.value = await fetchGetNamespace()
      header.value = await fetchHeaderParams(k8sObject)
      console.log('NAMESPACES: ', namespaces.value)
      body.value = namespaces.value.map((p) => ({
        name: p.Name,
        label: '',
        age: p.Age,
        status: p.Status,
      }))
      head.value = header.value.map((h) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    // await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: head,
      body: body,
    },
  }
  return response
}
export function gridBodyNodes(namespace) {
  const body = ref([])
  const header = ref([])
  const fetchData = async () => {
    try {
      body.value = await fetchGetNodes()
      header.value = await fetchHeaderParams(namespace)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })
  const response = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }
  return response
}
//# sourceMappingURL=GridComposable.js.map
