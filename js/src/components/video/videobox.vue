<script setup>
import router from "@/router";
import {ref} from "vue";
const props=defineProps({
    videoShortMessage:{
        type:Object,
        default:{
            video_id:null,
            information:"视频信息未加载",
            cover_url:"",
        }
    }
})



const playVideo= ()=>{
    console.log(props.videoShortMessage.video_id)
    router.push({path:"/home", query:{playVideoId:props.videoShortMessage.video_id}})
}

const getVideoInformation =async ()=>{
}
</script>

<template>
   <div id="box" @click="playVideo" >
       <img :src="props.videoShortMessage.cover_url" alt="" >
     <div id="videoname">{{props.videoShortMessage.information}}</div>
   </div>
</template>
<style lang="scss" scoped>
       #box{
           width: 20%;
           height: 200px;
           margin-left: 10px;
           margin-top: 2px;
           cursor: pointer;
           border-radius: 15px;

           img{
               width: 100%;
               height: 70%;
               border-radius: 15px;
           }
            #videoname{
                width: 100%;
                height: 30%;
                margin-top: -6px;
                color: white;
                font-size: 17px;
            }
       }
       #box:hover{
           transform: scale(1.02);
           background-color: rgba(255,192,203,0.2);
       }
       #videoPlay{
           position: absolute;
           margin: auto auto;
           width: 100vh;
           height: 100vh;
       }

</style>