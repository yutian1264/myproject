<template>
  <div style="display: block">
    <div class="head_titler" style="float: right">
          <select class="form-control" style="display: inline-flex" id="select_city">
          </select>
          <a class="btn m-b-xs btn-sm btn-default btn-addon mar-r-15"
             id="createSceneBtn" @click="createScene('btn')">
            <i class="fa fa-plus"></i>创建场景
          </a>
    </div>

    <div style="height: 680px;width:100%;padding:0 15px">
      <div id="allmap" ref="allmap">
      </div>
    </div>
    <dialog3 v-if="ismodalvisible" :contentCss="dialogContent"  @on-close="closeModal">
      <span slot="title">新建影集</span>
      <div style="height: 570px;padding: 10px; margin-top: 20px; overflow-y: auto">
        <div>
          <p><span>当前选中地址为：</span><span>{{currentAddr}}</span></p>
          <p><span>{{lng_span}}</span><span>{{lat_span}}</span></p>
          <p><span>目录：</span>
            <select v-model="dirSel" @change="changeDirectory">
            <option v-for="item in directoryList" :value="item.name">{{item.name}}</option>
          </select>
            新增目录：<input v-model="directory_Type"><button @click="addDirectory">添加</button></p>
          <p><span>类型：</span><input v-model="fileType"></p>
          <!-- v-show="dirSel!=''" -->
          <upload-comp :dir="dirSel" ref="upcomp" @on-completed="uploadCompleted" @begin-upload="beginUpload"></upload-comp>
        </div>
        <button @click="cancel_upload" >取消</button>
        <button @click="save_upload" :disabled="is_disabled">保存</button>
      </div>
    </dialog3>
  </div>

</template>

