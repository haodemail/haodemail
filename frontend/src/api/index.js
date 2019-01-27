import axios from "axios";
import router from "../router";


class APIManager {
  constructor() {
    let base
    if (process.env.NODE_ENV === 'development') {
      base = 'http://localhost:8000'
    } else {
      base = ''
    }

    const Axios = axios.create({
      baseURL: base,
      timeout: 5000,
    })

    // http request interceptors
    Axios.interceptors.request.use(
      config => {
        if (sessionStorage.JWT_TOKEN) {
          config.headers.Authorization = `token ${sessionStorage.JWT_TOKEN}`;
        }
        return config;
      },
      err => {
        return Promise.reject(err);
      });

    // http response interceptors
    Axios.interceptors.response.use(
      response => {
        return response;
      },
      error => {
        if (error.response) {
          switch (error.response.status) {
            case 401:
              // auth failed
              sessionStorage.commit('LOG_OUT');
              router.replace({
                path: 'Login',
                query: {redirect: router.currentRoute.fullPath}
              });
          }
        }
        return Promise.reject(error.response.data);
      });
    this.$http = Axios

  }

  login(params) {
    return this.$http.post("login", { params: params })
  }

  // account manager
  createDomain(params) {
    return this.$http.post("api/domain/create", { params: params })
  }
  listDomain(params) {
    return this.$http.post("api/domain/list", {params: params})
  }
  deleteDomain(params) {
    return this.$http.post("api/domain/delete", { params: params })
  }
  createUser(params) {
    return this.$http.post("api/user/create", { params: params })
  }
  listUser(params) {
    return this.$http.post("api/user/list", {params: params})
  }
  deleteUser(params) {
    return this.$http.post("api/user/delete", { params: params })
  }

}

export default new APIManager()
