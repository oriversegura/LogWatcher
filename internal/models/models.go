package models

import (
	"time"
)

type TelegramConfig struct {
	Token string
	IDChat string
}

type EmailConfig struct {
	ServidorSMTP string
	Puerto int
	Usuario string
	Contrasena string
}

type HostConfig struct {
	Usuario string
	Contrasena string
	Puerto int
	IPHost string
	RutaLogs string
}

type DBConfig struct {
	Uri string

}

type PatronDeLog struct {
	Regex string
	FrecuenciaDeLog int
	TiempoDeLog time.Duration
}

type Config struct {
	HostConfig HostConfig
	DBConfig DBConfig
	PatronDeLog []PatronDeLog
	TelegramConfig TelegramConfig
	EmailConfig EmailConfig
}