<template>
    <div>
      <button @click="getlist">get list</button>
      <ul style="display: grid">
        <li v-for="ele in imgList"> <item :imgPath="ele.path" :valname="ele.name"></item></li>

      </ul>

    </div>
</template>

<script>
  import item from "./../tpl/ItemImageComp.vue"
    export default {
        name: "setting",
        components:{
          item
        },
        data(){
          return{
            imgList:[],
          }
        },
        methods:{
          getlist(){
            var that=this
            test(function(r){
              if(r.length>0){
                that.imgList=r
              }
            })
          }
        }

    }

  function test(d) {
    https.get("http://192.168.110.143:8500/f/d/getAllList?type=",function(r){
      d(r)
    })
  }

</script>

<style scoped>
  ul{
    display: flex;
  }
  li{
    display:block;
    font-size:14px;

  }
</style>
