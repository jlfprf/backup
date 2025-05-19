package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
<<<<<<< HEAD
=======
	"time"
>>>>>>> e527404 (updating)

	"github.com/otiai10/copy"
)

func main() {
	data, err := os.ReadFile("backup.json")
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
	now := time.Now().Format("2006-01-02 15:04:05")
	dst := backupData.Destination + "/" + now

	for _, file := range backupData.Files {
		destPath := filepath.Join(dst, file)
		if err := copy.Copy(file, destPath); err != nil {
			fmt.Printf("Error backing up %s: %v\n", file, err)
		}
	}

	fmt.Println("Tasks completed!")
}
