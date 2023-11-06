import request from '@/utils/request'


//获取上传回调凭证
export const getPublishTokenService=()=>
    request.get('/api/video/publishtoken')

//获取上传视频列表
export const getpublishListService=(user_id)=>
     request.get(`api/video/publishlist`,{params:{user_id}})


//获取喜欢视频列表
export const getLikeVideoListService=(user_id)=>
    request.get(`api/video/favoritelist`,{params:{user_id}})

//获取收藏视频
export const getCollectVideoListService=(user_id)=>
    request.get(`api/video/collectvideolist`,{params:{user_id}})

//上传视频
export const publishVideoService=(key,avatar,tag,information,user_id,user_name)=>
    request.post(`api/video/publishvideo`,{key,avatar,tag,information,user_id,user_name})


//删除视频
export const deleteVideoService=(user_id,video_id)=>
    request.delete(`api/video/publishvideo`,{params:{user_id,video_id}})


//点赞视频
export const likeVideoService=(user_id,video_id,status)=>
    request.post(`api/video/like`,{user_id,video_id,status})

//获取视频列表
export const getVideoListService=(user_id,tag)=>
    request.get(`api/video/list`,{params:{user_id,tag}})




//通过videoid获取视频信息





















