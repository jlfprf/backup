package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

func main() {
	data, err := os.ReadFile("filestobackup.json")
	if err != nil {
		fmt.Println("Could not read JSON file:", err)
		return
	}

	backupData := struct {
		Files       []string `json:"filestobackup"`
		Destination string   `json:"destination"`
	}{}

	if err := json.Unmarshal(data, &backupData); err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}
	fmt.Println(backupData)

	for _, file := range backupData.Files {
		destPath := filepath.Join(backupData.Destination, file)
		if err := copy.Copy(file, destPath); err != nil {
			fmt.Printf("Error backing up %s: %v\n", file, err)
		}
	}

	fmt.Println("Tasks completed!")
}
