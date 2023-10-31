<script setup>
import {useUserStore} from "@/store/User";
import {Search} from "@element-plus/icons-vue";
const userStore = useUserStore();
</script>

<template>
  <div id="search">
                <el-input  v-model="userStore.user.test" id="in" >
                </el-input>
                <el-button><el-icon size="15px" style="margin-right: 10px" ><Search/></el-icon>搜索</el-button>
            </div>
</template>

<style scoped>
   /*修改el-input的背景颜色为透明*/

</style>