<template>
  <div
    class="welcome"
  >
    Hi {{ greetStr }}~ Enjoy Y-Tool
  </div>
</template>

<script setup>
import api from "../api";
import { ref } from "vue";

const greetStr = ref("waiting...");

api
  .get("/username")
  .then((res) => (greetStr.value = res.data))
  .catch((err) => {
    greetStr.value = "request failed";
    console.log(err);
  });
</script>

<style scoped>
@keyframes typing {
  from { width: 0 }
}
.welcome {
  width: 100%;
  text-align: center;
  line-height: 80vh;
  font-size: 3em;
  font-weight: bold;
  white-space: nowrap;
  overflow: hidden;
  animation: typing 5s;
}
</style>