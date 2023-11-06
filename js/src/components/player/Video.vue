<script setup>
import VideoPlayer from "@/components/player/VideoPlayer.vue";




</script>

<template>




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
