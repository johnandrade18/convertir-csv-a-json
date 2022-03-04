package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type OrganizationRecord struct {
	Organizacion string `json:"organizacion"`
	Usuarios     string `json:"usuario"`
	Rol          string `json:"rol"`
}

func createOrganizationList(data [][]string) []OrganizationRecord {
	var organozationList []OrganizationRecord
	for i, line := range data {
		if i > 0 {
			var record OrganizationRecord
			for j, field := range line {
				if j == 0 {
					record.Organizacion = field
				} else if j == 1 {
					record.Usuarios = field
				} else if j == 2 {
					record.Rol = field
				}
			}
			organozationList = append(organozationList, record)
		}
	}
	return organozationList
}

func main() {
	// Open the file
	csvfile, err := os.Open("organizaciones.csv")
	if err != nil {
		log.Fatalln("No se pudo abrir el archivo csv", err)
	}

	defer csvfile.Close()

	// Parse the file
	csvReader := csv.NewReader(csvfile)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("No se pudo analizar el archivo csv", err)
	}

	// Create the JSON
	organizationList := createOrganizationList(data)

	jsonData, err := json.MarshalIndent(organizationList, "", "  ")
	if err != nil {
		log.Fatalln("No se pudo crear el archivo json", err)
	}
	// Write the JSON to stdout
	fmt.Println(string(jsonData))
}
