<script setup>

</script>

<template>
   <div id="usermenu">
       <el-menu>
           <el-menu-item index="1">作品</el-menu-item>
           <el-menu-item index="2">喜欢</el-menu-item>
           <el-menu-item index="3">收藏</el-menu-item>
</el-menu>

   </div>
</template>

<style lang="scss" scoped>
   div{
       display: flex;
       /*background-color: skyblue;*/
       height: 70px;
       width: 100%;




       .el-menu{
           display: flex;
           flex-direction: row;
           background-color: transparent;
           border: none;


        .el-menu-item{
            font-size: 25px;
            height: 70px;
            color: white;
            border-radius: 20px;
        }
           .el-menu-item.is-active{
               background-color: rgba(255,192,203,0.2);
               color: rgba(255,192,203,1);
           }
           .el-menu-item:hover{
               background-color: rgba(255,192,203,0.2);
           }


       }
   }

</style>