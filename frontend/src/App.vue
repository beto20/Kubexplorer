<script lang="ts">
import { defineComponent } from 'vue';
import HomePage from "./pages/HomePage.vue";
import KsSidebar from "./layout/Sidebar.vue";

export default defineComponent({
    name: 'App',
    components: {HomePage, KsSidebar},
});
</script>

<template>
    <div class="app-container">
        <ks-sidebar></ks-sidebar>
        <div class="content">
            <router-view></router-view>
        </div>
    </div>
</template>


<style>
.app-container {
    display: flex;
    min-height: 100vh;
}

/*
.content {
    flex: 1;
    padding: 20px;
    background-color: #f5f6fa;
}
*/
</style>

