<script lang="ts">
import { defineComponent } from 'vue';
import HomePage from "./pages/HomePage.vue";

export default defineComponent({
    name: 'App',
    components: {HomePage},
});
</script>

<template>
    <home-page></home-page>
<!--    <div id="app">-->
<!--        <nav>-->
<!--            <router-link to="/">Home</router-link>-->
<!--            <router-link to="/login">login</router-link>-->
<!--            <router-link to="/contact">Contact</router-link>-->
<!--        </nav>-->
<!--        <router-view></router-view>-->
<!--    </div>-->
</template>


<style>
nav {
    display: flex;
    gap: 10px;
}
</style>