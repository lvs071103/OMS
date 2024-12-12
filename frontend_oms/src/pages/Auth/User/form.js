import React, { useState, useEffect } from 'react'
import {
  Form,
  Button,
  Space,
  Input,
  Select,
  Checkbox,
  message,
  Switch,
  InputNumber
} from 'antd'
import 'moment/locale/zh-cn'
import { http } from '@/tools/http'
import moment from 'moment'
import PermsApp from '../Group/perms'
import transformPermissions from '@/tools/handleTreeData'

const UserForm = (props) => {
  const { id, handleOk, timestamp } = props
  const form = React.useRef(null)
  const [superuserchecked, setSuperuserChecked] = useState(false)
  const [staffchecked, setStaffChecked] = useState(false)
  const [activechecked, setActiveChecked] = useState(false)
  const [treeData, setTreeData] = useState([])
  const [selectedKeys, setSelectedKeys] = useState([])
  const [checkedKeys, setCheckedKeys] = useState([])
  const [value, setValue] = useState([])
  const [groups, setGroups] = React.useState([])

  const handleSelect = (keys) => {
    setSelectedKeys(keys) // 更新状态以存储选中值
  }

  const handleCheck = (keys) => {
    setCheckedKeys(keys) // 更新状态以存储选中值
  }

  // 获取用户组列表
  const getGroups = async () => {
    const res = await http.get('/v1/group/list')
    const { groups } = res.data.data
    const groupList = groups.map(item => ({
      label: item.name,
      value: item.id
    }))
    setGroups(groupList)
  }

  const handlePerms = (checkedKeys) => {
    const permission_ids = []
    // 如果没有选中任何权限，则返回空列表
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

  const groupsSelectProps = {
    mode: 'multiple',
    style: { width: '100%' },
    value,
    options: groups,
    placeholder: 'Please select',
    maxTagCount: 'responsive',
    onChange: (newValue) => {
      setValue(newValue)
    }
  }

  const onReset = () => {
    form.current.resetFields()
    setSelectedKeys([])
    setCheckedKeys([])
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

  const submitURL = id ? `/v1/user/${id}` : '/v1/user/add'

  // 提交时请求后端增加group接口
  const onFinish = async (values) => {
    const perms = handlePerms(checkedKeys)
    const formData = {
      ...values,
      gender: values.gender ? 1 : 0,
      permissions: perms
    }
    console.log(formData)
    const resp = await http.post(submitURL, formData)
    if (resp.status === 200) {
      message.success('操作成功')
      // 关闭弹窗
      handleOk()
    } else {
      message.error('操作失败: ' + resp.data.error)
    }
  }

  useEffect(() => {
    const loadPrems = async () => {
      const res = await http.get('/v1/permission/list')
      const permissions = res.data.data
      setTreeData(transformPermissions(permissions))
    }

    const loadDetail = async () => {
      const res = await http.get(`/v1/user/${id}`)
      // user.groups必须是一个列表，里面就是一个id，不需要对象
      const users = res.data.data
      // console.log(users)
      // 后端传过来的是一个对象，需要转换成列表
      const user_perms = users?.permissions?.map(item => item.id)
      const selectedKeyList = user_perms?.map((item) => `c_${item}`)
      const groups_list = users?.groups?.map(item => item.id)
      setCheckedKeys(selectedKeyList)
      form.current.setFieldsValue({
        username: users.username,
        password: users.password,
        permissions: selectedKeys,
        groups: groups_list,
        is_superuser: users.is_superuser,
        is_active: users.is_active,
        is_staff: users.is_staff,
        first_name: users.first_name,
        last_name: users.last_name,
        last_login: users.last_login ? moment(users.last_login) : null,
        email: users.email,
        date_joined: users.date_joined ? moment(users.date_joined) : null,
        age: users.age,
        gender: users.gender,
        job: users.job,
        address: users.address
      })
    }


    if (id) {
      // 加载权限列表
      loadPrems()
      getGroups()
      loadDetail()
    } else {
      onReset()
      // 加载权限列表
      loadPrems()
      getGroups()
    }
  }, // eslint-disable-next-line
    [id, timestamp])

  return (
    <div>
      <Form
        ref={form}
        name="control-ref"
        onFinish={onFinish}
        defaultinitialvalues={{
          is_superuser: superuserchecked,
          is_staff: staffchecked,
          is_active: activechecked
        }}
        {...layout}
      >
        <Form.Item
          name="username"
          label="用户名"
          rules={[
            {
              required: true,
              max: 150,
              message: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
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
              message: 'Please input your password!'
            },
            {
              min: 6,
              max: 128,
              message: 'Password  must be minimum 6 characters.'
            },
          ]}
        >
          {id ?
            <Input.Password autoComplete="off" readOnly /> :
            <Input.Password autoComplete="off" />}
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

        <Form.Item
          name="groups"
          label="用户组"
        >
          <Select
            onDropdownVisibleChange={getGroups}
            {...groupsSelectProps}
          />
        </Form.Item>

        <Form.Item
          name="is_superuser"
          label="管理员"
          valuePropName="checked"
        >
          <Checkbox onChange={(e) => { setSuperuserChecked(e.target.checked) }}></Checkbox>
        </Form.Item>

        <Form.Item
          name="is_staff"
          label="员工"
          valuePropName="checked"
        >
          <Checkbox onChange={(e) => { setStaffChecked(e.target.checked) }}></Checkbox>
        </Form.Item>

        <Form.Item
          name="is_active"
          label="激活"
          valuePropName="checked"
        >
          <Checkbox onChange={(e) => { setActiveChecked(e.target.checked) }}></Checkbox>
        </Form.Item>

        <Form.Item
          name="first_name"
          label="名"
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="last_name"
          label="姓"
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="email"
          label="邮箱"
          rules={[
            {
              type: 'email',
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="age"
          label="年龄"
          rules={[
            {
              type: 'number',
              min: 0,
              max: 120,
              message: '请输入有效的年龄',
            },
          ]}
        >
          <InputNumber min={0} max={120} />
        </Form.Item>

        <Form.Item
          name="gender"
          label="性别"
          valuePropName="checked"
          getValueFromEvent={(checked) => (checked ? 1 : 0)}
        >

          <Switch checkedChildren="男" unCheckedChildren="女" />
        </Form.Item>

        <Form.Item
          name="job"
          label="职业"
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="address"
          label="地址"
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
    </div>
  )
}

export default UserForm
