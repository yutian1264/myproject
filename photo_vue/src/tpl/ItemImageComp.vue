<template>
    <div style="border: black 1px solid;display:block;padding: 5px;">
      <div>
        <input type="checkbox" style="width: 20px; height: 20px;">
        <img :src="imgPath"  width="180" height="250"/>
      </div>
      <div style="display: grid;margin: 5px">
        <span>name:</span> <input type="text" :value="valname" id="name"><br>
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
