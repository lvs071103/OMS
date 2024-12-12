// 封装 axios
// 实例化 请求拦截器 响应拦截器

import Info from '@/global'
import axios from 'axios'
import { history } from './history'
import { getToken } from './token'

const { apiURI } = Info

const http = axios.create({
  baseURL: apiURI,
  timeout: 5000,
})
// 添加请求拦截器
http.interceptors.request.use((config) => {
  // if not login and token
  // 获取本地存储的token
  const token = getToken()
  // 如果token存在，就在请求头中添加token
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  // 返回config
  return config
}, (error) => {
  // 对请求错误做些什么
  return Promise.reject(error)
})

// 添加响应拦截器
http.interceptors.response.use((response) => {
  // 2xx 范围内的状态码都会触发该函数。
  // 对响应数据做点什么
  return response
}, (error) => {
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  // console.dir(error?.response?.data?.detail)
  if (error.response.status === 401) {
    // 跳回到登录 reactRouter默认状态下 并不支持在组件之外完成路由跳转
    history.push('/login/')
  }
  // 返回错误
  // return Promise.reject(error)
  return error
})

export { http }

