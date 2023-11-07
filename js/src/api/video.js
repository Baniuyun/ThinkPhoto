import request from '@/utils/videorequest.js'


//获取上传回调凭证
export const getPublishTokenService=(user_id)=>
    request.get('/api/video/publishtoken',{user_id})

//获取上传视频列表
export const getpublishListService=(user_id)=>
     request.get(`api/video/publishlist`,{user_id})


//获取喜欢视频列表
export const getLikeVideoListService=(user_id)=>
    request.get(`api/video/favoritelist`,{user_id})

//获取收藏视频
export const getCollectVideoListService=(user_id)=>
    request.get(`api/video/collectvideolist`,{user_id})

//上传视频
export const publishVideoService=(key,avatar,tag,information,user_id,user_name)=>
    request.post(`api/video/publishvideo`,{key,avatar,tag,information,user_id,user_name})


//删除视频
export const deleteVideoService=(user_id,video_id)=>
    request.delete(`api/video/publishvideo`,{user_id,video_id})


//点赞视频
export const likeVideoService=(user_id,video_id,status)=>
    request.post(`api/video/like`,{user_id,video_id,status})

//获取视频列表
export const getVideoListService=(user_id,tag)=>
    request.get(`api/video/list`,{user_id,tag})




//通过videoid获取视频信息
export  const getVideoInfoService=(video_id)=>
    request.get(`api/video/info`,{video_id})




















