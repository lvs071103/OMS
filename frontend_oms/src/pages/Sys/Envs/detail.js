import { useEffect, useState } from "react"
import React from 'react'
import { http } from '@/tools/http'
import { Descriptions } from 'antd'


export default function Detail (props) {
  const { id, timestamp } = props
  const [details, setDetails] = useState({})

  useEffect(() => {
    const loadDetail = async () => {
      const resp = await http.get(`/v1/sys/config/env/${id}`)
      const { data } = resp.data
      setDetails(data)
    }
    if (id) {
      loadDetail()
    }
  }, [id, timestamp])

  return (
    <>
      <Descriptions title="基本信息" layout="vertical" bordered>
        <Descriptions.Item label="ID">{details?.id}</Descriptions.Item>
        <Descriptions.Item label="环境名">{details?.name}</Descriptions.Item>
        <Descriptions.Item label="标签">{details?.label}</Descriptions.Item>
        <Descriptions.Item label="描述">{details?.desc}</Descriptions.Item>
      </Descriptions>
    </>

  )
}
