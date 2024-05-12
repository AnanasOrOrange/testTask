package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server running")

	s := echo.New()

	s.GET("/status", Handler, MW)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(date)
	ans := fmt.Sprintf("Количество дней до 1го января 2025 года %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, ans)
	if err != nil {
		return err
	}
	return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("detected admin")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