<script>
  // import dialog from "./../tpl/components/dialog_comp.vue"
  // import dialog from "./../tpl/components/DialogModal.vue"
  import Dialog3 from "./components/dialog3";
  import UploadComp from "./components/upload-comp";

  export default {
    components: {
      UploadComp,
      Dialog3},
    name: "map-view"
    , data() {
      return {
        lon: "116.16280752777779",
        lat: "39.83879322222222",
        lng_span:0,
        lat_span:0,
        isShowPublish:false,
        topNum:50,
        ismodalvisible:false,
        currentAddr:"",
        upload_view:"",
        fileType:"",
        directory_Type:"",
        directoryList:[],//目录集合
        dirSel:"",
        dialogContent:"width:650px",
        is_disabled:true,
        totalFiles:0,
        fileMess:[]  //上传文件详细信息
      }
    },
    methods: {
      //创建地图
      createMap() {
        cmap(this.$refs.allmap,this.createScene)
      },
      //创建影集
      createScene(a,b,c){
        this.ismodalvisible=true;
        this.lng_span=a
        this.lat_span=b
        if(DataObject.isNotNull(c)){
          this.currentAddr=c.str;
        }
        this.getDirectoryList();
      },
      closeModal(){
        this.ismodalvisible=false;
      },
      changeDirectory(){
        this.dispatchChild()
      },
      dispatchChild(){
        this.$refs.upcomp.childFn(this.dirSel)
      },
      //文件上传OK 但没有保存到数据库中，需要点击保存才能添加记录
      uploadCompleted(status,res){
        // console.log("map====:"+status)
        this.is_disabled=false
        this.fileMess=res
      },
      beginUpload(res){
        this.totalFiles=res;
      },
      cancel_upload(){

      },//保存上传记录
      save_upload(){
        if(StringsUtils.isNull(this.dirSel)){
          alert("请选择目录！")
          return
        }
        var param={lat:this.lat_span,lng:this.lng_span,
          fileMsg:this.fileMess.success,user:"wyt",type: this.fileType,position:this.currentAddr,category:this.dirSel}
          https.postGo(url+"/f/d/addResource",param,function(r){

          })

      },
      getDirectoryList(){
        var that=this;
        getDirectoryList("wyt",function(r) {
          that.directoryList=[];
          that.directoryList=r
        })
      },
      //添加目录
      addDirectory(){
        var that=this;
        if(StringsUtils.isNull(this.directory_Type)){
          alert("目录名称不能为空！")
          return
        }
       var  param={name:this.directory_Type,user:"wyt"}
        https.postGo(url+"/f/d/addDirectory",param,function(r){
        //  var obj=json.parse(r)
          // if(obj.Error==1){
          console.log("新增目录完成")
          that.getDirectoryList();
          that.dirSel=that.directory_Type
          that.directory_Type=""
          // }
        })
      }
    },
    mounted() {
      this.createMap()
      createMarker()
      this.getDirectoryList();
    }
  }
  var map;
  var points=[]
  var markers=[]
  var sceneList=[]
  var lon="116.16280752777779"
  var lat="39.83879322222222"
  var opts = {
    width : 500,     // 信息窗口宽度
    title : "场景信息" , // 信息窗口标题
    enableMessage:true//设置允许信息窗发送短息
  };
  function cmap(document,openDialog) {
     map = new BMap.Map(document); // 创建Map实例
    map.centerAndZoom(new BMap.Point(116.404, 39.915), 11);// 初始化地图,设置中心点坐标和地图级别
    map.addControl(new BMap.MapTypeControl({//添加地图类型控件
      mapTypes: [
        BMAP_NORMAL_MAP,
        BMAP_HYBRID_MAP
      ]
    }));
    map.setCurrentCity("北京");// 设置地图显示的城市 此项是必须设置的
    map.enableScrollWheelZoom(true); //开启鼠标滚轮缩放
    var lng_temp=0
    var lat_temp=0
    var detailAddr={};
    var geoc = new BMap.Geocoder();
    map.addEventListener("rightclick",function(e){
       lng_temp=e.point.lng
       lat_temp=e.point.lat
      var pt = e.point;
      geoc.getLocation(pt, function(rs){
        var addComp = rs.addressComponents;
        detailAddr={
          city:addComp.city,
          province:addComp.province,
          district:addComp.district,
          street:addComp.street,
          streetNumber:addComp.streetNumber,
          str:rs.address
        }
       // alert(addComp.province + ", " + addComp.city + ", " + addComp.district + ", " + addComp.street + ", " + addComp.streetNumber);
      });
    })
    var menu = new BMap.ContextMenu();
    var txtMenuItem = [
      {
        text:'新建影集',
        callback:function(){
          openDialog(lng_temp,lat_temp,detailAddr)
        }
      }
    ];
    for(var i=0; i < txtMenuItem.length; i++){
      menu.addItem(new BMap.MenuItem(txtMenuItem[i].text,txtMenuItem[i].callback,100));
    }
    map.addContextMenu(menu);

  }
  // function openDialog(){
  //   console.log("open dialog")
  // }
  function createMarker() {
    commUtil.test("丽雅",function(r){
      sceneList=commUtil.collateData(r.data);
      getPointList(sceneList)
      addMarker(points)
    })
  }
  function getDirectoryList(user,dirListFunc){
    user="wyt"
    https.get(url+"/f/d/getDirectoryListByUser?user="+user,function(r){
      console.log(r)
      dirListFunc(r)
    })
  }


  function getPointList(sceneList){
    points=[];
    sceneList.forEach(function(scene){
      if(scene.lng==0){
        scene.lng=lon
      }
      if(scene.lat==0){
        scene.lat=lat
      }
      points.push(new BMap.Point(scene.lng,scene.lat));
    })
  }
   function addMarker(points){
    markers=[];
    points.forEach(function(point,index){
      var marker = new BMap.Marker(point);
      markers.push(marker);
      map.addOverlay(marker);
      addClickHandler(marker);
    })
  }
  function addClickHandler(marker){
    marker.addEventListener("click",function(e){
      var content=getContent(e.currentTarget);
      openInfo(content,e)

    });
  }
  function getContent(tMakers){
    var indexArr=[];
    var iconBounds= getXYWH(tMakers);
    var bool=false;
    for(var i=0;i< markers.length;i++){
      var e=markers[i];
      var iconBounds1=getXYWH(e)
      if (isOverlap(iconBounds,iconBounds1)){
        bool= true;
      } else {
        bool= false;
      }
      if(bool){//||(e.point.lng == tMakers.point.lng&&e.point.lat == tMakers.point.lat)
        indexArr.push(i);
      }
    }

    var content='<ul id="edit" style="overflow-y: scroll;max-height:500px">';
    var isCore="";
    for(var i=0;i<indexArr.length;i++) {
      var index = indexArr[i];
      var path = sceneList[index].thumb;
      if (sceneList[index].core_scene) {
        isCore = "是"
      } else {
        isCore = "否"
      }
      content = content +
        "<li class='li_info li_border'>" +
        "<div class='sc_left'><img src='" + path + "' class='imgDemo' alt='场景图片' width='96px' height='128px'> </div>" +
        "<div class='sc_left pad-l-15'>" +
        // "<p>地点：" +
        // "<span>" + $scope.sceneList[index].position + "</span></p>" +
        "<p>场景名称：</p>" +
        "<input  value='" + sceneList[index].name + "' class='scene_desc' disabled>" +
        "<p>场景描述：</p>" +
        "<textarea id='desc'  class='scene_desc' disabled>" + sceneList[index].desc + "</textarea>" +
        "<p><button class='btn btn-primary btn-xs edit' id='" + sceneList[index].id + "'  >编辑</button>" +
        "<button class='btn btn-primary btn-xs edit' param='" + sceneList[index].name + " name='cancel' style='margin-left: 50px;display: none'>取消</button>" +
        "<button class='btn btn-primary btn-xs edit' param='" + sceneList[index].id + "' name='delete' style='margin-left: 50px'>删除</button>"+
        "<button class='btn btn-primary btn-xs edit'  name='add' style='margin-left: 50px'>添加</button>"+
        "</p>" +
        "</div>" +
        "</li>";
    }

    content=content+"</ul>"+
      "</div>";
    return content;
  }
  function openInfo(content,e){
    var p = e.target;
    var lon=p.getPosition().lng;
    var lat=p.getPosition().lat;
    var point = new BMap.Point(p.getPosition().lng, p.getPosition().lat);
    var id="";
    sceneList.forEach(function (e1){
      if(e1.Longitude == lon&&e1.Latitude == lat){
        id=e1.id;
      }
    })
    var infoWindow = new BMap.InfoWindow(content,opts);  // 创建信息窗口对象ssss

    infoWindow.addEventListener("open",ta(id,infoWindow));
    infoWindow.addEventListener("close",function tt(){
      infoWindow.removeEventListener("open",ta)
      infoWindow.removeEventListener("close",tt)
    })
    map.openInfoWindow(infoWindow,point);
  }
  function isOverlap(rc1,rc2)
  {
    if (rc1.x + rc1.w  > rc2.x &&
      rc2.x + rc2.w  > rc1.x &&
      rc1.y + rc1.h > rc2.y &&
      rc2.y + rc2.h > rc1.y
    )
      return true;
    else
      return false;
  }

  function getXYWH(marker){

    var result={};
    var projection = map.getMapType().getProjection();
    var lngLat = marker.point;
    // var lngLatStr = "经纬度：" + lngLat.lng + ", " + lngLat.lat;
    var worldCoordinate = projection.lngLatToPoint(lngLat);
    //  var worldCoordStr = "<br />平面坐标：" + worldCoordinate.x + ", " + worldCoordinate.y;
    var iconAnchorPointOffset= marker.getIcon().size;
    // var overlayCoordStr= "<br />宽高：" + iconAnchorPointOffset.width + ", " + iconAnchorPointOffset.height;
    var pixelCoordinate = new BMap.Pixel(Math.floor(worldCoordinate.x * Math.pow(2,  map.getZoom() - 18)),
      Math.floor(worldCoordinate.y * Math.pow(2,  map.getZoom() - 18)));
    // var pixelCoordStr = "<br />像素坐标：" + pixelCoordinate.x + ", " + pixelCoordinate.y;
    // console.log(lngLatStr + worldCoordStr + pixelCoordStr);
    result["x"]= pixelCoordinate.x
    result["y"]= pixelCoordinate.y
    result["w"]= iconAnchorPointOffset.width
    result["h"]= iconAnchorPointOffset.height
    return result;
  }
  function ta(id,infoWindow){
    setTimeout(function(){
      infoWindow.redraw();
      // $('#edit').on('click',"button", function(e) {
      //   var arr = $(e.target).parent().parent().find("input");
      //   arr.attr("disabled", "disabled");
      //   var txt = $(e.target).parent().parent().find("textarea");
      //   txt.attr("disabled", "disabled");
      //   var p = $(e.target).parent().parent().find(".isCore");
      //
      //   if ($(e.target).attr("name") == "cancel") {
      //     var param = $(e.target).attr("param");
      //     $($(e.target).parent().find("button")[2]).css("display","inline-block");
      //     var a = $(this).parent().parent().find(".scene_desc_focus");
      //     if (StringsUtils.isNotNull(param)) {
      //       var arrList = param.split(",");
      //       var name = arrList[0];
      //       var isCore = arrList[1];
      //       var desc = arrList[2];
      //       $(a[0]).val(name);
      //       $(a[1]).val(desc);
      //
      //       var pText = "是否为核心场景：" + "<span class='core_span'>" + isCore + "</span>";
      //       p[0].innerHTML = pText;
      //     }
      //     a.blur();
      //     a.removeClass("scene_desc_focus");
      //     var btnArr = $(e.target).parent().find("button");
      //     $(btnArr[1]).css("display", "none")
      //     btnArr[0].innerHTML = "编辑";
      //   }else if($(e.target).attr("name") == "delete"){
      //     var id= $(e.target).attr("param");
      //     var btn=$(this)
      //     layer.confirm("删除后无法恢复,确定要删除?",{btn:["是","否"]},function() {
      //       https.get(baseUrl + "scene/deleteSceneById.do?id=" + id, function (r) {
      //         if (r == "true" || r) {
      //           var ulList=$(btn.parent().parent().parent().parent()).children()
      //           if(ulList.length>1){
      //             $(btn.parent().parent().parent()).remove()
      //             for(var i=$scope.sceneList.length-1;i>=0;i--) {
      //               if($scope.sceneList[i].id==id){
      //                 $scope.sceneList.splice(i,1);
      //               }
      //             }
      //             $scope.getPointList($scope.sceneList);
      //             $scope.addMarker($scope.points);
      //             $scope.addLabel($scope.markers);
      //           }else{
      //             $scope.map.clearOverlays();
      //             $scope.getScenesList();
      //           }
      //           layer.msg("删除成功", {icon: 0, time: 1000});
      //         } else {
      //           layer.msg("删除失败", {icon: 0, time: 1000});
      //         }
      //       })
      //     })
      //   }else {
      //     if (e.target.innerHTML == "保存") {
      //       $($(e.target).parent().find("button")[1]).css("display", "none");
      //       $($(e.target).parent().find("button")[2]).css("display", "inline-block");
      //       var core_scene = $(e.target).parent().parent().find("input[type=radio]:checked");
      //       var pText = "是否为核心场景：";
      //       if (core_scene[0].value == "0") {
      //         pText = pText + "<span class='core_span'>是</span>";
      //       } else {
      //         pText = pText + "<span class='core_span'>否</span>";
      //       }
      //       p[0].innerHTML = pText;
      //       var a = $(this).parent().parent().find(".scene_desc_focus")
      //       id = $(e.target).attr("id");
      //       a.blur();
      //       a.removeClass("scene_desc_focus");
      //
      //       $http.get(baseUrl + "scene/updateScene.do?id=" + id + "&isCore=" + core_scene[0].value + "&desc=" + $(txt).val() + "&name=" + $(arr[0]).val() + "&search=" + $(arr[1]).val()).success(function (data) {
      //         if (data.status == 1) {
      //           layer.msg("保存成功", {icon: 0, time: 1000});
      //           e.target.innerHTML = "编辑";
      //           $($(e.target).parent().find("button")[2]).css("display", "inline-block");
      //           $scope.getScenesList();
      //         } else {
      //           layer.msg(data.msg, {icon: 0, time: 1000});
      //         }
      //       })
      //     }else if(e.target.innerHTML == "添加"){
      //       $scope.createScene($(this).attr("param"))
      //     }else{
      //       $($(e.target).parent().find("button")[2]).css("display","none");
      //       $($(e.target).parent().find("button")[1]).css("display","inline-block")
      //       var spano=$(e.target).parent().parent().find(".core_span");
      //
      //       var pText="是否为核心场景：";
      //       if(spano[0].innerHTML=="是"){
      //         pText=pText+"<input type='radio' name='core_scene'  value='0'  checked/>"+
      //           "<span>是</span>"+
      //           "<input type='radio' name='core_scene'  value='1'  />"+
      //           "<span>否</span>";
      //       }else{
      //         pText=pText+"<input type='radio' name='core_scene'  value='0'  />"+
      //           "<span>是</span>"+
      //           "<input type='radio' name='core_scene'  value='1'  checked/>"+
      //           "<span>否</span>";
      //       }
      //
      //       p[0].innerHTML=pText;
      //       arr.removeAttr("disabled");
      //       txt.removeAttr("disabled");
      //       e.target.innerHTML="保存";
      //       var a=$(this).parent().parent().find(".scene_desc");
      //       a.focus();
      //       a.addClass("scene_desc_focus");
      //     }
      //   }

      // })
    },100)
  }
