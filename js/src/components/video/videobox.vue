<script setup>
const props=defineProps({
    videoShortMessage:{
        type:Object,
        default:{
            id: Number,
            information:String,
            cover_url:String,
        }
    }
})


const playVideo=()=>{
    router.push({path:'/video',query:{id:props.videoShortMessage.id}})
}

</script>

<template>
   <div id="box" @click="playVideo" >
       <img src="{{props.cover_url}}" alt="">

     <div id="videoname">{{props.information}}</div>
   </div>
</template>
<style lang="scss" scoped>
       #box{
           width: 20%;
           height: 200px;
           margin-left: 10px;
           margin-top: 2px;
           cursor: pointer;
           border-radius: 15px;

           img{
               width: 100%;
               height: 70%;
               border-radius: 15px;
           }
            #videoname{
                width: 100%;
                height: 30%;
                margin-top: -6px;
                color: white;
                font-size: 17px;
            }
       }
       #box:hover{
           transform: scale(1.02);
           background-color: rgba(255,192,203,0.2);
       }



</style>