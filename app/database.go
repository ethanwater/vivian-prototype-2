package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
)

func EstablishLinkDatabase(ctx context.Context, app *App) *sql.DB {
	toml, err := toml.LoadFile("config.toml")

	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	config := mysql.Config{
		User:   "root",
		Net:    "tcp",
		Addr:   toml.Get("database.address").(string),
		DBName: toml.Get("database.name").(string),
	}

	database, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		app.Logger(ctx).Error("vivian: cant open database", "db", config.DBName)
		os.Exit(1)
	}

	ping := database.Ping()
	if ping != nil {
		app.Logger(ctx).Error("vivian: cant connect to database", "db", config.DBName, err)
		os.Exit(2)
	} else {
		app.Logger(ctx).Debug("vivian: connected to database", "db", config.DBName, "address", config.Addr)
	}

	return database
}