import React, { useState, useEffect } from 'react'
import {
  Card,
  Button,
  Table,
  Modal,
  Input,
  Space,
  Popconfirm,
  Drawer,
  message,
  Tag,
  Tooltip
} from 'antd'
import {
  EditOutlined,
  DeleteOutlined,
  UnorderedListOutlined
} from '@ant-design/icons'
// import EnvForm from './form'
import { useStore } from '@/store'
import { http } from '@/tools/http'
import Detail from './detail'

const { Search } = Input

export default function JabTable () {
  const { userStore } = useStore()
  const isSuperuser = userStore?.userInfo?.is_superuser
  const [title, setTitle] = useState('')
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState([])
  const [open, setOpen] = useState(false)
  const [jobs, setJobs] = useState({
    list: [], //组列表
    count: 0 // 组数量
  })

  // 分页参数管理
  const [params, setParams] = useState({
    page: 1,
    pageSize: 10
  })

  // 弹出添加窗口
  const showAddModal = () => {
    setTitle('添加')
    setIsModalOpen(true)
  }

  // 搜索
  const onSearch = async (q) => {
    const res = await http.get(`/v1/app/release/job/list?q=${q}`)
    const { data } = res.data
    setJobs({
      list: data,
      count: data.length,
    })
  }

  // 取消处理，关闭弹窗
  const handleCancel = () => {
    setIsModalOpen(false)
  }

  const handleOk = async () => {
    // 关闭弹窗，并更新Groups
    setIsModalOpen(false)
    flushPage()
  }

  // 刷新一下列表
  const flushPage = async () => {
    const res = await http.get(`/v1/app/release/job/list`, { params })
    const { data } = res.data
    setJobs({
      list: data?.jobs,
      count: data?.total,
    })
  }

  // 定义游标
  const [cursor, setCursor] = useState({
    id: 0,
    name: ''
  })

  // 关闭详情页
  const onClose = async () => {
    setOpen(false)
  }

  // 切换当前页或者调整每页显示数量
  const pageChange = (page, newPageSize) => {
    setParams({
      ...params,
      page: params.pageSize !== newPageSize ? 1 : page,
      pageSize: newPageSize
    })
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

  // 删除
  const delRecord = async (data, page) => {
    const resp = await http.delete(`/v1/app/release/job/${data.id}`)
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

  const showDetail = (data) => {
    setOpen(true)
    setCursor(
      {
        id: data.id,
        name: data.name
      }
    )
  }


  // 当params发生变化时，显示刷新数据
  useEffect(() => {
    const loadList = async () => {
      const res = await http.get(`/v1/app/release/job/list`, { params })
      const { data } = res.data
      console.log(data)
      setJobs({
        list: data?.jobs,
        count: data?.total
      })
    }
    loadList()
  }, // eslint-disable-next-line
    [params])


  const onSelectChange = (newSelectedRowKeys) => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys)
    setSelectedRowKeys(newSelectedRowKeys)
  }
  const rowSelection = {
    selectedRowKeys,
    onChange: onSelectChange,
  }

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '地址',
      dataIndex: 'address',
    },
    {
      title: '认证类型',
      dataIndex: 'auth_type',
      onCell: () => {
        return {
          style: {
            overflow: 'hidden',
            whiteSpace: 'nowrap',
            textOverflow: 'ellipsis',
            cursor: 'pointer'
          }
        }
      },

      render: (_, record) => (
        <Tooltip placement='topLeft'>
          {record?.auth_type === true ?
            <Tag color="blue">Auth</Tag> : <Tag color="green">已完成</Tag>
          }
        </Tooltip>
      )
    },
    {
      title: '描述',
      dataIndex: 'desc',
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
              onConfirm={() => delRecord(data, params.page)}
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
        {/* <Modal
          title={title}
          open={isModalOpen}
          onOk={handleOk}
          onCancel={handleCancel}
          width={860}
          centered
          footer={null}
        >
          <EnvForm
            handleOk={handleOk}
            title={title}
            id={title === '编辑' ? cursor.id : null}
            timestamp={new Date().getTime()}
          />
        </Modal> */}
        <Table
          rowKey={record => record.id}
          columns={columns}
          dataSource={jobs.list}
          rowSelection={rowSelection}
          scroll={{ x: 'max-content' }}
          bordered
          pagination={
            {
              defaultCurrent: params.page,
              pageSize: params.pageSize,
              total: jobs.count,
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
