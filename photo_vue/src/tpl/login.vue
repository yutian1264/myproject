<template>
  <div>
    name:<input>
    pwd:<input>
    <button @click="login">登陆</button>
  </div>

</template>

<script>
    export default {
        name: "login",
      methods:{
        login(){
          console.log("login")
        }
      }
    }
</script>

<style scoped>

</style>
