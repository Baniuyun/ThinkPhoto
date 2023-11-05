<script setup>
import Video from "@/components/player/Video.vue";

import { ref, onMounted } from "vue";

const carouselRef = ref(null);
const loop=ref(false);
const VideoDown = () => {
  if (carouselRef.value) {
    carouselRef.value.next();
  }
};

const VideoUp = () => {
  if (carouselRef.value) {
    carouselRef.value.prev();
  }
};

onMounted(() => {
  // 在全局监听键盘事件
  window.addEventListener("keydown", (event) => {
    if (event.key === "ArrowDown") {
      VideoDown();
    } else if (event.key === "ArrowUp") {
      VideoUp();
    }
  });
  //鼠标滑轮滚动添加事件
    window.addEventListener("mousewheel", (event) => {
        if (event.wheelDelta < 0) {
        VideoDown();
        } else if (event.wheelDelta > 0) {
        VideoUp();
        }
    });
});







</script>
<template>

  <el-carousel height="100%"
               direction="vertical"
               :autoplay="false"
               ref="carouselRef"
               loop="false"

                    >

    <el-carousel-item >
        <Video></Video>
    </el-carousel-item>
      <el-carousel-item >
          <Video></Video>
      </el-carousel-item>
      <el-carousel-item >
          <Video></Video>
      </el-carousel-item>
      <el-carousel-item >
          <Video></Video>
      </el-carousel-item>

  </el-carousel>
</template>

<style>
 .el-carousel .el-carousel__button {
    display: none;
  }


</style>