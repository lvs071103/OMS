import React from 'react'
import GroupTable from './table'
import { Breadcrumb, theme } from 'antd'
import { Link } from 'react-router-dom'

export default function Group () {
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
          { title: '用户组' }]}
      />
      <div
        style={{
          background: colorBgContainer,
          borderRadius: borderRadiusLG,
        }}
      >
        <GroupTable />
      </div>
    </>
  )
}
