<template>
<div>
<footer>
    <div class="author">
        Official website:
        <a href="http://{{.Website}}">{{.Website}}</a> /
        Contact me:
        <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
</footer>
<div class="backdrop"></div>
<!--<script type="text/plain">
   var url= "http://"+window.location.host

    function abcPost(){
        https.post(url+"/",{"a":"222","b":"444"},function(r){
            var d=JSON.parse(r)

            console.log(r)
            console.log(d.id)
        })
    }
    function conndb(){
        https.post(url+"/server/def/conndb",{},function(r){

            console.log(r)
            if(r!=null){
                var arr=eval(r)
                console.log(arr[0].uname)
            }else{
                console.log("请登录!!!")
            }

        })
    }
   $("#file_upload").uploadify({		//绑定元素
       'fileObjName':'files',//html input标签的name属性的值吧。
       'debug':false,
       'auto':true,	//自动上传
       'buttonText':'Choose A File',
       'removeCompleted':false, //上传完成以后是否保存进度条
       'cancelImg':'/static/js/uploadify/uploadify-cancel.png',
       'swf':'/static/js/uploadify/uploadify.swf',	//必须设置  swf文件路径
       'uploader':'http://192.168.170.114:8080/bzy/g_upload/uploadzip.do',			//必须设置，上传文件触发的url
       'fileTypeDesc':'FileType',
       'fileTypeExts':'*.jpg;*.jpge;*.gif;*.png;*.zip;*.ZIP;*.rar;*.RAR',
       'overrideEvents':['onDialogClose'],
       'multi':true,
       'formData':{'url':window.location.search}	//这里我需要得到当前页面url的问号后面的值，所以就用 fromData 这个参数。
   })


</script>-->
<button onclick="abcPost()">post</button>
<button onclick="conndb()">连接数据库</button>
<form id="fform" action="http://192.168.170.114:8080/bzy/g_upload/uploadzip.do" method="POST" enctype="multipart/form-data">
    <input id="myfile" name="files" type="file" multiple="multiple" />
    <input type="submit" value="上传"  />
</form>
<!--<form id="fform" action="/server/def/upload" method="POST" enctype="multipart/form-data">
    <input id="myfile" name="myfile" type="file" multiple="multiple" />
    <input type="submit" value="上传"  />
</form>-->
</div>
</template>

export default {
  name: 'app',
  components: {
    // Hello
  }
}
</script>
