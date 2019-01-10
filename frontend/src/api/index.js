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
          console.log('axios:' + error.response.status);
          switch (error.response.status) {
            case 401:
              // auth failed
              store.commit('LOG_OUT');
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
    return this.$http.post(`/login`, { params: params })
  }
  logout(params) {
    return this.$http.post(`/logout`, { params: params })
  }
}

export default new APIManager()
