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

