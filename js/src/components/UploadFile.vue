<script setup>
import {ref} from "vue";
import {useUserStore} from "@/store/User";
import {ElMessage} from "element-plus";
import {Plus} from "@element-plus/icons-vue";
const userStore = useUserStore()
const uploadRef=ref(null);
const onUploadFile=(file)=>{
    const reader = new FileReader();
    reader.readAsDataURL(file.raw);
    reader.onload = () => {
        imgUrl.value=reader.result;
}}
const onUpdateAvatar = async () => {
  await userUploadAvatarService(imgUrl.value)
  await userStore.getUser()
  ElMessage({ type: 'success', message: '更换头像成功' })
}


defineExpose(
    {
        onUpdateAvatar
    })


</script>


<template>
  <el-upload
    ref="uploadRef"
    class="avatar-uploader"
    :auto-upload="false"
    :show-file-list="false"
    :on-change="onUploadFile"
  >
    <div class="image-container">
        <img v-if="imgUrl" :src="userStore.user.avatar" class="avatar" />
        <el-icon v-else> <Plus /></el-icon>
    </div>
  </el-upload>
</template>



<style lang="scss" scoped>
.avatar-uploader {
  display: inline-block;
  width: 100px;
  height: 100px;
  border-radius: 50%;

  .image-container {
    width: 100px;
    height: 100px;
    display: flex;
    justify-content: center; /* 水平居中对齐 */
    align-items: center; /* 垂直居中对齐 */
    //虚线
    border: 1px dashed var(--el-border-color);
      img{
            width: 100%;
            height: 100%;
            border-radius: 50%;
      }
  }

  :deep() {
    .avatar {
      width: 100px;
      height: 100px;
      display: block;
    }
    .el-upload {
      border: 1px dashed var(--el-border-color);
      border-radius: 6px;
      cursor: pointer;
      position: relative;
      overflow: hidden;
      transition: var(--el-transition-duration-fast);
    }
    .el-upload:hover {
      border-color: var(--el-color-primary);
    }
    .el-icon.avatar-uploader-icon {
      font-size: 28px;
      color: #8c939d;
      width: 278px;
      height: 278px;
      text-align: center;
    }
  }
}
</style>
