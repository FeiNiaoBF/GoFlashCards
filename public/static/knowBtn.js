document.getElementById('knownbtn').addEventListener('click', function (event) {
  event.preventDefault()

  const cardId = this.getAttribute('data-card-id')
  const url = `/card/know/${cardId}`
  const data = { know: true }

  fetch(url, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.ok) {
        return response.text()
      } else {
        throw new Error('Network response was not ok')
      }
    })
    .then((data) => {
      // 处理数据
      console.log(data)
    })
    .catch((error) => {
      // 处理错误
      console.error(
        'There has been a problem with your fetch operation:',
        error
      )
    })
})

document
  .getElementById('deleteBtn')
  .addEventListener('click', function (event) {
    event.preventDefault()

    const cardId = this.getAttribute('data-card-id')
    const url = `/card/delete/${cardId}`

    // 使用window.alert来提醒用户确认删除
    if (window.confirm('Are you sure you want to delete this card?')) {
      const data = { delete: true }
      fetch(url, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data), // 将数据转换为JSON字符串
      })
        .then((response) => {
          if (response.ok) {
            return response.text() // 或者 response.text()、response.blob() 等
          } else {
            throw new Error('Network response was not ok')
          }
        })
        .then((data) => {
          // 处理数据
          console.log(data)
          // 删除成功后可以添加其他操作，如重定向或更新页面
        })
        .catch((error) => {
          // 处理错误
          console.error(
            'There has been a problem with your fetch operation:',
            error
          )
        })
    }
  })
