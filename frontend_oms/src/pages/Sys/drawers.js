import React from 'react'
import { Drawer } from 'antd'
import Detail from './detail'

export default function DrawerApp ({ doc, open, setOpen, title }) {
  const onClose = () => {
    setOpen(false)
  }
  return (
    <>
      <Drawer
        title="详情"
        placement="right"
        size='large'
        open={open}
        onClose={onClose}
        width='53%'
      >
        <Detail doc={doc} />
      </Drawer>
    </>
  )
}
