import axios from 'axios'
import { useUserStore } from "@/store/User";
import { ElMessage } from "element-plus";


const loginorregisterinstance = axios.create({
   baseURL: "http://42.194.236.154:8889/",
    timeout: 10000,
    headers: {
     'Content-Type': 'application/json',
   'Accept': '*/*',
 },
})

 //请求拦截器
loginorregisterinstance.interceptors.request.use(
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

export default loginorregisterinstance