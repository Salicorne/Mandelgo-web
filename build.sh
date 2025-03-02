GOOS=js GOARCH=wasm go build -o mandelgo-web.wasm
cp mandelgo-web.wasm ../salicorne/static/mandelgo/mandelgo-web.wasm
cp index.html ../salicorne/static/mandelgo/index.html
cp index.css ../salicorne/static/mandelgo/index.css
cp wasm_exec.js ../salicorne/static/mandelgo/wasm_exec.js
