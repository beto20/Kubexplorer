<script lang="ts">
import {computed, defineComponent, onMounted, reactive, ref, toRefs} from 'vue'
import { gridGeneralComposable} from "../composables/GridGeneralComposable";
import {gridBodyPodsComposable} from "../composables/GridWorkloadComposable";
import KsSidebarDetail from "./SidebarDetail.vue";

interface GridState {
    // header: Array <{
    //    title: string,
    //    key: string,
    //    align: string,
    //    sortable: boolean
    // }>;
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
    filterNamespace: string | null;
    filterStatus: string | null;
    sortBy: any;
}

interface HeadState {
    header: Array <any>;
}

export default defineComponent({
    name: "ks-grid-body",
    components: {KsSidebarDetail},
    props: {
        k8sObject: {
            type: String,
            // required: true
        },
        namespace: {
            type: String,
            // required: true
        }
    },
    setup(props) {
        const editPod = (item: any) => {
            console.log("EDIT", item);
        };

        const deletePod = (item: any) => {
            console.log("DELETE", item);
        };

        const { pods, headers, fetchData } = gridBodyPodsComposable("mock", "pod");
        // const { nodes, namespaces, headers, fetchData } = gridGeneralComposable(props.k8sObject);

        const namespaces = ['ns-local', 'ns-dev'];
        const statuses = ['Alive', 'Inactive'];

        const state = reactive<GridState>({
            // header: [],
            items: [],
            search: '',
            filterNamespace: null,
            filterStatus: null,
            sortBy: [{ key: 'name', order: 'asc' }],
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
                cpu: p.Container.Limit.Cpu + '/' + p.Container.Request.Cpu,
                memory: p.Container.Limit.Memory + '/' + p.Container.Request.Memory,
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
            header,
            namespaces,
            statuses,
            ...toRefs(state),
            editPod,
            deletePod,
            filteredItems,
            isSidebarVisible,
            selectedRow,
            onRowClick,
        }
    },

})
</script>

<template>
    <v-container>
        <v-card>
            <v-card-title>
                <v-spacer></v-spacer>
                <v-text-field
                    v-model="search"
                    label="Search"
                    class="mx-4"
                    clearable
                    aria-label="Search table"
                ></v-text-field>
            </v-card-title>

            <v-data-table-virtual
                :headers="header.header"
                :items="filteredItems"
                :search="search"
                :sort-by="sortBy"
                height="720"
                item-value="name"
                density="compact"
                fixed-header
                @click:row="onRowClick"
                hover
            >
                <template v-slot:top>
                    <v-toolbar flat>
                        <v-toolbar-title>Filters</v-toolbar-title>
                        <v-spacer></v-spacer>
                        <v-select
                            v-model="filterNamespace"
                            :items="namespaces"
                            label="Namespace"
                            clearable
                            class="mx-5"
                        ></v-select>
                        <v-select
                            v-model="filterStatus"
                            :items="statuses"
                            label="Status"
                            clearable
                            class="mx-5"
                        ></v-select>
                    </v-toolbar>
                </template>

                <template v-slot:item.status="{ item }">
                    <v-chip :color="item.status === 'Active' ? 'green' : 'red'" dark>
                        {{ item.status }}
                    </v-chip>
                </template>

                <template v-slot:item.actions="{ item }">
                    <v-btn @click.stop="editPod(item)" icon >
                        <v-icon icon="mdi-pencil" />
                    </v-btn>
                    <v-btn @click.stop="deletePod(item)" icon>
                        <v-icon icon="mdi-delete" />
                    </v-btn>
                </template>
            </v-data-table-virtual>
        </v-card>

        <v-card>
            <v-layout>
                <v-navigation-drawer
                    v-model="isSidebarVisible"
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
                                    @click="isSidebarVisible = false"
                                ></v-btn>
                                <v-card-title style="color: black">nombre de objecto</v-card-title>
                            </template>
                        <v-divider></v-divider>
                        <v-card-text>
                            <p>Name: {{ selectedRow?.name }}</p>
                            <p>Namespace: {{ selectedRow?.namespace }}</p>
                            <p>Replicas: {{ selectedRow?.replicas }}</p>
                            <p>CPU: {{ selectedRow?.cpu }}</p>
                            <p>Memory: {{ selectedRow?.memory }}</p>
                            <p>Age: {{ selectedRow?.age }}</p>
                            <p>Status: {{ selectedRow?.status }}</p>
                        </v-card-text>
                    </v-card>
                </v-navigation-drawer>
            </v-layout>
        </v-card>

<!--        <ks-sidebar-detail></ks-sidebar-detail>-->

    </v-container>
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