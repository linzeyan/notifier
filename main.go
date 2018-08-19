package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gotoolkit/bot"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Specification struct {
	Debug       bool
	Port        int
	BotToken    string
	EnableAdmin bool
	Admins      []string
	Groups      []string
}

type Messager struct {
	Message string `json:"message"`
}

func main() {
	var s Specification
	err := envconfig.Process("notifier", &s)
	if err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}

	// start bot
	var ids []int64
	groupIds, err := getIds(s.Groups)
	if err != nil {
		log.Fatalf("Failed to parse group ids : %v", err)
	}
	ids = append(ids, groupIds...)
	if s.EnableAdmin {
		adminIds, err := getIds(s.Admins)
		if err != nil {
			log.Fatalf("Failed to parse admin ids : %v", err)
		}
		ids = append(ids, adminIds...)
	}
	withIDs := bot.WithChatIDs(ids...)
	tb, err := bot.NewTelegramBot(s.BotToken, bot.DebugMode(s.Debug), withIDs)
	if err != nil {
		log.Fatalf("Failed to create telegram bot: %v", err)
	}

	// start server
	e := echo.New()

	// add Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/send", sendMessage(tb))

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", s.Port)); err != nil {
			e.Logger.Info("Shutting down notifier server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func sendMessage(tb *bot.Bot) echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(Messager)
		if err := c.Bind(m); err != nil {
			return err
		}
		tb.SendMessage(m.Message)
		return c.JSON(http.StatusOK, echo.Map{
			"status": true,
			"code":   0,
		})
	}
}

func getIds(sa []string) ([]int64, error) {
	si := make([]int64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return si, err
		}
		si = append(si, int64(i))
	}
	return si, nil
}
