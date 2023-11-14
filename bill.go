package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(n string) bill {
	b := bill{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n\n"
	var total float64 = 0
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}
	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%v\n\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(n string, p float64) {
	b.items[n] = p
}

func checkSaveDestination() bool {
	// Specify the folder name you want to check
	folderName := "bills"

	// Get the absolute path to the current directory (root of the project)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return false
	}

	// Construct the absolute path to the folder
	folderPath := filepath.Join(dir, folderName)

	// Check if the folder exists
	if _, fileStatusCheckError := os.Stat(folderPath); fileStatusCheckError == nil {
		return true
	} else if os.IsNotExist(fileStatusCheckError) {
		// Folder does not exist, so create it
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			panic(err)
		}
		fmt.Println("Folder created:", folderPath)
	} else {
		fmt.Println("Error checking folder:", fileStatusCheckError)
		panic(fileStatusCheckError)
	}

	return true
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	checkSaveDestination()

	writeFileError := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if writeFileError != nil {
		panic(writeFileError)
	}
	fmt.Println("file saved")
}
