import React, { useEffect } from 'react'
import { Form, Input, Button, Space, message, Select, Switch } from 'antd'
import { http } from '@/tools/http'

export default function InstancesForm (props) {
  const { id, handleOk, timestamp } = props
  const formRef = React.useRef(null)
  const [envs, setEnvs] = React.useState([])

  const layout = {
    labelCol: {
      span: 5,
    },
    wrapperCol: {
      span: 16,
    },
  }

  const tailLayout = {
    wrapperCol: {
      offset: 5,
      span: 16,
    },
  }

  const getEnvs = async () => {
    const res = await http.get('/v1/sys/config/env/list')
    const { envs } = res.data
    const envList = envs?.map(item => ({
      label: item.name,
      value: item.id
    }))
    setEnvs(envList)
  }

  const submitURL = id ? `/v1/app/release/jenkins/${id}` : '/v1/app/release/jenkins/add'

  // 提交时请求后端增加group接口
  const onFinish = async (values) => {
    const formData = {
      ...values,
      auth_type: values.auth_type ? 1 : 0, // 将 Switch 的值转换为 1 或 0
    }
    // 提交数据
    const response = await http.post(submitURL, formData)
    if (response.status === 200) {
      message.success('操作成功')
      // 关闭弹窗
      handleOk()
    } else {
      message.error('操作失败')
    }
  }

  // 重置form表单
  const onReset = () => {
    formRef.current.resetFields()
  }

  // 编辑回填数据
  useEffect(() => {
    // 加载详情
    const loadDetail = async () => {
      const res = await http.get(submitURL)
      console.log(res)
      const { data } = res.data
      console.log(data)
      setEnvs(data.envs)
      // 回填数据至表单
      formRef.current.setFieldsValue({
        name: data?.name,
        label: data?.label,
        desc: data?.desc,
      })
    }

    if (id) {
      // 编辑时加载详情回填数据
      loadDetail()
    } else {
      // 重置表单
      onReset()
    }
  },
    // eslint-disable-next-line 
    [id, formRef, timestamp])

  return (
    <Form
      {...layout}
      ref={formRef}
      name="control-ref"
      onFinish={onFinish}
    >
      <Form.Item
        name="name"
        label="名称"
        rules={[
          {
            required: true,
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="address"
        label="地址"
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="选择环境"
        name="env_id"
        rul
      >
        <Select
          placeholder="选择环境"
          onDropdownVisibleChange={getEnvs}
          showSearch
          optionFilterProp="children"
          filterOption={(input, option) => (option?.label ?? '').includes(input)}
          filterSort={(optionA, optionB) =>
            (optionA?.label ?? '').toLowerCase().localeCompare((optionB?.label ?? '').toLowerCase())
          }
          options={envs}
        />
      </Form.Item>


      <Form.Item
        name="auth_type"
        label="认证类型"
        valuePropName="checked"
        getValueFromEvent={(checked) => checked ? 'Auth' : 'Token'}
      >
        <Switch checkedChildren="Auth" unCheckedChildren="Token" />
      </Form.Item>

      <Form.Item
        name="username"
        label="用户名"
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="password"
        label="密码"
      >
        <Input type='password' />
      </Form.Item>

      <Form.Item
        name="desc"
        label="描述"
      >
        <Input />
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Space>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
          <Button htmlType="button" onClick={onReset}>
            Reset
          </Button>
        </Space>
      </Form.Item>
    </Form>
  )
}
