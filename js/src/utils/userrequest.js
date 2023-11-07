import axios from 'axios'
import { useUserStore } from "@/store/User";
import { ElMessage } from "element-plus";


const userinstance = axios.create({
   baseURL: "http://42.194.236.154:8887",
    timeout: 10000,
    headers: {
     'Content-Type': 'application/json',
   'Accept': '*/*',
 },
})

 //请求拦截器
userinstance.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = token
    }
    return config;
  },
  (err) => Promise.reject(err)
)

/* 响应拦截器 */
/*instance.interceptors.response.use(
  (res) => {
    if (res.data.code === 0) {
      return res
    }
    ElMessage({ message: res.data.message || "服务异常", type: 'error' });
    return Promise.reject(res.data);
  },
  (err) => {
    ElMessage({ message: err.response.data.message || '服务异常', type: 'error' });
    console.log(err); // 移动到错误处理代码块内
    if (err.response?.status === 401) {
      /!* 弹出登录框或处理其他逻辑 *!/
    }
    return Promise.reject(err);
  }
)*/

export default userinstance
