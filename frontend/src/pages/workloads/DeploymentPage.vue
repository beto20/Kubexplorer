<script lang="ts">
import {defineComponent} from 'vue'
import KsGridBody from "../../layout/GridBody.vue";
import KsGridBodyV2 from "../../layout/GridBody2.vue";
import KsNavBar from "../../layout/Navbar.vue";

export default defineComponent({
    name: "DeploymentPage" ,
    components: {KsNavBar, KsGridBody, KsGridBodyV2 },
    data() {
        return {
            k8sObject: 'deployment',
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