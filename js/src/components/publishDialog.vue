<script setup>
    import {ref} from "vue"
    import {Plus} from "@element-plus/icons-vue";
    import {ElMessage} from "element-plus";
    import {getPublishTokenService,publishVideoService} from "@/api/video";
    import * as qiniu from 'qiniu-js'
    import {useUserStore} from "@/store/User";
    import {generateRandomString} from "@/utils/file";
    import router from "@/router";
    const userStore = useUserStore()
    const dialogVisible = ref(false)
    const uploadFormRef = ref(null)
const changeVisible = () => {
    dialogVisible.value = !dialogVisible.value
}
const uploadVideoButtonRef = ref(null)
const titleId = ref("titleId")
const titleClass = ref("titleClass")
const uploadRef = ref(null)
const videoUrl = ref(null)
const videoFile=ref(null)
let videoKey=""
const onUploadFile =  (file) => {
    console.log(file)
    videoUrl.value = URL.createObjectURL(file.raw);
     videoFile.value = file;
    //上传视频到七牛云,并获取key

}

const summitForm  =async ()=>{
      await uploadFormRef.value.validate()
      await getPublishTokenService().then(res=>{
    const token = res.data.token
    console.log(token)
    videoKey=generateRandomString()
      const putExtra = {
        fname: "",
        params: {},
        mimeType: ["video/mp4", "video/x-m4v", "video/*"],
    };
    const config = {
        useCdnDomain: true,
        region: qiniu.region.z0,
    };
    const observable = qiniu.upload(videoFile.value, videoKey.value, token, putExtra, config);
    const observer = {
        next(res) {
            console.log(res)
        },
        error(err) {
            console.log(err)
        },
        complete(res) {
            console.log("返回的key值是："+res.data.key)
            console.log(res)
            videoKey=res.data.key
        },
    };
    const  subscription = observable.subscribe(observer);
})
     await   publishVideoService(videoKey,userStore.getAvatar(),uploadForm.value.classification.toString(),uploadForm.value.information,userStore.getUserId().toString(),userStore.getUsername())
              .then(res=>{
                  console.log(res)
                  ElMessage.success("上传成功")
                  changeVisible()
              }).catch(err=>{
                  console.log(err)
                  ElMessage.error("上传失败")
              }
        )
    }



//进行视频上传前的校验
const beforeUpload = (file) => {
    const isVideo = file.type === 'video/mp4';
    const isLt2M = file.size / 1024 / 1024 < 100;
    if (!isVideo) {
        ElMessage.error('上传视频只能是 MP4 格式!');
    }
    if (!isLt2M) {
        ElMessage.error('上传视频大小不能超过 100MB!');
    }
    return isVideo && isLt2M;
}
const uploadVideo = () => {
    uploadRef.value.click()
}
const uploadVideoButtonClick = () => {
    uploadVideoButtonRef.value.click()
}
const uploadForm = ref({
    information: '',
    classification: 1,
})

const uploadFormRules = ref({
    information: [
        {required: true, message: '请输入视频描述', trigger: 'blur'},
        {min: 3, max: 16, message: '长度在 3 到 16 个字符', trigger: 'blur'}
    ],
})





defineExpose(
    {
        changeVisible,
    }
    )





</script>
<template>
  <el-dialog
    :custom-class="['publish-dialog-class']"
    v-model="dialogVisible"
    width="300px"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <!-- 放置logo和一些提示       -->
    <template #header="{ logo, message }">
      <div id="dialog-header">
        <img src="@/assets/logo.png" alt="">
        <h4 :id="titleId" :class="titleClass"></h4>
      </div>
    </template>

    <el-form v-model="uploadForm"  :rules="uploadFormRules"  ref="uploadFormRef">
         <el-form-item label="视频预览">
        <div v-if="videoUrl" class="video-container" >
          <video :src="videoUrl" controls="controls" style="width: 100%; height: 100%"></video>
        </div>
             <div v-else id="img-container" @click="uploadVideoButtonClick">
                 <el-icon> <Plus /></el-icon>
             </div>
      </el-form-item>


      <!-- 生成一个上传视频的form元素，并且在上传的时候显示仪表盘型进度条,上传完成后在播放器中获取视频地址并且可以预览   -->
      <el-form-item >
        <el-upload
          ref="uploadRef"
          class="video-uploader"
          :auto-upload="false"
          :show-file-list="false"
          :on-change="onUploadFile"
        >
              <el-button type="primary" @click="uploadVideo" id="uploadVideoButton" ref="uploadVideoButtonRef">上传视频</el-button>
        </el-upload>

      </el-form-item>
    <el-form-item label="视频简介">
         <el-input
            v-model="uploadForm.information"
            placeholder="请输入视频描述"
        ></el-input>
    </el-form-item>
    <el-form-item label="视频分类">
        <el-radio-group v-model="uploadForm.classification" id="class"   fill="#FD6585" >
            <el-radio-button label="1" name="sports">体育</el-radio-button>
            <el-radio-button label="2" name="type">二次元</el-radio-button>
            <el-radio-button label="3" name="type">知识</el-radio-button>
            <el-radio-button label="4" name="type">娱乐</el-radio-button>
            <el-radio-button label="5" name="type">新闻</el-radio-button>
            <el-radio-button label="6" name="type">音乐</el-radio-button>
          </el-radio-group>
    </el-form-item>
        <el-form-item>
            <el-button @click="summitForm" >提交</el-button>
        </el-form-item>


    </el-form>

    <template #footer>
      <span class="dialog-footer"></span>
    </template>
  </el-dialog>
</template>



<style lang="scss" >
.publish-dialog-class {
        width: 600px;
        height: 700px;
     background-image: linear-gradient( 135deg, #FFD3A5 -30%, #FD6585 100%);
        position: absolute;
        left: 33%;
        border-radius: 15px;

    #dialog-header{
        top: 10px;
        position: relative;
        display: flex;
        justify-content: center;

        img{
            width: 60px;
            height: 60px;
            margin-left: 15px;
        }

    }



    .el-form{
        padding-top: 20px;
        background: white;
        width: 100%;
        height: 500px;
        border-radius: 15px;

        .el-form-item{
            position: relative;
            left:8%;
            width: 80%;

            .video-container{
                height: 220px;
                margin: auto auto;
                position: relative;
                z-index: 2;
            }
            #uploadVideoButton{
                position: relative;
                z-index: 3;
                /*居中*/
                left: 200px;
                margin-top: -50px;
            }
            #img-container{
                position: relative;
                z-index: 2;
                width: 100%;
                height: 200px;
                background-color: #f5f5f5;
                display: flex;
                justify-content: center;
                align-items: center;
                border-radius: 10px;
            }
            #class{
                position: relative;
                width: 100%;
                height: 50px;
                display: flex;
                justify-content: center;
                align-items: center;
                border-radius: 10px;
            }


            .el-input{
                height: 40px;
                border-radius: 10px;
                border: none;
                box-shadow: 0 0 10px 0 rgba(0,0,0,0.1);
                line-height: 40px;
            }
            .el-input:focus{
                border: none;

            }

            .el-button{
                width: 80%;
                height: 40px;
                border-radius: 10px;
                border: none;
                box-shadow: 0 0 10px 0 rgba(0,0,0,0.1);
                line-height: 40px;
                background: linear-gradient( 135deg, #FFD3A5 -30%, #FD6585 100%);
                margin: 0 auto;
                 display: flex;
                justify-content: center;
                align-items: center;
            }


        }

    }
    .switchLoginAndRegister{
        text-align: center;
    }


}
</style>