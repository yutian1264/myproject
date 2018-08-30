<template>
    <div style="border: black 1px solid;display: inline-flex;padding: 5px;">
      <input type="checkbox" style="width: 20px; height: 20px;">
      <img :src="imgPath" width="200" height="250"/>
      <div style="display: grid;margin: 5px">
        <span>name:</span> <input type="text" :value="valname" id="name">
        <span>desc:</span> <input type="text" id="desc">
      </div>

    </div>
</template>

<script>
    export default {
      props:['imgPath','valname'],
        name: "item-image-comp",
        data(){
          return{
          }
        }
    }
</script>

<style scoped>

</style>
