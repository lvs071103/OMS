import { Breadcrumb, theme } from 'antd'
import 'moment/locale/zh-cn'
import { Link } from 'react-router-dom'
import UserTable from './table'

const User = () => {

  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()

  return (
    <>
      <Breadcrumb
        style={{
          margin: '16px 0',
        }}
        items={[
          { title: <Link to='/'>首页</Link> },
          { title: '用户管理' },
          { title: '用户列表' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <UserTable />
      </div>
    </>
  )
}

export default User