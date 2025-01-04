<script lang="ts">

import {defineComponent, ref} from "vue";
import {database} from "../../../wailsjs/go/models";
import CommonParameterDto = database.CommonParameterDto;
import {useLayoutComposableExample} from "../../composables/useLayoutComposableExample";
import KsNavBar from "../../layout/Navbar.vue";


export default defineComponent({
    name: "GeneralPage",
    components: {KsNavBar},
    data() {
        return {
            followingPage: 'General'
        }
    },
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
    <div>
        <ks-nav-bar :content="followingPage"></ks-nav-bar>
    </div>
    <div>
        <router-view></router-view>
    </div>

</template>


<style scoped>

</style>