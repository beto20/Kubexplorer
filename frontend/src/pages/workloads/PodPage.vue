<script lang="ts">
import {defineComponent} from 'vue'
import KsGridBody from "../../layout/GridBody.vue";
import KsNavBar from "../../layout/Navbar.vue";
import KsGridBodyV2 from "../../layout/GridBody2.vue";

export default defineComponent({
    name: "PodPage" ,
    components: {KsNavBar, KsGridBody, KsGridBodyV2 },
    data() {
        return {
            k8sObject: 'pod',
            namespace: 'mock'
        }
    },
})
</script>

<template>
    <ks-grid-body-v2 :namespace="namespace" :k8sObject="k8sObject" ></ks-grid-body-v2>
</template>

<style scoped>

</style>