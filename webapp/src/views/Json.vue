<script setup>
import {ref} from "vue";
import api from "../api";
import {message} from "ant-design-vue";

const indent = ref("  ");
const input = ref("");
const output = ref("");

function formatJson() {
  if (input.value === "") {
    message.error("请在输入框中输入需要格式化的 JSON");
    return;
  }
  api
      .post("/json", {indent: indent.value, content: input.value})
      .then((resp) => {
        if (resp.data.code === 0) {
          output.value = resp.data.data;
        } else {
          output.value = resp.data.msg;
        }
      })
      .catch((err) => (output.value = err.toJSON));
}
</script>

<template>
  <a-row :gutter="16" justify="space-between">
    <a-col :span="12">
      <a-input placeholder="空白则不缩进并且不换行" v-model:value="indent" :allow-clear="true">
        <template #prefix>
          <span>使用</span>
        </template>
        <template #suffix>
          <span>作为缩进</span>
          <a-button type="primary" @click="formatJson">格式化</a-button>
        </template>
      </a-input>
      <h2>输入：</h2>
      <a-textarea
          placeholder="输入 JSON 字符串"
          v-model:value="input"
          :show-count="true"
          :allowClear="true"
          :auto-size="{ minRows: 10, maxRows: 20 }"
      />
    </a-col>
    <a-col :span="12">
      <h2>输出：</h2>
      <a-textarea
          v-if="output !== ''"
          placeholder="结果"
          v-model:value="output"
          :show-count="true"
          :allowClear="true"
          :auto-size="true"
      />
    </a-col>
  </a-row>
</template>
