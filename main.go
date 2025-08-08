package main

import (
	"LogWatcher/config"
	"LogWatcher/internal/models"
	"fmt"
	"os"
)


func main() {
	var cfg *models.Config

	if len(os.Args) > 1 {
		// El primer argumento (índice 1) es la cadena deseada
		filepath := os.Args[1]
		cfg, err := config.CargarConfig(filepath)
		if err != nil {
			fmt.Println("Error al cargar la configuracion", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("No se proporcionó ninguna ruta!")
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)
}
