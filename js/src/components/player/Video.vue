<script setup>
import VedioPlayer from "@/components/player/VedioPlayer.vue";




</script>

<template>
<VedioPlayer></VedioPlayer>
    <div id="VideoMenu" >
         <div class="block">
             <div id="avatar">
          <el-avatar :size="50" :src="circleUrl"/>  </div>
             <img src="../../assets/icon/like1.svg" style="height: 50px;width: 50px" >
             <img src="../../assets/icon/collect.svg" alt="" style="width: 57px;height: 57px">
             <img src="../../assets/icon/share.svg" style="width: 45px;height: 45px">
        </div>
    </div>

    <div id="VedioMessage">
        <div id="creater"><strong>我是作者名字</strong></div>
        <div  id="videoname">我是视频名字</div>
    </div>



</template>

<style lang="scss" scoped>
VedioPlayer{
    z-index: 1;
    width: 200px;
}
#VideoMenu{
    display: flex;
    text-align: center;
    position: absolute;
    top: 200px;
    right: 40px;
    width: 60px;
    height: 300px;
    z-index: 2;
    background-color: transparent;

    #avatar,img{
        margin-bottom: 20px;
    }


}




#VedioMessage{
    color: #ffffff;
    position: absolute;
    bottom: 180px;
    left: 20px;
   width: 300px;
    height: 60px;
    background-color: transparent;
    z-index: 2;

    #creater{
          font-size: 30px;
          font-weight: 600;
          margin-bottom: 5px;
    }
    #videoname{
        font-size: 20px;
        font-weight: 200;
    }


}
</style>