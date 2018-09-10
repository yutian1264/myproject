<template>
  <div id="c">
    <ul>
      <li><a @click="headerClick('first')">首页</a></li>
      <li><a  @click="headerClick('setting')">setting</a></li>
      <li><a  @click="headerClick('map')">map</a></li>
      <li><a>帮助</a></li>
      <li><a>关于我们</a></li>
    </ul>
    <span style="float: right;">欢迎{{user}}</span>
    <!--<button @click.stop="additem">additem</button>-->
    <!--<input v-model="myvalue" @keydown.enter="testkeydown">-->
   <!--<com-a :number="myvalue" @my-event="myEvent">-->
     <!--<p slot="header">header</p>-->
     <!--<p>test</p>-->
     <!--<p slot="footer">footer</p>-->
   <!--</com-a>-->
    <!--{{tar}}-->
  </div>
</template>

<script>
    import Vue from "vue"
    import comA from "./../tpl/compA.vue"
    export default {
      name:"app",
      components:{
        comA
      },
      data(){
        return{

          user:"wyt"
        }
      },
      methods:{
        headerClick(s){
          this.$emit("changetab",s)
        }
      }
    }

    // function test() {
    //   $.ajax({
    //     url:"http://192.168.110.143:8080/bzy/user/logins.do",
    //     type:"post",
    //     data:JSON.stringify({name:"wyt",pwd:"123",loginType:0}),
    //     dataType : "json",
    //     contentType : "application/json; charset=utf-8",
    //     success:function(d){
    //      console.log(d)
    //     },
    //     error:function(e){
    //       console.log("获取远程数据异常")
    //     }
    //   })
    // }

</script>

<style scoped>
  ul {
    list-style-type: none;
    padding: 0;
  }
a{
  cursor: pointer;
}
  li{
    display: inline-block;
    margin: 0 10px;
  }
  #c{
    display:inline-flex;
    width: 100%;
    display: flex;

  }
</style>
