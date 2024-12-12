import React from 'react'
import { Breadcrumb, Layout, theme } from 'antd'



export default function HomeApp () {
  const breadcrumbItems = [
    { title: '仪表盘' },
  ]
  const { Content } = Layout
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken()

  return (
    <>
      <Content
        style={{
          margin: '0 16px',
        }}
      >
        <Breadcrumb
          style={{
            margin: '16px 0',
          }}
          items={breadcrumbItems}
        >
        </Breadcrumb>
        <div
          style={{
            padding: 24,
            minHeight: 360,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          仪表盘
        </div>
      </Content>
    </>
  )
}
