import Vue from 'vue'
import Vuex from 'vuex'
//import configuration from '../config/config'

Vue.use(Vuex)


export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {

    parseCode ({commit}, drawflowData){
      debugger
      Vue.axios.post("http://localhost:3000/parse",{Data: JSON.stringify(drawflowData)},{}).then((response) => {
        console.log(response.data)
      })
      .catch(error => {
        debugger
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    },
    save ({commit}, drawflowData){
      debugger
      Vue.axios.post("http://localhost:3000/save",{Data: JSON.stringify(drawflowData)},{}).then((response) => {
        console.log(response.data)
      })
      .catch(error => {
        debugger
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    }
  },
  modules: {
  }
})
