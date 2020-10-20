/*
 * File         : app.js
 * Project      : NetMon
 * Created Date : Wednesday, Oct 7th 2020, 1:20:14 AM
 * Author       : Pramod Devireddy
 * 
 * Last Modified: Thursday, 8th October 2020 8:25:59 am
 * Modified By  : Pramod Devireddy
 * 
 * Copyright (c)2020 - Pramod Devireddy
 * ************************* Description *************************
 *  
 * ***************************************************************
 */


const home_page = httpVueLoader('js/HomePage.vue')
const log_data_page = httpVueLoader('js/LogDataPage.vue')


const routes = [{
  path: '/netmon',
  component: home_page
}, {
  path: '/netmon/log',
  component: log_data_page
}, ]

const router = new VueRouter({
  routes: routes,
  mode: 'history'
})


new Vue({
  el: "#app",
  vuetify: new Vuetify(),
  router,
  data: {
    drawer: false,
    items: [
      {
        title: 'Log',
        icon: 'mdi-network',
        path: './log'
      },
      {
        title: 'About',
        icon: 'mdi-account',
        path: './about'
      },
    ],
  },
  created: function () {


  }
});
