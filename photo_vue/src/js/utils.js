//icon 1 √  2×  3? 4锁  5难过 6笑脸  7警告
var indexMask=0;
function openLoad() {
    indexMask=layer.load(1, {
        shade: [0.1,'#fff'] //0.1透明度的白色背景
    });
}
function closeLoad(time) {
    if(StringsUtils.isNull(time)){
        time=500;
    }
    setTimeout(function(){
        layer.closeAll()
    },time)
}

var loginUserName = "";
var Constants={
    //项目进度状态
    UNKNOWN:"unknown",
    PREPARE_RES:"prepare_res",
    STARTING:"starting",
    STARTED:"started",
    COMMENTS:"comments",
    WARNING:"warning",
    END:"end",
    FAILURE:"failure",
    //设计用例删除边框和颜色颜色
    DELETECOLOR:"#46ff0f",
    ADDCOLOR:"#ff1d0e",
    UPDATECOLOR:"#181bff",
    //用例管理颜色状态
    CASEGRAY:"#727070",
    CASEYELLOR:"#e4e653",
    CASEGREEN:"#30e604",
    MANAGER:"manager",
    LEADER:"leader",
    TESTER:"tester",
    REQUIREMENT:"requirement",
    EXECUTE:"execute",
    // REMOVECASESETFORCREATEPROJECTLOCALSTORAGE:"setRemoveList",//用例集中删除掉的
    // PROJECTUSER:"projectSelectUser",//选择的用户
    // PROJECTRESULT:"projectSetResult",//用例集选择结果
    PROJECTCASE_TYPE:"case_type",//当前选择用例类型/用例集还是用例
    // PROJECTCASERESULT:"caseNameCreateProgram",//用例选择结果
    PROJECTLOCALDATA:"projectLocalData",//用例选择结果
    USER_REGION:"厂商"//用户域
}

//获取通用接口数据

var common_t={
    /**
     * 获取系统设置常量
     * @param baseUrl url
     * @param 常量type 常量类型
     */
    getSetting:function(baseUrl,type,callback){
        https.get(baseUrl+"settings/getListByType.do?type="+type,function(d){
            if(DataObject.isNotNull(d)){
                callback(d);
            }else{
                Alert.show("获取setting常量失败");
                callback(null);
            }
        })
    }
}
var https={
    get:function(url,callback){
      var request= $.ajax({
            url:url,
            type:"get",
            dataType : "json",
            contentType : "application/json; charset=utf-8",
            success:function(d){
                callback(d);
            },
            error:function(e){

                callback(null);
                console.log("获取远程数据异常")
            }
        })
        return request;
    },
    getAsync:function(url,async,callback){
        $.ajax({
            url:url,
            type:"get",
            dataType : "json",
            async:async,
            contentType : "application/json; charset=utf-8",
            success:function(d){
                callback(d);
            },
            error:function(e){

                callback(null);
                console.log("获取远程数据异常")
            }
        })

    },
    post:function(url,param,callback){
       var request= $.ajax({
            url:url,
            type:"post",
            data:JSON.stringify(param),
            dataType : "json",
            contentType : "application/json; charset=utf-8",
            success:function(d){
                callback(d);
            },
            error:function(e){

                callback(null);
                console.log("获取远程数据异常")
            }
        })
        return request;
    },
    postAsync:function(url,param,async,callback){
        $.ajax({
            url:url,
            type:"post",
            data:JSON.stringify(param),
            dataType : "json",
            async:async,
            contentType : "application/json; charset=utf-8",
            success:function(d){
                callback(d);
            },
            error:function(e){

                callback(null);
                console.log("获取远程数据异常")
            }
        })
    },
    postGo:function(url,param,callback) {
        var request = $.ajax({
            url: url,
            type: "post",
            data: {"param": JSON.stringify(param)},
            dataType: "text",
            success: function (d) {
                callback(d);
            },
            error: function (e) {

                callback(null);
                console.log("获取远程数据异常")
            }
        })
        return request;
    },
    updateItemCaseChildren: function (url,data){

    $.ajax({
        type:"post",
        url:url,
        dataType : "json",
        contentType : "application/json; charset=utf-8",
        async:false,
        data:JSON.stringify(data),
        success:function(datas){
            console.log("更新表中字段")
        },
        error:function(r){
            console.log("get composite case error")
        }
    });
}

}

