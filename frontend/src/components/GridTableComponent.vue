<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
    name: "KsGridTable",
    props: {
        headers: Array<any>,
        items: Array<any>,
        search: String,
        sortBy: Array<any>,
    },
    emits: ["rowClick", "edit", "delete"],
    setup() {
        const editPod = (item: any) => {
            console.log("EDIT", item);
        };

        const deletePod = (item: any) => {
            console.log("DELETE", item);
        };

        return {
            editPod,
            deletePod,
        }
    }
})
</script>

<template>
    <v-data-table-virtual
        :headers="headers"
        :items="items"
        :search="search"
        :sort-by="sortBy"
        height="720"
        item-value="name"
        density="compact"
        fixed-header
        hover
    >
        <template v-slot:item.status="{ item }">
            <v-chip :color="item.status === 'Active' ? 'green' : 'red'" dark>
                {{ item.status }}
            </v-chip>
        </template>

        <template v-slot:item.actions="{ item }">
            <v-btn @click.stop="editPod(item)" icon>
                <v-icon icon="mdi-pencil" />
            </v-btn>
            <v-btn @click.stop="deletePod(item)" icon>
                <v-icon icon="mdi-delete" />
            </v-btn>
        </template>
    </v-data-table-virtual>
</template>

<style scoped>

</style>