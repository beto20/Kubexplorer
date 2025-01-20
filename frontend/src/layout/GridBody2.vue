<script lang="ts">
import {computed, defineComponent, onMounted, reactive, ref, toRefs} from 'vue'
import {gridBodyPodsComposable} from "../composables/GridWorkloadComposable";
import KsSidebarDetail from "./SidebarDetail.vue";
import KsGridTable from "../components/GridTableComponent.vue";
import KsGridHeader from "../components/GridHeaderComponent.vue";

interface GridState {
    items: Array<{
        name: string;
        namespace: string;
        replicas: number;
        cpu: string;
        memory: string;
        age: string;
        status: string;
    }>;
    search: string;
    filterNamespace: string ;
    filterStatus: string ;
    sortBy: any;
}

interface HeadState {
    header: Array <any>;
}

export default defineComponent({
    name: "KsGridBodyV2",
    components: { KsGridHeader, KsGridTable, KsSidebarDetail },
    props: {
        k8sObject: { type: String, required: true },
        namespace: { type: String, required: true },
    },
    setup(props) {
        const { pods, headers, fetchData } = gridBodyPodsComposable("mock", props.k8sObject);
        const namespaces = ["ns-local", "ns-dev"];
        const statuses = ["Alive", "Inactive"];

        const state = reactive<GridState>({
            items: [],
            search: "",
            filterNamespace: "",
            filterStatus: "",
            sortBy: [{ key: "name", order: "asc" }],
        });

        const header = reactive<HeadState>({
            header: [],
        });

        const filteredItems = computed(() => {
            return state.items.filter((item) => {
                const namespaceMatch = !state.filterNamespace || item.namespace.includes(state.filterNamespace);
                const statusMatch = !state.filterStatus || item.status.includes(state.filterStatus);
                const searchMatch = !state.search || item.name.toLowerCase().includes(state.search.toLowerCase());
                return namespaceMatch && statusMatch && searchMatch;
            });
        });

        const isSidebarVisible = ref(false);
        const selectedRow = ref<any>(null);

        const onRowClick = (cellData: any, item: any) => {
            selectedRow.value = item.item;
            console.log("selectedRow.value", item.item)

            if (isSidebarVisible.value) {
                console.log("FALSE")
                isSidebarVisible.value = false
                setTimeout(() => {
                    console.log("Delayed");
                    isSidebarVisible.value = true;
                }, 100);
            } else {
                console.log("FALSE")
                isSidebarVisible.value = true;
            }
        };

        onMounted(async () => {
            await fetchData();
            state.items = pods.value.map((p) => ({
                name: p.Name,
                namespace: p.Namespace,
                replicas: p.Replicas,
                cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
                memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
                age: p.Age,
                status: p.Status,
            }));

            header.header = headers.value.map((h) => ({
                title: h.Title,
                key: h.Key,
                align: h.Align,
                sortable: h.Sortable,
            }));
        });

        return {
            namespaces,
            header,
            statuses,
            ...toRefs(state),
            filteredItems,
            isSidebarVisible,
            selectedRow,
            onRowClick,
        };
    },
});
</script>

<template>
    <v-container>
        <v-card>
            <KsGridHeader
                v-model:search="search"
                v-model:filterNamespace="filterNamespace"
                v-model:filterStatus="filterStatus"
                :namespaces="namespaces"
                :statuses="statuses"
            />
            <KsGridTable
                :headers="header.header"
                :items="filteredItems"
                :search="search"
                :sortBy="sortBy"
                @click:row="onRowClick"
            />
        </v-card>
        <v-card>
            <KsSidebarDetail
                :isVisible="isSidebarVisible"
                :selectedRow="selectedRow"
                @close="isSidebarVisible = false"
            />
        </v-card>
    </v-container>
</template>