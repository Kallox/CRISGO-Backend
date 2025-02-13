package db

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseConfig() map[string]string {
	var config = make(map[string]string)
	cfg, err := ini.Load("config/server.cfg")
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		return nil
	}
	config["host"] = cfg.Section("database").Key("db_host").String()
	config["port"] = cfg.Section("database").Key("db_port").String()
	config["user"] = cfg.Section("database").Key("db_user").String()
	config["password"] = cfg.Section("database").Key("db_password").String()
	config["name"] = cfg.Section("database").Key("db_name").String()

	return config
}

func DBConnection() {

	config := getDatabaseConfig()
	var dsn = "host=" + config["host"] + " user=" + config["user"] + " password=" + config["password"] + " dbname=" + config["name"] + " port=" + config["port"]
	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
		// This condition is added to avoid the not used error... Remove this line after adding the code to use the DB variable
		if DB != nil {
			fmt.Println("Connection Established")
		}
	}

}
