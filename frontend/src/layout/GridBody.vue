<script lang="ts">
import {defineComponent, ref} from 'vue'

export default defineComponent({
    name: "ks-grid-body",
    data() {
        return {
            search: "",
            filterNamespace: null,
            filterStatus: null,
            headers: ref([
                { title: "Name", key: "name", align: 'start', sortable: true },
                { title: "Namespace", key: "namespace", align: 'start', sortable: false  },
                { title: "Replicas", key: "replicas", align: 'start', sortable: false  },
                { title: "Cpu", key: "cpu", align: 'start', sortable: false  },
                { title: "Memory", key: "memory", align: 'start', sortable: false  },
                { title: "Age", key: "age", align: 'start', sortable: true  },
                { title: "Status", key: "status", align: 'start', sortable: false  },
                { title: "Actions", key: "actions", align: 'start', sortable: false },
            ] as const),
            items: Array.from({ length: 100 }).map((v, k) => (
                {
                    name: `pod-${k}`,
                    namespace: "ns-local",
                    replicas: "1/1",
                    cpu: "500mb/2Gi",
                    memory: "500Mb/1500Mb",
                    age: "50min ago",
                    status: "Active",
                }
            )),
            namespace: ["ns-local", "ns-dev"],
            statuses: ["Active", "Inactive"],
            sortBy: [
                { key: 'name', order: 'asc' }
            ] as const,
            start: 0,
            timeout: null || 0,
            rowHeight: 24,
            perPage: 25
        };
    },
    computed: {
        filteredItems() {
            return this.items.filter((item) => {
                const namespaceMatch = !this.filterNamespace || item.namespace.includes(this.filterNamespace);
                const statusMatch = !this.filterStatus || item.status.includes(this.filterStatus);
                return namespaceMatch && statusMatch;
            });
        },
    },
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
    },
    methods: {
        // onRowClick(...args: any) {
        //     // this will log all arguments, as array
        //     console.log(args);
        // }
        onRowClick(cellData: any, item: any) {
            console.log("cellData", cellData);
            console.log("item", item);
            console.log("item", item.item.name);
        },

    }
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
                :headers="headers"
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
                            :items="namespace"
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
    </v-container>
</template>


<style scoped>

</style>