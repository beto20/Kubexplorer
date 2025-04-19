<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
    name: "KsGridHeader",
    props: {
        search: String,
        filterNamespace: String,
        filterStatus: String,
        namespaces: Array,
        statuses: Array,
    },
    emits: ["update:search", "update:filterNamespace", "update:filterStatus"],
})
</script>

<template>
    <v-toolbar flat>
        <v-toolbar-title>Filters</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-text-field
            :model-value="search"
            label="Search"
            clearable
            class="mx-4"
            @update:model-value="$emit('update:search', $event)"
        ></v-text-field>
        <v-select
            :model-value="filterNamespace"
            :items="namespaces"
            label="Namespace"
            clearable
            class="mx-5"
            @update:model-value="$emit('update:filterNamespace', $event)"
        ></v-select>
        <v-select
            :model-value="filterStatus"
            :items="statuses"
            label="Status"
            clearable
            class="mx-5"
            @update:model-value="$emit('update:filterStatus', $event)"
        ></v-select>
    </v-toolbar>
</template>

<style scoped>

</style>