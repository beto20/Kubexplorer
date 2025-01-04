<script lang="ts">
import {defineComponent} from 'vue'
import KsGridBody from "../../layout/GridBody.vue";

export default defineComponent({
    name: "NamespacePage",
    components: { KsGridBody }
})
console.log("PRUEBA CARGA NS")
</script>

<template>
    <ks-grid-body></ks-grid-body>
</template>

<style scoped>

</style>