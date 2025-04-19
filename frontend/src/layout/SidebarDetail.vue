<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
    name: "ks-sidebar-detail",
    props: {
        isVisible: Boolean,
        selectedRow: Object,
    },
    emits: ["close"],
})
</script>

<template>
    <v-layout>
        <v-navigation-drawer
            v-model="isVisible"
            location="right"
            :width="750"
            temporary
            class="custom-sidebar"
        >
            <v-card flat>
                <template v-slot:prepend>
                    <v-btn
                        icon="mdi-chevron-left"
                        variant="text"
                        @click="$emit('close')"
                    ></v-btn>
                    <v-card-title style="color: black">Object Details</v-card-title>
                </template>
                <v-divider></v-divider>
                <v-card-text v-if="selectedRow">
                    <p>Name: {{ selectedRow.name }}</p>
                    <p>Namespace: {{ selectedRow.namespace }}</p>
                    <p>Replicas: {{ selectedRow.replicas }}</p>
                    <p>CPU: {{ selectedRow.cpu }}</p>
                    <p>Memory: {{ selectedRow.memory }}</p>
                    <p>Age: {{ selectedRow.age }}</p>
                    <p>Status: {{ selectedRow.status }}</p>
                </v-card-text>
            </v-card>
        </v-navigation-drawer>
    </v-layout>
</template>

<style scoped>
.custom-sidebar {
    position: fixed; /* Ensure it's above all other content */
    top: 0;
    right: 0;
    height: 100%;
    z-index: 1100; /* Higher than navbar */
}
</style>