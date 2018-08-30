<template>
  <div>
    <button @click="getlist('up')">上一页</button>
    <button @click="getlist('down')">下一页</button>
    <div style="height: 600px; overflow-y: auto">
      <ul>
        <li v-for="ele in imgList">
          <item :imgPath="ele.path" :valname="ele.name"></item>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  // var url = "http://192.168.110.143:8500"
  var url = "http://192.168.2.100:8500"
  import item from "./../tpl/ItemImageComp.vue"

  export default {
    name: "setting",
    components: {
      item
    },
    data() {
      return {
        imgList: [],
      }
    },
    mounted(){
      console.log("页面加载完成")
      var that = this
      test(function (res) {
        var r=res.data
        if (r.length > 0) {
          $.each(r, function (i, e) {
            e.path = e.path.replace("J://", url + "/j/")
            e.path = e.path.replace("H://", url + "/h/")
            e.path = e.path.replace("D://", url + "/d/")
            e.path = e.path.replace("C://", url + "/c/")
            e.path = e.path.replace("E://", url + "/e/")
            e.path = e.path.replace("G://", url + "/g/")
          })
          that.imgList = r
        }
      })
    },
    methods: {
      getlist(type) {
        var that = this
        this.imgList=[];
        test1(type,function (res) {
          var r=res.data
          if (r.length > 0) {
            $.each(r, function (i, e) {
              e.path = e.path.replace("J://", url + "/j/")
              e.path = e.path.replace("H://", url + "/h/")
              e.path = e.path.replace("D://", url + "/d/")
              e.path = e.path.replace("C://", url + "/c/")
              e.path = e.path.replace("E://", url + "/e/")
              e.path = e.path.replace("G://", url + "/g/")
            })
            that.imgList = r
          }
        })
      }
    }

  }

  var pageSize=0;
  var pageCount=20;
  function test(d) {
    https.get(url + "/f/d/getListByPage?type=&pageSize="+pageSize+"&pageCount="+pageCount, function (r) {
      d(r)
    })
  }
  function test1(type,d) {
    if(type=="up"){
      if(pageSize>0){
        pageSize--
      }else{
        pageSize=0
      }
    }else{
      pageSize++
    }

    https.get(url + "/f/d/getListByPage?type=&pageSize="+pageSize+"&pageCount="+pageCount, function (r) {
      d(r)
    })
  }

</script>

<style scoped>
  ul {
    position: relative;
  }

  li {
    display: block;
    font-size: 14px;
    float: left;
    margin-right: 5px;
    margin-bottom: 5px;

  }
</style>
