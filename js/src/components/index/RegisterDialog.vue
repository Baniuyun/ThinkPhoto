<script setup>
    import {useUserStore} from "@/store/User";
    const userStore = useUserStore()
    import {ref} from "vue"
    const dialogVisible = ref(false)
    function changeVisable(){
        dialogVisible.value = !dialogVisible.value
    }
    defineExpose(
        {
            changeVisable,
        }
    )







</script>
<template>
    <el-dialog

    v-model="dialogVisible"
    width="300px"
    element-loading-background="rgba(0, 0, 0, 0.8)"
    title="登录">
<!--    :before-close="handleClose">-->
    <span>

       <el-form>
        <el-form-item label="账号">
            <el-input v-model="userStore.user.account" placeholder="请输入账号" size="large"></el-input>
        </el-form-item>
        <el-form-item label="密码">
            <el-input type="password" v-model="userStore.user.password" placeholder="请输入密码"  show-password></el-input>
        </el-form-item>
    </el-form>

    </span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">登录</el-button>
      </span>
    </template>
  </el-dialog>
</template>


<style>
.my-info-dialog .el-dialog__body {
        background-color: green;
    }


</style>
