import { defineComponent } from 'vue'
import { GetDeployments } from '../../wailsjs/go/middleware/WorkloadMiddleware'
export default defineComponent({
  name: 'GridComponent',
  data() {
    return {
      items: ['Card 1', 'Card 2', 'Card 3', 'Card 4'],
    }
  },
})
function getDeployments() {
  GetDeployments()
    .then((response) => {
      console.log(response)
    })
    .catch((error) => {
      console.log(error)
    })
}
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
const __VLS_0 = {}.VApp
/** @type {[typeof __VLS_components.VApp, typeof __VLS_components.vApp, typeof __VLS_components.VApp, typeof __VLS_components.vApp, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({}))
const __VLS_2 = __VLS_1({}, ...__VLS_functionalComponentArgsRest(__VLS_1))
var __VLS_4 = {}
__VLS_3.slots.default
var __VLS_3
var __VLS_dollars
let __VLS_self
//# sourceMappingURL=GridDeployment.vue.js.map
