import React, { useEffect, useState } from 'react'
import {
  Card,
  Button,
  Table,
  Modal,
  Input,
  Space,
  Popconfirm,
  Drawer,
  message
} from 'antd'
import {
  EditOutlined,
  DeleteOutlined,
  UnorderedListOutlined
} from '@ant-design/icons'
import { useStore } from '@/store'
import GroupForm from './form'
import { http } from '@/tools/http'
import Detail from './detail'



const { Search } = Input

export default function GroupTable () {

  const { userStore } = useStore()
  const isSuperuser = userStore?.userInfo?.is_superuser

  const [groups, setGroups] = useState({
    list: [], //组列表
    count: 0 // 组数量
  })

  const [open, setOpen] = useState(false)

  // 分页参数管理
  const [params, setParams] = useState({
    page: 1,
    pageSize: 10
  })

  // 定义标题
  const [title, setTitle] = useState('')

  // 定义游标
  const [cursor, setCursor] = useState({
    id: 0,
    name: ''
  })

  // 当params发生变化时，显示刷新数据
  useEffect(() => {
    const loadList = async () => {
      const res = await http.get(`/v1/group/list`, { params })
      console.log(res.data)
      const { data } = res.data
      setGroups({
        list: data?.groups,
        count: data?.total
      })
    }
    loadList()
  }, // eslint-disable-next-line
    [params])


  // 切换当前页或者调整每页显示数量
  const pageChange = (page, newPageSize) => {
    setParams({
      ...params,
      page: params.pageSize !== newPageSize ? 1 : page,
      pageSize: newPageSize
    })
  }

  // 删除
  const delGroup = async (data, page) => {
    const resp = await http.delete(`/v1/group/${data.id}`)
    if (resp.status === 200) {
      message.success('删除成功')
    } else {
      message.error('删除失败')
    }
    // 刷新一下列表
    setParams({
      ...params,
      page
    })
  }

  const [isModalOpen, setIsModalOpen] = useState(false)

  // 弹出添加窗口
  const showAddModal = () => {
    setTitle('添加')
    setIsModalOpen(true)
  }

  // 弹出编辑窗口
  const showEditModal = (data) => {
    setTitle('编辑')
    setIsModalOpen(true)
    setCursor(
      {
        id: data.id,
        name: data.name
      }
    )
  }

  const showDetail = (data) => {
    setOpen(true)
    setCursor(
      {
        id: data.id,
        name: data.name
      }
    )
  }

  // 刷新一下列表
  const flushPage = async () => {
    const res = await http.get(`/v1/group/list`, { params })
    const { data } = res.data
    // 刷新当前页面
    // window.location.reload()
    setGroups({
      list: data.groups,
      count: data.total,
    })
  }

  const handleOk = async () => {
    // 关闭弹窗，并更新Groups
    setIsModalOpen(false)
    flushPage()
  }

  // 搜索
  const onSearch = async (data) => {
    const res = await http.get(`/v1/group/list?q=${data}`)
    const { groups, count } = res.data
    setGroups({
      list: groups,
      count: count,
    })
  }

  // 关闭详情页
  const onClose = async () => {
    setOpen(false)
  }


  const handleCancel = () => {
    setIsModalOpen(false)
  }

  const columns = [
    {
      title: 'id',
      dataIndex: 'id',
    },
    {
      title: '用户组名',
      dataIndex: 'name',
    },
    {
      title: '权限',
      key: 'permissions',
      dataIndex: 'permissions',
      // 调整列宽，超出部分隐藏并用...显示
      width: '55%',
      onCell: () => {
        return {
          style: {
            maxWidth: 100,
            overflow: 'hidden',
            whiteSpace: 'nowrap',
            textOverflow: 'ellipsis',
            cursor: 'pointer',
          },
        }
      },
      // 显示一列数组值的显示
      render: (_, record) => (
        <>
          {record?.permissions?.map((item) => {
            return (
              <span key={item.id}> {item.name} </span>
            )
          })}
        </>
      )
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
              onClick={() => { showEditModal(data) }}
            /> : null}

            {isSuperuser ? <Popconfirm
              title="确认删除该条记录吗?"
              onConfirm={() => delGroup(data, params.page)}
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
          <Search id='groupSearch' placeholder="input search text" onSearch={onSearch} enterButton />
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
          <GroupForm
            handleOk={handleOk}
            title={title}
            id={title === '编辑' ? cursor.id : null}
            timestamp={new Date().getTime()}
          />
        </Modal>
        <Table
          rowKey={record => record.id}
          columns={columns}
          dataSource={groups.list}
          scroll={{ x: 'max-content' }}
          bordered
          pagination={
            {
              defaultCurrent: params.page,
              pageSize: params.pageSize,
              total: groups.count,
              onChange: pageChange,
              className: "pagination",
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
