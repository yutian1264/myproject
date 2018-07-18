
var https= {
    get: function (url, callback) {
        var request = $.ajax({
            url: url,
            type: "get",
            dataType: "json",
            contentType: "application/json; charset=utf-8",
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
    getAsync: function (url, async, callback) {
        $.ajax({
            url: url,
            type: "get",
            dataType: "json",
            async: async,
            contentType: "application/json; charset=utf-8",
            success: function (d) {
                callback(d);
            },
            error: function (e) {

                callback(null);
                console.log("获取远程数据异常")
            }
        })

    },
    post: function (url, param, callback) {
        var request = $.ajax({
            url: url,
            type: "post",
            data: {"param":JSON.stringify(param)},
            dataType: "text",
            // contentType: "application/json; charset=utf-8",
            success: function (d) {
                callback(d);
            },
            error: function (e) {

                if(e.status==405){
                    console.log("44444444444444444444444444405")
                    callback(null);
                }else{
                    callback(null);
                    console.log("获取远程数据异常")
                }

            }
        })
        return request;
    },
    postAsync: function (url, param, async, callback) {
        $.ajax({
            url: url,
            type: "post",
            data: JSON.stringify(param),
            dataType: "json",
            async: async,
            contentType: "application/json; charset=utf-8",
            success: function (d) {
                callback(d);
            },
            error: function (e) {

                callback(null);
                console.log("获取远程数据异常")
            }
        })
    },
    updateItemCaseChildren: function (url, data) {

        $.ajax({
            type: "post",
            url: url,
            dataType: "json",
            contentType: "application/json; charset=utf-8",
            async: false,
            data: JSON.stringify(data),
            success: function (datas) {
                console.log("更新表中字段")
            },
            error: function (r) {
                console.log("get composite case error")
            }
        });
    }
}


