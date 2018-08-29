import Vue from 'vue'
import App from './tpl/App.vue'
import headerView from './tpl/header.vue'
import setting from './tpl/Setting.vue'

new Vue({
  el: '#app',
  components:{
    header_view:headerView,
    app:App,
    setting:setting
  },

})

