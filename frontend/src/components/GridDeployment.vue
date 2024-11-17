<script lang="ts">
import { defineComponent } from 'vue';
import {GetDeployments} from "../../wailsjs/go/middleware/DeploymentMiddleware";

export default defineComponent({
    name: 'GridComponent',
        data() {
            return {
                items: ['Card 1', 'Card 2', 'Card 3', 'Card 4'] as string[],
            };
    },
});

function getDeployments() {
    GetDeployments("assi").then((response) => {
        console.log(response);

    }).catch((error) => {
        console.log(error);
    })
}
</script>


<template>
    <v-app>
<!--        <v-container>-->
<!--            <v-row>-->
<!--                <v-col cols="12" md="6">-->
<!--                    <v-card>-->
<!--                        <v-card-title>Hello, Vuetify!</v-card-title>-->
<!--                        <v-card-text>This is a basic Vuetify setup.</v-card-text>-->
<!--                        <v-card-actions>-->
<!--                            <v-btn color="primary">Click Me</v-btn>-->
<!--                        </v-card-actions>-->
<!--                    </v-card>-->
<!--                </v-col>-->
<!--            </v-row>-->
<!--        </v-container>-->

<!--        <v-container>-->
<!--            <v-row>-->
<!--                <v-col-->
<!--                    v-for="(item, index) in items"-->
<!--                    :key="index"-->
<!--                    cols="12"-->
<!--                    md="4"-->
<!--                >-->
<!--                    <v-card>-->
<!--                        <v-card-title>{{ item }}</v-card-title>-->
<!--                    </v-card>-->
<!--                </v-col>-->
<!--            </v-row>-->
<!--            <v-btn>-->
<!--                Button-->
<!--            </v-btn>-->
<!--        </v-container>-->

    </v-app>
</template>
