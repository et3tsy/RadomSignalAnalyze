//导入vue
import Vue from 'vue';
import VueRouter from 'vue-router';
//导入组件
import Main from "../views/Main";

//使用
Vue.use(VueRouter);
//导出
export default new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Main
    },
  ]

})
