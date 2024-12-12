import { http } from '@/tools/http'
import { Button, Form, Input, message, Space } from 'antd'
import React, { useEffect, useState } from 'react'
import PermsApp from './perms'
import transformPermissions from '@/tools/handleTreeData'


const GroupForm = (props) => {

  const { id, handleOk, timestamp } = props
  const formRef = React.useRef(null)
  const [treeData, setTreeData] = useState([])
  const [selectedKeys, setSelectedKeys] = useState([])
  const [checkedKeys, setCheckedKeys] = useState([])

  const handleSelect = (keys) => {
    setSelectedKeys(keys) // 更新状态以存储选中值
  }

  const handleCheck = (keys) => {
    setCheckedKeys(keys) // 更新状态以存储选中值
  }

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

  const handlePerms = (checkedKeys) => {
    const permission_ids = []
    if (checkedKeys === undefined || checkedKeys.length === 0) {
      return permission_ids
    } else {
      checkedKeys.forEach((item) => {
        if (item.includes('p')) {
          // 如果含有p则为全选, 不添加到permission_ids中
          return
        } else {
          // 去除c_前缀，添加到permission_ids中
          permission_ids.push(item.replace('c_', ''))
        }
      })
    }
    return permission_ids
  }

  const submitURL = id ? `/v1/group/${id}` : '/v1/group/add'

  // 提交时请求后端增加group接口
  const onFinish = async (values) => {
    // 将选中值添加到表单提交的数据中
    const perms = handlePerms(checkedKeys)
    console.log(perms)
    const formData = {
      ...values,
      permissions: perms
    }
    console.log(formData)
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
    setSelectedKeys([])
    setCheckedKeys([])
  }

  // 编辑回填数据
  useEffect(() => {
    // 加载权限列表
    const loadPermissions = async () => {
      const res = await http.get('/v1/permission/list')
      const permissions = res.data.data
      // 将权限列表转换为树形结构
      const treeData = transformPermissions(permissions)
      setTreeData(treeData)
    }
    // 加载详情
    const loadDetail = async () => {
      const res = await http.get(`/v1/group/${id}`)
      const { data } = res.data
      // 获取已选权限列表
      const permsSelectdList = data.permissions?.map((item) => item.id)
      // 选中的权限加上前缀
      const selectedKeyList = permsSelectdList?.map((item) => `c_${item}`)
      setCheckedKeys(selectedKeyList)
      // 回填数据至表单
      formRef.current.setFieldsValue({
        name: data?.name,
        desc: data?.desc,
        permissions: selectedKeys
      })
    }

    if (id) {
      // 加载权限列表
      loadPermissions()
      // 编辑时加载详情回填数据
      loadDetail()
    } else {
      // 新增时加载权限列表
      loadPermissions()
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
        label="组名"
        rules={[
          {
            required: true,
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="desc"
        label="描述"
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="permissions"
        label="权限"
      >
        <PermsApp
          treeData={treeData}
          selectedKeys={selectedKeys}
          checkedKeys={checkedKeys}
          onSelect={handleSelect}
          onCheck={handleCheck}
        />
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

export default GroupForm