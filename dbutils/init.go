package dbutils

import (
	"fmt"
	"log"
	"os"

	"gaurav.kapil/tigerhall/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func Inititialize() (err error) {

	pass := os.Getenv("PGPASS")
	user := os.Getenv("PGUSER")
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		user,
		pass)
	DbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if os.Getenv("MIGRATE_1") == "1" {
		Migration_1()
	}
	if os.Getenv("MIGRATE_DROPALL") == "1" {
		Migration_Dropall()
	}
	return
}

func Migration_1() {
	log.Printf("(Migration_1) is configured for this run.")
	DbConn.Migrator().CreateTable(model.TigerData{})
	DbConn.Migrator().CreateTable(model.Sighting{})
	DbConn.Migrator().CreateTable(model.LoginData{})
	DbConn.Migrator().CreateTable(model.UserDataWithPassword{})
	log.Printf("(Migration_1) has finished.")
}

func Migration_Dropall() {
	log.Printf("(Migration_2) is configured for this run.")
	DbConn.Migrator().DropTable(model.TigerData{})
	DbConn.Migrator().DropTable(model.Sighting{})
	DbConn.Migrator().DropTable(model.LoginData{})
	DbConn.Migrator().DropTable(model.UserDataWithPassword{})
	log.Printf("(Migration_2) has finished.")
}
