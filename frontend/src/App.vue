<script lang="ts">
import { defineComponent } from 'vue';
import ksSidebarV2 from "./layout/SidebarV2.vue";
// import GeneralPage from "./GeneralPage.vue"

export default defineComponent({
    name: 'App',
    components: {
        // GeneralPage,
        ksSidebarV2
    },
});
</script>

<template>
    <div class="app-container">
        <ks-sidebar-v2></ks-sidebar-v2>
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

