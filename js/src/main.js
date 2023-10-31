import './assets/main.css'
import pinia from './store'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {ElementPlus} from "@element-plus/icons-vue";

//导入Vue3VedioPlayer
import vue3videoPlay from 'vue3-video-play' // 引入组件
import 'vue3-video-play/dist/style.css'


const app=createApp(App)
    app.use(ElementPlus)
    app.use(router)
    app.use(pinia)
    app.use(vue3videoPlay)
    app.mount('#app')
