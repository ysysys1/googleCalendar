package main

import (
	"context"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/calendar-open/environments"
	"github.com/calendar-open/outbounds"
	"github.com/calendar-open/usecase"
	"github.com/calendar-open/web_controller"
)

func main() {
	calendarClient := outbounds.NewCalendarClient()
	envs := environments.NewEnvironments()
	usecase := usecase.NewCalendarBatchUsecase(calendarClient, envs)
	controller := web_controller.NewOpenReservationFrameWebController(usecase)

	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", controller.OpenReservationFrame); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
