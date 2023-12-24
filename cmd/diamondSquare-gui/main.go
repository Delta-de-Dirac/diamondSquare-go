package main

import (
	_ "errors"
	"image/color"
	"log"
	"strconv"
	"strings"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/Delta-de-Dirac/diamondSquare-go/internal/utils"
	"github.com/Delta-de-Dirac/diamondSquare-go/pkg/heightmap"
)

const startingScreenWidth = 800
const startingScreenHeight = 400

const minScreenWidth = 800
const minScreenHeight = 400

func main() {
	log.Println("Starting...")

	//bgColor := color.RGBA{0x27,0x38,0x73,0xFF}
	//fgColor := color.RGBA{0xED,0x43,0x8D,0xFF}

	bgColor := color.RGBA{0xDD,0xDD,0xDD,0xFF}
	fgColor := color.RGBA{0x11,0x11,0x11,0xFF}

	// middle vertical line
	middleVerticalLineStartPos := rl.Vector2{ X: startingScreenWidth/2, Y: 10}
	middleVerticalLineEndPos := rl.Vector2{ X: startingScreenWidth/2, Y: startingScreenHeight - 10}


	// generate button
	buttonGenerateRectangle := rl.Rectangle{X: 10.0, Y: 120.0, Width: 380.0, Height: 30.0}
	buttonGenerateText := "Button"

	// save button
	buttonSaveRectangle := rl.Rectangle{X: 10.0, Y: 180.0, Width: 380.0, Height: 30.0}
	buttonSaveText := "Button"

	// size vbox
	vboxSizeRectangle := rl.Rectangle{X: 50.0, Y: 0.0, Width: 340.0, Height: 30.0}
	sizeInputActive := false
	sizeInput := ""
	sizeValid := false
	intSize := 257

	sizeLabelRectangle := rl.Rectangle{X: 10.0, Y: 0.0, Width: 40.0, Height: 30.0}
	sizeLabelString := "Size"

	sizeWarningRectangle := rl.Rectangle{X: 50.0, Y: 30.0, Width: 340.0, Height: 30.0}
	sizeWarningString := ""

	// h vbox
	vboxHRectangle := rl.Rectangle{X: 50.0, Y: 60.0, Width: 340.0, Height: 30.0}
	hInputActive := false
	hInput := ""
	hValid := false
	floatH := 0.9

	hLabelRectangle := rl.Rectangle{X: 10.0, Y: 60.0, Width: 40.0, Height: 30.0}
	hLabelString := "h"

	hWarningRectangle := rl.Rectangle{X: 50.0, Y: 90.0, Width: 340.0, Height: 30.0}
	hWarningString := ""

	// filename textbox
	//tboxFilenameRectangle := rl.Rectangle{X: 10.0, Y: 90.0, Width: 125.0, Height: 30.0}


	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(startingScreenWidth, startingScreenHeight, "Diamond Square")
	rl.SetWindowMinSize(minScreenWidth, minScreenHeight)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	gui.LoadStyleDefault()

	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 20)

	displayMap, err := heightmap.NewHeightmap(257)
	if err != nil{
		log.Fatal(err)
	}

	mapImage := rl.NewImageFromImage(displayMap.GetGrayImage())
	mapTexture := rl.LoadTextureFromImage(mapImage)
	rl.UnloadImage(mapImage)

	for ;!rl.WindowShouldClose();{
		if rl.IsMouseButtonPressed(rl.MouseLeftButton){
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), vboxSizeRectangle){
				sizeInputActive = true
				hInputActive = false
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), vboxHRectangle){
				sizeInputActive = false
				hInputActive = true
			} else{
				sizeInputActive = false
				hInputActive = false
			}
		}
		if rl.IsWindowResized(){
			middleVerticalLineStartPos = rl.Vector2{ X: float32(rl.GetScreenWidth())/2, Y: 10}
			middleVerticalLineEndPos = rl.Vector2{ X: float32(rl.GetScreenWidth())/2, Y:float32( rl.GetScreenHeight()) - 10}
		}

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.ClearBackground(bgColor)
		rl.DrawLineEx(middleVerticalLineStartPos, middleVerticalLineEndPos, 3.0, fgColor)

		//displayMap
		rl.DrawTexture(mapTexture,
				410,
				10,
				rl.White)


		//labels
		gui.Label(sizeLabelRectangle, sizeLabelString)
		gui.Label(sizeWarningRectangle, sizeWarningString)
		gui.Label(hLabelRectangle, hLabelString)
		gui.Label(hWarningRectangle, hWarningString)

		//vboxes
		tmpCmp := strings.Clone(sizeInput)
		gui.TextBox(vboxSizeRectangle, &sizeInput, 10, sizeInputActive)
		sizeInput = utils.FilterString(sizeInput, "1234567890")
		if sizeInput != tmpCmp{
			sizeValid = false
			num, err := strconv.Atoi(sizeInput)
			if err != nil{
				sizeWarningString = "Size must be a number"
			} else if !utils.IsPowerOf2(num-1){
				sizeWarningString = "Size must be power of 2 + 1"
			} else {
				sizeValid = true
				intSize = num
				sizeWarningString = ""
			}
		}

		tmpCmp = strings.Clone(hInput)
		gui.TextBox(vboxHRectangle, &hInput, 10, hInputActive)
		hInput = utils.FilterString(hInput, ".1234567890")
		if hInput != tmpCmp{
			hValid = false
			f, err := strconv.ParseFloat(hInput, 64)
			if err != nil{
				hWarningString = "h must be a number"
			} else if f < 0 || f > 1{
				hWarningString = "h must be between 0 and 1"
			} else {
				hValid = true
				floatH = f
				hWarningString = ""
			}
		}

		//buttons
		if gui.Button(buttonGenerateRectangle, buttonGenerateText){
			if hValid && sizeValid{
				displayMap, err = heightmap.NewHeightmap(intSize)
				if err != nil{
					log.Fatal(err)
				}
				displayMap.GenMapP(floatH)
				mapImage = rl.NewImageFromImage(displayMap.GetGrayImage())
				rl.ImageResizeNN(mapImage, 380, 380)
				mapTexture = rl.LoadTextureFromImage(mapImage)
				rl.UnloadImage(mapImage)
			}
		}
		if gui.Button(buttonSaveRectangle, buttonSaveText){
			log.Println("Hello save!")
		}

		rl.EndDrawing()
	}
}

