package main

import (
	"context"
	"finance/database"
	"finance/internal/handlers"
	"finance/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/register", handlers.HandleRegister)
	mux.HandleFunc("GET /api/currentUser", handlers.HandleGetCurrentUser)
	mux.HandleFunc("POST /api/login", handlers.HandleLogin)
	mux.Handle("GET /api/allTransactions", middleware.CheckAuth(http.HandlerFunc(handlers.HandleGetAllTransactions)))
	mux.HandleFunc("POST /api/logout", handlers.HandleLogout)
	mux.HandleFunc("POST /api/createTransaction", handlers.HandleCreateTransaction)
	mux.HandleFunc("POST /api/createBudget", handlers.HandleCreateBudget)
	mux.HandleFunc("GET /api/allBudgets", handlers.HandleGetAllBudgets)
	mux.HandleFunc("POST /api/createExpense", handlers.HandleCreateExpense)
	mux.HandleFunc("GET /api/allExpenses", handlers.HandleGetAllExpenses)
	mux.HandleFunc("DELETE /api/deleteTransaction/{id}", handlers.HandleDeleteTransaction)

}

func main() {

	database.ConnectDB()
	mux := http.NewServeMux()
	handler := middleware.CorsMiddleware(mux)

	setupRoutes(mux)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	serverPort := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))

	server := &http.Server{
		Addr:    serverPort,
		Handler: handler,
	}

	go func() {
		fmt.Printf("Starting server on:%v", serverPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shut down: %v\n", err)
	}

}
