const transformPermissions = (permissions) => {
  const treeData = []

  permissions.forEach(permission => {
    const { id, content_types, name } = permission
    const { id: contentTypeId, model } = content_types

    // 查找或创建父节点
    let parentNode = treeData.find(node => node.key === `p_${contentTypeId}`)
    if (!parentNode) {
      parentNode = {
        key: `p_${contentTypeId}`,
        title: model,
        children: []
      }
      treeData.push(parentNode)
    }

    // 添加子节点
    parentNode.children.push({
      key: `c_${id}`,
      title: name
    })
  })
  return treeData
}

export default transformPermissions