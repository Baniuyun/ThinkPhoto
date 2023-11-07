import request from '@/utils/userrequest'



//获取搜索结果
export const getSearchResultService=(keyword)=>
    request.get('/api/video/search',{params:{keyword}})
