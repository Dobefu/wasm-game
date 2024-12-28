'use strict'

// @ts-check

function createCanvas() {
  const canvas = document.createElement('canvas')
  canvas.id = 'game'

  document.body.appendChild(canvas)

  return canvas
}

;(() => {
  const canvas = createCanvas()
  console.log(canvas)
})()
