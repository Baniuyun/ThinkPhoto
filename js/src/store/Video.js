import  {defineStore} from "pinia";
import {reactive} from "vue";


/*导出一个Vedio仓库。通过该仓库的调用，可以获取当前观看的视频信息，上一条视频信息，下一条视频信息，以此类推*/
export const useVideoStore = defineStore("Video",
    ()=>{
         const video = reactive({
                id: "",
                title: "",
                play_url: "",
                cover_url: "",
                author_id: "",
                author_name: "",
                favorite_count: "",
                is_favorite:"",
                comment_count:"",
                tag:[],
                time:"",
    });
             const videoList = reactive({
             videoList:[],
             videoListIndex:0,
        videoListLength:0,
    })
        const getNextVideo = ()=>{
            if (videoList.videoListIndex<videoList.videoListLength){
                videoList.videoListIndex++
            }
            return videoList.videoList[videoList.videoListIndex]
        }

        const getPrevVideo =()=>{
            if (videoList.videoListIndex>0){
                videoList.videoListIndex--
            }
            return videoList.videoList[videoList.videoListIndex]
         }
    }

