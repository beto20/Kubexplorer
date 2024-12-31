<script lang="ts">

import {defineComponent, ref} from "vue";
import {database} from "../../../wailsjs/go/models";
import CommonParameterDto = database.CommonParameterDto;
import {useLayoutComposableExample} from "../../composables/useLayoutComposableExample";


export default defineComponent({
    name: "GeneralPage",
    setup() {
        const response = ref<CommonParameterDto[]>([]);
        const { result } = useLayoutComposableExample();

        const callMiddleware = async () => {
            try {
                response.value = result.value.map((r) => ({
                    Name: r.Name,
                    Link: r.Link,
                    Icon: r.Icon,
                }));
                console.log('Result from Go:', response.value);

            } catch (error) {
                console.error('Error calling Go function:', error);
            }
        };

        return {
            callMiddleware,
            response,
        };
    }
})

</script>


<template>
    <nav class="navbar">
        <div class="nav-content">
            <div class="nav-links">
                <router-link to="/pod" class="nav-link">Pod</router-link>
                <router-link to="/deployment" class="nav-link">Deployment</router-link>
            </div>
        </div>
    </nav>
    <div>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
        <h1>asdasdas</h1>
    </div>
</template>


<style scoped>
.navbar {
    position: sticky;
    top: 0;
    left: 0;
    right: 0;
    background-color: rgba(255, 255, 255, 0.95);
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0);
    transition: all 0.3s ease;
    z-index: 1000;
}

.navbar-scrolled {
    background-color: rgba(255, 255, 255, 0.98);
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.nav-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 20px;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.nav-links {
    display: flex;
    gap: 2rem;
}

.nav-link {
    color: #333;
    text-decoration: none;
    font-weight: 500;
    padding: 0.5rem 1rem;
    transition: color 0.3s ease;
}

.nav-link:hover {
    color: #007bff;
}

.nav-link.router-link-active {
    color: #007bff;
}

</style>