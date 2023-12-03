<script setup>
import { ref } from "vue";
import {
  Button,
  Layout,
  LayoutHeader,
  LayoutSider,
  LayoutContent,
  Menu,
  MenuItem,
  SubMenu,
} from "ant-design-vue";
import { AppstoreOutlined, BugOutlined } from "@ant-design/icons-vue";

const collapsed = ref(false);
const selectedKeys = ref([]);
const openKeys = ref(["tools"]);
</script>

<template>
  <Layout style="min-height: 100vh; min-width: 100vh">
    <LayoutSider v-model="collapsed" collapsible>
      <h1 class="menu-title">
        <router-link to="/welcome">Y-Tool</router-link>
      </h1>
      <Menu
        v-model:selected-keys="selectedKeys"
        v-model:open-keys="openKeys"
        theme="dark"
        mode="inline"
      >
        <SubMenu key="tools">
          <template #title>
            <span>
              <AppstoreOutlined />
              <span>实用工具</span>
            </span>
          </template>
          <MenuItem key="tools-http">
            <span
              ><router-link to="/http" /><router-link to="/http"
                >HTTP请求工具</router-link
              ></span
            >
          </MenuItem>
          <MenuItem key="tools-json">
            <span
              ><router-link to="/json" /><router-link to="/json"
                >JSON工具</router-link
              ></span
            >
          </MenuItem>
        </SubMenu>
        <SubMenu key="debug">
          <template #title>
            <span>
              <BugOutlined />
              <span>调测工具</span>
            </span>
          </template>
        </SubMenu>
      </Menu>
    </LayoutSider>
    <Layout style="padding: 0">
      <LayoutContent :style="{ margin: '24px 16px', background: '#fff' }">
        <router-view />
      </LayoutContent>
    </Layout>
  </Layout>
</template>

<style scoped>
.menu-title {
  color: #fff;
  font-size: 20px;
  text-align: center;
}
</style>
