'use strict'

// @ts-check

function createCanvas() {
  const canvas = document.createElement('canvas')

  document.body.appendChild(canvas)

  return canvas
}

;(() => {
  const canvas = createCanvas()
  console.log(canvas)
})()
