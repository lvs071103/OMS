import { Breadcrumb, theme } from 'antd'
import 'moment/locale/zh-cn'
import { Link } from 'react-router-dom'
import JobTable from './table'

const Jobs = () => {

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
          { title: '任务列表' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <JobTable />
      </div>
    </>
  )
}

export default Jobs