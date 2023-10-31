<script setup>
//导入
import Register from "@/components/index/RegisterDialog.vue";
import Aside from "@/components/index/indexAside.vue";
import SearchBox from "@/components/index/search.vue";
import HeadMenu from "@/components/index/headMenu.vue";
import {ref} from "vue";


//登录对话框部分
const RegisterDialog=ref(null)
const changeVisible= ()=>{
    RegisterDialog.value.changeVisable();
}



</script>




<template id="body">
  <Register ref="RegisterDialog"></Register>


  <div class="common-layout">
    <el-container id="left">
        <el-aside>
     <Aside></Aside>
        </el-aside>


      <el-container id="right">
        <el-header id="head">
         <SearchBox> </SearchBox>

             <HeadMenu @changeVisible="changeVisible"></HeadMenu>

        </el-header>


<!--主视图-->
        <el-main id="main">
            <router-view style="width: 100%;height: 100%"></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>



<style lang="scss" scoped>



.common-layout{
   text-align: center;
    .el-aside{
    position: absolute;
        display: flex;
        flex-direction: column;
        align-items: center;
    left: 0;
    top: 0;
    width: 150px;
    height: 100%;

}



#right{
    position: absolute;
    left: 150px;
    top: 0;
    width: calc(100% - 150px);
    height: 100%;
    text-align: left;

    #head{
        height: 80px;
        width: 100%;
        //background-color: #409EFF;
        display: flex;
        align-items: center;
        font-size: 20px;
        font-weight: bold;
        justify-content: space-between;

        #search{
            width: 400px;
            /*设置在中间*/
            margin: 0 auto;
            display: flex;
        }
        #head-menu{
            position: relative;
            display: flex;
            align-items: center;
            right: 20px;
            top: 0;
            height: 20px;
            color: red;
        .head-menu-button {
          display: flex;
            flex-direction: column;
          align-items: center;
            width: 40px;
            height: 30px;
            background-color: transparent;
          text-align: center;
            margin: 20px auto;
             el-icon {
            font-size: 24px; /* 调整图标大小 */
}
}



            .el-dropdown-link{
                //color: white;
                cursor: pointer;
            }



        }



    }
    .el-main{
        height: calc(100% - 60px);
        width: 100%;
    }
}

}
</style>