package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to file App")
	filePrompt := "Enter Your File name to create a a file"
	fileName := KeyboardPrompt(filePrompt)
	file := CreateFile(fileName)
	contentPrompt := "Enter Content you want to save in " + fileName
	content := KeyboardPrompt(contentPrompt)
	savedResult := saveTextReader(file, content)

	fmt.Printf("Number of bytes written: %d\n", savedResult)
	defer file.Close()

}

func CreateFile(filename string) os.File {
	myfilename := filename + ".txt"
	file, err := os.Create(myfilename)
	if err != nil {
		panic(err)
	}
	return *file
}

func KeyboardPrompt(dialogMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(dialogMessage)
	filename, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	filename = strings.TrimSpace(filename)

	return filename
}

func saveTextReader(file os.File, content string) int {
	result, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
	return result

}

// file, err := os.Create("myFile.csv")
// if err != nil {
// 	panic(err)
// }
// n, err := file.WriteString(content)
// if err != nil {
// 	panic(err)
// }
