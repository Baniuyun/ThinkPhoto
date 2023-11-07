import  {defineStore} from "pinia";
import {reactive} from "vue";
import {getUserInfoService} from "@/api/user";


export const useUserStore = defineStore("User",

    ()=>{
    const user = reactive({
        user_id: 0,
        username: "未登录",
        avatar: "",
        token:"",
        following_count:0,
        follower_count:0
    })
   const isLogin=()=> {
       return localStorage.getItem("token")
   }

   const setUser_id=(user_id)=> {
       user.user_id = user_id
       localStorage.setItem("user_id", user_id)
   }
   const updateUser=async ()=>{
        let id=parseInt(getUserId(),10)
        console.log(id+typeof(id))
        getUserInfoService(id).then(
            (res)=>{
                console.log(res)
                setUserInfo(res.data.user_id, res.data.user_name, res.data.avatar,res.data.following_count,res.data.follower_count)

            }

        )
   }

 const getUserId=()=>
     localStorage.getItem("user_id")

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
    const setUserInfo=(user_id, username, avatar,fowllowing_count,follower_count)=> {
        user.user_id = user_id
        localStorage.setItem("user_id", user_id)
        user.username = username
        user.avatar = avatar
        user.following_count=fowllowing_count
        user.follower_count=follower_count
    }
    const loginOut=()=> {
        user.user_id = 0
        user.username = "未登录"
        user.avatar = ""
        user.token = ""
        localStorage.removeItem("token")
    }

    return {setToken,loginOut,isLogin,setUserInfo,setUser_id,getUserId,getUsername,getAvatar,getToken,updateUser}

},
{
    persist:true
}

    )
