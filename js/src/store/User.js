import  {defineStore} from "pinia";
import {reactive} from "vue";

export const useUserStore = defineStore("User",

    ()=>{
    const user = reactive({
        user_id: "",
        fans:"",
        avatar: "",
        name: "",
        favorite_count:"" ,
        like:"",
        works:"",
        token:"",
    })

    function  setUser() {
    }

    function printUser() {
        alert(user.account)
    }


    return {user, setUser,printUser}


},
{
    persist: true
}
)