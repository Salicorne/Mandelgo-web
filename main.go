package main

import (
	"fmt"
	"log"
	"syscall/js"
)

var (
	ctx    js.Value
	window js.Value
	sizeX  int
	sizeY  int
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

	ctx := canvas.Call("getContext", "2d")
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
}
