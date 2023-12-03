<script setup>
import {ref} from "vue";
import {message} from "ant-design-vue";

// 解析请求头
const showImportReqInfo = ref(false)
const importReqInfo = ref("")

// 请求信息
const methodInput = ref("GET")
const urlInput = ref("")
const headerInput = ref("")
const reqParamInput = ref("")

// 发送请求
function sendRequest() {
  message.info("发送请求...")
}

// 导入请求
function parseRequestInfo() {
  showImportReqInfo.value = false
}
</script>

<template>
  <a-row :gutter="16" justify="space-between">
    <a-col :span="12">
      <a-input-group compact>
        <h2 style="width: 100%">Url：</h2>
        <a-select v-model:value="methodInput" style="width: 20%">
          <a-select-option value="GET">GET</a-select-option>
          <a-select-option value="POST">POST</a-select-option>
        </a-select>
        <a-input placeholder="请求地址" v-model:value="urlInput"
                 style="width: 80%" allow-clear/>
      </a-input-group>
      <a-input-group>
        <h2>请求头:</h2>
        <a-textarea v-model:value="headerInput" placeholder="请输入请求头信息" :rows="4"/>
      </a-input-group>
      <a-input-group>
        <a-row justify="space-between" align="bottom">
          <a-col>
            <h2>请求参数：</h2>
          </a-col>
          <a-col>
            <a-button type="primary" @click="() => showImportReqInfo=true">导入请求</a-button>
            <a-modal v-model:open="showImportReqInfo" title="从请求头中导入请求信息" @ok="parseRequestInfo">
              <a-textarea v-model:value="importReqInfo" :rows="5" allow-clear/>
            </a-modal>
            <a-divider type="vertical"/>
            <a-button type="primary" @click="sendRequest">发送请求</a-button>
          </a-col>
        </a-row>
        <a-textarea placeholder="请输入请求参数"
                    v-model:value="reqParamInput"
                    :auto-size="{minRows: 10, maxRows: 20}"
                    allow-clear/>
      </a-input-group>
    </a-col>
    <a-col :span="12">
      <h2>请求结果：</h2>
    </a-col>
  </a-row>
</template>