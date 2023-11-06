<script setup>

import {
    Basketball,
    Headset,
    House,
    MagicStick,
    Notebook,
    PictureRounded,
    Position,
    Star,
    User,
    View
} from "@element-plus/icons-vue";
</script>

<template>
          <div class="aside_logo" >
              <img src="@/assets/logo.png"  alt=""> <div>享看视频</div>
          </div>
<!--使用elementplus的menu生成侧边菜单栏，包括首页，推荐，我的，热门，体育，二次元，知识，娱乐，音乐-->
          <el-menu
             router>

           <el-menu-item index="home">
          <el-icon><House/></el-icon>
          <span>首页</span>
        </el-menu-item>


              <el-menu-item index="recommend">
                <el-icon><Star /></el-icon>
                  <span>推荐</span>
              </el-menu-item>


              <el-menu-item index="user">
                 <el-icon><User /></el-icon>
                  <span>我的</span>
              </el-menu-item>


              <el-menu-item index="hot">
             <el-icon><Position/></el-icon>
                  <span>热门</span>
              </el-menu-item>


              <el-menu-item index="sport">
                <el-icon><Basketball /></el-icon>
                  <span>体育</span>
              </el-menu-item>



              <el-menu-item index="cartoon">
                <el-icon><MagicStick/></el-icon>
                  <span>二次元</span>
              </el-menu-item>


          <el-menu-item index="knowledge">
            <el-icon><Notebook /></el-icon>
              <span>知识</span>
          </el-menu-item>


          <el-menu-item index="entertainment">
            <el-icon><PictureRounded/></el-icon>
              <span>娱乐</span>
          </el-menu-item>


          <el-menu-item index="news">
         <el-icon><View /></el-icon>
              <span>新闻</span>
          </el-menu-item>

               <el-menu-item index="music">
         <el-icon><Headset /></el-icon>
              <span>音乐</span>
                   </el-menu-item>
          </el-menu>

</template>

<style lang="scss" scoped>
    .aside_logo{
          position: relative;
        top:30px;
        display: flex;
        justify-content: center;
         margin-bottom: 40px;

        img{
            width: 40px;
            height: 40px;

        }
        div{
            position: relative;
            width: 80px;
            font-size: 20px;
            height: 40px;
            line-height: 40px;
            color: white;
            font-weight: bold;
            margin-left: 10px;
        }

    }


    .el-menu
    {
         width:100%;
         background-color: transparent;
         top: 30px;
        border: none;
        .el-menu-item{
            color: white;
            height: 60px;
            border-radius: 20px;
            display: flex;
            align-items: center;

            .el-icon,span{
                margin-left: 10px;
            }

        }

        .el-menu-item:hover {
            background-color: rgba(255, 192, 203, 0.2);
        }


         .el-menu-item.is-active{

            background-color: rgba(255,192,203,0.2);
            color: rgba(255,192,203,1);
     }


    }



</style>