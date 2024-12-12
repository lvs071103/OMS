import { Breadcrumb, theme } from 'antd'
import 'moment/locale/zh-cn'
import { Link } from 'react-router-dom'
import EnvTable from './table'

const EnvApp = () => {

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
          { title: '系统设置' },
          { title: '环境设置' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <EnvTable />
      </div>
    </>
  )
}

export default EnvApp