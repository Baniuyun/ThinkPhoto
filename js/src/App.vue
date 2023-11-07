<script setup>
import { useRoute,useRouter} from "vue-router";
const router = useRouter();
const route = useRoute();

</script>


<template>
    <router-view></router-view>
</template>


<style>
    html,body{
        background: linear-gradient(220.55deg, #FF9D7E 0%, #4D6AD0 100%);
    /* background: linear-gradient(220.55deg, #DDE4FF 0%, #8DA2EE 100%);*/
       /* background: linear-gradient(220.55deg, #4063BC 0%, #6B0013 100%);*/
   }
    *{
        margin: 0;
        padding: 0;
    }
    .el-message__wrapper {
    z-index: 100; /* 设置一个足够高的 z-index 值 */
}


</style>