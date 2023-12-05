package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Delta-de-Dirac/diamondSquare-go/pkg/heightmap"
)


func main() {
	if len(os.Args) != 4{
		log.Printf("Wrong number of arguments... expected 3 arguments, but received %d", len(os.Args)-1)
		log.Fatal("usage: diamondSquare-cli <size> <h> <output filename>")
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

	if strings.HasSuffix(fileName,".png"){
		outputFormat = "png"
	}
	if strings.HasSuffix(fileName,".gif"){
		outputFormat = "gif"
	}
	if strings.HasSuffix(fileName,".jpeg"){
		outputFormat = "jpeg"
	}
	if strings.HasSuffix(fileName,".jpg"){
		outputFormat = "jpeg"
	}

	if outputFormat == ""{
		log.Fatal("argument <output filename> must end in .png .jpg .jpeg or .gif")
	}

	hmap, err := heightmap.NewHeightmap(size)
	if err != nil{
		log.Fatal(err)
	}

	err = hmap.GenMap(h)
	if err != nil{
		log.Fatal(err)
	}

	err = hmap.SaveMap(fileName, outputFormat)
	if err != nil{
		log.Fatal(err)
	}

}

