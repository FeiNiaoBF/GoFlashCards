document
  .getElementById('putButton')
  .addEventListener('click', function (event) {
    event.preventDefault() // 阻止按钮的默认行为，虽然这里没有默认行为

    const cardId = this.getAttribute('data-card-id')
    const url = `/card/know/${cardId}` // 请求的URL
    const data = { know: true } // 请求体数据，根据需要修改

    fetch(url, {
      method: 'PUT', // 使用PUT方法
      headers: {
        'Content-Type': 'application/json', // 如果发送JSON数据
      },
      body: JSON.stringify(data), // 将数据转换为JSON字符串
    })
      .then((response) => {
        if (response.ok) {
          return response.text // 或者 response.text()、response.blob() 等
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
