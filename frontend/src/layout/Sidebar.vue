<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
    name: 'ks-sidebar',
    data() {
        return {
            isCollapsed: false,
            menuItems: [
                { text: 'Home', link: 'home', icon: '🏠' },
                { text: 'Overview', link: 'overview', icon: '📊' },
                { text: 'General', link: 'general', icon: '📊' },
                { text: 'Workload', link: 'workload', icon: '📊' },
                { text: 'Network', link: 'network', icon: '📊' },
                { text: 'Storage', link: 'storage', icon: '📊' },
                { text: 'Settings', link: 'settings', icon: '⚙️' },
            ]
        }
    },
    methods: {
        toggleSidebar() {
            this.isCollapsed = !this.isCollapsed;
        }
    }
});
</script>

<template>
    <div :class="['sidebar', { 'sidebar--collapsed': isCollapsed }]">
        <div class="sidebar-header">
            <button class="toggle-btn" @click="toggleSidebar">
                {{ isCollapsed ? '→' : 'Kubessistant ←' }}
            </button>
        </div>
        <ul class="nav-list">
            <li v-for="(item, index) in menuItems" :key="index">
                <a :href="item.link" class="nav-link">
                    <span class="icon">{{ item.icon }}</span>
                    <span class="link-text" :class="{ 'hidden': isCollapsed }">
                        {{ item.text }}
                    </span>
                </a>
            </li>
        </ul>
    </div>
</template>

<style scoped>
.sidebar {
    width: 250px;
    background-color: #2c3e50;
    color: #ecf0f1;
    transition: all 0.3s ease;
    display: flex;
    flex-direction: column;
    position: relative;
}

.sidebar--collapsed {
    width: 60px;
}

.sidebar-header {
    padding: 15px;
    display: flex;
    justify-content: flex-end;
    border-bottom: 1px solid #34495e;
}

.toggle-btn {
    background: none;
    border: none;
    color: #ecf0f1;
    cursor: pointer;
    padding: 5px 10px;
    font-size: 1.2em;
    border-radius: 4px;
}

.toggle-btn:hover {
    background-color: #34495e;
}

.nav-list {
    list-style: none;
    padding: 0;
    margin: 0;
    overflow-y: auto;
}

.nav-link {
    color: #ecf0f1;
    text-decoration: none;
    padding: 12px 15px;
    display: flex;
    align-items: center;
    transition: background-color 0.2s;
}

.nav-link:hover {
    background-color: #34495e;
}

.icon {
    width: 24px;
    text-align: center;
    margin-right: 10px;
}

.link-text {
    transition: opacity 0.3s;
    white-space: nowrap;
}

.hidden {
    opacity: 0;
    width: 0;
    display: none;
}

/* Add smooth transitions for text */
.sidebar .link-text {
    transition: opacity 0.3s ease;
}

/* Ensure sidebar content doesn't overflow */
.sidebar {
    overflow-x: hidden;
}

/* Make scrollbar less intrusive */
.nav-list::-webkit-scrollbar {
    width: 5px;
}

.nav-list::-webkit-scrollbar-track {
    background: #2c3e50;
}

.nav-list::-webkit-scrollbar-thumb {
    background: #34495e;
    border-radius: 3px;
}
</style>
