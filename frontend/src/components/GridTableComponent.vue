<script lang="ts">
import {defineComponent, onMounted} from 'vue'
import {gridGeneralComposable} from "../composables/GridTableComposable";


interface State {
  isRestarted?: any;
}

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
        const response = gridGeneralComposable("")

        const editPod = (item: any) => {
            console.log("EDIT", item);
        };

        const state: State = {
          isRestarted: null
        }

        onMounted(async () => {
            await response.fetchData()
            state.isRestarted = response.isRestarted
        })

        const deletePod = (item: any) => {
            // console.log("DELETE", item);
            console.log("RESTARTED", state.isRestarted.value)
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