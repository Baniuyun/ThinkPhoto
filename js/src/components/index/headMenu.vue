<script setup>
import {useUserStore} from "@/store/User";
import {ChatRound, DocumentAdd, Money, Notification, User} from "@element-plus/icons-vue";
const emit = defineEmits(["changeRegisterVisible"]);
const userStore = useUserStore();
const loginDialogShow = () => {
    emit("changeRegisterVisible");
};
const publishDialogShow = () => {
    emit("changePublishVisible");
};
</script>

<template>

            <div id="head-menu">
    <el-button class="head-menu-button" @click="publishDialogShow"><el-icon><DocumentAdd/></el-icon>投稿</el-button>
     <el-button type="primary"  round    id="loginbutton"   @click="loginDialogShow" v-if="!userStore.isLogin()">
      <el-icon style="font-size: 30px"  > <User/></el-icon>登录</el-button>
<!--      element下拉栏          -->
        <el-dropdown v-if="userStore.isLogin()">
          <span class="el-dropdown-link">
             <el-avatar src="userStore.getAvatar"></el-avatar>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="userStore.loginOut()">
                  退出登录
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>

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
    outline: none;
    .head-menu-button{
        margin-right: 20px;
        font-size: 20px;
        line-height: 50px;
        color: #ffffff;
        border: none;
        background-color: transparent;

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
       background-color: rgba(147,228,247,0.9);
        width:110px;
        height: 50px;
        text-align: center;
    }
}


</style>