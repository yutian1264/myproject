<template>
    <div style="border: black 1px solid;display:block;padding: 5px;">
      <div>
        <input type="checkbox" @click="itemClick"
               id="checkedBox" style="width: 20px; height: 20px;"
               :value="id"
               v-model="mybox" name="mybox">
        <img :src="imgPath"  :value="id"   width="180" height="250" @click="itemClick"/>
      </div>
      <div style="display: grid;margin: 5px">
        <span>name:</span> <input type="text" :value="valname" id="name"><br>
        <span>desc:</span> <input type="text" id="desc">
      </div>

    </div>
</template>

<script>
    export default {
      props:['imgPath','valname','id'],
        name: "item-image-comp",
        data(){
          return{
            mybox:[]
          }
        },
      methods:{
        itemClick(){
          var flg=false;
          if( this.mybox.length==0){
            this.mybox.push(this.id)
            flg=true
          }else{
            flg=false;
            this.mybox=[];
          }
          this.$emit('changedItem',this.id,flg);
        }
      }
    }
</script>

<style scoped>

</style>
