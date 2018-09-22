import axios from 'axios'

var headers = {}
const token = localStorage.getItem('bazam-token')

var baseUrl = process.env.NODE_ENV === 'development' ? 'http://localhost:9090/v1' : ''

if (token !== null && token !== undefined && token !== '') {
  headers['Authorization'] = token
}

var instance = axios.create({
  baseURL: baseUrl,
  headers: headers
})
export default instance
