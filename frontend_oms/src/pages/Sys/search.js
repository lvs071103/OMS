import React from 'react'
import {
  Breadcrumb,
  Button,
  Card,
  Form,
  Input,
  Row,
  Col,
  DatePicker,
  Select
} from 'antd'
import { Link } from 'react-router-dom'
import { http } from '@/tools/http'

const { RangePicker } = DatePicker

export default function SearchApp (props) {
  const { onFinish } = props

  const [indices, setIndices] = React.useState([])

  const getIndices = async () => {
    const res = await http.get('/v1/es/_cat/indices')
    const indices = res.data
    const idx_list = indices.map((item) => {
      return {
        value: item.index,
        label: item.index
      }
    })
    setIndices(idx_list)
  }

  const handleClear = () => {
    console.log('Selection cleared!')
  }

  // const handleChangeIdx = async (index) => {
  //   if (index === undefined) {
  //     return
  //   }
  //   console.log(index)
  // }

  const rangeConfig = {
    rules: [
      {
        type: 'array',
        message: 'Please select time!',
      },
    ],
  }

  return (
    <Card
      title={
        <Breadcrumb items={[
          { title: <Link to='/'>首页</Link> },
          { title: 'SQL检索' }, { title: '操作列表' }]}
        />
      }
      style={{ marginBottom: 20 }}
    >
      <Form
        onFinish={onFinish}
      >
        <Row>
          <Col flex="auto">
            <Form.Item
              label="索引"
              name="index"
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <Select
                allowClear
                placeholder="------------请选择索引------------"
                showSearch
                optionFilterProp="children"
                onDropdownVisibleChange={getIndices}
                // onChange={handleChangeIdx}
                onClear={handleClear}  // 处理清除操作
                filterOption={(input, option) => (option?.label ?? '').includes(input)}
                filterSort={(optionA, optionB) =>
                  (optionA?.label ?? '').toLowerCase().localeCompare((optionB?.label ?? '').toLowerCase())
                }
                options={indices}
              />
            </Form.Item>


            <Form.Item
              name="rangeTimePicker"
              label="时间范围"
              {...rangeConfig}
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <RangePicker showTime format="YYYY-MM-DD HH:mm:ss" />
            </Form.Item>
          </Col>
        </Row>

        <Row>
          <Col flex="auto">
            <Form.Item
              label="database"
              name="database"
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <Input placeholder="database name" autoComplete='on' style={{ width: '400px' }} />
            </Form.Item>

            <Form.Item
              label="table"
              name="table"
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <Input placeholder="table name" autoComplete='on' style={{ width: '400px' }} />
            </Form.Item>
          </Col>
        </Row>

        <Row>
          <Col flex="auto">
            <Form.Item
              label="类型"
              name="type"
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <Input placeholder="UPDATE or INSERT or DELETE" autoComplete='on' style={{ width: '400px' }} />
            </Form.Item>

            <Form.Item
              label="关键字"
              name="key"
              style={{ display: 'inline-flex', width: 'calc(45% - 4px)' }}
            >
              <Input placeholder="table name" autoComplete='on' style={{ width: '400px' }} />
            </Form.Item>
          </Col>
        </Row>

        <Form.Item>
          <Button type="primary" htmlType="submit" style={{ marginLeft: 80 }}>
            筛选
          </Button>
        </Form.Item>
      </Form>
    </Card>
  )
}
