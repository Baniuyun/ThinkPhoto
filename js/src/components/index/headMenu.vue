<script setup>

import {ChatRound, DocumentAdd, Money, Notification, User} from "@element-plus/icons-vue";
const emit = defineEmits(["changeVisible"]);
const loginDialoShow = () => {
    emit("changeVisible");
};
</script>

<template>

            <div id="head-menu">
               <el-dropdown>
        <span class="head-menu-button" ><el-icon><Money/></el-icon>充值</span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>华为</el-dropdown-item>
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>中国</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

<!-- 通知下拉菜单栏     -->
                <el-dropdown>
    <span class="head-menu-button"><el-icon><Notification/></el-icon>通知</span>
      <template #dropdown>
        <el-dropdown-menu >
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>华为</el-dropdown-item>
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>中国</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

<!--  私信下拉菜单栏 -->
          <el-dropdown>
    <span class="head-menu-button"><el-icon><ChatRound/></el-icon>私信</span>
      <template #dropdown>
        <el-dropdown-menu >
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>华为</el-dropdown-item>
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>中国</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
<!-- 投稿-->
         <el-dropdown>
    <span class="head-menu-button"><el-icon><DocumentAdd/></el-icon>投稿</span>
      <template #dropdown>
        <el-dropdown-menu >
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>华为</el-dropdown-item>
          <el-dropdown-item>加油</el-dropdown-item>
          <el-dropdown-item>中国</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>








     <el-button type="primary" style="background-color: red;width:110px;height: 50px;text-align: center;line-height: 50px;font-size: 20px;" round    id="loginbutton"   @click="loginDialoShow">
      <el-icon style="font-size: 30px"  > <User/></el-icon>登录</el-button>
            </div>
</template>

<style lang="scss" scoped>
#head-menu{
    display: flex;
    position: absolute;
    top: 20px;
    right: 0;
    height: 50px;
    background-color: transparent;
    z-index: 2;
    .head-menu-button{
        margin-right: 20px;
        font-size: 20px;
        line-height: 50px;
        color: #ffffff;
        border: none;

        .el-icon{
            position: relative;
            top: 3px;
        }



    }
    #loginbutton{
        margin-right: 0;
        font-size: 20px;
        line-height: 50px;
        color: #ffffff;
        border: none;
    }
}


</style>