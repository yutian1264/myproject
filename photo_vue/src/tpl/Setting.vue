<template>
  <div>
    <button @click="getlist('up')">上一页</button>
    <button @click="getlist('down')">下一页</button>
    <button @click="opendialog">dialog</button>

    条件：<select v-model="restype1" @change="changeType">
    <option v-for="item in ownerList" :value="item.name">{{item.name}}</option>
  </select>
    name:
    所有者：<select v-model="ownerSel">
      <option v-for="item in ownerList" :value="item.name">{{item.name}}</option>
    </select>
    分类已有：<select v-model="restype2" >
    <option v-for="item in typeList1" :value="item.name">{{item.name}}</option>
    </select>
    分类：<input type="text" v-model="res_type"/>
    <button @click="save">保存</button>
    <div style="height: 600px; overflow-y: auto">
      <ul id="list">
        <li v-for="ele in imgList">
          <item :choose="choose" :imgPath="ele.thumb" @changedItem="changedItem" :valname="ele.path" :id="ele.id"></item>
        </li>
      </ul>
    </div>
    <!--<DialogModal class="dialg" v-show="ismodalvisible" @close="closeModal" @save="save"></DialogModal>-->
    <dialog3 :is-show="ismodalvisible" @on-close="closeModal" @save="save"></dialog3>
  </div>
</template>

<script>
  // var url = "http://192.168.110.143:8500"
  var url = "http://192.168.2.100:8500"
  import item from "./../tpl/ItemImageComp.vue"
  import DialogModal  from "./../tpl/components/DialogModal.vue"
  import dialog3  from "./../tpl/components/dialog3.vue"

  export default {
    name: "setting",
    components: {
      item,
      DialogModal,
      dialog3
    },
    data() {
      return {
        imgList: [],
        typeList:[],
        typeList1:[],
        ownerList:[],
        changedList:[],
        ownerSel:"",
        restype1:"",
        res_type:"",
        restype2:"",
        ismodalvisible:false,
        choose:false
      }
    },
    mounted(){
      console.log("页面加载完成")
      var that = this
      test(this.restype1,function (res) {
        var r=res.data
        that.imgList= collateData(r)
      })
      getResList(function(r){
          that.typeList1=r
      })
      getOwnweList(function(r){
        that.ownerList=r
      })
    },
    methods: {
      closeModal(){
        this.ismodalvisible=false
      },
      opendialog(){
        this.ismodalvisible=true
      },
      save(){
        alert("save")
      },
      changeType(){
        var that = this
        test(this.restype1,function (res) {
          var r=res.data
          that.imgList= collateData(r)
        })
      },
      changedItem(id,flg){
        if(flg){
          this.changedList.push(id);
        }else{
          this.changedList.remove(id)
        }
      },
      save(){
        if(StringsUtils.isNull(this.ownerSel)){
          alert("请选择归属！")
          return
        }
        if(this.changedList.length==0){
           alert("请选择资源！")
            //this.$message("请选择资源！");
           return
        }
        var res_type="other"
        if(StringsUtils.isNull(this.res_type)||StringsUtils.isNotNull( this.restype2)){
          res_type=this.restype2;
        }else{
          res_type =StringsUtils.isNull(this.res_type)?"other":this.res_type
        }
        var param={owner:this.ownerSel,type:res_type,ids:this.changedList}
        var that = this
        UpdateRess(param,function(r1){
          that.imgList=[];
          that.changedList=[]
          test(that.restype1,function (res) {
            var r=res.data
            that.imgList= collateData(r)
          })
        })

      },
      getlist(type) {
        var that = this
        this.imgList=[];
        test1(type,this.restype1,function (res) {
          var r=res.data
          that.imgList= collateData(r)
        })
      }
    }

  }
  //整理数据
  function collateData(r){
    if (r.length > 0) {
      $.each(r, function (i, e) {
        e.path = e.path.replace("J:/", url + "/j/")
        e.path = e.path.replace("H:/", url + "/h/")
        e.path = e.path.replace("D:/", url + "/d/")
        e.path = e.path.replace("C:/", url + "/c/")
        e.path = e.path.replace("E:/", url + "/e/")
        e.path = e.path.replace("G:/", url + "/g/")
        e.thumb = e.thumb.replace("J:/", url + "/j/")
        e.thumb = e.thumb.replace("H:/", url + "/h/")
        e.thumb = e.thumb.replace("D:/", url + "/d/")
        e.thumb = e.thumb.replace("C:/", url + "/c/")
        e.thumb = e.thumb.replace("E:/", url + "/e/")
        e.thumb = e.thumb.replace("G:/", url + "/g/")
      })
      return r
    }
  }
  var pageSize=0;
  var pageCount=100;
  function test(retype,d) {
    retype=StringsUtils.isNull(retype)?'':retype
    var url1=url + "/f/d/getListByPage?type="+retype+"&pageSize="+pageSize+"&pageCount="+pageCount
    console.log(url1)
    https.get(url1, function (r) {
      d(r)
    })
  }
  function test1(type,owner,d) {
    if(type=="up"){
      if(pageSize>0){
        pageSize--
      }else{
        pageSize=0
      }
    }else{
      pageSize++
    }

    https.get(url + "/f/d/getListByPage?type="+owner+"&pageSize="+pageSize+"&pageCount="+pageCount, function (r) {
      d(r)
    })
  }
function getResList(d) {
    https.get(url + "/f/d/getResTypeList", function (r) {
      d(r)
    })
  }
  function getOwnweList(d) {
    https.get(url + "/f/d/getResOwnerList", function (r) {
      d(r)
    })
  }
  function UpdateRess(param,d) {
    https.postGo(url + "/f/d/UpdateRess", param,function (r) {
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
