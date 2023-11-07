<script setup>
import VideoCarousel from "@/components/player/VideoCarousel.vue";
import {reactive, ref} from "vue";
import {onMounted} from "vue";
import router from "@/router";
import {getVideoInfoService} from "@/api/video";


const videoMessage=ref({
       id:1,
            play_url:"https://stream7.iqilu.com/10339/upload_transcode/202002/18/20200218093206z8V1JuPlpe.mp4",
            cover_url:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",
            title:"视频2",
            favorite_count:0,
            comment_count:0,
            is_favorite:false,
            information:"视频2的信息",
            tags:["标签1","标签2"],
            user_id:1,
            is_follow:false,
            user_name:"用户1",
            avatar:"https://img.iqilu.com/data/attachment/forum/202002/18/111415j0q0q0q0q0q0q0q0.jpg",}
)

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
const getVideoInfo=async ()=>{
    const res=await getVideoInfoService(videoMessage.id)
    videoMessage.value=res.data
}

//在渲染完成后console你好
onMounted(() => {
    //获取 router中的参数
    const videoId=router.currentRoute.value.query.playVideoId
    if(videoId)
    {
        console.log("id为"+router.currentRoute.value.query.playVideoId)
        videos.unshift(getVideoInfo(videoId))
    }
    else {
    }

})




</script>
<template>
<VideoCarousel  :videos="videos"></VideoCarousel>
</template>

<style>



</style>