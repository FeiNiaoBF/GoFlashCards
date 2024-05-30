document.addEventListener('DOMContentLoaded', function () {
  const canvas = document.getElementById('gameCanvas')
  const ctx = canvas.getContext('2d')
  const gridSize = 20 // 每个单元格的大小
  const gridWidth = canvas.width / gridSize
  const gridHeight = canvas.height / gridSize
  let grid = []
  let intervalId = null

  function initGrid() {
    for (let y = 0; y < gridHeight; y++) {
      grid[y] = []
      for (let x = 0; x < gridWidth; x++) {
        grid[y][x] = Math.random() > 0.7 ? 1 : 0 // 随机初始化
      }
    }
  }

  function drawGrid() {
    for (let y = 0; y < gridHeight; y++) {
      for (let x = 0; x < gridWidth; x++) {
        const cell = grid[y][x]
        const xPos = x * gridSize
        const yPos = y * gridSize
        ctx.fillStyle = cell ? 'black' : 'white'
        ctx.fillRect(xPos, yPos, gridSize, gridSize)
        ctx.strokeRect(xPos, yPos, gridSize, gridSize)
      }
    }
  }

  function countNeighbors(x, y) {
    let count = 0
    for (let i = -1; i <= 1; i++) {
      for (let j = -1; j <= 1; j++) {
        if (i === 0 && j === 0) continue
        const nx = (x + i + gridWidth) % gridWidth
        const ny = (y + j + gridHeight) % gridHeight
        count += grid[ny][nx]
      }
    }
    return count
  }

  function updateGrid() {
    const newGrid = []
    for (let y = 0; y < gridHeight; y++) {
      newGrid[y] = []
      for (let x = 0; x < gridWidth; x++) {
        const neighbors = countNeighbors(x, y)
        const isAlive = grid[y][x]
        if (isAlive && (neighbors < 2 || neighbors > 3)) {
          newGrid[y][x] = 0
        } else if (!isAlive && neighbors === 3) {
          newGrid[y][x] = 1
        } else {
          newGrid[y][x] = grid[y][x]
        }
      }
    }
    grid = newGrid
  }

  function startGame() {
    if (!intervalId) {
      intervalId = setInterval(() => {
        updateGrid()
        drawGrid()
      }, 500)
    }
  }

  function stopGame() {
    clearInterval(intervalId)
    intervalId = null
  }

  function clearGame() {
    stopGame()
    initGrid()
    drawGrid()
  }

  document.getElementById('startButton').addEventListener('click', startGame)
  document.getElementById('stopButton').addEventListener('click', stopGame)
  document.getElementById('clearButton').addEventListener('click', clearGame)

  initGrid()
  drawGrid()
})
