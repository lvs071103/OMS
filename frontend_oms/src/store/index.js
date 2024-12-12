// 把所有的模块做统一处理
// 导出一个统一的方法: useStore
import LoginStore from "./login.Store"
import React from "react"
import UserStore from "./user.Store"

class RootStore {
  constructor() {
    this.loginStore = new LoginStore()
    this.userStore = new UserStore()
  }
}

// 实例化根
// 导出useStore context
const rootStore = new RootStore()
const context = React.createContext(rootStore)
const useStore = () => React.useContext(context)

export { useStore }