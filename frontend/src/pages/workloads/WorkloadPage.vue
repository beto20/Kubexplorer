<script lang="ts">
import {defineComponent} from "vue";
import KsNavBar from "../../layout/Navbar.vue";

export default defineComponent({
    name: "WorkloadPage" ,
    components: {KsNavBar}
})
</script>
<template>
    <div>
        <ks-nav-bar></ks-nav-bar>
    </div>
    <div>
        <router-view></router-view>
        <h2>Workload page</h2>
    </div>
</template>

<style scoped>

</style>