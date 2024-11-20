<script lang="ts">
import { defineComponent, ref } from 'vue';
export default defineComponent({
    name: 'ks-sidebar',
    setup() {
        // Reactive state for sidebar collapse
        const isCollapsed = ref(false);

        // Method to toggle the sidebar state
        const toggleSidebar = () => {
            isCollapsed.value = !isCollapsed.value;
        };

        // Expose variables and methods
        return {
            isCollapsed,
            toggleSidebar,
        };
    },
});

</script>

<template>
    <div class="app-container">
        <div :class="['sidebar', { 'sidebar--collapsed': isCollapsed }]">
            <button @click="toggleSidebar">Toggle</button>
            <ul>
                <li><a href="#home">Home</a></li>
                <li><a href="#about">About</a></li>
                <li><a href="#contact">Contact</a></li>
            </ul>
        </div>
        <div class="content">
            <h1>Content Area</h1>
            <p>This is the main content area.</p>
        </div>
    </div>
</template>

<style scoped>
.app-container {
    display: flex;
}

.sidebar {
    width: 250px;
    background-color: #2c3e50;
    color: #ecf0f1;
    padding: 15px;
    transition: width 0.3s;
    overflow: hidden;
}

.sidebar--collapsed {
    width: 60px;
}

.sidebar ul {
    list-style: none;
    padding: 0;
}

.sidebar li {
    margin: 10px 0;
}

.sidebar a {
    color: #ecf0f1;
    text-decoration: none;
}

.content {
    flex: 1;
    padding: 20px;
}

</style>
