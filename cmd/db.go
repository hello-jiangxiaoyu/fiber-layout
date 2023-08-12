package cmd

import (
	"fiber/pkg/global"
	"fmt"
	"os"
	"strings"
)

func autoMigrateDB() {
	fmt.Println("Start Gorm auto migrate database.")
	migrateList := []any{}

	if err := global.DB.Debug().AutoMigrate(migrateList...); err != nil {
		os.Exit(ExitMigrate)
		return
	}

	fmt.Println("Gorm auto migrate ok")
}

func createDbTables() {
	if err := initGlobal(); err != nil {
		return
	}
	sqlBytes, err := os.ReadFile("./deploy/sql/create.sql")
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitReadFile)
	}

	statements := strings.Split(string(sqlBytes), ";")
	for _, sql := range statements {
		if strings.TrimSpace(sql) != "" {
			if err = global.DB.Exec(sql).Error; err != nil {
				fmt.Println(err)
				os.Exit(ExitExecSql)
			}
		}
	}
}
