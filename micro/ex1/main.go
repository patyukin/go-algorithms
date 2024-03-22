package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt) // Обработка сигнала SIGINT

	go func() {
		<-interrupt
		fmt.Println("Received SIGINT signal")
		cancel() // Останавливаем контекст для graceful shutdown
	}()

	// Логика работы сервера
	server := startServer(ctx)

	<-ctx.Done()                // Ждем сигнала закрытия контекста
	shutdownServer(server)      // Останавливаем сервер
	time.Sleep(1 * time.Second) // Дополнительная задержка (например, для завершения обработки запросов)
	fmt.Println("Graceful shutdown completed")
}

func startServer(ctx context.Context) *http.Server {
	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	return server
}

func shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during server shutdown: %v\n", err)
	}
}
