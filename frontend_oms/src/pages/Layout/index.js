import React, { useState, useEffect } from 'react'
import {
  TeamOutlined,
  SettingOutlined,
  LogoutOutlined,
  UserOutlined,
  DashboardFilled,
  CodeOutlined
} from '@ant-design/icons'
import { Layout, Menu, theme, Popconfirm, Avatar } from 'antd'
import { useNavigate, Outlet, Link, useLocation } from 'react-router-dom'
import Info from '@/global'
import { useStore } from '@/store'
import { http } from '@/tools/http'
import { getLocalUser } from '@/tools/token'
import { observer } from 'mobx-react'


const { Header, Footer, Sider } = Layout

function getItem (label, key, icon, children) {
  return {
    key,
    icon,
    children,
    label,
  }
}

const items = [
  getItem(<Link to={'/'}>仪表盘</Link>, '/', <DashboardFilled />),
  getItem('用户管理', 'sub1', <TeamOutlined />, [
    getItem(<Link to={'/v1/group/list'}>用户组</Link>, '/v1/group/list'),
    getItem(<Link to={'/v1/user/list'}>用户列表</Link>, '/v1/user/list'),
  ]),
  getItem('应用发布', 'sub2', <CodeOutlined />, [
    getItem(<Link to={'/v1/release/summary'}>概要</Link>, '/v1/release/summary'),
    getItem(<Link to={'/v1/job/list'}>任务管理</Link>, '/v1/job/list'),
    getItem(<Link to={'/v1/queue/list'}>发布管理</Link>, '/v1/queue/list'),
    getItem(<Link to={'/v1/app/release/jenkins/list'}>应用接入</Link>, '/v1/app/release/jenkins/list'),
  ]),
  getItem('系统设置', 'sub3', <SettingOutlined />, [
    getItem(<Link to={'/v1/sys/config/env/list'}>环境设置</Link>, '/v1/sys/config/env/list'),
  ]),
]

const LayoutApp = observer(() => {

  const {
    loginStore,
    userStore
  } = useStore()

  const [username, setUserName] = useState()
  const [collapsed, setCollapsed] = useState(false)
  const {
    token: { colorBgContainer },
  } = theme.useToken()

  // 确定退出
  const navigate = useNavigate()
  const onLogout = async () => {
    // 退出登陆 删除token 跳回到登录
    const res = await http.post('/v1/logout')
    if (res.status === 200) {
      loginStore.loginOut()
      navigate('/login/')
    }
  }

  // 获取当前路径
  const { pathname } = useLocation()

  // 获取用户数据
  useEffect(() => {
    try {
      // 刷新当前页面时，重新获取用户信息
      userStore.getUserInfo()
      // 刷新页面，可防止header中的用户信息丢失
      setUserName(getLocalUser())
    } catch (error) {
      console.log(error)
    }
  }, // eslint-disable-next-line
    [])

  return (
    <Layout
      style={{
        minHeight: '100vh',
      }}
    >
      <Sider collapsible collapsed={collapsed}
        onCollapse={(value) => setCollapsed(value)}>
        <div className="demo-logo-vertical"
          style={{
            height: '64px',
            lineHeight: '48px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            textAlign: 'center',
            color: '#fff',
            fontSize: '20px',
            fontWeight: 'bold',
            backgroundColor: '#002140', // 添加背景色
          }}
        >
          {collapsed ? 'O+' : 'OMS'}
        </div>
        <Menu
          theme="dark"
          defaultSelectedKeys={[pathname]}
          selectedKeys={[pathname]}
          mode="inline"
          items={items}
        />
      </Sider>
      <Layout>
        <Header
          style={{
            padding: '0 16px',
            background: colorBgContainer,
            display: 'flex',
            justifyContent: 'flex-end',
            alignItems: 'center',
          }}
        >
          <div className="userinfo" style={{ display: 'flex', alignItems: 'center' }}>
            <Avatar icon={<UserOutlined />} style={{ marginRight: 8 }} />
            <span className="username">欢迎使用{Info.title} | {username}</span>
            <span className="user-logout" style={{ marginLeft: '8px' }}>
              <Popconfirm
                onConfirm={onLogout}
                title="是否确认退出？"
                okText="退出"
                cancelText="取消">
                <LogoutOutlined /> 退出
              </Popconfirm>
            </span>
          </div>
        </Header>

        <Layout className="layout-content" style={{ padding: 20 }}>
          {/* 二级路由出口 */}
          <Outlet />
        </Layout>
        <Footer
          style={{
            textAlign: 'center',
          }}
        >
          Dominos Design ©{new Date().getFullYear()} Created by IT
        </Footer>
      </Layout>
    </Layout>
  )
})

export default LayoutApp