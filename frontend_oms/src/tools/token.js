// 封装localStorage存取token

const key = 'oms'

const setToken = (token) => {
  return window.localStorage.setItem(key, token)
}

const getToken = () => {
  return window.localStorage.getItem(key)
}

const setLocalUser = (username) => {
  return window.localStorage.setItem('username', username)
}

const setLocalUid = (userInfo) => {
  return window.localStorage.setItem('uid', userInfo)
}

const getLocalUser = () => {
  return window.localStorage.getItem('username')
}

const getLocalUid = () => {
  return window.localStorage.getItem('uid')
}

const removeToken = () => {
  return window.localStorage.removeItem(key)
}

const removeLocalUser = () => {
  return window.localStorage.removeItem('username')
}

export {
  setToken,
  getToken,
  removeToken,
  setLocalUid,
  setLocalUser,
  getLocalUser,
  getLocalUid,
  removeLocalUser
}