import Vue from 'vue'
import App from './App.vue'

import VueRouter from 'vue-router';
Vue.use(VueRouter);

Vue.config.productionTip = false

import Images from './components/Images.vue'
import Index from './components/Index.vue'
import Delete from './components/Delete.vue'
import Upload from './components/Upload.vue'
import NotFound from './components/NotFound.vue'

import 'bootstrap';

import 'bootstrap/dist/css/bootstrap.css';

import titleMixin from './mixins/titleMixin';
Vue.mixin(titleMixin)


const routes = [
  {
    path: '/',
    redirect: '/images'
  },
  {
    name: 'Images',
    path: '/images',
    component: Images
  },
  {
    name: 'Index',
    path: '/images/:id',
    component: Index
  },
  {
    name: 'Delete',
    path: '/delete/:id',
    component: Delete
  },
  {
    name: 'Upload',
    path: '/upload',
    component: Upload
  },
  {
    name: 'Not Found',
    path: '*',
    component: NotFound
  }
  
]

const router = new VueRouter({ mode: "history",
routes: routes})

new Vue(Vue.util.extend({ router },
  App)).$mount('#app');
