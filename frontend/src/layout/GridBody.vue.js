import { computed, defineComponent, onMounted, reactive, ref, toRefs } from 'vue'
import KsSidebarDetail from './SidebarDetail.vue'
export default defineComponent({
  name: 'ks-grid-body',
  components: { KsSidebarDetail },
  props: {
    k8sObject: {
      type: String,
      // required: true
    },
    namespace: {
      type: String,
      // required: true
    },
  },
  setup(props) {
    const editPod = (item) => {
      console.log('EDIT', item)
    }
    const deletePod = (item) => {
      console.log('DELETE', item)
    }
    // const { pods, headers, fetchData } = gridBodyPodsComposable("mock", "pod");
    // const { nodes, namespaces, headers, fetchData } = gridGeneralComposable(props.k8sObject);
    const namespaces = ['ns-local', 'ns-dev']
    const statuses = ['Alive', 'Inactive']
    const state = reactive({
      // header: [],
      items: [],
      search: '',
      filterNamespace: null,
      filterStatus: null,
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
      // await fetchData();
      // state.items = pods.value.map((p) => ({
      //     name: p.Name,
      //     namespace: p.Namespace,
      //     replicas: p.Replicas,
      //     cpu: p.Container.Limit.Cpu + '/' + p.Container.Request.Cpu,
      //     memory: p.Container.Limit.Memory + '/' + p.Container.Request.Memory,
      //     age: p.Age,
      //     status: p.Status,
      // }));
      // header.header = headers.value.map((h) => ({
      //     title: h.Title,
      //     key: h.Key,
      //     align: h.Align,
      //     sortable: h.Sortable,
      // }));
    })
    return {
      header,
      namespaces,
      statuses,
      ...toRefs(state),
      editPod,
      deletePod,
      filteredItems,
      isSidebarVisible,
      selectedRow,
      onRowClick,
    }
  },
})
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
const __VLS_componentsOption = { KsSidebarDetail }
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
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
const __VLS_9 = {}.VCardTitle
/** @type {[typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, ]} */ // @ts-ignore
const __VLS_10 = __VLS_asFunctionalComponent(__VLS_9, new __VLS_9({}))
const __VLS_11 = __VLS_10({}, ...__VLS_functionalComponentArgsRest(__VLS_10))
__VLS_12.slots.default
const __VLS_13 = {}.VSpacer
/** @type {[typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, ]} */ // @ts-ignore
const __VLS_14 = __VLS_asFunctionalComponent(__VLS_13, new __VLS_13({}))
const __VLS_15 = __VLS_14({}, ...__VLS_functionalComponentArgsRest(__VLS_14))
const __VLS_17 = {}.VTextField
/** @type {[typeof __VLS_components.VTextField, typeof __VLS_components.vTextField, typeof __VLS_components.VTextField, typeof __VLS_components.vTextField, ]} */ // @ts-ignore
const __VLS_18 = __VLS_asFunctionalComponent(
  __VLS_17,
  new __VLS_17({
    modelValue: __VLS_ctx.search,
    label: 'Search',
    ...{ class: 'mx-4' },
    clearable: true,
    'aria-label': 'Search table',
  }),
)
const __VLS_19 = __VLS_18(
  {
    modelValue: __VLS_ctx.search,
    label: 'Search',
    ...{ class: 'mx-4' },
    clearable: true,
    'aria-label': 'Search table',
  },
  ...__VLS_functionalComponentArgsRest(__VLS_18),
)
var __VLS_12
const __VLS_21 = {}.VDataTableVirtual
/** @type {[typeof __VLS_components.VDataTableVirtual, typeof __VLS_components.vDataTableVirtual, typeof __VLS_components.VDataTableVirtual, typeof __VLS_components.vDataTableVirtual, ]} */ // @ts-ignore
const __VLS_22 = __VLS_asFunctionalComponent(
  __VLS_21,
  new __VLS_21({
    ...{ 'onClick:row': {} },
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
    height: '720',
    itemValue: 'name',
    density: 'compact',
    fixedHeader: true,
    hover: true,
  }),
)
const __VLS_23 = __VLS_22(
  {
    ...{ 'onClick:row': {} },
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
    height: '720',
    itemValue: 'name',
    density: 'compact',
    fixedHeader: true,
    hover: true,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_22),
)
let __VLS_25
let __VLS_26
let __VLS_27
const __VLS_28 = {
  'onClick:row': __VLS_ctx.onRowClick,
}
__VLS_24.slots.default
{
  const { top: __VLS_thisSlot } = __VLS_24.slots
  const __VLS_29 = {}.VToolbar
  /** @type {[typeof __VLS_components.VToolbar, typeof __VLS_components.vToolbar, typeof __VLS_components.VToolbar, typeof __VLS_components.vToolbar, ]} */ // @ts-ignore
  const __VLS_30 = __VLS_asFunctionalComponent(
    __VLS_29,
    new __VLS_29({
      flat: true,
    }),
  )
  const __VLS_31 = __VLS_30(
    {
      flat: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_30),
  )
  __VLS_32.slots.default
  const __VLS_33 = {}.VToolbarTitle
  /** @type {[typeof __VLS_components.VToolbarTitle, typeof __VLS_components.vToolbarTitle, typeof __VLS_components.VToolbarTitle, typeof __VLS_components.vToolbarTitle, ]} */ // @ts-ignore
  const __VLS_34 = __VLS_asFunctionalComponent(__VLS_33, new __VLS_33({}))
  const __VLS_35 = __VLS_34({}, ...__VLS_functionalComponentArgsRest(__VLS_34))
  __VLS_36.slots.default
  var __VLS_36
  const __VLS_37 = {}.VSpacer
  /** @type {[typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, ]} */ // @ts-ignore
  const __VLS_38 = __VLS_asFunctionalComponent(__VLS_37, new __VLS_37({}))
  const __VLS_39 = __VLS_38({}, ...__VLS_functionalComponentArgsRest(__VLS_38))
  const __VLS_41 = {}.VSelect
  /** @type {[typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, ]} */ // @ts-ignore
  const __VLS_42 = __VLS_asFunctionalComponent(
    __VLS_41,
    new __VLS_41({
      modelValue: __VLS_ctx.filterNamespace,
      items: __VLS_ctx.namespaces,
      label: 'Namespace',
      clearable: true,
      ...{ class: 'mx-5' },
    }),
  )
  const __VLS_43 = __VLS_42(
    {
      modelValue: __VLS_ctx.filterNamespace,
      items: __VLS_ctx.namespaces,
      label: 'Namespace',
      clearable: true,
      ...{ class: 'mx-5' },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_42),
  )
  const __VLS_45 = {}.VSelect
  /** @type {[typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, ]} */ // @ts-ignore
  const __VLS_46 = __VLS_asFunctionalComponent(
    __VLS_45,
    new __VLS_45({
      modelValue: __VLS_ctx.filterStatus,
      items: __VLS_ctx.statuses,
      label: 'Status',
      clearable: true,
      ...{ class: 'mx-5' },
    }),
  )
  const __VLS_47 = __VLS_46(
    {
      modelValue: __VLS_ctx.filterStatus,
      items: __VLS_ctx.statuses,
      label: 'Status',
      clearable: true,
      ...{ class: 'mx-5' },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_46),
  )
  var __VLS_32
}
{
  const { 'item.status': __VLS_thisSlot } = __VLS_24.slots
  const [{ item }] = __VLS_getSlotParams(__VLS_thisSlot)
  const __VLS_49 = {}.VChip
  /** @type {[typeof __VLS_components.VChip, typeof __VLS_components.vChip, typeof __VLS_components.VChip, typeof __VLS_components.vChip, ]} */ // @ts-ignore
  const __VLS_50 = __VLS_asFunctionalComponent(
    __VLS_49,
    new __VLS_49({
      color: item.status === 'Active' ? 'green' : 'red',
      dark: true,
    }),
  )
  const __VLS_51 = __VLS_50(
    {
      color: item.status === 'Active' ? 'green' : 'red',
      dark: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_50),
  )
  __VLS_52.slots.default
  item.status
  var __VLS_52
}
{
  const { 'item.actions': __VLS_thisSlot } = __VLS_24.slots
  const [{ item }] = __VLS_getSlotParams(__VLS_thisSlot)
  const __VLS_53 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_54 = __VLS_asFunctionalComponent(
    __VLS_53,
    new __VLS_53({
      ...{ onClick: {} },
      icon: true,
    }),
  )
  const __VLS_55 = __VLS_54(
    {
      ...{ onClick: {} },
      icon: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_54),
  )
  let __VLS_57
  let __VLS_58
  let __VLS_59
  const __VLS_60 = {
    onClick: (...[$event]) => {
      __VLS_ctx.editPod(item)
    },
  }
  __VLS_56.slots.default
  const __VLS_61 = {}.VIcon
  /** @type {[typeof __VLS_components.VIcon, typeof __VLS_components.vIcon, ]} */ // @ts-ignore
  const __VLS_62 = __VLS_asFunctionalComponent(
    __VLS_61,
    new __VLS_61({
      icon: 'mdi-pencil',
    }),
  )
  const __VLS_63 = __VLS_62(
    {
      icon: 'mdi-pencil',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_62),
  )
  var __VLS_56
  const __VLS_65 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_66 = __VLS_asFunctionalComponent(
    __VLS_65,
    new __VLS_65({
      ...{ onClick: {} },
      icon: true,
    }),
  )
  const __VLS_67 = __VLS_66(
    {
      ...{ onClick: {} },
      icon: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_66),
  )
  let __VLS_69
  let __VLS_70
  let __VLS_71
  const __VLS_72 = {
    onClick: (...[$event]) => {
      __VLS_ctx.deletePod(item)
    },
  }
  __VLS_68.slots.default
  const __VLS_73 = {}.VIcon
  /** @type {[typeof __VLS_components.VIcon, typeof __VLS_components.vIcon, ]} */ // @ts-ignore
  const __VLS_74 = __VLS_asFunctionalComponent(
    __VLS_73,
    new __VLS_73({
      icon: 'mdi-delete',
    }),
  )
  const __VLS_75 = __VLS_74(
    {
      icon: 'mdi-delete',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_74),
  )
  var __VLS_68
}
var __VLS_24
var __VLS_8
const __VLS_77 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_78 = __VLS_asFunctionalComponent(__VLS_77, new __VLS_77({}))
const __VLS_79 = __VLS_78({}, ...__VLS_functionalComponentArgsRest(__VLS_78))
__VLS_80.slots.default
const __VLS_81 = {}.VLayout
/** @type {[typeof __VLS_components.VLayout, typeof __VLS_components.vLayout, typeof __VLS_components.VLayout, typeof __VLS_components.vLayout, ]} */ // @ts-ignore
const __VLS_82 = __VLS_asFunctionalComponent(__VLS_81, new __VLS_81({}))
const __VLS_83 = __VLS_82({}, ...__VLS_functionalComponentArgsRest(__VLS_82))
__VLS_84.slots.default
const __VLS_85 = {}.VNavigationDrawer
/** @type {[typeof __VLS_components.VNavigationDrawer, typeof __VLS_components.vNavigationDrawer, typeof __VLS_components.VNavigationDrawer, typeof __VLS_components.vNavigationDrawer, ]} */ // @ts-ignore
const __VLS_86 = __VLS_asFunctionalComponent(
  __VLS_85,
  new __VLS_85({
    modelValue: __VLS_ctx.isSidebarVisible,
    location: 'right',
    width: 750,
    temporary: true,
    ...{ class: 'custom-sidebar' },
  }),
)
const __VLS_87 = __VLS_86(
  {
    modelValue: __VLS_ctx.isSidebarVisible,
    location: 'right',
    width: 750,
    temporary: true,
    ...{ class: 'custom-sidebar' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_86),
)
__VLS_88.slots.default
const __VLS_89 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_90 = __VLS_asFunctionalComponent(
  __VLS_89,
  new __VLS_89({
    flat: true,
  }),
)
const __VLS_91 = __VLS_90(
  {
    flat: true,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_90),
)
__VLS_92.slots.default
{
  const { prepend: __VLS_thisSlot } = __VLS_92.slots
  const __VLS_93 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_94 = __VLS_asFunctionalComponent(
    __VLS_93,
    new __VLS_93({
      ...{ onClick: {} },
      icon: 'mdi-chevron-left',
      variant: 'text',
    }),
  )
  const __VLS_95 = __VLS_94(
    {
      ...{ onClick: {} },
      icon: 'mdi-chevron-left',
      variant: 'text',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_94),
  )
  let __VLS_97
  let __VLS_98
  let __VLS_99
  const __VLS_100 = {
    onClick: (...[$event]) => {
      __VLS_ctx.isSidebarVisible = false
    },
  }
  var __VLS_96
  const __VLS_101 = {}.VCardTitle
  /** @type {[typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, ]} */ // @ts-ignore
  const __VLS_102 = __VLS_asFunctionalComponent(
    __VLS_101,
    new __VLS_101({
      ...{ style: {} },
    }),
  )
  const __VLS_103 = __VLS_102(
    {
      ...{ style: {} },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_102),
  )
  __VLS_104.slots.default
  var __VLS_104
}
const __VLS_105 = {}.VDivider
/** @type {[typeof __VLS_components.VDivider, typeof __VLS_components.vDivider, typeof __VLS_components.VDivider, typeof __VLS_components.vDivider, ]} */ // @ts-ignore
const __VLS_106 = __VLS_asFunctionalComponent(__VLS_105, new __VLS_105({}))
const __VLS_107 = __VLS_106({}, ...__VLS_functionalComponentArgsRest(__VLS_106))
const __VLS_109 = {}.VCardText
/** @type {[typeof __VLS_components.VCardText, typeof __VLS_components.vCardText, typeof __VLS_components.VCardText, typeof __VLS_components.vCardText, ]} */ // @ts-ignore
const __VLS_110 = __VLS_asFunctionalComponent(__VLS_109, new __VLS_109({}))
const __VLS_111 = __VLS_110({}, ...__VLS_functionalComponentArgsRest(__VLS_110))
__VLS_112.slots.default
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.name
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.namespace
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.replicas
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.cpu
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.memory
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.age
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.selectedRow?.status
var __VLS_112
var __VLS_92
var __VLS_88
var __VLS_84
var __VLS_80
var __VLS_3
/** @type {__VLS_StyleScopedClasses['mx-4']} */ /** @type {__VLS_StyleScopedClasses['mx-5']} */ /** @type {__VLS_StyleScopedClasses['mx-5']} */ /** @type {__VLS_StyleScopedClasses['custom-sidebar']} */ var __VLS_dollars
let __VLS_self
//# sourceMappingURL=GridBody.vue.js.map
