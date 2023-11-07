<script setup>
import {ref} from "vue";
import  EditDialog from "@/components/user/EditDialog.vue";
import {useUserStore} from "@/store/User";
const userStore = useUserStore()
const isLogin=ref(false);
const isFollow=ref(false);
const isSelf=ref(false);
const editDialog=ref(null);
const EditMessage=()=>{
    console.log("编辑资料");
    changeEditVisible();

}
const Follow=()=>{
    console.log("关注");
}
const UnFollow=()=>{
    console.log("取消关注");
}



const changeEditVisible= ()=>{
    editDialog.value.changeEditVisible();
}





const EditVisible=()=>{
    return isSelf.value;
}
const FollowVisible=()=>{
    return !isSelf.value&&!isFollow.value;
}
const UnFollowVisible=()=>{
    return !isSelf.value&&isFollow.value;
}

</script>

<template>
    <EditDialog ref="editDialog"></EditDialog>


<div id="head">
    <el-avatar src="@" id="avatar"></el-avatar>
    <div id="message">
      <div id="username">
          {{userStore.getUsername()}}
      </div>


      <div id="others">
        <div id="follow">关注</div>
        <div id="fans">粉丝</div>
        <div id="like">获赞</div>
      </div>
    </div>

   <el-button class="bt" v-if=" true"  @click="changeEditVisible">编辑资料</el-button>
    <el-button class="bt" v-if="false"  @click="Follow">关注</el-button>
    <el-button   class="bt" v-if="false"  @click="UnFollow">取消关注</el-button>

  </div>
</template>

<style lang="scss" scoped>
#head {
  width: 100%;
  height: 200px;
  background-color: transparent;
  display: flex;
  align-items: center;
  text-align: center;

  #avatar {
    width: 150px;
    height: 150px;
  }

  #message {
      color: white;
    background-color: transparent;
    position: relative;
    bottom: 20px;
    left: 40px;
    display: flex;
    flex-direction: column;
    height: 50px;
    width: 300px;

    #username {
      font-size: 30px;
        text-align: left;
    }

    #others {
        font-size: 20px;
      background-color: transparent;
      display: flex;
      flex-direction: row;
        justify-content: left;
        margin-top: 20px;
      div{
          width: 100px;
          text-align: left;

      }
    }

  }
    .bt{
        position: relative;
        left: 50%;
        top: 20px;
        width: 100px;
        height: 50px;
        background-color: skyblue;
        color: white;
        line-height: 50px;
        border: none;
        border-radius: 15px;
    }



}
</style>