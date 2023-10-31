<script setup>

</script>

<template>
   <div id="box">
    <img src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg">
     <div id="videoname">我是视频名字</div>

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