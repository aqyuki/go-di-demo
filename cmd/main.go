package main

import (
	"context"
	"database/sql"
	"di/model"
	"di/repository"
	"di/service"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	create_table = `
	CREATE TABLE task (
		title VARCHAR(50) NOT NULL,
		note VARCHAR(100)
	);
	`
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// mysqlへの接続の作成
	db, err := sql.Open("mysql", os.Getenv("dsn"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if os.Getenv("create") == "true" {
		// 初期設定
		_, err = db.Exec(create_table)
		if err != nil {
			log.Fatal(err)
		}
	}

	tr := repository.NewTaskRepository(db)
	ts := service.NewTaskService(tr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ts.CreateTask(ctx, &model.Task{Title: "Sample1", Description: "Sample task"})
}
