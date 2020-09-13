package main

import (
	"context"
	"database/sql"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	config2 "ted-processor/pkg/domain/infra/config"
	"ted-processor/pkg/domain/infra/errors"
	infra "ted-processor/pkg/domain/infra/logger"
	"ted-processor/pkg/domain/ted"
	"ted-processor/pkg/domain/ted/manager"
)

func main() {
	defer infra.Logger.Sync()
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		infra.Logger.Fatalw("failed to create protocol", "error",errors.Wrap(err))
	}

	// read configuration
	config := &config2.Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		BAHost: os.Getenv("BANK_ACCOUNT_HOST"),
	}

	c, err := cloudevents.NewClient(p)
	if err != nil {
		infra.Logger.Fatalw("failed to create cloudevents client", "error",errors.Wrap(err))
	}

	// database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBDatabase)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		infra.Logger.Panicw("error to connect database", "error", err)
		panic(err)
	}
	defer db.Close()

	// factories
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	pm := ted.NewPaymentNotification(client,config)
	repo := ted.NewTedPostgresRepo(db)
	tm := manager.NewTedManager(repo,pm)

	infra.Logger.Infow("start to listen application", "port","8080")
	errSrv := c.StartReceiver(ctx,tm.Receive)
	if errSrv != nil {
		infra.Logger.Fatalw("failed to start cloud events receiver", "error",errors.Wrap(errSrv))
	}

}
