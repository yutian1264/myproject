<template>
  <div style="display: block">
    <div class="head_titler">
      <ul style="display: flex">
        <li>
          <h1 class="mar-l-15 font-thin h3">资源管理</h1>
        </li>
        <li>
          <a class="btn m-b-xs btn-sm btn-default btn-addon mar-r-15"
             id="createSceneBtn" @click="createScene()">
            <i class="fa fa-plus"></i>创建场景
          </a>
          <!--<button class="btn btn-primary btn-sm pull-right" >创建场景</button>-->
        </li>
      </ul>
    </div>
    <div>
      <div class="mar-15" style="float: left">
        <select class="form-control " id="select_city">

        </select>
      </div>
    </div>
    <div style="height: 680px;width:100%;padding:0 15px">
      <div id="allmap" ref="allmap">
      </div>
    </div>
    <!--<dialog :is-show="isShowPublish" :top-distance="topNum"   @on-close="closeDialog">-->
    <!--</dialog>-->
    <!--<dialog v-show="ismodalvisible" @close="closeModal"></dialog>-->
  </div>

</template>

<script>
  // import dialog from "./../tpl/components/dialog_comp.vue"
  // import dialog from "./../tpl/components/DialogModal.vue"
  export default {
    name: "map-view",
    componenets:{
      // dialog,
      // dialog
    }
    , data() {
      return {
        lon: "116.16280752777779",
        lat: "39.83879322222222",
        isShowPublish:false,
        topNum:50,
        ismodalvisible:false
      }
    },
    methods: {
      map() {
        let map = new BMap.Map(this.$refs.allmap); // 创建Map实例
        map.centerAndZoom(new BMap.Point(116.404, 39.915), 11);// 初始化地图,设置中心点坐标和地图级别
        map.addControl(new BMap.MapTypeControl({//添加地图类型控件
          mapTypes: [
            BMAP_NORMAL_MAP,
            BMAP_HYBRID_MAP
          ]
        }));
        map.setCurrentCity("北京");// 设置地图显示的城市 此项是必须设置的
        map.enableScrollWheelZoom(true); //开启鼠标滚轮缩放
      },
      closeDialog(){
         this.isShowPublish=false;
        //把绑定的弹窗数组 设为false即可关闭弹窗
      },
      createScene(){
        //this.isShowPublish=true;
        this.ismodalvisible=true;
      },
      closeModal(){
        this.ismodalvisible=false;
      }
    },
    mounted() {
      this.map()
    }
  }

</script>

<style scoped>

  #allmap {
    height: 700px;
    width: 100%;
    overflow: hidden;
  }

  .head_titler ul {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 0px;
    margin-bottom: 10px;
  }

  .mar-l-15 {
    margin-left: 15px;
  }

  .mar-r-15 {
    margin-right: 15px;
  }

  .mar-t-15 {
    margin-top: 15px;
  }

  .li_city {
    width: 100px;
    display: inline-block;
  }
</style>
