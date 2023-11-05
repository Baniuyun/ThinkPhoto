<script setup>

import Head from "@/components/user/head.vue";
import Menu from "@/components/user/menu.vue";
import Worklist from "@/components/user/worklist.vue";
</script>

<template>
    <div>
        <Head></Head>
    <Menu></Menu>
    <Worklist></Worklist>
    </div>

</template>

<style lang="scss" scoped>
  div{
      //overflow-y: hidden;
  }
  div::-webkit-scrollbar{
      display: none;
  }
</style>