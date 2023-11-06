<script setup>
    import {useUserStore} from "@/store/User";
    const userStore = useUserStore()
    import {ref} from "vue"
    import UploadFile from "@/components/UploadFile.vue";
    const EditdialogVisible = ref(false)


    const changeEditVisible=()=>{
        EditdialogVisible.value = !EditdialogVisible.value;
    }







defineExpose(
        {
            changeEditVisible,
        }
    )





</script>
<template>
   <el-dialog
    :custom-class="['profit-dialog-class']"
    v-model="EditdialogVisible"
    width="300px"
    element-loading-background="rgba(0, 0, 0, 0.8)"
    >
<!-- 放置logo和一些提示       -->
      <template #header="{logo,message }">
      <div id="dialog-header">
          <img src="../../assets/logo.png" alt="">
        <h4 :id="titleId" :class="titleClass"></h4>
      </div>
    </template>

 <el-form  >
      <el-form-item prop="avatar" >
          <div id="uploadAvatar">
             <UploadFile></UploadFile>
              <br>
              <el-button>点击上传头像</el-button>
          </div>

      </el-form-item>
      <el-form-item  label="昵称">
        <el-input  placeholder="请输入昵称"></el-input>
           <el-button type="primary" >保存</el-button>
      </el-form-item>
    </el-form>








    <template #footer>
      <span class="dialog-footer">
      </span>
    </template>
  </el-dialog>
</template>


<style lang="scss" >
.profit-dialog-class {
        width: 500px;
        height: 500px;
     background-image: linear-gradient( 135deg, #FFD3A5 -30%, #FD6585 100%);
        position: absolute;
        top: 12%;
        left: 37%;
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
        height: 350px;
        border-radius: 15px;

        .el-form-item{
            position: relative;
            left:8%;
            width: 80%;
            /*使uploadFile居中对齐*/
            #uploadAvatar{
                display: block;
                margin-left: 40%;
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
                margin: 10px auto;
            }


        }

    }
    .switchLoginAndRegister{
        text-align: center;
    }


}


</style>