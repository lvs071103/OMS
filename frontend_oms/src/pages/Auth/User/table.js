import { http } from '@/tools/http'
import { DeleteOutlined, EditOutlined, UnorderedListOutlined } from '@ant-design/icons'
import {
  Button,
  Card,
  Input,
  message,
  Modal,
  Popconfirm,
  Space,
  Table,
  Tag,
  Drawer
} from 'antd'
import 'moment/locale/zh-cn'
import { useEffect, useState } from 'react'
import { useStore } from '@/store'
import UserForm from './form'
import moment from 'moment'
import Detail from './detail'


const { Search } = Input

const User = () => {

  // 用户列表管理 统一管理数据 将来修入给setList传对象
  const [users, setUsers] = useState({
    list: [], //用户列表
    count: 0 // 用户数量
  })

  const [selectedRowKeys, setSelectedRowKeys] = useState([])
  const [open, setOpen] = useState(false)

  // 分页参数管理
  const [params, setParams] = useState({
    page: 1,
    pageSize: 10
  })

  // 定义游标
  const [cursor, setCursor] = useState({
    id: 0,
    name: ''
  })

  const { userStore } = useStore()
  const isSuperuser = userStore?.userInfo?.is_superuser
  // const currentUser = localStorage.getItem('username')

  const [isModalOpen, setIsModalOpen] = useState(false)
  const [title, setTitle] = useState('')
  const [row, setRow] = useState({})

  // 如果异步请求函数需要依赖一些数据的变化而重新执行
  // 推荐把它写在内部
  // 统一不抽离函数到外面 只要涉及到异步请求的函数 都放到useEffect内部
  // 本质区别：写到外面每次组件更新都会重新进行函数初始化 这本身就是一次性能消耗
  // 而写到useEffect中，只会在依赖项发生变化的时候 函数才会进行重新初始化
  // 避免性能损失

  const reqURL = '/v1/user/list'

  useEffect(() => {
    const loadList = async () => {
      const res = await http.get(reqURL, { params })
      if (res.status !== 200) {
        message.error('获取用户列表失败, 报错：' + res.response.data.error)
      } else {
        const { data } = res
        setUsers({
          list: data.users,
          count: data.total
        })
      }
    }
    loadList()
  },// eslint-disable-next-line 
    [params])


  // 弹出添加窗口
  const showAddModal = () => {
    setTitle('添加')
    setIsModalOpen(true)
  }

  // 切换当前页触发
  const pageChange = (page, newPageSize) => {
    setParams({
      ...params,
      page: params.pageSize !== newPageSize ? 1 : page,
      pageSize: newPageSize
    })
  }

  // 删除
  const delUser = async (data, page) => {
    const resp = await http.delete(`/v1/user/${data.id}`)
    if (resp.status === 200) {
      message.success("删除成功")
      // 刷新一下列表
      setParams({
        ...params,
        page
      })
    } else {
      message.error("删除失败")
    }

  }

  // 编辑
  const goUpdate = (data) => {
    setTitle('编辑')
    setIsModalOpen(true)
    setRow(data)
  }

  //弹出详情窗口
  const showDetail = (data) => {
    setOpen(true)
    setCursor(
      {
        id: data.id,
        name: data.name
      }
    )
  }

  const handleOk = async () => {
    // 关闭弹窗，并更新users
    setIsModalOpen(false)
    const resp = await http.get(reqURL, { params })
    const { data } = resp
    setUsers(
      {
        list: data.users,
        count: data.total
      }
    )
  }

  const handleCancel = () => {
    setIsModalOpen(false)
  }

  const onSelectChange = (newSelectedRowKeys) => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys)
    setSelectedRowKeys(newSelectedRowKeys)
  }
  const rowSelection = {
    selectedRowKeys,
    onChange: onSelectChange,
  }

  // 搜索
  const onSearch = async (data) => {
    const res = await http.get(`${reqURL}?q=${data}`)
    setUsers({
      list: res.data.data,
      count: res.count
    })
  }

  // 关闭详情页
  const onClose = async () => {
    setOpen(false)
  }

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: '用户名',
      dataIndex: 'username',
    },
    {
      title: '加入时间',
      dataIndex: 'date_joined',
      render: (text, _) => {
        return moment(text).format('YYYY-MM-DD HH:mm:ss')
      }
    },
    {
      title: '最近登陆',
      dataIndex: 'last_login',
      render: (text, _) => {
        return text ? moment(text).format('YYYY-MM-DD HH:mm:ss') : '-'
      }
    },
    {
      title: '邮箱',
      dataIndex: 'email'
    },
    {
      title: '管理员',
      dataIndex: 'is_superuser',
      render: (text, _) => {
        return `${text}` === 'true' ?
          <Tag color="green">是</Tag> :
          <Tag color="yellow">否</Tag>
      }
    },
    {
      title: '员工',
      dataIndex: 'is_staff',
      render: (text, _) => {
        return `${text}` === 'true' ?
          <Tag color="green">是</Tag> :
          <Tag color="yellow">否</Tag>
      }
    },
    {
      title: '状态',
      dataIndex: 'is_active',
      render: (text, _) => {
        return `${text}` === 'true' ?
          <Tag color="green">激活</Tag> :
          <Tag color="yellow">未激活</Tag>
      }
    },
    {
      title: '操作',
      render: data => {
        return (
          <Space size="middle">
            {isSuperuser ? <Button
              type="primary"
              shape="circle"
              icon={<EditOutlined />}
              onClick={() => goUpdate(data)}
            /> : null}
            {isSuperuser ? <Popconfirm
              title="确定删除该条记录吗?"
              onConfirm={() => delUser(data, params.page)}
              okText="确认"
              cancelText="取消"
            >
              <Button
                type="primary"
                danger
                shape="circle"
                icon={<DeleteOutlined />}
              />
            </Popconfirm> : null}
            <Button
              type="primary"
              ghost
              shape="circle"
              icon={<UnorderedListOutlined />}
              onClick={() => { showDetail(data) }}
            />
          </Space>
        )
      }
    }
  ]

  return (
    <div>
      <Card
        title={
          <Button type="primary" onClick={showAddModal}>Add</Button>
        }
        extra={
          <Search id='userSearch' placeholder="input search text" onSearch={onSearch} enterButton />
        }
      >
        <Modal
          title={title}
          open={isModalOpen}
          onOk={handleOk}
          onCancel={handleCancel}
          width={860}
          centered
          footer={null}
        >
          <UserForm
            handleOk={handleOk}
            title={title}
            id={title === '编辑' || title === '详情' ? row.id : null}
          />
        </Modal>
        <Table
          rowKey={row => (row.id)}
          columns={columns}
          dataSource={users.list}
          rowSelection={rowSelection}
          scroll={{ x: 'max-content' }}
          bordered
          pagination={
            {
              defaultCurrent: params.page,
              pageSize: params.pageSize,
              total: users.count,
              onChange: pageChange
            }
          }
        />

        {/* 详情页面 */}
        <Drawer
          title="详情"
          placement="right"
          size='large'
          open={open}
          onClose={onClose}
          width='53%'
        >
          <Detail id={cursor.id} timestamp={new Date().getTime()} />
        </Drawer>
      </Card>
    </div>
  )
}

export default User