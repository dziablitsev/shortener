package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/dziablitsev/shortener/internal/app"
	"github.com/dziablitsev/shortener/internal/logger"
)

func main() {
	appErr := app.Run()
	if loggerErr := logger.Initialize(zap.DebugLevel.String()); loggerErr != nil {
		panic(loggerErr)
	}
	logger.Log.Fatal(fmt.Sprint(appErr))
}
