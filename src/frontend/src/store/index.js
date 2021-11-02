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
    program: {
      Id: '',
      Title: '',
      Description: '',
      DrawFlowData: ''
    }
  },
  mutations: {
    setCode: (state, code) => (state.code = code),
    setOutput: (state, output) => (state.output = output),
    setError: (state, code) => (state.code = code),
    clearCode: (state) => (state.code = ''),
    clearOutput: (state) => (state.output = ''),
    setProgramList: (state, programs) => (state.programs = programs),
    setProgram: (state, program) => (state.program = program)
  },
  getters: {
    allPrograms: (state) => {
      return state.programs
    },
    getProgram: (state) => {
      return state.program
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
    getDetails({commit}, uid){
      Vue.axios.get(config.devServer.api + "programs/" + uid).then((response) => {
        if(response.data.StatusCode === "200"){
          let p = JSON.parse(response.data.Result).programByUid[0];
          
          commit('setProgram', {
            Id: p.uid,
            Title: p.title,
            Description: p.description,
            DrawFlowData: p.drawFlowData
          });
        }else{
          commit('setError', response.data.Error)
        }
      })
      .catch(error => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    },
    programs ({commit}){
      Vue.axios.get(config.devServer.api + "programs").then((response) => {
        if(response.data.StatusCode === "200"){
          commit('setProgramList', JSON.parse(response.data.Result).programs);
        }else{
          commit('setError', response.data.Error)
        }
      })
      .catch(error => {
        this.errorMessage = error.message;
        console.error("There was an error!", error);
      });
    },
    parseCode ({commit}, drawflowData){
      Vue.axios.post(config.devServer.api + "program/parse",{Data: JSON.stringify(drawflowData)},{}).then((response) => {
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
    add ({commit}, drawflowData){
      Vue.axios.post(config.devServer.api + "program/add",drawflowData,{}).then((response) => {
        if(response.data.StatusCode === "200"){
          this.$router.push('programs');
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
      Vue.axios.post(config.devServer.api + "program/run",{Data: sourceCode},{}).then((response) => {
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
    }
  },
  modules: {
  }
})