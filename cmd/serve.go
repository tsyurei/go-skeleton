package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"go-skeleton/app"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveAPP(ctx context.Context, app *app.App) {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router := mux.NewRouter()

	app.InitRouter(router)

	logFile, err := os.OpenFile("skeleton.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	defer logFile.Close()
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	server := &http.Server{
		Addr:        fmt.Sprintf(":%v", app.Config.Port),
		Handler:     cors(router),
		ReadTimeout: 10 * time.Second,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("Serving Api at http://127.0.0.1:%v", app.Config.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := app.New()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			// this will wait until cancel signal is sent to the app
			<-ch
			logrus.Info("Signal received. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer cancel()
			serveAPP(ctx, a)
		}()

		wg.Wait()
		return nil
	},
}
