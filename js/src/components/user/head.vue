<script setup>

</script>

<template>
<div id="head">
    <el-avatar src="" id="avatar"></el-avatar>
    <div id="message">
      <div id="username">
        我是用户名字
      </div>


      <div id="others">
        <div id="follow">关注</div>
        <div id="fans">粉丝</div>
        <div id="like">获赞</div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
#head {
  width: 100%;
  height: 200px;
  background-color: transparent;
  display: flex;
  align-items: center;
  text-align: center;

  #avatar {
    width: 150px;
    height: 150px;
  }

  #message {
      color: white;
    background-color: transparent;
    position: relative;
    bottom: 20px;
    left: 40px;
    display: flex;
    flex-direction: column;
    height: 50px;
    width: 300px;

    #username {
      font-size: 30px;
        text-align: left;
    }

    #others {
        font-size: 20px;
      background-color: transparent;
      display: flex;
      flex-direction: row;
        justify-content: left;
        margin-top: 20px;
      div{
          width: 100px;
          text-align: left;

      }
    }

  }
}
</style>