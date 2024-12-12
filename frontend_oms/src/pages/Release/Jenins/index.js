import { Breadcrumb, theme } from 'antd'
import 'moment/locale/zh-cn'
import { Link } from 'react-router-dom'
import InstancesTable from './table'

const JeninsInstances = () => {

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
          { title: '应用发布' },
          { title: '应用接入' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <InstancesTable />
      </div>
    </>
  )
}

export default JeninsInstances