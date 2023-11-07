<script setup>
import VideoPlayer from "@/components/player/VideoPlayer.vue";
import { ref, onMounted,reactive}from "vue";

const carouselRef = ref(null);
const loop=ref(false);
const interval=300;
let timer=ref(null);
const curactive=ref(0);
const videoRefs=ref([]);
const props=defineProps({

    videos:{
        type:Array,
        default:()=>[]
    }
})
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