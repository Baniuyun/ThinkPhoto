<script setup >
import { ref } from 'vue';
import { reactive } from 'vue';
import { videoPlay } from 'vue3-video-play/lib/index';
import {useUserStore} from "@/store/User";
import {likeVideoService} from "@/api/video";
const userStore = useUserStore()

const player = ref(null);
 const props=defineProps({
     videoMessage:{
         type:Object,
         default:{
                    id: Number,
                    information: String,
                    play_url: String,
                    cover_url: String,
                    user_name: String,
                    avatar:String,
                    is_favorite: Boolean,
                    favorite_count: Number,
                    tag: Array,
                    time: Number,
                    user_id: Number,
                    is_follow: String,
                }},})


const options = reactive({
  width: '100%',
    fit: 'contain',
  color: "#ffc0cb", //主题色
  muted: false, //静音
  webFullScreen: false,
  speedRate: ["0.75", "1.0", "1.25", "1.5", "2.0"], //播放倍速
  autoPlay: false, //自动播放
  loop: true, //循环播放
  mirror: false, //镜像画面
  lightOff: false,  //关灯模式
  volume: 1, //默认音量大小
  control: true, //是否显示控制器
  title: props.videoMessage.information, //视频名称
  src: props.videoMessage.play_url, //视频源
  poster: props.videoMessage.cover_url, //封面
    preload: "auto", //预加载
    keyboard: {
        ArrowLeft: 0,
        ArrowRight: 0,
        ArrowUp: 0,
        ArrowDown: 0,
    },
    //取消上下按键对音量的控制
})

const play = () => {
  player.value.play();
}
const pause = () => {
  player.value.pause();
}

const onPlay = () => {
  console.log('播放')
}
const onPause = () => {
  console.log('暂停')
}

const onTimeupdate = () => {
  console.log('时间更新')
}
const onCanplay = () => {
  console.log('可以播放')
}


const likeVideo=async ()=>{
    console.log("喜欢")
    await likeVideoService(userStore.getUserId(),props.videoMessage.id).then(
        res=>{
            console.log(res)
        }
    )

}
const collectVideo=()=>{
    console.log("收藏")
}
const shareVideo=()=>{
    console.log("分享")
}





defineExpose({
  play,
  pause,
})


</script>
<template>

  <div>
    <videoPlay
    ref="player"
    v-bind="options"
    @play="onPlay"
    @pause="onPause"
    @timeupdate="onTimeupdate"
    @canplay="onCanplay" />
  </div>


    <div id="videoMenu" >
         <div class="block">
             <div id="avatar">
          <el-avatar :size="50" :src="props.videoMessage.avatar" />  </div>
             <img src="../../assets/icon/like1.svg" style="height: 50px;width: 50px"  @click="likeVideo" >
             <img src="../../assets/icon/collect.svg" alt="" style="width: 57px;height: 57px" @click="collectVideo">
             <img src="../../assets/icon/share.svg" style="width: 45px;height: 45px"  @click="shareVideo">
        </div>
    </div>

    <div id="videoMessage">
        <div id="authorName"><strong>{{props.videoMessage.user_name}}</strong></div>
        <div  id="videoName">{{props.videoMessage.information}}</div>
    </div>


</template>

















<style lang="scss" scoped>
video-play{
    z-index: 1;
    width: 200px;
}



#videoMenu{
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

#videoMessage{
    color: #ffffff;
    position: absolute;
    bottom: 180px;
    left: 20px;
   width: 300px;
    height: 60px;
    background-color: transparent;
    z-index: 2;

    #authorName{
          font-size: 30px;
          font-weight: 600;
          margin-bottom: 5px;
    }
    #videoName{
        font-size: 20px;
        font-weight: 200;
    }
}
</style>