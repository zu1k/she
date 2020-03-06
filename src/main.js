import Vue from 'vue'
import App from './App.vue'
import ZeitUI from '@zeit-ui/vue'
import '@zeit-ui/vue/dist/zeit-ui.css' // require style
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false
Vue.use(VueAxios, axios)
Vue.use(ZeitUI)

// ZeitUI.theme.enableLight()
// ZeitUI.theme.enableDark()

new Vue({
  render: h => h(App),
}).$mount('#app')
