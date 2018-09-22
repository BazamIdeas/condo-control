import axios from '../../axios.js'

/**
 * Route
 */

const route = '/portfolios'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Nombre', value: 'name' },
  { text: 'Descripción', value: 'description' },
  { text: 'Cliente', value: 'client' },
  { text: 'Locación', value: 'location' },
  { text: 'Servicio', value: 'service' },
  { text: 'Actividad', value: 'activity' },
  {text: 'Acciones', align: 'center', value: ''}
]
const options = { headers: { 'Content-Type': 'multipart/form-data' } }

var getForm = (data) => {
  let formData = new FormData()
  let images
  if (data.files) {
    images = data.files
  } else if (data.file) {
    images = data.file
  }
  let length = images.length

  formData.append('name', data.name)
  formData.append('description', data.description)
  formData.append('client', data.client)
  formData.append('priority', data.priority)
  formData.append('location[id]', data.location.id)
  formData.append('service[id]', data.service.id)
  formData.append('activity[id]', data.activity.id)

  // Set Images
  for (let i = 0; i < length; i++) {
    const element = images[i]
    formData.append('images[]', element)
  }

  return formData
}

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {}, trashed: [], struct: struct },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    },
    CREATE (state, data) {
      state.all.push(data.res)
      alert('El elemento fue creado')
    },
    UPDATE_ONE (state, data) {
      state.all[data.item.in] = data.item
      alert('El elemento fue actualizado')
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
    },
    async create ({ commit }, params) {
      let formData = getForm(params.item)
      let item = params.item

      try {
        let res = await axios.post('/' + params.state, formData, options)
        params.res = res.data
        params.res.location = item.location
        params.res.service = item.service
        params.res.activity = item.activity
        commit('CREATE', params)
      } catch (error) {
        console.log(error)
      }
    },
    async updateOne ({ commit }, params) {
      let formData = getForm(params.item)

      try {
        let res = await axios.put('/' + params.state + '/' + params.item.id, formData, options)
        params.res = res.data
        commit('UPDATE_ONE', params)
      } catch (error) {
        console.log(error)
      }
    }
  }
}
