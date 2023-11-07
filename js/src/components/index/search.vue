<script setup>
import {Search} from "@element-plus/icons-vue";
import router from "@/router";
import {useHomeStore} from "@/store/Home";
const HomeStore=useHomeStore();



const getSearchContent = to => {
    router.push({
        path: "/search",
    });};

</script>

<template>
  <div id="search">
              <input id="searchinput" v-model="HomeStore.searchKeyword">
                <el-button @click="getSearchContent"><el-icon size="15px" style="margin-right: 10px" ><Search/></el-icon>搜索</el-button>
            </div>
</template>

<style lang="scss" scoped>
  #search {
     border: none;
     display: flex;
     position: relative;
     left: 100px;
      font-size:17px;
     #searchinput {
         background-color: rgba(255, 192, 203, 0.5);
         width: 400px;
         height: 50px;
         border: none;
         border-radius: 15px 0 0 15px;
         outline: none; // 去除默认轮廓样式
         color: white;
         text-indent: 20px;
         line-height: 60px;
         font-size: 16px;

     }
     .el-button {
         height: 50px;
         width: 100px;
         background-color: rgba(255, 192, 203, 0.5);
         border: none;
         border-radius: 0 15px 15px 0;
         color: gray;
         font-size: 17px;
         line-height: 50px;

     }

     &:hover .el-button {
         background-color: rgba(255,192,203,0.5);
         color: white;
     }
 }
</style>