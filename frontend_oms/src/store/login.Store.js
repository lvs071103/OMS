// login module
import {
  getToken,
  removeToken,
  setToken,
  removeLocalUser,
  setLocalUid
} from '@/tools/token'
import { http } from '@/tools/http'
import { setLocalUser } from '@/tools/token'
import { makeAutoObservable, runInAction } from 'mobx'

class LoginStore {
  token = getToken() || ''
  user = null
  constructor() {
    // 响应式
    makeAutoObservable(this)
  }

  getToken = async ({ username, password }) => {
    // 调用登陆接口
    const response = await http.post('/v1/login', {
      "username": username, "password": password
    })
    // 存入token
    this.token = response.data.data.token
    // 存入localStorage
    setToken(this.token)
    // 存入用户名
    setLocalUser(username)
    // 存储uid
    setLocalUid(response.data.data.uid)
    // 获取用户信息
    // this.fetchUserDetails(response.data.data.uid)
  }

  fetchUserDetails = async (uid) => {
    try {
      const response = await http.get(`/v1/user/${uid}`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      })
      runInAction(() => {
        this.user = response.data
      })
    } catch (error) {
      console.error('Failed to fetch user details:', error)
    }
  };

  // 退出
  loginOut = () => {
    // 移除token
    removeToken()
    removeLocalUser()
    this.user = null
  }
}

export default LoginStore