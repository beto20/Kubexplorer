<script setup lang="ts">

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

</style>