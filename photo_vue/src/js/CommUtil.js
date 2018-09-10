var pageSize = 0;
var pageCount = 100;
// var url = "http://192.168.110.143:8500"
var url = "http://192.168.2.100:8500"
var commUtil = {
  //整理数据
  collateData: function (r) {
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
  },
  getResourcByCity:function(){

  },

  test: function (retype, d) {
    retype = StringsUtils.isNull(retype) ? '' : retype
    var url1 = url + "/f/d/getListByPage?type=" + retype + "&pageSize=" + pageSize + "&pageCount=" + pageCount
    console.log(url1)
    https.get(url1, function (r) {
      d(r)
    })
  },
  test1: function (type, owner, d) {
    if (type == "up") {
      if (pageSize > 0) {
        pageSize--
      } else {
        pageSize = 0
      }
    } else {
      pageSize++
    }

    https.get(url + "/f/d/getListByPage?type=" + owner + "&pageSize=" + pageSize + "&pageCount=" + pageCount, function (r) {
      d(r)
    })
  },
  getResList: function (d) {
    https.get(url + "/f/d/getResTypeList", function (r) {
      d(r)
    })
  },
  getOwnweList: function (d) {
    https.get(url + "/f/d/getResOwnerList", function (r) {
      d(r)
    })
  },
  UpdateRess: function (param, d) {
    https.postGo(url + "/f/d/UpdateRess", param, function (r) {
      d(r)
    })
  },
  getAddrByPoints:function(lng,lat){

  }
}
