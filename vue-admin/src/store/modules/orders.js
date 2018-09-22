import axios from '../../axios.js'

/**
 * Route
 */

const route = '/orders'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Valor Inicial', value: 'initial_value' },
  { text: 'Valor Final', value: 'final_value' },
  { text: 'Descuento', value: 'discount' },
  { text: 'Estado', value: 'status' },
  { text: 'Cliente', value: 'clients' },
  { text: 'Pasarela', value: 'gateways' },
  { text: 'Pais', value: 'countries' },
  { text: 'Id de Pago', value: 'payment_id' },
  { text: 'Acciones', align: 'center', value: '' }
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
