import React, { useState } from 'react'
import DataTable from './table'
import DrawerApp from './drawers'
import SearchApp from './search'


const Group = () => {
  const [doc, setDoc] = useState()
  const [open, setOpen] = useState(false)

  const childRef = React.createRef()
  const onFinish = (values) => {
    const { index, database, table, rangeTimePicker, key, type } = values
    //数据处理
    const _params = {}
    if (rangeTimePicker) {
      const rangeTime = rangeTimePicker
      const startTime = rangeTime[0].valueOf()
      const endTime = rangeTime[1].valueOf()
      _params.startTime = JSON.stringify(startTime)
      _params.endTime = JSON.stringify(endTime)
    }
    if (index) {
      _params.index = index
    }
    if (database) {
      _params.database = database
    }
    if (table) {
      _params.table = table
    }
    if (key) {
      _params.key = key
    }
    if (type) {
      _params.type = type
    }
    // 修改params数据， 引起接口重新发送，对象的合并是一个整体覆盖，改了对像的整体引用
    // 调用子组件方法
    childRef.current.func(_params)
  }

  return (
    <div>
      {/* 简介 */}
      <SearchApp onFinish={onFinish} />
      <DataTable setDoc={setDoc} setOpen={setOpen} onRef={childRef} />
      <DrawerApp doc={doc} open={open} setOpen={setOpen} />
    </div>
  )
}

export default Group
