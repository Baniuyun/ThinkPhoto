<script setup>
import  {ref} from "vue"
import {useUserStore} from "@/store/User";
import worklist from "@/components/user/worklist.vue";
const userStore = useUserStore()

</script>

<template>
<worklist></worklist>



</template>

<style scoped>

</style>