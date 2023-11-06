import  request from '@/utils/request'


//修改头像
export const userUploadAvatarService = (avatar) =>
    request.patch('/my/update/avatar', { avatar })



//注册并且使用json格式传递数据
export const userRegisterService = (user_name, password)=>{
    return request({
        url: '/v1/tpvideo/register',
        method: 'post',
        data: {
            'user_name': user_name,
            'password': password
        }
    })



}


//登录
export const userLoginService = (user_name, password) =>{
    return request({
        url: '/v1/tpvideo/login',
        method: 'post',
        data: {
            'user_name': user_name,
            'password': password
        }
    })
}