var time={
    currentTime:function() {
        var date = new Date();
        var seperator1 = "-";
        var seperator2 = ":";
        var month = date.getMonth() + 1;
        var strDate = date.getDate();
        if (month >= 1 && month <= 9) {
            month = "0" + month;
        }
        if (strDate >= 0 && strDate <= 9) {
            strDate = "0" + strDate;
        }
        var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
            + " " + date.getHours() + seperator2 + date.getMinutes()
            + seperator2 + date.getSeconds();
        return currentdate;
    },
    long2date:function(timeString){
        var d = new Date(Number(timeString)) ;
        return d;
    },
    //20171207111310
    stringFormat:function(s){
        if(StringsUtils.isNull(s)){
            return ""
        }
        return s.substr(0,4)+"-"+s.substr(4,2)+"-"+s.substr(6,2)+" "+s.substr(8,2)+":"+s.substr(10,2)+":"+s.substr(12,2)
    },
    string2Long:function(s){
        var date = new Date(s);
        return date.getTime()
    }

}

var StringsUtils={
    isNull:function(s) {
      s=$.trim(s)
        if(s==undefined || s=="undefined" || s=="" || s==null || s=="null") {
            return true;
        }
        return false;
    },
    isNotNull:function(s) {
        s=$.trim(s)
        if(s!=undefined && s!="undefined" && s!="" && s!=null && s!="null") {
            return true;
        }
        return false;
    }
}

//localStorage
var ls={
    setStorage:function(key ,value){
        localStorage.setItem(key,value);
    },
    getStorageByKey:function(key){
        with(document) {
          return localStorage.getItem(key);
        }
    }
}
Array.prototype.deleteNone=function(){
    if(this==null || this.length==0){
        return [];
    }
    var result=[];
    for(var i=0;i<this.length;i++){
        if(!StringsUtils.isNull(this[i])){
            result.push(this[i]);
        }
    }
    return result;
}
Array.prototype.distinct = function(){
	var res = [this[0]];
	for(var i=0; i<this.length; i++){
		var repeat = false;
		for(var j=0; j<res.length; j++){
			if(this[i] == res[j]){
				repeat = true;
				break;
			}
		}
		if(!repeat){
			res.push(this[i]);
		}
	}
	return res;
}
Array.prototype.distinctObject = function(key){
    var res = [this[0]];
    for(var i=0; i<this.length; i++){
        var repeat = false;
        for(var j=0; j<res.length; j++){
            if(this[i][key] == res[j][key]){
                repeat = true;
                break;
            }
        }
        if(!repeat){
            res.push(this[i]);
        }
    }
    return res;
}
Array.prototype.delObject = function(filter){
    var idx = filter;
    if(typeof filter == 'function'){
        for(var i=0;i<this.length;i++){
            if(filter(this[i],i)) idx = i;
        }
        this.splice(idx,1)
    }else{
        var index = this.indexOf(idx);
        if (index > -1) {
            this.splice(index, 1);
        }
    }
}
//去掉undefined
Array.prototype.delUndefined = function(){
    var result=[];
  if(DataObject.arrNotNull(this)){
      for(var i=0;i<this.length;i++){
          if(this[i]!=undefined){
              result.push(this[i])
          }
      }
  }
  return result;
}
//是否存在重复数据
Array.prototype.hasRepeatData = function(){
    var repeat=false;
    for(var i=0; i<this.length; i++){
         repeat = false;
        for(var j=i+1; j<this.length; j++){
            if(this[i] == this[j] && !StringsUtils.isNull(this[i])){
                repeat = true;
                break;
            }
        }
        if(repeat){
         break;
        }
    }
    return repeat;
}
var Alert={
    show:function(msg,icon,time){
        if(StringsUtils.isNull(icon)){
            icon=7;
        }
        if(StringsUtils.isNull(time)){
            time=3;
        }
        time=time*1000;
        layer.msg(msg, {icon: icon,time:time});
    },
    loading_:function(){
        return layer.load(1, {
                shade: [0.1,'#fff']
            });
    }
}
//去掉空数据
Array.prototype.takeOutNull = function(){
    var res = [];
    for(var i=0; i<this.length; i++){
       if(!StringsUtils.isNull(this[i])){
           res.push(this[i]);
       }
    }
    return res;
}

