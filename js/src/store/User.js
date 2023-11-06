import  {defineStore} from "pinia";
import {reactive} from "vue";

export const useUserStore = defineStore("User",

    ()=>{
    const user = reactive({
        user_id: 0,
        username: "未登录",
        avatar: "",
        token:"",
    })
   const isLogin=()=> {
       return localStorage.getItem("token")
   }

   const setUser_id=(user_id)=> {
       user.user_id = user_id
   }

 const getUserId=()=>
      user.user_id

 const getUsername=()=>
      user.username


    const getAvatar=()=>
         user.avatar
    const getToken=()=>
        user.token
    const setToken=(token)=> {
        user.token = token
        localStorage.setItem("token", token)
    }
    const setUserInfo=(user_id, username, avatar)=> {
        user.user_id = user_id
        user.username = username
        user.avatar = avatar
    }
    const loginOut=()=> {
        user.user_id = 0
        user.username = "未登录"
        user.avatar = ""
        user.token = ""
        localStorage.removeItem("token")
    }

    return {setToken,loginOut,isLogin,setUserInfo,setUser_id,getUserId,getUsername,getAvatar,getToken}

},
{
    persist:true
}

    )
