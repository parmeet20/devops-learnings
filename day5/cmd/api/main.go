package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parmeet20/bloghive/internal/config"
	"github.com/parmeet20/bloghive/internal/db"
	"github.com/parmeet20/bloghive/internal/handler"
	"github.com/parmeet20/bloghive/internal/repository"
	"github.com/parmeet20/bloghive/internal/service"
)

func main() {
	cfg := config.Load()

	client, err := db.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(cfg.DBName)

	repo := repository.NewBlogRepository(database)
	svc := service.NewBlogService(repo)

	router := gin.Default()
	handler.NewBlogHandler(router, svc)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Server running on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
