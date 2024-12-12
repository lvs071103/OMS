import { http } from "@/tools/http"
import { getLocalUid } from "@/tools/token"
import { makeAutoObservable, runInAction } from "mobx"
import { message } from "antd"

class UserStore {
  userInfo = null
  // 响应式
  //constructor函数是类的特殊方法，用于创建和初始化类的实例
  constructor() {
    // 是 MobX 提供的一个函数，用于将类的属性和方法自动转换为响应式的。
    makeAutoObservable(this)
  }
  // 获取用户信息
  getUserInfo = async () => {
    try {
      // 获取当前用户信息
      const resp = await http.get(`/v1/user/${getLocalUid()}`)
      // 存入用户信息
      runInAction(() => {
        this.userInfo = resp.data.data
      })
      if (resp.status !== 200) {
        message.error(resp.response.data.error)
        return
      }
    } catch (error) {
      console.error('Failed to fetch user details:', error)
    }
  }
}

export default UserStore