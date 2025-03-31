package router

import "log"

type App struct {
	Port           string
	Info, ErrorLog *log.Logger
}
