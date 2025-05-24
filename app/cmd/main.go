package main

import (
	"SmartHouseAPI/config"
	"SmartHouseAPI/handler"
	"SmartHouseAPI/repository/postgres"
	"SmartHouseAPI/route"
	"SmartHouseAPI/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// todo: add initial setup migration in SQL
// todo: fix some mismatches device_id->device_name
// todo: finish handler layer
// todo: deprecate all unusable function in repo layer

func main() {
	conf, err := config.LoadConfigFile(".env")
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки конфигурации: %v", err)
	}

	repo, err := postgres.NewRepository(conf.DB)
	if err != nil {
		log.Fatalf("❌ Ошибка создания репозитория: %v", err)
	}

	s := service.NewService(fmt.Sprintf("tcp://%s:%s", conf.MQTT.Host, conf.MQTT.Port), repo, conf.MQTT)
	err = s.Start()
	if err != nil {
		log.Fatal("❌ Ошибка создания сервиса:", err)
	}

	h := handler.NewHandler(s)
	engine := route.SetRouter(h)

	address := fmt.Sprintf("%s:%d", conf.S.Host, conf.S.Port)

	srv := &http.Server{
		Addr:    address,
		Handler: engine,
	}

	go func() {
		log.Printf("🚀 Запуск сервера на %s\n", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Ошибка сервера: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("⚠️ Получен сигнал остановки, выключаем сервер...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Ошибка при завершении сервера: %v", err)
	}

	log.Println("✅ Сервер успешно остановлен")
}
