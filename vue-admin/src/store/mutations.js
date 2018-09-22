export const STORAGE_KEY = 'bazam'

// for testing
if (navigator.userAgent.indexOf('PhantomJS') > -1) {
  window.localStorage.clear()
}

export const mutations = {
  LOGIN (state, data) {
    state.user = data.res
    localStorage.setItem('bazam-token', state.user.token)
    localStorage.setItem('bazam-user', JSON.stringify(state.user))
    alert(data.message)
  },
  GET_ALL (state, data) {
    state[data.state].all = data.res
  },
  GET_ONE (state, data) {
    state[data.state].editedItem = data.res
  },
  CREATE (state, data) {
    state[data.state].all.unshift(data.res)
    alert('El elemento fue creado')
  },
  UPDATE_ONE (state, data) {
    state[data.state].all[data.item.in] = data.item
    console.log(state[data.state].all[data.item.in])
    alert('El elemento fue actualizado')
  },
  DELETE_ONE (state, data) {
    state[data.state].all.splice(data.item.index, 1)
    alert('El elemento fue eliminado')
  },
  GET_ALL_TRASHED (state, data) {
    state[data.state].trashed = data.res
  },
  RESTORE (state, data) {
    state[data.state].trashed.splice(data.item.index, 1)
    alert('El elemento fue restaurado')
  },
  TRASH (state, data) {
    state[data.state].trashed.splice(data.item.index, 1)
    alert('El elemento fue eliminado permanentemente')
  }
}
