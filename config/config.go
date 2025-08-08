package config

import (
	"LogWatcher/internal/models"
	"encoding/json"
	"os"
)



func CargarConfig(filepath string) (*models.Config, error) {

	data, err := os.ReadFile(filepath)
	
	if err != nil {
		return nil, err
	}

	config := &models.Config{}

	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

