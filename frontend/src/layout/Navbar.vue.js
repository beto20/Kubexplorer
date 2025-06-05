import { defineComponent, onMounted, reactive, toRefs } from 'vue'
import { navbarComposable } from '../composables/NavbarComposable'
export default defineComponent({
  name: 'ks-nav-bar',
  props: {
    content: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const { objects, fetchData } = navbarComposable()
    const state = reactive({
      menu: [],
      objects: [],
    })
    onMounted(async () => {
      await fetchData()
      console.log('OBJECTS', objects.value)
      const dto = objects.value.map((o) => {
        return {
          Name: o.Name,
          IsVisible: o.IsVisible,
          IsEditable: o.IsEditable,
          K8sObject: o.K8sObject.map((k) => {
            return {
              Name: k.Name,
              Link: k.Link,
              IsEditable: k.IsEditable,
              IsVisible: k.IsVisible,
            }
          }),
        }
      })
      const type = dto.find((x) => x.Name == props.content)
      state.menu = dto
      state.objects = type.K8sObject
    })
    return {
      ...toRefs(state),
    }
  },
})
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
/** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ // CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.nav,
  __VLS_intrinsicElements.nav,
)({
  ...{ class: 'navbar' },
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: 'nav-content' },
})
for (const [item, index] of __VLS_getVForSourceType(__VLS_ctx.objects)) {
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.div,
    __VLS_intrinsicElements.div,
  )({
    ...{ class: 'nav-links' },
    key: index,
  })
  const __VLS_0 = {}.RouterLink
  /** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
  const __VLS_1 = __VLS_asFunctionalComponent(
    __VLS_0,
    new __VLS_0({
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    }),
  )
  const __VLS_2 = __VLS_1(
    {
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_1),
  )
  __VLS_3.slots.default
  item.Name
  var __VLS_3
}
/** @type {__VLS_StyleScopedClasses['navbar']} */ /** @type {__VLS_StyleScopedClasses['nav-content']} */ /** @type {__VLS_StyleScopedClasses['nav-links']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ var __VLS_dollars
let __VLS_self
//# sourceMappingURL=Navbar.vue.js.map
