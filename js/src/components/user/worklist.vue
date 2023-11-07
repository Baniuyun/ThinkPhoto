<script setup>
import { ref, computed } from 'vue';
import Videobox from "@/components/video/videobox.vue";
const loading = ref(false);
const noMore = ref(false);
const canLoadMore = computed(() => !loading.value && !noMore.value);
let videoRow=0



 const props = defineProps({
     //获取视频信息的函数
      getMoreVideoFunction: {
        type: Function,
        default: () => () => []
      }
    });
//视频信息
 const videoMessages = ref([])




const loadMore = () => {
  if (canLoadMore.value) {
    // 模拟加载更多数据
    const lastVideoIndex = videoMessages.value.length
    console.log(lastVideoIndex)
    const moreVideo=props.getMoreVideoFunction();
    console.log(moreVideo)
    if(moreVideo){
      for(let i=lastVideoIndex;i<moreVideo.length;i++){
          videoMessages.value.push(moreVideo[i])
          videoRow=Math.floor(videoMessages.value.length/5)+1
      }
    }

  } else {
    stopLoad();
  }
};

const stopLoad = () => {
  loading.value = false;
  noMore.value = true;
};

</script>

<template>
  <div class="container" v-infinite-scroll="loadMore" infinite-scroll-distance="40px" :infinite-scroll-disabled="!canLoadMore">
    <ul v-if="videoMessages.length!==0">

        <li v-for="index in videoRow" :key="index">
            <Videobox v-for="item in videoMessages.slice((index-1)*5,index*5)"  :videoShortMessage="item"></Videobox>
        </li>
    </ul>
    <div class="loading" v-if="canLoadMore.value">加载中...</div>
    <div class="nomore" v-if="!canLoadMore.value">暂无更多</div>
  </div>
</template>

<style lang="scss" scoped>
.container {
  height: 96%;
  overflow: auto;
  width: 100%;
  position: relative;
}

ul li{
    display: flex;
    margin-bottom: 10px;
}



.loading, .nomore {
  position: relative;
  bottom: 10px;
  width: 100%;
  text-align: center;
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
}

.loading {
  bottom: 40px;
}

/*隐藏滚动条*/
.container::-webkit-scrollbar {
  display: none;
}
</style>


