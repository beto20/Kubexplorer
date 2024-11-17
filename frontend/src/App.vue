<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'App',
});
</script>

<template>
    <div id="app">
        <nav>
            <router-link to="/">Home</router-link>
            <router-link to="/login">login</router-link>
            <router-link to="/contact">Contact</router-link>
        </nav>
        <router-view></router-view>
    </div>
</template>


<style>
nav {
    display: flex;
    gap: 10px;
}
</style>