</script>

<style scoped>

  .dialogContent{
    width: 400px;
  }

  #allmap {width: 100%;height: 100%;overflow: hidden;margin:0;font-family:"微软雅黑";}
  .mar-l-15{
    margin-left:15px;
  }
  .mar-r-15{
    margin-right:15px;
  }
  .mar-t-15{
    margin-top:15px;
  }
  .li_city{
    width: 100px;
    display: inline-block;
  }
  label{
    max-width: 1000px;
  }
  img{
    max-width:none;
  }
  .sc_left{
    float:left;

  }
  .li_info{
    float: left;
    padding: 5px 0;
    width: 450px;
  }
  .li_border{
    border-bottom: 1px solid #03A9F4;

  }
  .pad-l-15{
    padding-left:15px;
  }
  .scene_desc{
    border:1px solid #fff;
    padding:2px;
    border-radius:2px;
    background-color: #fff;
    width:330px;
  }
  .scene_desc_focus{
    border:1px solid #bbb;
  }
  textarea{
    resize:none;
    height:50px;
  }
  #select_city{
    width:150px;
  }
  .mar-15{
    margin:15px;
  }
  .anchorBL{
    display:none;
  }
  .core_span{
    margin-right:30px;
  }
  #edit li:last-child{
    border-bottom:none;
  }
</style>
