<script setup>
import { ref, computed } from 'vue';
import  vidiebox from "@/components/video/videobox.vue";
import Videobox from "@/components/video/videobox.vue";
const numbers = ref([1, 2, 3, 4, 5]);
const loading = ref(false);
const noMore = ref(false);

const canLoadMore = computed(() => !loading.value && !noMore.value && numbers.value.length < 40);

const loadMore = () => {
  if (canLoadMore.value) {
    // 模拟加载更多数据
    const lastNumber = numbers.value[numbers.value.length - 1];
    for (let i = 1; i <= 3; i++) {
      numbers.value.push(lastNumber + i);
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
    <ul>
      <li v-for="index in 8"><videobox v-for="index in 5"></videobox></li>
    </ul>
    <div class="loading" v-if="loading">加载中...</div>
    <div class="nomore" v-if="noMore">暂无更多</div>
  </div>
</template>

<style lang="scss" scoped>
.container {
  height: calc(100% - 270px);
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
  bottom: 40px;
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

