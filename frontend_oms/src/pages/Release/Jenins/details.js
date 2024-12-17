import { useEffect, useState } from "react"
import React from 'react'
import { http } from '@/tools/http'
import { Descriptions, Tag, Tooltip } from 'antd'


export default function Detail (props) {
  const { id, timestamp } = props
  const [details, setDetails] = useState({})

  useEffect(() => {
    const loadDetail = async () => {
      const resp = await http.get(`/v1/app/release/jenkins/${id}`)
      console.log(resp)
      const { data } = resp.data
      setDetails(data)
    }
    if (id) {
      loadDetail()
    }
  }, [id, timestamp])

  const maskPassword = (password) => {
    return password ? '*'.repeat(password.length) : ''
  }

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text)
    alert('密码已复制到剪贴板')
  }

  return (
    <>
      <Descriptions title="基本信息" layout="vertical" bordered>
        <Descriptions.Item label="ID">{details?.id}</Descriptions.Item>
        <Descriptions.Item label="环境名">{details?.name}</Descriptions.Item>
        <Descriptions.Item label="环境ID">{details?.env_id}</Descriptions.Item>
        <Descriptions.Item label="环境名">{details?.env_name}</Descriptions.Item>
        <Descriptions.Item label="验证方式">{
          details?.auth_type ?
            <Tag color="blue">Auth</Tag> :
            <Tag color="green">Token</Tag>
        }</Descriptions.Item>
        <Descriptions.Item label="用户名">{details?.username}</Descriptions.Item>
        <Descriptions.Item label="密码">
          <Tooltip title="点击复制密码">
            <span onClick={() => copyToClipboard(details?.password)}>
              {maskPassword(details?.password)}
            </span>
          </Tooltip>
        </Descriptions.Item>
        <Descriptions.Item label="地址">{details?.address}</Descriptions.Item>
        <Descriptions.Item label="描述">{details?.desc}</Descriptions.Item>
      </Descriptions>
    </>

  )
}
