const home = { template: '<div id="uploader" class="wu-example">' +
    '    <div id="thelist" class="uploader-list"></div>' +
    '    <div class="btns">' +
    '        <div id="picker">选择文件</div>' +
    '        <button id="ctlBtn" class="btn btn-default">开始上传    </button>' +
    '    </div>' +
    '</div>' }
const about = { template: '/static/htmls/morUpload.vue' }
// import Vue from "vue";
// import VueRouter from "vue-router"

//
// import home from "/static/htmls/t.vue";
// import about from '/static/htmls/morUpload.vue';

const routes = [
    { path: '/home', component: home },
    // { path: '/about', component: about }
]

const router = new VueRouter({
    routes
})

const app = new Vue({
    router
}).$mount('#app')

