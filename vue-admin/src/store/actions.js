import axios from '../axios.js'

export default {
  async login ({ commit }, params) {
    try {
      let res = await axios.post('/users/login', params.item)
      params.res = res.data
      commit('LOGIN', params)
    } catch (error) {
      console.log(error)
    }
  },
  async register ({ commit }, params) {
    try {
      let res = await axios.post('/users', params.item)
      params.res = res.data
      commit('LOGIN', params)
    } catch (error) {
      console.log(error)
    }
  },
  async getAll ({ commit }, params) {
    try {
      let res = await axios.get('/' + params.state)
      params.res = res.data
      commit('GET_ALL', params)
    } catch (error) {
      console.log(error)
    }
  },
  async getOne ({ commit }, params) {
    try {
      let res = await axios.get('/' + params.state + '/' + params.item.id)
      params.res = res.data
      commit('GET_ONE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async create ({ commit }, params) {
    try {
      let res = await axios.post('/' + params.state, params.item)
      params.res = res.data
      commit('CREATE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async updateOne ({ commit }, params) {
    try {
      let res = await axios.put('/' + params.state + '/' + params.item.id, params.item)
      params.res = res.data
      commit('UPDATE_ONE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async deleteOne ({ commit }, params) {
    try {
      let res = await axios.delete('/' + params.state + '/' + params.item.id)
      params.res = res.data
      commit('DELETE_ONE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async getAllTrashed ({ commit }, params) {
    try {
      let res = await axios.patch('/' + params.state + '/trashed')
      params.res = res.data
      commit('GET_ALL_TRASHED', params)
    } catch (error) {
      console.log(error)
    }
  },
  async restore ({ commit }, params) {
    try {
      let res = await axios.delete('/' + params.state + '/' + params.item.id + '/restore')
      params.res = res.data
      commit('RESTORE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async trash ({ commit }, params) {
    try {
      let res = await axios.delete('/' + params.state + '/' + params.item.id + '?trash=true')
      params.res = res.data
      commit('TRASH', params)
    } catch (error) {
      console.log(error)
    }
  }
}
