import 'antd/dist/reset.css' // 正确导入 antd 的 CSS 文件
import { unstable_HistoryRouter as HistoryRouter, Route, Routes } from 'react-router-dom'
import { history } from '@/tools/history'
import { Suspense, lazy } from 'react'
import AuthRoute from './components/AuthRoute'
import LayoutApp from 'src/pages/Layout'


const Login = lazy(() => import('@/pages/Auth/Login'))
const Home = lazy(() => import('@/pages/Home'))
const User = lazy(() => import('@/pages/Auth/User'))
const Group = lazy(() => import('@/pages/Auth/Group'))
const EnvApp = lazy(() => import('@/pages/Sys/Envs'))
const SignUp = lazy(() => import('@/pages/Auth/SignUp'))
const ReleaseSummary = lazy(() => import('@/pages/Release/index'))
const JeninsInstances = lazy(() => import('@/pages/Release/Jenins/index'))
const Jobs = lazy(() => import('@/pages/Release/Job/index'))
const Queue = lazy(() => import('@/pages/Release/Queue/index'))

function App () {
  return (
    // 路由配置
    <HistoryRouter history={history}>
      <Suspense
        fallback={
          <div
            style={{
              textAlign: 'center',
              marginTop: 400
            }}
          >
            loading...
          </div>
        }
      >
        <Routes>
          {/* 注册 */}
          <Route path='/SignUp' element={<SignUp />}></Route>
          {/* 创建路由path和组件对应关系 */}
          {/* Layout需要鉴权处理 */}
          <Route path='/' element={
            <AuthRoute>
              <LayoutApp />
            </AuthRoute>
          }>
            <Route index element={<Home />}></Route>
            <Route path='/v1/user/list' element={<User />}></Route>
            <Route path='/v1/group/list' element={<Group />}></Route>
            <Route path='/v1/release/summary' element={<ReleaseSummary />}></Route>
            <Route path='/v1/sys/config/env/list' element={<EnvApp />}></Route>
            <Route path='/v1/app/release/jenkins/list' element={<JeninsInstances />}></Route>
            <Route path='/v1/app/release/job/list' element={<Jobs />}></Route>
            <Route path='/v1/app/release/queue/list' element={<Queue />}></Route>
          </Route>
          <Route path='/login/' element={<Login />}></Route>
        </Routes>
      </Suspense>
    </HistoryRouter>
  )
}

export default App
