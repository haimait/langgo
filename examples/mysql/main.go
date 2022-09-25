package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
)

func main() {
	langgo.Run(&mysql.Instance{})
	if core.EnvName == core.Development {
		err := mysql.Main().AutoMigrate(&Account{})
		if err != nil {
			log.Logger("app", "main").Error().Err(err).Interface("err", err).Send()
		}
	}
	acc := Account{
		Name: "chihuo",
	}
	mysql.Main().Create(&acc)
	acc.Name = "famingjia"
	mysql.Main().Save(&acc)
	newAcc := Account{}
	mysql.Main().First(&newAcc, "id=?", acc.ID)
	log.Logger("app", "main").Info().Interface("newAcc", newAcc).Send()
	mysql.Main().Unscoped().Delete(&Account{}, newAcc.ID)
}
