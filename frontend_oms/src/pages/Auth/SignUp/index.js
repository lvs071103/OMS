import React from 'react'
import {
  Button,
  Checkbox,
  Col,
  Form,
  Input,
  message,
  Row
} from 'antd'
import './index.css'
import axios from 'axios'

const formItemLayout = {
  labelCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 6,
    },
  },
  wrapperCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 16,
    },
  },
}
const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
}
const SignUpApp = () => {
  const [form] = Form.useForm()
  const onFinish = async (values) => {
    console.log('Received values of form: ', values)
    try {
      const response = await axios.post('http://localhost:8000/api/v1/signup', values, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
      console.log(response)
      if (response.status === 200) {
        message.success('Register successfully!')
        // 跳转到登录页面
        setTimeout(() => {
          window.location.href = '/login'
        }, 1000)
      } else {
        message.error('Failed to register!')
      }
    } catch (error) {
      console.error('Error:', error)
      message.error(error.response.data.error)
    }
  }

  const getCaptcha = async () => {
    // 检查是否已经输入了邮箱
    const email = form.getFieldValue('email')
    if (!email) {
      message.error('Please input your email first!')
      return
    }
    const params = {
      "email": email
    }
    try {
      const response = await axios.post('http://localhost:8000/api/v1/captcha', params, {
        headers: {
          'Content-Type': 'application/json',
        },
      })

      if (response.status === 200) {
        message.success('Captcha has been sent to your email!')
      } else {
        message.error('Failed to send captcha!')
      }
    } catch (error) {
      console.error('Error:', error)
      message.error('Failed to send captcha!')
    }
  }


  return (
    <div className='full-page-container'>
      <Form
        {...formItemLayout}
        form={form}
        name="signup"
        onFinish={onFinish}
        style={{
          maxWidth: 600,
          margin: 'center',
          padding: '30px',
          paddingTop: '60px',
          borderRadius: '8px',
          background: '#fff',
        }}
      >

        <Form.Item
          name="username"
          label="用户名"
          tooltip="What do you want others to call you?"
          rules={[
            {
              required: true,
              message: 'Please input your username!',
              whitespace: true,
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="email"
          label="邮箱"
          rules={[
            {
              type: 'email',
              message: 'The input is not valid E-mail!',
            },
            {
              required: true,
              message: 'Please input your E-mail!',
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="password"
          label="密码"
          rules={[
            {
              required: true,
              message: 'Please input your password!',
            },
          ]}
          hasFeedback
        >
          <Input.Password />
        </Form.Item>

        <Form.Item
          name="confirm_password"
          label="确认密码"
          dependencies={['password']}
          hasFeedback
          rules={[
            {
              required: true,
              message: 'Please confirm your password!',
            },
            ({ getFieldValue }) => ({
              validator (_, value) {
                if (!value || getFieldValue('password') === value) {
                  return Promise.resolve()
                }
                return Promise.reject(new Error('The new password that you entered do not match!'))
              },
            }),
          ]}
        >
          <Input.Password />
        </Form.Item>


        <Form.Item label="Captcha">
          <Row gutter={8}>
            <Col span={12}>
              <Form.Item
                name="verify_code"
                noStyle
                rules={[
                  {
                    required: true,
                    message: 'Please input the captcha you got!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Button onClick={getCaptcha}>Get captcha</Button>
            </Col>
          </Row>
        </Form.Item>

        <Form.Item
          name="agreement"
          valuePropName="checked"
          rules={[
            {
              validator: (_, value) =>
                value ? Promise.resolve() : Promise.reject(new Error('Should accept agreement')),
            },
          ]}
          {...tailFormItemLayout}
        >
          <Checkbox>
            I have read the agreement
          </Checkbox>
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          <Button type="primary" htmlType="submit">
            Register
          </Button>
        </Form.Item>
      </Form>
    </div>

  )
}
export default SignUpApp