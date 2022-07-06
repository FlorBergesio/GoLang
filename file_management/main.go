package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile_ioutil()
	readFile_os()
	writeFile_os()
	appendTextToFile()
}

/* Lectura básica de un archivo */
func readFile_ioutil() {
	file, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(file))
	}
}

/* Manejo de archivo mediante funciones de sistema operativo */
func readFile_os() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	file.Close()
}

func writeFile_os() {
	file, err := os.Create("./new_file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(file, "This is a new hello world!")

	file.Close()
}

func appendTextToFile() {
	file := "./new_file.txt"
	if Append(file, "\nThis is a new line!") == false {
		fmt.Println("Error en la 2da linea")
	}
}

func Append(file string, text string) bool {
	openedFile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = openedFile.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	openedFile.Close()
	return true
}
