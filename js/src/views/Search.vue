<script setup>
import { useRouter } from 'vue-router';
import {ref,reactive} from "vue";
import Worklist from "@/components/user/worklist.vue";
import Menu from "@/components/search/menu.vue";
const router = useRouter([]);
import {getPublishTokenService} from "@/api/video";
import {useUserStore} from "@/store/User";
const userStore = useUserStore();

const test1 = async () => {
    await getPublishTokenService().then((res)=>{
        alert(res.data.token)
    })
}


</script>

<template>
<test></test>
    {{userStore.getToken()+userStore.getUserId()}}
    <el-button @click="test1">alert</el-button>
<!--    <div id="container">
    <Menu></Menu>
    <Worklist   :worklist="worklist"></Worklist>
        </div>-->
</template>

<style >
#container{
   overflow-y: scroll;
}

#container::-webkit-scrollbar {
  display: none;
}
</style>