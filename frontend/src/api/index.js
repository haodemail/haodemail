import axios from 'axios'

let base
if (process.env.NODE_ENV === 'development') {
    base = 'http://localhost:8000'
} else {
    base = ''
}

export const listDomainAPI = params => { return axios.get(`${base}/api/domain/list`, { params: params }) }
export const addDomainAPI = params => { return axios.post(`${base}/api/domain/add`, { params: params }) }
