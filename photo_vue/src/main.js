import Vue from 'vue'
import headerView from './tpl/header.vue'
import container from './tpl/container.vue'
new Vue({
  el: '#app',
  components:{
    header_view:headerView,
    container
  },
  data(){
    return {
      tab:"setting"
    }
  },
  methods:{
    header_change(s){
      this.tab=s;
    },
    onclick_test(){
      alert("click")
    }
  }

})

