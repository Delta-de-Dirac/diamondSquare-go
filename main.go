package main

import (
	"log"
	"os"
	"strconv"
	"math/rand"
	"math"
	"image"
	"image/png"
	"image/gif"
	"image/jpeg"
	"image/color"
	"errors"
)

func isPowerOf2(x int) bool {
	if x <= 0{
		return false
	}
	for x % 2 == 0{
		x /= 2
		if x == 1{
			return true
		}
	}
	return false
}

func initializeCorners(hmap [][]float64){
	hmap[0][0] = 2*(rand.Float64()-0.5)
	hmap[0][len(hmap)-1] = 2*(rand.Float64()-0.5)
	hmap[len(hmap)-1][0] = 2*(rand.Float64()-0.5)
	hmap[len(hmap)-1][len(hmap)-1] = 2*(rand.Float64()-0.5)
}

func diamondStep(hmap [][]float64, scale float64, depth int){
	begin := (len(hmap)-1)>>(depth+1)
	step := (len(hmap)-1)>>depth
	for i:=0;i<1<<depth;i++{
		for j:=0;j<1<<depth;j++{
			hmap[begin + i*step][begin + j*step] = (
				hmap[i*step][j*step] +
				hmap[i*step][2*begin + j*step] +
				hmap[2*begin + i*step][j*step] +
				hmap[2*begin + i*step][2*begin + j*step])/4
			hmap[begin + i*step][begin + j*step] += scale*2*(rand.Float64()-0.5)
		}
	}
}

func squareStep(hmap [][]float64, scale float64, depth int){
	begin := (len(hmap)-1)>>(depth+1)
	step := 2*begin

	//Edge cases for first and last line of nodes
	for column:=begin;column<len(hmap);column+=step{
		hmap[0][column] = (
			hmap[0][column-begin] +
			hmap[0][column+begin] +
			hmap[begin][column])/3
		hmap[0][column] += scale*2*(rand.Float64()-0.5)

		hmap[len(hmap)-1][column] = (
			hmap[len(hmap)-1-begin][column] +
			hmap[len(hmap)-1][column+begin] +
			hmap[len(hmap)-1][column-begin])/3
		hmap[len(hmap)-1][column] += scale*2*(rand.Float64()-0.5)
	}

	//Edge cases for first and last column of nodes
	for line:=begin;line<len(hmap);line+=step{
		hmap[line][0] = (
			hmap[line-begin][0] +
			hmap[line][begin] +
			hmap[line+begin][0])/3
		hmap[line][0] += scale*2*(rand.Float64()-0.5)

		hmap[line][len(hmap)-1] = (
			hmap[line-begin][len(hmap)-1] +
			hmap[line+begin][len(hmap)-1] +
			hmap[line][len(hmap)-1-begin])/3
		hmap[line][len(hmap)-1] += scale*2*(rand.Float64()-0.5)
	}

	for line:=step;line<len(hmap)-step;line+=step{
		for column:=begin;column<len(hmap);column+=step{
			hmap[line][column] = (
				hmap[line-begin][column] +
				hmap[line+begin][column] +
				hmap[line][column-begin] +
				hmap[line][column+begin])/4

			hmap[line][column] += scale*2*(rand.Float64()-0.5)
		}
	}




	for line:=begin;line<len(hmap);line+=step{
		for column:=step;column<len(hmap)-step;column+=step{
			hmap[line][column] = (
				hmap[line-begin][column] +
				hmap[line+begin][column] +
				hmap[line][column-begin] +
				hmap[line][column+begin])/4

			hmap[line][column] += scale*2*(rand.Float64()-0.5)
		}
	}
}

func normalizeHmap(hmap [][]float64){
	minVal := hmap[0][0]
	maxVal := hmap[0][0]
	for i := range hmap{
		for j := range hmap[i]{
			if hmap[i][j] > maxVal{
				maxVal = hmap[i][j]
			}
			if hmap[i][j] < minVal{
				minVal = hmap[i][j]
			}
		}
	}
	for i := range hmap{
		for j := range hmap[i]{
			hmap[i][j] -= minVal
			if maxVal-minVal == 0{
				hmap[i][j] /= 1
				continue
			}
			hmap[i][j] /= maxVal-minVal
		}
	}
}


func genMap(hmap [][]float64, h float64){
	initializeCorners(hmap)
	factor := math.Pow(2, -h)
	scale := 1.0
	for i:=0;(len(hmap)-1)>>i>1;i++{
		diamondStep(hmap, scale, i)
		squareStep(hmap, scale, i)
		scale *= factor
	}
	normalizeHmap(hmap)
}

func saveMap(hmap [][]float64, fileName string, outputFormat string) error{
	file, err := os.Create(fileName)
	if err != nil{
		return err
	}
	defer file.Close()
	outputImage := image.NewGray(image.Rect(0,0,len(hmap),len(hmap)))
	for i := range hmap{
		for j := range hmap[i]{
			outputImage.Set(i,j, color.Gray{
				Y: (uint8)(hmap[i][j]*255),
			})
		}
	}
	switch outputFormat{
		case "png":
			png.Encode(file, outputImage)
		case "jpeg":
			jpeg.Encode(file, outputImage, &jpeg.Options{Quality: 100})
		case "gif":
			gif.Encode(file, outputImage, &gif.Options{NumColors: 256, Quantizer: nil ,Drawer: nil})
		default:
			return errors.New("Error saving result")
	}
	return nil
}

func main() {
	log.Printf("Starting...")

	if len(os.Args) != 4{
		log.Printf("Wrong number of arguments... expected 3 arguments, but received %d", len(os.Args)-1)
		log.Fatal("usage: diamondSquare-go <size> <h> <output filename>")
	}
	size, err := strconv.Atoi(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}
	h, err := strconv.ParseFloat(os.Args[2],64)
	if err != nil{
		log.Fatal(err)
	}
	fileName := os.Args[3]

	if len(fileName) < 4 {
		log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
	}

	outputFormat := ""

	switch fileName[len(fileName)-4:]{
		case ".png":
			outputFormat = "png"
		case ".jpg":
			outputFormat = "jpeg"
		case ".gif":
			outputFormat = "gif"
	}

	if outputFormat == ""{
		if len(fileName) < 5 {
			log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
		}

		if !(fileName[len(fileName)-5:] == ".jpeg"){
			log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
		}
		outputFormat = "jpeg"
	}

	if outputFormat == ""{
		log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
	}

	if !isPowerOf2(size - 1){
		log.Fatal("argument <size> must positive integer (2^n)+1 where n natural. Example: 3, 9, 17, 33...")
	}
	if !isPowerOf2(size - 1){
		log.Fatal("argument <size> must positive integer (2^n)+1 where n natural. Example: 3, 9, 17, 33...")
	}
	if h<0 || h>1{
		log.Fatal("argument <h> must be float between 0 and 1")
	}

	hmap := make([][]float64, size)

	for i := range hmap{
		hmap[i] = make([]float64, size)
	}

	genMap(hmap, h)

	err = saveMap(hmap, fileName, outputFormat)
	if err != nil{
		log.Fatal(err)
	}

}

