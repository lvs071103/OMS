import React, {
  useState,
  useEffect,
  useImperativeHandle
} from 'react'
import {
  Table,
  Tooltip,
  Space,
  Button,
  Card,
  Input,
  Flex,
  message
} from 'antd'
import {
  EditOutlined,
  DownloadOutlined
} from '@ant-design/icons'
import { http } from '@/tools/http'
import DropdownComponent from './dropdown'


const { Search } = Input


const DataTable = ({ setDoc, setOpen, onRef }) => {
  const defaultCheckedList = [
    "database",
    "table",
    "type",
    "old",
    "new",
    "action"
  ]
  const [selectedRowKeys, setSelectedRowKeys] = useState([])
  const [loading, setLoading] = useState(false)
  const [dataSource, setDataSource] = useState([]) // 新增状态来存储表格数据
  const [total, setTotal] = useState(0)
  const [checkedList, setCheckedList] = useState(defaultCheckedList)

  const columns = [
    {
      title: '_index',
      dataIndex: '_index',
      key: '_index',
    },
    {
      title: '_id',
      dataIndex: '_id',
      key: '_id',
    },
    {
      title: 'database',
      dataIndex: 'database',
      key: 'database',
      render: (_, item) => (
        <Tooltip placement='topLeft'
          title={item._source.database || '-'}
        >
          {item._source.database}
        </Tooltip>
      )
    },
    {
      title: 'table',
      dataIndex: 'table',
      key: 'table',
      render: (_, item) => (
        <Tooltip placement='topLeft'
          title={item._source.table || '-'}
        >
          {item._source.table}
        </Tooltip>
      )
    },
    {
      title: 'isDdl',
      dataIndex: 'isDdl',
      key: 'isDdl',
      render: (_, item) => (
        <Tooltip placement='topLeft'>
          {JSON.stringify(item._source.isDdl)}
        </Tooltip>
      )
    },
    {
      title: 'old',
      dataIndex: 'old',
      key: 'old',
      onCell: () => {
        return {
          style: {
            maxWidth: 500,
            overflow: 'hidden',
            whiteSpace: 'nowrap',
            textOverflow: 'ellipsis',
            cursor: 'pointer'
          }
        }
      },
      render: (_, item) => (
        <Tooltip placement='topLeft'>
          {item._source.old}
        </Tooltip>
      )
    },
    {
      title: 'type',
      dataIndex: 'type',
      key: 'type',
      render: (_, item) => (
        <Tooltip placement='topLeft'
          title={item._source.type || '-'}
        >
          {item._source.type}
        </Tooltip>
      )
    },
    {
      title: 'pkNames',
      dataIndex: 'pkNames',
      key: 'pkNames',
      render: (_, item) => (
        <Tooltip placement='topLeft'>
          {JSON.stringify(item._source.pkNames)}
        </Tooltip>
      )
    },
    {
      title: 'ts',
      dataIndex: 'ts',
      key: 'ts',
      render: (_, item) => (
        <Tooltip placement='topLeft'
          title={item._source.ts || '-'}
        >
          {item._source.ts}
        </Tooltip>
      )
    },
    {
      title: 'es',
      dataIndex: 'es',
      key: 'es',
      render: (_, item) => (
        <Tooltip placement='topLeft'
          title={item._source.es || '-'}
        >
          {item._source.es}
        </Tooltip>
      )
    },
    {
      title: 'new',
      dataIndex: 'new',
      key: 'new',
      onCell: () => {
        return {
          style: {
            maxWidth: 500,
            overflow: 'hidden',
            whiteSpace: 'nowrap',
            textOverflow: 'ellipsis',
            cursor: 'pointer'
          }
        }
      },
      render: (_, item) => (
        <Tooltip placement='topLeft'>
          {item._source.new}
        </Tooltip>
      )
    },
    {
      title: '操作',
      key: 'action',
      fixed: 'right',
      render: data => {
        return (
          <Space size="middle">
            <Button
              type="primary"
              shape="circle"
              icon={<EditOutlined />}
              onClick={() => { handlerDisplay(data) }}
            />
            <Button
              type="primary"
              shape="circle"
              icon={<DownloadOutlined />}
              onClick={() => { downloadRollBackSQL(data) }}
            />
          </Space>
        )
      }
    }
  ]

  const start = () => {
    setLoading(true)
    // ajax request after empty completing
    setTimeout(() => {
      // setSelectedRowKeys([])
      setLoading(false)
    }, 1000)
    console.log(selectedRowKeys)
    selectedRowKeys.forEach(element => {
      const data = element.split(":")
      const dataObj = {
        _index: data[0],
        _id: data[1]
      }
      console.log(dataObj)
      downloadRollBackSQL(dataObj)
    })
  }

  const onSelectChange = (newSelectedRowKeys) => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys)
    setSelectedRowKeys(newSelectedRowKeys)
  }
  const rowSelection = {
    selectedRowKeys,
    onChange: onSelectChange,
  }
  const hasSelected = selectedRowKeys.length > 0

  // 分页参数管理
  const [params, setParams] = useState({
    page: 1,
    pageSize: 10
  })

  // 切换当前页或者调整每页显示数量
  const pageChange = (page, newPageSize) => {
    setParams({
      ...params,
      page: params.pageSize !== newPageSize ? 1 : page,
      pageSize: newPageSize
    })
  }

  // 当params发生变化时，显示刷新数据
  useEffect(() => {
    const loadList = async () => {
      const res = await http.get(`/v1/es/list`, { params })
      const { results, count } = res.data
      // 为每个元素增加 key 字段
      const resultsWithKey = results.map(item => ({
        ...item,
        key: item._index + ':' + item._id, // 或者使用其他唯一标识符
      }))
      setDataSource(resultsWithKey)
      setTotal(count)
    }
    loadList()
  }, // eslint-disable-next-line
    [params])


  const handlerDisplay = (data) => {
    setDoc(data)
    setOpen(true)
  }

  const downloadRollBackSQL = async (data) => {
    try {
      const res = await http.get(`/v1/es/${data._index}/_doc/download/${data._id}`, {
        responseType: 'blob', // 确保响应类型为 blob
      })
      const blob = new Blob([res.data], { type: 'application/octet-stream' })
      const url = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = url
      link.setAttribute('download', `${data._id}.sql`) // 设置下载文件名
      document.body.appendChild(link)
      link.click()
      link.parentNode.removeChild(link)
      message.success('File downloaded successfully!')
    } catch (error) {
      console.error('Error downloading file:', error)
      message.error('Failed to download file.')
    }
  }

  const subFilter = (array, value) => {
    for (let i in array) {
      if (Object.values(array[i]).join(',').includes(value)) {
        return true
      }
    }
  }

  // 当前页搜索
  const onSearch = async (value) => {
    // 当前搜索逻辑
    const filteredData = dataSource.filter(item =>
      (item._source.new && subFilter(item._source.new, value)) ||
      (item._source.type && item._source.type.includes(value)) ||
      (item._source.database && item._source.database.includes(value)) ||
      (item._source.table && item._source.table.includes(value))
    )
    setDataSource(filteredData)
  }

  useImperativeHandle(onRef, () => {
    return {
      func: searchAction
    }
  })

  function searchAction (_params) {
    // 当提交查询时，更新params状态
    setParams({
      ...params,
      index: _params?.index,
      database: _params?.database,
      table: _params?.table,
      startTime: _params?.startTime,
      endTime: _params?.endTime,
      key: _params?.key,
      type: _params?.type,
    })
  }

  const newColumns = columns.map((item) => ({
    ...item,
    hidden: !checkedList.includes(item.key),
  }))

  return (
    <Card
      extra={
        <Space>
          <Search id='currentPageSearch' placeholder="search current page" enterButton onSearch={onSearch} />
          <DropdownComponent
            columns={columns}
            checkedList={checkedList}
            setCheckedList={setCheckedList}
          />
        </Space>
      }
    >
      <Flex gap="middle" vertical>
        <Flex align="center" gap="middle">
          <Button type="primary" onClick={start}
            disabled={!hasSelected} loading={loading}>
            批量下载
          </Button>
          {hasSelected ? `Selected ${selectedRowKeys.length} items` : null}

        </Flex>
        <Table
          // 此处便用key，是便于复选框取值_index:_id拼接而成
          rowKey={item => item.key}
          rowSelection={rowSelection}
          columns={newColumns}
          dataSource={dataSource}
          scroll={{ x: 'max-content' }}
          bordered
          pagination={
            {
              defaultCurrent: params.page,
              pageSize: params.pageSize,
              total: total,
              onChange: pageChange,
            }
          }
        />
      </Flex>
    </Card>
  )
}
export default DataTable