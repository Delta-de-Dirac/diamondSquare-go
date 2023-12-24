package main

import (
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

	//layout
	margin := float32(10)

	// generate button
	buttonGenerateRectangle := rl.Rectangle{X: margin, Y: float32(startingScreenHeight)/3.0, Width: float32(startingScreenWidth)/2 - margin - 10, Height: float32(startingScreenHeight)/6.0}
	buttonGenerateText := "Generate!"

	// save button
	buttonSaveRectangle := rl.Rectangle{X: margin, Y: 2.0*float32(startingScreenHeight)/3.0, Width: float32(startingScreenWidth)/2 - margin - 10, Height: float32(startingScreenHeight)/6.0}
	buttonSaveText := "Save!"

	// size tbox
	tboxSizeRectangle := rl.Rectangle{X: margin+40, Y: 0.0, Width: float32(startingScreenWidth)/2 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	sizeInputActive := false
	sizeInput := ""
	sizeValid := false
	intSize := 257

	sizeLabelRectangle := rl.Rectangle{X: margin, Y: 0.0, Width: float32(startingScreenWidth)/20.0, Height: float32(startingScreenHeight)/12.0}
	sizeLabelString := "Size"

	sizeWarningRectangle := rl.Rectangle{X: margin+40, Y: float32(startingScreenHeight)/12.0, Width: float32(startingScreenWidth)/2 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	sizeWarningString := ""

	// h tbox
	tboxHRectangle := rl.Rectangle{X: margin+40, Y: float32(startingScreenHeight)/6.0, Width: float32(startingScreenWidth)/2.0 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	hInputActive := false
	hInput := ""
	hValid := false
	floatH := 0.9

	hLabelRectangle := rl.Rectangle{X: margin, Y: float32(startingScreenHeight)/6.0, Width: float32(startingScreenWidth)/20.0, Height: float32(startingScreenHeight)/12.0}
	hLabelString := "h"

	hWarningRectangle := rl.Rectangle{X: margin+40, Y: float32(startingScreenHeight)/4.0, Width: float32(startingScreenWidth)/2.0 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	hWarningString := ""

	// filename tbox
	tboxFilenameRectangle := rl.Rectangle{X: margin+40, Y: float32(startingScreenHeight)/2.0, Width: float32(startingScreenWidth)/2.0 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	filenameInputActive := false
	filenameInput := ""
	filenameValid := false
	outputFormat := ""

	filenameLabelRectangle := rl.Rectangle{X: margin, Y: float32(startingScreenHeight)/2.0, Width: float32(startingScreenWidth)/20.0, Height: float32(startingScreenHeight)/12.0}
	filenameLabelString := "file"

	filenameWarningRectangle := rl.Rectangle{X: margin+40, Y: 7*float32(startingScreenHeight)/12.0, Width: float32(startingScreenWidth)/2.0 - margin - 50, Height: float32(startingScreenHeight)/12.0}
	filenameWarningString := ""


	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(startingScreenWidth, startingScreenHeight, "Diamond Square")
	rl.SetWindowMinSize(minScreenWidth, minScreenHeight)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	gui.LoadStyleDefault()

	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, 20)

	//displayMap
	displayMapX := int32(startingScreenWidth/2) + int32(margin)
	displayMapY := int32(margin)
	displayMapSize := utils.Min(int32(startingScreenWidth/2) - int32(2*margin), int32(startingScreenHeight) - int32(2*margin))
	displayMap, err := heightmap.NewHeightmap(257)
	if err != nil{
		log.Fatal(err)
	}

	mapImage := rl.NewImageFromImage(displayMap.GetGrayImage())
	rl.ImageResizeNN(mapImage, displayMapSize, displayMapSize)
	mapTexture := rl.LoadTextureFromImage(mapImage)
	rl.UnloadImage(mapImage)

	for ;!rl.WindowShouldClose();{
		if rl.IsMouseButtonPressed(rl.MouseLeftButton){
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), tboxSizeRectangle){
				sizeInputActive = true
				hInputActive = false
				filenameInputActive = false
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), tboxHRectangle){
				sizeInputActive = false
				hInputActive = true
				filenameInputActive = false
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), tboxFilenameRectangle){
				sizeInputActive = false
				hInputActive = false
				filenameInputActive = true
			} else{
				sizeInputActive = false
				hInputActive = false
				filenameInputActive = false
			}
		}
		if rl.IsWindowResized(){
			newWidth := float32(rl.GetScreenWidth())
			newHeight := float32(rl.GetScreenHeight())

			//update display map
			displayMapX = int32(newWidth/2) + int32(margin)
			displayMapY = int32(margin)
			displayMapSize = utils.Min(int32(newWidth/2) - int32(2*margin), int32(newHeight) - int32(2*margin))

			mapImage = rl.NewImageFromImage(displayMap.GetGrayImage())
			rl.ImageResizeNN(mapImage, displayMapSize, displayMapSize)
			mapTexture = rl.LoadTextureFromImage(mapImage)
			rl.UnloadImage(mapImage)

			//update middle bar
			middleVerticalLineStartPos = rl.Vector2{ X: newWidth/2, Y: 10}
			middleVerticalLineEndPos = rl.Vector2{ X: newWidth/2, Y: newHeight - 10}

			//update buttons
			buttonGenerateRectangle = rl.Rectangle{X: margin, Y: newHeight/3.0, Width: newWidth/2 - margin - 10, Height: newHeight/6.0}

			buttonSaveRectangle = rl.Rectangle{X: margin, Y: 2.0*newHeight/3.0, Width: newWidth/2 - margin - 10, Height: newHeight/6.0}

			//update tboxes
			tboxSizeRectangle = rl.Rectangle{X: margin+40, Y: 0.0, Width: newWidth/2 - margin - 50, Height: newHeight/12.0}

			sizeLabelRectangle = rl.Rectangle{X: margin, Y: 0.0, Width: newWidth/20.0, Height: newHeight/12.0}

			sizeWarningRectangle = rl.Rectangle{X: margin+40, Y: newHeight/12.0, Width: newWidth/2 - margin - 50, Height: newHeight/12.0}

			tboxHRectangle = rl.Rectangle{X: margin+40, Y: newHeight/6.0, Width: newWidth/2.0 - margin - 50, Height: newHeight/12.0}

			hLabelRectangle = rl.Rectangle{X: margin, Y: newHeight/6.0, Width: newWidth/20.0, Height: newHeight/12.0}

			hWarningRectangle = rl.Rectangle{X: margin+40, Y: newHeight/4.0, Width: newWidth/2.0 - margin - 50, Height: newHeight/12.0}

			tboxFilenameRectangle = rl.Rectangle{X: margin+40, Y: newHeight/2.0, Width: newWidth/2.0 - margin - 50, Height: newHeight/12.0}

			filenameLabelRectangle = rl.Rectangle{X: margin, Y: newHeight/2.0, Width: newWidth/20.0, Height: newHeight/12.0}

			filenameWarningRectangle = rl.Rectangle{X: margin+40, Y: 7*newHeight/12.0, Width: newWidth/2.0 - margin - 50, Height: newHeight/12.0}

		}

		rl.BeginDrawing()
		rl.ClearScreenBuffers()
		rl.ClearBackground(bgColor)
		rl.DrawLineEx(middleVerticalLineStartPos, middleVerticalLineEndPos, 3.0, fgColor)

		//displayMap
		rl.DrawTexture(mapTexture,
				displayMapX,
				displayMapY,
				rl.White)


		//labels
		gui.Label(sizeLabelRectangle, sizeLabelString)
		gui.Label(sizeWarningRectangle, sizeWarningString)
		gui.Label(hLabelRectangle, hLabelString)
		gui.Label(hWarningRectangle, hWarningString)
		gui.Label(filenameLabelRectangle, filenameLabelString)
		gui.Label(filenameWarningRectangle, filenameWarningString)


		//tboxes
		tmpCmp := strings.Clone(sizeInput)
		gui.TextBox(tboxSizeRectangle, &sizeInput, 10, sizeInputActive)
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
		gui.TextBox(tboxHRectangle, &hInput, 10, hInputActive)
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

		tmpCmp = strings.Clone(filenameInput)
		gui.TextBox(tboxFilenameRectangle, &filenameInput, 30, filenameInputActive)
		if filenameInput != tmpCmp{
			filenameValid = true
			outputFormat = ""
			log.Println(filenameInput)
			if strings.HasSuffix(filenameInput,".png"){
				outputFormat = "png"
				filenameWarningString = ""
			}else if strings.HasSuffix(filenameInput,".gif"){
				outputFormat = "gif"
				filenameWarningString = ""
			}else if strings.HasSuffix(filenameInput,".jpeg"){
				outputFormat = "jpeg"
				filenameWarningString = ""
			}else if strings.HasSuffix(filenameInput,".jpg"){
				outputFormat = "jpeg"
				filenameWarningString = ""
			}else{
				filenameValid = false
				filenameWarningString = "Must end in .png .jpg .jpeg or .gif"
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
				rl.ImageResizeNN(mapImage, displayMapSize, displayMapSize)
				mapTexture = rl.LoadTextureFromImage(mapImage)
				rl.UnloadImage(mapImage)
			}
		}
		if gui.Button(buttonSaveRectangle, buttonSaveText){
			if filenameValid{
				err := displayMap.SaveMap(filenameInput, outputFormat)
				if err != nil{
					log.Println(err)
				}
			} else {
				filenameWarningString = "Must end in .png .jpg .jpeg or .gif"
			}
		}

		rl.EndDrawing()
	}
}

