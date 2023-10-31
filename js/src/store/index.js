//导入pinia并使用
import  {createPinia} from 'pinia'
//导入pinia插件
import persist from 'pinia-plugin-persistedstate'
const pinia = createPinia()
pinia.use(persist)

export default pinia

