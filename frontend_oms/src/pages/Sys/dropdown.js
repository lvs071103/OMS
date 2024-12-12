import React from 'react'
import { Checkbox, Popover, Button, Space } from 'antd'
import { SettingOutlined } from '@ant-design/icons'

const DropdownComponent = ({ columns, checkedList, setCheckedList }) => {
  const options = columns.map(({ key, title }) => ({
    label: title,
    value: key,
  }))

  const checkboxRender = () => {
    return (
      <Checkbox.Group
        value={checkedList}
        options={options}
        onChange={(value) => {
          setCheckedList(value)
        }}
      />
    )
  }

  return (
    <Space wrap>
      <Popover content={checkboxRender} title='调整字段' trigger="click">
        <Button type='primary' icon={<SettingOutlined />}></Button>
      </Popover>
    </Space>
  )
}

export default DropdownComponent