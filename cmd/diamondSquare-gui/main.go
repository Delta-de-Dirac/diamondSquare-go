package main

import (
	"errors"
	"image"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Delta-de-Dirac/diamondSquare-go/internal/utils"
	"github.com/Delta-de-Dirac/diamondSquare-go/pkg/heightmap"
)

func main() {
	a := app.New()
	w := a.NewWindow("Diamond Square")

	sizeEntry := widget.NewEntry()
	sizeEntry.Text = "257"

	hEntry := widget.NewEntry()
	hEntry.Text = "0.95"

	fileNameEntry := widget.NewEntry()
	fileNameEntry.Text = ""
	fileNameEntry.Disable()

	sizeEntry.Validator = func(s string) error{
		i, err := strconv.Atoi(s)
		if err != nil{
			return err
		}
		if !utils.IsPowerOf2(i-1){
			err := errors.New("Size must be positive integer (2^n)+1 where n natural. Example: 3, 9, 17, 33...")
			return err
		}
		return nil
	}

	hEntry.Validator = func(s string) error{
		f, err := strconv.ParseFloat(s, 64)
		if err != nil{
			return err
		}
		if f < 0 || f > 1{
			return errors.New("Parameter h must be float between 0 and 1")
		}
		return nil
	}

	fileNameEntry.Validator = func(s string) error{
		if strings.HasSuffix(s,".png"){
			return nil
		}
		if strings.HasSuffix(s,".gif"){
			return nil
		}
		if strings.HasSuffix(s,".jpeg"){
			return nil
		}
		if strings.HasSuffix(s,".jpg"){
			return nil
		}
		return errors.New("filename must end in .png .jpg .jpeg or .gif")
	}


	mainCanvas := canvas.NewImageFromImage(image.NewGray(image.Rect(0,0,513,513)))
	mainCanvas.SetMinSize(fyne.NewSize(400,400))
	mainCanvas.FillMode = canvas.ImageFill(canvas.ImageFillContain)

	displayMap, err  := heightmap.NewHeightmap(257)

	if err != nil{
		log.Fatal(err)
	}

	parametersForm := widget.NewForm()
	savingForm := widget.NewForm()

	saveButton := widget.NewButton("Save image to file",func(){
		outputFormat := ""
		if strings.HasSuffix(fileNameEntry.Text,".png"){
			outputFormat = "png"
		}
		if strings.HasSuffix(fileNameEntry.Text,".gif"){
			outputFormat = "gif"
		}
		if strings.HasSuffix(fileNameEntry.Text,".jpeg"){
			outputFormat = "jpeg"
		}
		if strings.HasSuffix(fileNameEntry.Text,".jpg"){
			outputFormat = "jpeg"
		}
		if outputFormat == ""{
			log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
		}
		displayMap.SaveMap(fileNameEntry.Text, outputFormat)
	})
	saveButton.Disable()

	generateButton := widget.NewButton("Generate",func(){
		size, err := strconv.Atoi(sizeEntry.Text)
		if err != nil{
			log.Fatal(err)
		}
		hmap, err := heightmap.NewHeightmap(size)
		if err != nil{
			log.Fatal(err)
		}
		hf, err := strconv.ParseFloat(hEntry.Text, 64)
		if err != nil{
			log.Fatal(err)
		}
		hmap.GenMap(hf)
		mainCanvas.Image = hmap.GetGrayImage()
		mainCanvas.Refresh()
		displayMap = hmap
		fileNameEntry.Enable()
		savingForm.Refresh()
	})


	parametersForm.Append("Size", sizeEntry)
	parametersForm.Append("Parameter h", hEntry)
	parametersForm.SetOnValidationChanged(func(err error){
		if err != nil{
			generateButton.Disable()
			return
		}
		generateButton.Enable()
	})

	savingForm.Append("Output FileName", fileNameEntry)

	savingForm.SetOnValidationChanged(func(err error){
		if err != nil{
			saveButton.Disable()
			return
		}
		saveButton.Enable()
	})


	leftContainer := container.NewVBox(parametersForm,
					   generateButton,
					   savingForm,
					   saveButton)


	hc := container.NewGridWithColumns(2, leftContainer, mainCanvas)

	w.SetContent(hc)
	w.ShowAndRun()
}

