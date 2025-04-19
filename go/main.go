package main

import (
	"log"

	"github.com/foyez/go/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// // Load environment variables from .env file
	// err := godotenv.Load("app.env")
	// if err != nil {
	// 	fmt.Println("Error loading app.env file")
	// }
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	// http.HandleFunc("/", api.HandleRequest)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "3000"
	// }

	// fmt.Printf("Server is running in %s mode on port %s\n", os.Getenv("ENVIRONMENT"), port)
	// panic(http.ListenAndServe(":"+port, nil))
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migration up:", err)
	}

	log.Println("db migrated successfully")
}
