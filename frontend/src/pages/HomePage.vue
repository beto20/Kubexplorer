<script lang="ts">
import { defineComponent } from 'vue';
import KsSidebar from "../layout/Sidebar.vue";

export default defineComponent({
    name: 'HomePage',
    components: {KsSidebar},
});
</script>

<template>
    <h2>Home page</h2>
    <ks-sidebar></ks-sidebar>

</template>

<style scoped>

h2 {
    text-align: center;
    color: rebeccapurple;
}

</style>