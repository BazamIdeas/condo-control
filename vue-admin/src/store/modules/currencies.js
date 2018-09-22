import axios from '../../axios.js'

/**
 * Route
 */

const route = '/currencies'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Nombre', value: 'name' },
  { text: 'Iso', value: 'iso' },
  { text: 'Simbolo', value: 'symbol' },
  { text: 'Pais', value: 'countries[0].name' },
  {text: 'Acciones', align: 'center', value: ''}
]

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {}, trashed: [], struct: struct },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    }
  },
  actions: {
    async getAll ({commit}) {
      try {
        let res = await axios.get(route)
        commit('GET_ALL', res.data)
      } catch (error) {
        console.log(error)
      }
    }
  },
  getters: { }
}
