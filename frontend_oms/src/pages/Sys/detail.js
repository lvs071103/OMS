import React, { useEffect, useState } from 'react'
import {
  Descriptions,
  Button,
  message
} from 'antd'
import { http } from '@/tools/http'
import JSONInput from 'react-json-editor-ajrm'
import locale from 'react-json-editor-ajrm'
import { CopyToClipboard } from 'react-copy-to-clipboard'


export default function Detail ({ doc }) {
  const [detail, setDetail] = useState()
  useEffect(() => {
    const reqURL = `/v1/es/${doc._index}/_doc/detail/${doc._id}`
    const loadDetail = async () => {
      const res = await http.get(reqURL)
      setDetail(JSON.stringify(res.data, null, 2)) // 格式化 JSON 数据
    }
    if (doc._id) {
      loadDetail()
    }
  }, [doc._id, doc._index])

  return (
    <>
      <Descriptions layout="vertical" bordered>
        <Descriptions.Item label="_index">
          {doc._index}
        </Descriptions.Item>
        <Descriptions.Item label="_id">
          {doc._id}
        </Descriptions.Item>
      </Descriptions>
      <br />
      <Descriptions layout="vertical" column={1} bordered>
        <Descriptions.Item label="JSON Data">
          {detail ? (
            <>
              <CopyToClipboard text={detail} onCopy={() => message.success('JSON copied to clipboard!')}>
                <Button type="primary" style={{ marginBottom: '10px' }}>Copy JSON</Button>
              </CopyToClipboard>
              <JSONInput
                id="json-editor"
                placeholder={JSON.parse(detail)}
                locale={locale}
                height="auto"
                width="100%"
                viewOnly={true} // 设置为只读模式
                style={{
                  body: {
                    fontSize: '14px'
                  }
                }}
              />
            </>
          ) : (
            'Loading...'
          )}
        </Descriptions.Item>
      </Descriptions>
    </>
  )
}
