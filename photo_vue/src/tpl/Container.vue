<template>
    <div style="width: 100%;height: 500px;">
      <setting></setting>
      <!--<map-view></map-view>-->
    </div>
</template>

<script>
  import setting from './../tpl/Setting.vue'
   import MapView from './../tpl/MapView.vue'
    export default {
        name: "container",
      components:{
        setting,
        MapView:MapView
      }
    }
</script>

<style scoped>

</style>
