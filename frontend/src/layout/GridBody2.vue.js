import { computed, defineComponent, onMounted, reactive, ref, toRefs } from 'vue'
import { gridComposable } from '../composables/GridComposable'
import KsSidebarDetail from './SidebarDetail.vue'
import KsGridTable from '../components/GridTableComponent.vue'
import KsGridHeader from '../components/GridHeaderComponent.vue'
export default defineComponent({
  name: 'KsGridBodyV2',
  components: { KsGridHeader, KsGridTable, KsSidebarDetail },
  props: {
    k8sObject: { type: String, required: true },
    namespace: { type: String, required: true },
  },
  setup(props) {
    const response = gridComposable(props.k8sObject)
    const namespaces = ['ns-local', 'ns-dev']
    const statuses = ['Alive', 'Inactive']
    const state = reactive({
      items: [],
      search: '',
      filterNamespace: '',
      filterStatus: '',
      sortBy: [{ key: 'name', order: 'asc' }],
    })
    const header = reactive({
      header: [],
    })
    const filteredItems = computed(() => {
      return state.items.filter((item) => {
        const namespaceMatch =
          !state.filterNamespace || item.namespace.includes(state.filterNamespace)
        const statusMatch = !state.filterStatus || item.status.includes(state.filterStatus)
        const searchMatch =
          !state.search || item.name.toLowerCase().includes(state.search.toLowerCase())
        return namespaceMatch && statusMatch && searchMatch
      })
    })
    const isSidebarVisible = ref(false)
    const selectedRow = ref(null)
    const onRowClick = (cellData, item) => {
      selectedRow.value = item.item
      console.log('selectedRow.value', item.item)
      if (isSidebarVisible.value) {
        console.log('FALSE')
        isSidebarVisible.value = false
        setTimeout(() => {
          console.log('Delayed')
          isSidebarVisible.value = true
        }, 100)
      } else {
        console.log('FALSE')
        isSidebarVisible.value = true
      }
    }
    onMounted(async () => {
      await response.fetchData()
      header.header = response.content?.head.value
      state.items = response.content?.body.value
    })
    return {
      namespaces,
      header,
      statuses,
      ...toRefs(state),
      filteredItems,
      isSidebarVisible,
      selectedRow,
      onRowClick,
    }
  },
})
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
const __VLS_componentsOption = { KsGridHeader, KsGridTable, KsSidebarDetail }
let __VLS_components
let __VLS_directives
const __VLS_0 = {}.VContainer
/** @type {[typeof __VLS_components.VContainer, typeof __VLS_components.vContainer, typeof __VLS_components.VContainer, typeof __VLS_components.vContainer, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({}))
const __VLS_2 = __VLS_1({}, ...__VLS_functionalComponentArgsRest(__VLS_1))
var __VLS_4 = {}
__VLS_3.slots.default
const __VLS_5 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_6 = __VLS_asFunctionalComponent(__VLS_5, new __VLS_5({}))
const __VLS_7 = __VLS_6({}, ...__VLS_functionalComponentArgsRest(__VLS_6))
__VLS_8.slots.default
const __VLS_9 = {}.KsGridHeader
/** @type {[typeof __VLS_components.KsGridHeader, ]} */ // @ts-ignore
const __VLS_10 = __VLS_asFunctionalComponent(
  __VLS_9,
  new __VLS_9({
    search: __VLS_ctx.search,
    filterNamespace: __VLS_ctx.filterNamespace,
    filterStatus: __VLS_ctx.filterStatus,
    namespaces: __VLS_ctx.namespaces,
    statuses: __VLS_ctx.statuses,
  }),
)
const __VLS_11 = __VLS_10(
  {
    search: __VLS_ctx.search,
    filterNamespace: __VLS_ctx.filterNamespace,
    filterStatus: __VLS_ctx.filterStatus,
    namespaces: __VLS_ctx.namespaces,
    statuses: __VLS_ctx.statuses,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_10),
)
const __VLS_13 = {}.KsGridTable
/** @type {[typeof __VLS_components.KsGridTable, ]} */ // @ts-ignore
const __VLS_14 = __VLS_asFunctionalComponent(
  __VLS_13,
  new __VLS_13({
    ...{ 'onClick:row': {} },
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
  }),
)
const __VLS_15 = __VLS_14(
  {
    ...{ 'onClick:row': {} },
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_14),
)
let __VLS_17
let __VLS_18
let __VLS_19
const __VLS_20 = {
  'onClick:row': __VLS_ctx.onRowClick,
}
var __VLS_16
var __VLS_8
const __VLS_21 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_22 = __VLS_asFunctionalComponent(__VLS_21, new __VLS_21({}))
const __VLS_23 = __VLS_22({}, ...__VLS_functionalComponentArgsRest(__VLS_22))
__VLS_24.slots.default
const __VLS_25 = {}.KsSidebarDetail
/** @type {[typeof __VLS_components.KsSidebarDetail, ]} */ // @ts-ignore
const __VLS_26 = __VLS_asFunctionalComponent(
  __VLS_25,
  new __VLS_25({
    ...{ onClose: {} },
    isVisible: __VLS_ctx.isSidebarVisible,
    selectedRow: __VLS_ctx.selectedRow,
  }),
)
const __VLS_27 = __VLS_26(
  {
    ...{ onClose: {} },
    isVisible: __VLS_ctx.isSidebarVisible,
    selectedRow: __VLS_ctx.selectedRow,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_26),
)
let __VLS_29
let __VLS_30
let __VLS_31
const __VLS_32 = {
  onClose: (...[$event]) => {
    __VLS_ctx.isSidebarVisible = false
  },
}
var __VLS_28
var __VLS_24
var __VLS_3
var __VLS_dollars
let __VLS_self
//# sourceMappingURL=GridBody2.vue.js.map
