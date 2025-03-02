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
	imgDataHolder := ctx.Call("getImageData", 0, 0, sizeX, sizeY)
	if !imgDataHolder.Truthy() {
		log.Fatal("Failed to get canvas image data holder")
	}

	imgData := imgDataHolder.Get("data")
	if !imgData.Truthy() {
		log.Fatal("Failed to get canvas image data")
	}
	imgDataW := imgDataHolder.Get("width")
	if !imgData.Truthy() {
		log.Fatal("Failed to get canvas image data width")
	}
	sizeX := imgDataW.Int()
	imgDataH := imgDataHolder.Get("height")
	if !imgData.Truthy() {
		log.Fatal("Failed to get canvas image data height")
	}
	sizeY := imgDataH.Int()

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			r, g, b, a := mandelgo.GetColor(float64(x)/float64(sizeX)*(virt_x1-virt_x0)+virt_x0, float64(y)/float64(sizeY)*(virt_y1-virt_y0)+virt_y0).RGBA()
			imgData.SetIndex((y*sizeX+x)*4, uint8(r))
			imgData.SetIndex((y*sizeX+x)*4+1, uint8(g))
			imgData.SetIndex((y*sizeX+x)*4+2, uint8(b))
			imgData.SetIndex((y*sizeX+x)*4+3, uint8(a))
		}
	}

	ctx.Call("putImageData", imgDataHolder, 0, 0)
}
}
