import { useEffect, useState } from "react"
import React from 'react'
import { http } from '@/tools/http'
import { Descriptions } from 'antd'
import transformPermissions from "@/tools/handleTreeData"
import { Tree } from 'antd'
import { DownOutlined } from '@ant-design/icons'
import moment from 'moment'


export default function Detail (props) {
  const { id, timestamp } = props
  const [details, setDetails] = useState({})
  const [treeData, setTreeData] = useState([])

  useEffect(() => {
    const loadDetail = async () => {
      const resp = await http.get(`/v1/user/${id}`)
      console.log(resp)
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

  // 将密码转换为 * 号
  const maskPassword = (password) => {
    return password ? '*'.repeat(password.length) : ''
  }

  return (
    <>
      <Descriptions title="基本信息" layout="vertical" bordered>
        <Descriptions.Item label="编号">{details?.id}</Descriptions.Item>
        <Descriptions.Item label="用户名">{details?.username}</Descriptions.Item>
        <Descriptions.Item label="姓">{details?.last_name ? details.last_name : '-'}</Descriptions.Item>
        <Descriptions.Item label="名">{details?.first_name ? details?.first_name : '-'}</Descriptions.Item>
        <Descriptions.Item label="管理员">{details?.is_superuser ? '是' : '否'}</Descriptions.Item>
        <Descriptions.Item label="是否员工">{details?.is_staff ? '是' : '否'}</Descriptions.Item>
        <Descriptions.Item label="是否激活">{details?.is_active ? '是' : '否'}</Descriptions.Item>
        <Descriptions.Item label="职位">{details?.job ? details.job : '-'}</Descriptions.Item>
        <Descriptions.Item label="邮箱">{details?.email ? details?.email : '-'}</Descriptions.Item>
        <Descriptions.Item label="性别">{details?.gender === 1 ? "男" : "女"}</Descriptions.Item>
        <Descriptions.Item label="年龄">{details?.age ? details.age : '-'}</Descriptions.Item>
        <Descriptions.Item label="加入时间">{
          moment(details?.date_joined).format('YYYY-MM-DD HH:mm:ss')
        }</Descriptions.Item>
        <Descriptions.Item label="最近登陆">{
          moment(details?.last_login).format('YYYY-MM-DD HH:mm:ss')
        }</Descriptions.Item>
      </Descriptions>

      <Descriptions layout="vertical" column={1} bordered>
        <Descriptions.Item label="描述">{details?.desc?.valid ? details.String : '-'}</Descriptions.Item>
        <Descriptions.Item label="组">
          {details?.groups ? details?.groups?.map((item) => item.name).join(',') : '-'}
        </Descriptions.Item>
        {/* 密码显示星号 */}
        <Descriptions.Item label="密码">{maskPassword(details?.password)}</Descriptions.Item>
        <Descriptions.Item label="地址">{details?.address}</Descriptions.Item>
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
            treeData={treeData ? treeData : []}
          />
        </Descriptions.Item>
      </Descriptions>
    </>

  )
}
