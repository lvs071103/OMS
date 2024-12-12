import React from 'react'
import { Link } from 'react-router-dom'
import { observer } from 'mobx-react'
import { Card, Form, Input, Checkbox, Button, message } from 'antd'
import { useNavigate } from 'react-router-dom'
// 导入logo.png
// import logo from '@/assets/logo.png'
// 导入样式文件
import { useStore } from '@/store'
import './index.css'



const Login = observer(() => {
  const { loginStore } = useStore()
  const navigate = useNavigate()
  async function onFinish (values) {
    // 清空旧的 token
    localStorage.removeItem('oms')
    // values： 放置的是所有表单项中用户输入的内容
    // todo: login
    try {
      await loginStore.getToken({
        username: values.username,
        password: values.password
      })
      // 跳转首页
      navigate('/', { replace: true })
      // 提示用户
      message.success('登陆成功')
    } catch (error) {
      message.error('登录失败: 用户名或密码错误')
    }

  }

  function onFinishFailed (values) {
    console.log(values)
  }

  return (
    <div className='login'>
      <Card className='full-page-container'>
        {/* <img className='login-logo' src={logo} alt="" /> */}
        {/* 登录表单 */}
        {/* 子项用到的触发事件 需要在Form中都声明一下 */}
        <Form
          validateTrigger={['onBlur', 'onChange']}
          initialValues={{
            remember: true,
          }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          style={{
            maxWidth: '600px',
            margin: 'center',
            padding: '30px',
            paddingTop: '60px',
            borderRadius: '8px',
            boxShadow: '0 0 10px rgba(0, 0, 0, 0.2)',
            background: '#fff',
          }}
        >
          <Form.Item
            name="username"
            rules={[
              {
                required: true,
                message: 'Please input your username!',
              },
              {
                validateTrigger: 'onBlur'
              }
            ]}
          >
            <Input size="large" placeholder="请输入域帐号" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              {
                required: true,
                message: 'Please input your password!',
              },
              {
                min: 6,
                message: '密码长度不能低于６位'
              }
            ]}
          >
            <Input size="large" type='password' placeholder="请输入密码" autoComplete='off' />
          </Form.Item>
          <Form.Item
            name="remember"
            valuePropName="checked"
            rules={[
              {
                validator: (_, value) =>
                  value ? Promise.resolve() : Promise.reject('应该同意用户协议和隐私条款'),
              },
            ]}
          >
            <Checkbox className="login-checkbox-label">
              我已阅读并同意「用户协议」和「隐私条款」
            </Checkbox>
          </Form.Item>

          <Form.Item>
            <Link to="/signup" className="signup-link">
              还没有账号？立即注册
            </Link>
          </Form.Item>

          <Form.Item>
            {/* 渲染Button组件为submit按钮 */}
            <Button type="primary" htmlType="submit" size="large" block>
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div >
  )
})

export default Login
