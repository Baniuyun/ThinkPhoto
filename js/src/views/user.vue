<script setup>

import Head from "@/components/user/head.vue";
import Menu from "@/components/user/userindexmenu.vue";
import Worklist from "@/components/user/worklist.vue";
import {getpublishListService,getLikeVideoListService,getCollectVideoListService} from "@/api/video";
const getMoreVideoFunction=()=>{
    return [
     {video_id: 1,information: "1",cover_url:'https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg'},
        {video_id: 1,information: "1",cover_url:'https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg'},
            {video_id: 1,information: "1",cover_url:'https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg'},
            {video_id: 1,information: "1",cover_url:'https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg'},
            {video_id: 1,information: "1",cover_url:'https://i.loli.net/2021/04/14/nNly8EdXJ2aHYTe.jpg'},
            {video_id: 1,information: "1",cover_url:'http://b.hiphotos.baidu.com/image/pic/item/9d82d158ccbf6c81b94575cfb93eb13533fa40a2.jpg'},
            {video_id: 1,information: "1",cover_url:'http://e.hiphotos.baidu.com/image/pic/item/4bed2e738bd4b31c1badd5a685d6277f9e2ff81e.jpg'},
        ]
}


const  getpublishVideos=async ()=>{
    await getpublishListService(user).then(res=>{
        console.log(res)
    })
}
const getLikeVideo=async ()=>{
    await getLikeVideoListService().then(res=>{
        console.log(res)
    })
}

const getCollectVideo=async ()=>{
    await  getCollectVideoListService().then(res=>{
        console.log(res)
    })
}


</script>

<template>
    <div>
        <Head></Head>
    <Menu></Menu>
    <Worklist :getMoreVideoFunction="getpublishVideos"></Worklist>
        <Worklist :getMoreVideoFunction="getpublishListService"></Worklist>
        <Worklist :getMoreVideoFunction=""></Worklist>
    </div>

</template>

<style lang="scss" scoped>
  div{
      overflow-y: hidden;
  }
  div::-webkit-scrollbar{
      display: none;
  }
</style>