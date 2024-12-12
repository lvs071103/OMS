import React, { useEffect } from 'react'
import { Form, Input, Button, Space, message } from 'antd'
import { http } from '@/tools/http'

export default function EnvForm (props) {
  const { id, handleOk, timestamp } = props
  const formRef = React.useRef(null)

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

  const submitURL = id ? `/v1/sys/config/env/${id}` : '/v1/sys/config/env/add'

  // 提交时请求后端增加group接口
  const onFinish = async (values) => {
    // 提交数据
    const response = await http.post(submitURL, values)
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
        name="label"
        label="标签"
      >
        <Input />
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
