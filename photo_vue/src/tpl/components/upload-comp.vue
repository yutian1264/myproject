<template>
    <div>
      <div id="easyContainer"></div>
      <span style="display: none">{{dir}}</span>
    </div>
</template>

<script>
    export default {
      name: "upload-comp",
      props:["dir"],
      components:{
      },
      data(){
        return{
          fileLeng:0
        }

      },
      methods:{
        childFn(ref){
            this.directory=ref;
          upload_comp_View.setDirectory(this.directory)
        },
        uploadFinished(status,res){
          if(status==2){
            this.fileLeng=res;
            this.$emit("begin-upload",res)
          }else{
            console.log(status+"==upload dispatcherEvent=="+res.toString())
            if(res.error.length==this.fileLeng ||
              res.success.length==this.fileLeng||
              (res.success.length+res.error.length)==this.fileLeng){
              this.$emit("on-completed",status,res)
            }
          }
        }
      },
      mounted(){
        upload_comp_View.createUploadView(this.uploadFinished,this.directory)
      }

    }
    var uploadResultList=[]

    var upload_comp_View={
      setDirectory:function(dir){
        $('#easyContainer').easyUpload.setDir(dir)
      },
      createUploadView: function (callback,dir){
        $('#easyContainer').easyUpload({
          dir:dir,
          allowFileTypes: '*.jpg;*.png',//允许上传文件类型，格式';*.doc;*.pdf'
          allowFileSize: 10000000,//允许上传文件大小(KB)
          selectText: '选择文件',//选择文件按钮文案
          multi: true,//是否允许多文件上传
          //多文件上传时允许的文件数 multiNum: 5,
          showNote: true,//是否展示文件上传说明
          note: '提示：png、jpg',//文件上传说明
          showPreview: true,//是否显示文件预览
          url: url+'/f/d/addphoto',//上传文件地址
          fileName: 'file',//文件filename配置参数
          // formParam: {
          //     token: $.cookie('token_cookie')//不需要验证token时可以去掉
          // },//文件filename以外的配置参数，格式：{key1:value1,key2:value2}
          timeout: 30000,//请求超时时间
          okCode: 200,//与后端返回数据code值一致时执行成功回调，不配置默认200
          //formParam:{},
          successFunc: function(res){
            callback(1,res)
          },//上传成功回调函数
          errorFunc: function(res) {
            callback(0,res)
           // console.log('失败回调', res);
          },//上传失败回调函数
          deleteFunc: function(res) {
            console.log('删除回调', res);
          },//删除文件回调函数
          beginFunc:function(res){
            callback(2,res)
          }
        });
      }
    }
</script>

<style scoped>

</style>
