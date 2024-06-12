package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const OutputDirectoryName = "text"

type Header struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
	Date        string `xml:"date"`
	Author      string `xml:"author"`
	URL         string `xml:"url"`
}

type Game struct {
	Name         string `xml:"name,attr"`
	Description  string `xml:"description"`
	Year         string `xml:"year"`
	Manufacturer string `xml:"manufacturer"`
	ROM          ROM    `xml:"rom"`
}

type ROM struct {
	Name string `xml:"name,attr"`
	Size string `xml:"size,attr"`
}

type Datafile struct {
	XMLName xml.Name `xml:"datafile"`
	Header  Header   `xml:"header"`
	Game    []Game   `xml:"game"`
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting the current directory: %v\n", err)
		return
	}

	log.Printf("Current operation directory: %v\n", dir)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return nil
		}

		// Read only files with .xml and .dat extensions (.dat files are produced by Skraper)
		if info.Mode().IsRegular() && (strings.HasSuffix(info.Name(), ".xml") || strings.HasSuffix(info.Name(), ".dat")) {
			log.Printf("Reading file: %v\n", path)
			fileData, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading XML / DAT file: %v\n", err)
				return nil
			}
			generateGameDescriptionsForMuOS(fileData)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error traversing the directory: %v\n", err)
	}
}

// generateGameDescriptionsForMuOS generates game descriptions based on XML data and writes them to files.
// It takes in a byte array containing XML data as input.
// The function reads the XML data, creates a directory if necessary, and creates a description file for each game.
// The description files are named based on the ROM name of each game.
// Each description file includes the game's name, release date, manufacturer, and description.
// Encountered errors stop the execution.
func generateGameDescriptionsForMuOS(xmlData []byte) {
	processedCtr := 0

	// Reading XML data
	var datafile Datafile
	err := xml.Unmarshal(xmlData, &datafile)
	if err != nil {
		log.Fatalf("Error when parsing XML contents: %v\n", err)
	}

	// Checking if directory exists and creating if necessary
	if _, err := os.Stat(OutputDirectoryName); os.IsNotExist(err) {
		err := os.MkdirAll(OutputDirectoryName, 0755)
		if err != nil {
			log.Fatalf("Error creating %v directory: %v\n", OutputDirectoryName, err)
		}
	}

	// Creating descriptions and writing to a file (based on the ROM name)
	for _, game := range datafile.Game {
		log.Printf("Generating description for: %v\n", game.Name)
		gameDescFilePath := fmt.Sprintf("%v/%v.txt", OutputDirectoryName, game.ROM.Name)

		gameDescFile, err := os.Create(gameDescFilePath)
		if err != nil {
			log.Fatalf("Cannot create file %v: %v\n", gameDescFilePath, err)
		}

		_, err = fmt.Fprintf(gameDescFile, "Name: %v\nRelease Date: %v\nManufacturer: %v\n\nDescription: %v\n", game.Name, game.Year, game.Manufacturer, game.Description)
		if err != nil {
			log.Fatalf("Cannot write to file %v: %v\n", gameDescFilePath, err)
		}

		processedCtr++

		err = gameDescFile.Close()
		if err != nil {
			log.Fatalf("Cannot close file %v: %v\n", gameDescFilePath, err)
		}

		log.Printf("Generated description for %v in: %v\n", game.Name, gameDescFilePath)
	}

	log.Printf("Finished generating description for %v games\n", processedCtr)
}
