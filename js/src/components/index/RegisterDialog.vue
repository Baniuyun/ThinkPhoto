<script setup>
    import {useUserStore} from "@/store/User";
    const userStore = useUserStore()
    import {ref} from "vue"
    import {userRegisterService,userLoginService} from "@/api/user";
    import router from "@/router";
    const dialogVisible = ref(false)
    const loginformVisible = ref(true)
    const registerformVisible = ref(false)
    const loginFormRef = ref(null)
    const registerFormRef = ref(null)



    const changeVisable=()=>{
        dialogVisible.value = !dialogVisible.value;
    }












    /*登录表单*/
    const loginForm = ref({
        username: '',
        password: '',
    })
    const loginFormRules = ref({
        username: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 5, max: 16, message: '长度在 5 到 16 个字符', trigger: 'blur'}
        ],
        password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {pattern: /^[a-zA-Z0-9_-]{8,16}$/, message: '密码必须是8-16位的字母或数字', trigger: 'blur'}
        ],
    })

    /*切换注册表单*/




    /*注册表单*/
    const registerForm = ref({
        username: '',
        password: '',
        repassword: '',
    })

    const registerFormRules = ref({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 5, max: 16, message: '长度在 5 到 16 个字符', trigger: 'blur' }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { pattern: /^[a-zA-Z0-9_-]{8,16}$/, message: '密码必须是8-16位的字母或数字', trigger: 'blur' }
    ],
    repassword: [
        {
            validator: (rule, value, callback) => {
                if (value !== registerForm.value.password) {
                    callback(new Error('两次输入密码不一致!'))
                } else {
                    callback()
                }
            },
            trigger: 'blur'
        }
    ]
})
const switchRegister = () => {
        loginformVisible.value = false
        registerformVisible.value = true
        loginForm.value.username="";
        loginForm.value.password="";
    }
    /*切换登录表单*/
    const switchLogin = () => {
        loginformVisible.value = true
        registerformVisible.value = false
        registerForm.value.username="";
        registerForm.value.password="";
        registerForm.value.repassword="";
    }


    const login =  async () => {
        await loginFormRef.value.validate()
        await userLoginService(loginForm.value.username,loginForm.value.password).then((res)=>{
                console.log(res)
                userStore.setToken(res.data.token.accss_token)
                userStore.setUser_id(res.data.user_id)
            }
        )



  }





    const register = async () => {
        await registerFormRef.value.validate()
        await userRegisterService(registerForm.value.username,registerForm.value.password).then((res)=>{
                userStore.setToken(res.data.token.access_token)
                userStore.setUser_id(res.data.user_id)
            }
        )
         changeVisable()
         router.push({path: "/search"})
    }






defineExpose(
        {
            changeVisable,
        }
    )





</script>
<template>
    <el-dialog
    :custom-class="['register-dialog-class']"
    v-model="dialogVisible"
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



<!--    创建登录表单    -->
        <span v-if="loginformVisible">
        <el-form
        :model="loginForm"
        ref="loginFormRef"
        :rules="loginFormRules"
        >
            <el-form-item
            prop="username"
            label="账号"
            >
                <el-input
                v-model="loginForm.username"
                placeholder="请输入用户名"
                >

                </el-input>
            </el-form-item>
            <el-form-item
            label="密码"
            prop="password"
            >
                <el-input
                v-model="loginForm.password"
                placeholder="请输入密码"
                show-password
                ></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary"   @click="login">登录</el-button>
            </el-form-item>

        </el-form>

<!--   注册账号的字样    -->
        <div  @click="switchRegister" class="switchLoginAndRegister">
            <span>注册账号</span>
        </div>

    </span>

<!--    创建注册表单    -->
    <span v-if="registerformVisible">
        <el-form
        :model="registerForm"
        ref="registerFormRef"
        :rules="registerFormRules"
        >
            <el-form-item
            prop="username"
            label="账号"
            >
                <el-input
                v-model="registerForm.username"
                placeholder="请输入用户名"
                >

                </el-input>
            </el-form-item>
            <el-form-item
            label="密码"
            prop="password"
            >
                <el-input
                v-model="registerForm.password"
                placeholder="请输入密码"
                show-password
                ></el-input>
            </el-form-item>
            <el-form-item
            label="确认密码"
            prop="repassword"
            >
                <el-input
                v-model="registerForm.repassword"
                placeholder="请再次输入密码"
                show-password
                ></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="register">注册</el-button>
            </el-form-item>

        </el-form>

         <div  class="switchLoginAndRegister" @click="switchLogin">
            <span  >已有账号，我要登录</span>
        </div>

        </span>








    <template #footer>
      <span class="dialog-footer">
      </span>
    </template>
  </el-dialog>
</template>


<style lang="scss" >
/*设置登录框的style样式*/
.register-dialog-class {
        width: 400px;
        height: 400px;
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
        height: 250px;
        border-radius: 15px;

        .el-form-item{
            position: relative;
            left:8%;
            width: 80%;

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
            }


        }

    }
    .switchLoginAndRegister{
        text-align: center;
    }


}



</style>
