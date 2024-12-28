'use strict'

// @ts-check
;(() => {
  const go = new Go()

  WebAssembly.instantiateStreaming(
    fetch('wasm/game.wasm'),
    go.importObject,
  ).then((result) => {
    go.run(result.instance)
  })
})()
