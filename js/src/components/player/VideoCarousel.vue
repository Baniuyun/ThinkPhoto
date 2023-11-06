<script setup>
import VideoPlayer from "@/components/player/VideoPlayer.vue";
import { ref, onMounted,reactive}from "vue";

const carouselRef = ref(null);
const loop=ref(false);
const interval=300;
let timer=ref(null);
const curactive=ref(0);
const videoRefs=ref([]);
const VideoDown = () => {
  if (carouselRef.value) {
      carouselRef.value.next();
      curactive.value=(curactive.value+1+videos.length)%videos.length;
      onChangeCarousel((curactive.value-1+videos.length)%videos.length,curactive.value);
        videoRefs.value[curactive.value].focus();
  }
};

const VideoUp = () => {
  if (carouselRef.value) {
    carouselRef.value.prev();

      curactive.value=(curactive.value-1+videos.length)%videos.length;
      onChangeCarousel((curactive.value+1+videos.length)%videos.length,curactive.value);
        videoRefs.value[curactive.value].focus();

  }
};




onMounted(() => {
    videoRefs.value[0].play();

    window.addEventListener("keydown", (e) => {
        if (timer) {
            clearTimeout(timer);
        }
        timer = setTimeout(() => {
            if (e.key === "ArrowDown") {
                VideoDown();
            } else if (e.key === "ArrowUp") {
                VideoUp();
            }
        }, interval)
    })


    window.addEventListener("mousewheel", (e) => {
        if (timer) {
            clearTimeout(timer);
        }
        timer = setTimeout(() => {
            if (e.wheelDelta < 0) {
                VideoDown();
            } else if (e.wheelDelta > 0) {
                VideoUp();
            }
        }, interval)
    })
})






const onChangeCarousel = (prevIndex, nextIndex) => {
    videoRefs.value[prevIndex].pause();
    videoRefs.value[nextIndex].play();
};




const  videos= reactive([
        {   id:1,
            play_url:"https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218114723HDu3hhxqIT.mp4",
            cover_url:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
            title:"视频1",
            favorite_count:0,
            comment_count:0,
            is_favorite:false,
            information:"视频1的信息",
            tags:["标签1","标签2"],
            user_id:1,
            is_follow:false,
            user_name:"用户1",
            avatar:"https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdM.img",
        },

          {   id:1,
           play_url:"https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218093206z8V1JuPlpe.mp4",
            cover_url:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
            title:"视频1",
            favorite_count:0,
            comment_count:0,
            is_favorite:false,
            information:"视频1的信息",
            tags:["标签1","标签2"],
            user_id:1,
            is_follow:false,
            user_name:"用户1",
            avatar:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
        },

    {   id:1,
    play_url:"https://stream7.iqilu.com/10339/article/202002/18/2fca1c77730e54c7b500573c2437003f.mp4",
        cover_url:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
        title:"视频1",
        favorite_count:0,
        comment_count:0,
        is_favorite:false,
        information:"视频1的信息",
        tags:["标签1","标签2"],
        user_id:1,
        is_follow:false,
        user_name:"用户1",
        avatar:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
    },


])
</script>

<template>
<el-carousel height="100%"  autofocus
               direction="vertical"
               :autoplay="false"
               ref="carouselRef"
               :loop="true"
                    >
    <el-carousel-item v-for="(video,index) in videos" :key="index" >
        <VideoPlayer :videoMessage="video"   ref="videoRefs"></VideoPlayer>
        </el-carousel-item>

</el-carousel>

</template>

<style >
.el-carousel .el-carousel__button {
    display: none;
  }


</style>