Date.prototype.Format = function (fmt) { //author: meizz
    var o = {
        "M+": this.getMonth() + 1, //月份
        "d+": this.getDate(), //日
        "h+": this.getHours(), //小时
        "m+": this.getMinutes(), //分
        "s+": this.getSeconds(), //秒
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度
        "S": this.getMilliseconds() //毫秒
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}
var sortByKey = function(key){
    return function(o, p){
        var a, b;
        if (typeof o === "object" && typeof p === "object" && o && p) {
            a = o[key];
            b = p[key];
            if (a === b) {
                return 0;
            }
            if (typeof a === typeof b) {
                return a < b ? -1 : 1;
            }
            return typeof a < typeof b ? -1 : 1;
        }
        else {
            throw ("error");
        }
    }
}
var arrayUtils={
    /**
     *对比两个对象数组是否相同,并且标识出不同点,合并到第二个参数中.
     * 标识出不同对象不同类型 update delete add
     * @param caPreList
     * @param enPreList
     * @returns {*}
     */
    caseStepConcat:function(caPreList,enPreList){
        var result=[];
        if(caPreList && caPreList.length>0) {
            if(enPreList && enPreList.length>0) {
                $.each(caPreList,function(n,item){

                    //是否存在
                    var isExist=false;
                    //是否变更
                    var isChanged=false;
                    //变更索引
                    var changedIndex=0;
                    for(var i=0;i<enPreList.length;i++){
                        if(item.index_key==enPreList[i].index_key){
                            isExist=true;
                            if(item.case_step!=enPreList[i].case_step) {
                                isChanged=true;
                                changedIndex=i;
                            }
                        }
                    }
                    //存在
                    if(isExist){
                        if(isChanged){
                            //存在变更
                            item.parentUpdate="update";
                            item.oldStep=enPreList[changedIndex].case_step
                            result.push(item);

                        }else {
                            //未变更
                            result.push(item);
                        }

                    }else {
                        item.parentUpdate="add";
                        result.push(item);
                    }
                });
            }else {
                $.each(caPreList,function(i,v){
                    v.parentUpdate="add";
                    result.push(v);
                })
            }
        } else {
            if(enPreList && enPreList.length>0) {
                for(var i=0;i<enPreList.length;i++){
                    enPreList[i].parentUpdate="delete";
                }
                result=enPreList;
            }
        }
        //对比两个数组长度 添加delete
        if(result && result.length>0){
            if(enPreList && enPreList.length>0){
                $.each(enPreList,function(i,item){
                    var isExist=false;
                    $.each(result,function(n,v){
                        if(v.index_key==item.index_key)
                        {
                            isExist=true;
                        }
                    })
                    if(!isExist){
                        item.parentUpdate="delete";
                        result.push(item);
                    }
                })
            }
        }

        return result;
    },
    deleteObj:function(data){
        var result=[];
        if(data && data.length>0){
            for(var i=data.length-1;i>=0;i--){
                if(data[i].parentUpdate!="delete") {
                    delete data[i].parentUpdate
                    result.push( data[i]);
                }
            }
        }
        return result.reverse();
    }
}

//如果查询信息为空，显示1s无信息
function message_null(data,index_timeout){
    clearTimeout(index_timeout);
    layer.closeAll();
    if(data == "" || data == null || data == undefined || data.length == 0){
        layer.msg('查无数据！', {icon: 7, time: 1000});
    }
}
var Warnning=function(msg,yCallBack,nCallBack){
    var indexconfirm=layer.confirm(msg, {
        btn: ['是', '否']
    }, function () {
        layer.close(indexconfirm);
        yCallBack();
    },function(){
        layer.close(indexconfirm);
        nCallBack();

    });
}

// send websocket parameter

var PopupManager=function(tag,w,h,title){
    if(StringsUtils.isNull(w)){
        w=300;
    }
    if(StringsUtils.isNull(h)){
        h=300;
    }
    $(tag).dialog({
        autoOpen : false,
        title : title,
        modal : true,
        width : w,
        height : h,
        close:function () {
            $(this).dialog("destroy");
        }
    });
    $('.ui-dialog').css('z-index','9999');
    $(tag ).dialog("open");
}
var DataObject={
    isNotNull:function(data){
        if(data!=undefined && data!=null) {
            return true;
        }
        return false;
    },
    isNotNullResult:function(data){
        if(data!=undefined && data!=null && data.hasOwnProperty("status") && data.status==1 && DataObject.arrNotNull(data.data)){
            return true;
        }
        return false;
    },
    arrNotNull:function(data){
        if(data!=undefined && data!=null && data.length>0){
            return true;
        }
        return false;
    },
    checkedResult:function(data){
        if(data!=undefined && data!=null && data.status==1 && data.data!=undefined && data.data!=null){
            return true;
        }
        return false;
    }
}
Array.prototype.contains = function ( needle ) {
    for (i in this) {
        if (this[i] == needle) return true;
    }
    return false;
}
Array.prototype.remove = function(val) {
    var index = this.indexOf(val);
    if (index > -1) {
        this.splice(index, 1);
    }
};




var ObjectTools = {
    clone: function (obj){
        if (obj === null) return null;
        var o = obj.constructor === Array ? [] : {};
        for(var i in obj){
            o[i] = (obj[i] instanceof Date) ? new Date(obj[i].getTime()) : (typeof obj[i] === "object" ? arguments.callee(obj[i]) : obj[i]);
        }
        return o;
    },
    eqs: function(str1, str2) {
        return str1.toLowerCase() === str2.toLowerCase();
    },
    isArray: function(arr) {
        return Object.prototype.toString.apply(arr) === "[object Array]";
    },
    apply:function(){

    }
}

var base64Util = {
    encode:function(code){//编码
        var str = '';
        if (code != undefined && code.length > 0) {
            str = BASE64.encoder(code);//返回unicode 编码
        }
        // return str;
        return code;
    },
    decode:function(code){//解码
        var str = '';
        if (code != undefined && code.length > 0) {
            var unicode = BASE64.decoder(code);//返回会解码后的unicode码数组。
            //可由下面的代码编码为string
            for (var i = 0, len = unicode.length; i < len; ++i) {
                str += String.fromCharCode(unicode[i]);
            }
        }
        // return str;
        return code;
    }
}
var User_t={};
//全局当前登录用户
var UserUtil={
    setUser:function(name){
        loginUserName = name;
    },
    getUser:function(){
        return loginUserName;
    }
};
//checkbox复选框 全选按钮问题
var CheckboxUtil = {
     checkAll:function(o){
         var count=0,num=0;
         for(var i=0;i<o.length-1;i++)
         {
             if(o[i+1].checked==true)
             {
                 count++;
             }
             if(o[i+1].checked==false)
             {
                 num++;
             }
         }
         if(count==o.length-1)
         {
             o[0].checked=true;
         }
         if(num>0)
         {
             if(o[0].checked==true)
             {
                 o[0].checked=false;
             }
         }
         return o;
     }
};


//项目状态
var ProjectUtil = {
    UNKNOWN:0,
    PREPARING_RES:1,
    PREPARING_RES_WAIT_FOR_CARRIER:2,
    PAUSED:3,
    STARTING:4,
    STARTED:5,
    RUNNING:6,
    COMMENTS:7,
    UPLOADING:8,
    FAILURE:9,
    WARNING:10,
    END:11,
    getStatus : function(status) {
        if ("unknown"==status) {
            return ProjectUtil.UNKNOWN;
        }
        if ("preparing_res"==status) {
            return ProjectUtil.PREPARING_RES;
        }
        if ("preparing_res_wait_for_carrier"==status) {
            return ProjectUtil.PREPARING_RES_WAIT_FOR_CARRIER;
        }
        if ("paused"==status) {
            return ProjectUtil.PAUSED;
        }
        if ("starting"==status) {
            return UnitStatus.STARTING;
        }
        if ("started"==status) {
            return ProjectUtil.STARTED;
        }
        if ("running"==status) {
            return ProjectUtil.RUNNING;
        }
        if ("comments"==status) {
            return ProjectUtil.COMMENTS;
        }
        if ("failure"==status) {
            return ProjectUtil.FAILURE;
        }
        if ("end"==status){
            return ProjectUtil.END;
        }
        if ("uploading"==status) {
            return ProjectUtil.UPLOADING;
        }
        if ("warning"==status) {
            return ProjectUtil.WARNING;
        }
        return -1;
    },
    getCaseOperateType:function(num){
        switch (num){
            case 1:
                return "replay";
            case 2:
                return "del";
            case 3:
                return "restart";
            case 4:
                return "add";
        }
    }
}
var wsParameter=function(cmd,msg,other) {
    if(StringsUtils.isNull(msg)){
        msg="";
    }
    return JSON.stringify({"cmd":cmd,"data":msg,"other":other});
}
/**
 * 校验用例步骤/前置是否有变更
 * @param data1原始数据
 * @param data2新数据
 * @param type(true/false) (前置/步骤)
 * @returns {Array}
 */
function checkDifferent(data,data2,type){
    var result=[];
    if(DataObject.arrNotNull(data)){
        if(DataObject.arrNotNull(data2)){
            $.each(data2,function(i,v){
                $.each(data,function(ii,vv){
                    if(v.index_key==vv.index_key){
                        if(type){
                            if(StringsUtils.isNotNull(v.case_step)&&v.case_step!=vv.case_step){
                                result.push({"key1":v.case_step,"key2":"","oldKey1":vv.case_step,"oldKey2":""})
                            }
                        }else{
                            if( StringsUtils.isNotNull(v.case_step)&&(v.case_step!=vv.case_step || v.case_expect!=vv.case_expect)){
                                result.push({"key1":v.case_step,"key2":v.case_expect,"oldKey1":vv.case_step,"oldKey2":vv.case_expect})
                            }
                        }
                    }
                })
            })
        }
    }
    return result;
}
//计算与截止日期的差值
function countDate(abortDate){
	var s,min,day,month,year;
	var ms = Math.floor((new Date(abortDate - (new Date()).getTime())) / 1000);
	if(ms >= 0){
		s = Math.floor(ms%60);
		min = Math.floor(ms/60%60);
		hour = Math.floor(ms/60/60%60);
		day = Math.floor(ms/24/60/60%24);
		month = Math.floor(ms/30/24/60/60%30);
		year = Math.floor(ms/12/30/24/60/60%12);
	}else{
		s = Math.ceil(ms%60);
		min = Math.ceil(ms/60%60);
		hour = Math.ceil(ms/60/60%60);
		day = Math.ceil(ms/24/60/60%24);
		month = Math.ceil(ms/30/24/60/60%30);
		year = Math.ceil(ms/12/30/24/60/60%12);
	}

	console.log(year+"年"+month+"月"+day+"日"+hour+"时"+min+"分"+s+"秒");
	var html =  '<span class="count_year" times="'+ year +'">'+ Math.abs(year) +'年</span>'+
				'<span class="count_month" times="'+ month +'">'+ Math.abs(month) +'月</span>'+
				'<span class="count_day" times="'+ day +'">'+ Math.abs(day) +'日</span>'+
				'<span class="count_hour" times="'+ hour +'">'+ Math.abs(hour) +'时</span>'+
				'<span class="count_min" times="'+ min +'">'+ Math.abs(min) +'分</span>'+
				'<span class="count_s" times="'+ s +'">'+ Math.abs(s) +'秒</span>';
	return html;
}
