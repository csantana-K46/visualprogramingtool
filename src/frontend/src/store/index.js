import Vue from 'vue'
import Vuex from 'vuex'
import config from '../config/config'

Vue.use(Vuex)


export default new Vuex.Store({
  state: {
    code: '',
    output: '',
    error: '',
    programs: [],
  },
  mutations: {
    setCode: (state, code) => (state.code = code),
    setOutput: (state, output) => (state.output = output),
    setError: (state, code) => (state.code = code),
    clearCode: (state) => (state.code = ''),
    clearOutput: (state) => (state.output = '')
  },
  getters: {
    allPrograms: (state) => {
      return state.programs
    },
    getCode : (state) => {
      return state.code
    }
  },
  actions: {
    clear ({commit}){
      commit('clearCode')
      commit('clearOutput')
    },
    parseCode ({commit}, drawflowData){
      Vue.axios.post(config.devServer.api + "parse",{Data: JSON.stringify(drawflowData)},{}).then((response) => {
        if(response.data.StatusCode === "200"){
          commit('setCode', response.data.Result);
        }else{
          commit('setError', response.data.Error)
        }
      })
      .catch(error => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    },
    save ({commit}, drawflowData){
      Vue.axios.post(config.devServer.api + "save",{Data: JSON.stringify(drawflowData)},{}).then((response) => {
        debugger
        if(response.data.StatusCode === "200"){
          debugger
          commit('setCode', response.data.Result);
        }else{
          commit('setError', response.data.Error)
        }
      })
      .catch(error => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    },
    runCode ({commit}, sourceCode){
      Vue.axios.post(config.devServer.api + "run",{Data: sourceCode},{}).then((response) => {
        if(response.data.StatusCode === "200"){
          commit('setOutput', response.data.Result);
        }else{
          commit('setError', response.data.Error)
        }
      })
      .catch(error => {
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
