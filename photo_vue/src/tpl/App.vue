<template>
  <div id="app">
    <input v-model="myvalue" type="text">
    {{myvalue}}
    apple:<input type="checkbox" v-model="mybox" name="mybox" value="apple" checked/>
    orange:<input type="checkbox" v-model="mybox" name="mybox" value="orange"/>
    banana:<input type="checkbox" v-model="mybox" name="mybox" value="banana"/>
    xg:<input type="checkbox" v-model="mybox" name="mybox" value="xg"/>
    {{mybox}}
    first:<input v-model="radio" type="radio" name="myradio" value="first">
    second:<input  v-model="radio" type="radio" name="myradio" value="second">
    {{radio}}
    <p>
      <select v-model="sel" @change="selChange">
        <option v-for="item in selecttion" :value="item.value">{{item.text}}</option>

      </select>
      {{sel}}
    </p>
   <keep-alive>
      <p :is="currentView"> </p>
    </keep-alive>
    <com-a v-if="sel==0"></com-a>
    <auto-comp v-if="sel==1"></auto-comp>



  </div>
</template>

<script>
  import Vue from "vue"
  import AutoComp from "./../tpl/AutoComponent.vue"
  import ComA from "./../tpl/compA.vue"
export default {
  name: 'app',
  components:{
    AutoComp,
   ComA
  },

  data () {
    return {
      selecttion:[
        {text:"coma",value:0},
        {text:"auto",value:1}
      ],
      msg: 'Welcome to Your Vue.js App',
      myvalue:'',
      mybox:[],
      radio:[],
      sel:0,
      currentView:"com-a"
    }
  },
  methods:{
    selChange(){
      console.log("change"+$(event.target).val())
      switch ($(event.target).val()){
        case 0:
          this.currentView="com-a"
        //  Vue.set(this.currentView,"com-a")
          break;
        case 1:
          this.currentView="auto-comp"
         // Vue.set(this.currentView,"auto-comp")
          break;
      }
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #cc1d38;
}
</style>
