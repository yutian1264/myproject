import Vue from 'vue'
import App from './tpl/App.vue'
import headerView from './tpl/header.vue'

new Vue({
  el: '#app',
  components:{
    header_view:headerView,
    app:App
  },

})

