<template>
    <div style="width: 100%;height: 500px;">
      <setting v-if="tab=='setting'"></setting>

      <map-view v-else-if="tab=='map'"></map-view>
    </div>
</template>

<script>
  import setting from './../tpl/Setting.vue'
   import MapView from './../tpl/MapView.vue'
    export default {
      props: ["tab"],
        name: "container",
      data(){
       return{
         setting:false,
         map:true
       }

      },
      components:{
        setting,
        MapView:MapView
      }
    }
</script>

<style scoped>

</style>
