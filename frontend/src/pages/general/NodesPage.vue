<script lang="ts">
import {defineComponent} from 'vue'
import KsGridBody from "../../layout/GridBody.vue";
import KsNavBar from "../../layout/Navbar.vue";

export default defineComponent({
    name: "NodesPage",
    components: {KsNavBar, KsGridBody },
    data() {
        return {
            followingPage: 'General'
        }
    },
})

</script>

<template>
    <ks-grid-body></ks-grid-body>

</template>

<style scoped>

</style>