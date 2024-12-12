import React from 'react'
import { Tree } from 'antd'

const PermsApp = (props) => {
  const { treeData, onSelect, onCheck, selectedKeys, checkedKeys } = props
  return (
    <Tree
      checkable
      treeData={treeData}
      selectedKeys={selectedKeys}
      checkedKeys={checkedKeys}
      onSelect={onSelect}
      onCheck={onCheck}
    />
  )
}
export default PermsApp