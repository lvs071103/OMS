import { Breadcrumb, theme } from 'antd'
import 'moment/locale/zh-cn'
import { Link } from 'react-router-dom'
import QueueTable from './table'

const Queue = () => {

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
          { title: '任务队列' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <QueueTable />
      </div>
    </>
  )
}

export default Queue