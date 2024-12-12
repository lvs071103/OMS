import { useEffect, useState } from "react"
import React from 'react'
import { http } from '@/tools/http'
import { Descriptions } from 'antd'
import transformPermissions from "@/tools/handleTreeData"
import { Tree } from 'antd'
import { DownOutlined } from '@ant-design/icons'


export default function Detail (props) {
  const { id, timestamp } = props
  const [details, setDetails] = useState({})
  const [treeData, setTreeData] = useState([])

  useEffect(() => {
    const loadDetail = async () => {
      const resp = await http.get(`/v1/group/${id}`)
      const { data } = resp.data
      const { permissions } = data
      if (!permissions) {
        setTreeData([])
      } else {
        setTreeData(transformPermissions(permissions))
      }
      setDetails(data)
    }
    if (id) {
      loadDetail()
    }
  }, [id, timestamp])

  const onSelect = (selectedKeys, info) => {
    console.log('selected', selectedKeys, info)
  }

  return (
    <>
      <Descriptions title="基本信息" layout="vertical" bordered>
        <Descriptions.Item label="组编号">{details?.id}</Descriptions.Item>
        <Descriptions.Item label="组名">{details?.name}</Descriptions.Item>
        <Descriptions.Item label="描述">{details?.desc}</Descriptions.Item>
      </Descriptions>
      <br />
      {/* 权限 */}
      <Descriptions layout="vertical" column={1} bordered>
        <Descriptions.Item label="权限">
          <Tree
            showLine
            switcherIcon={<DownOutlined />}
            // 展开所有节点
            defaultExpandAll
            // 点击树节点触发
            onSelect={onSelect}
            // treedata 传入的是一个数组
            treeData={treeData}
          />
        </Descriptions.Item>
      </Descriptions>
    </>

  )
}
