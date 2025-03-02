package main

import (
	"fmt"
	"log"
	"syscall/js"

	mandelgo "github.com/Salicorne/mandelgo/mandelgo"
)

var (
	ctx   js.Value
	sizeX int
	sizeY int
)

func main() {
	fmt.Println("Initializing drawing context...")

	jsWindow := js.Global().Get("window")
	if !jsWindow.Truthy() {
		log.Fatal("Failed to initialize DOM window")
	}

	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		log.Fatal("Failed to initialize DOM document")
	}

	canvas := jsDoc.Call("getElementById", "canvas")
	if !canvas.Truthy() {
		log.Fatal("Failed to initialize DOM canvas")
	}

	ctx = canvas.Call("getContext", "2d")
	if !ctx.Truthy() {
		log.Fatal("Failed to initialize canvas context")
	}

	rect := canvas.Call("getBoundingClientRect")
	if !rect.Truthy() {
		log.Fatal("Failed to initialize canvas bounding rect")
	}

	log.Printf("Canvas has size %dx%d", rect.Get("width").Int(), rect.Get("height").Int())

	canvas.Set("width", rect.Get("width").Int())
	canvas.Set("height", rect.Get("height").Int())

	sizeX = rect.Get("width").Int()
	sizeY = rect.Get("height").Int()

	log.Println(mandelgo.GetColor(0, 0))

	plot()
}

func plot() {
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			ctx.Set("fillStyle", fmt.Sprintf("#%02x%02x00", int(float64(x)/float64(sizeX)*255), int(float64(y)/float64(sizeY)*255)))
			ctx.Call("fillRect", x, y, 1, 1)
		}
	}
}
