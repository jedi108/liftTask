package main

import (
	"lift/coreLift"
	"os"
	"fmt"
)

func main() {
	lift := startHomeLift(os.Args)          // Инициализация дома и лифта
	cc := coreLift.StartConsoleController() // Инициализация консольного контроллера (в будующем возможно нам понадобятся другие протоколы, например для тех поддержки лифтов)
	cc.PrintHomeLiftInfo(lift)              // Информация о доме и лифте
	cc.PrintCommands()                      // Информация о коммандах
	cc.InputLift(lift)                      // Инициализация панели управления лифтом
}

func startHomeLift(args []string) *coreLift.RealLift {

	// Инифиализируем и проверяем входящие параметры дома и лифта
	homeConfig := coreLift.NewHomeConfig().SetHomeOptions(args).SetLiftOptions(args)
	if homeConfig.HasErrors == true {
		for _, v := range homeConfig.Errors {
			fmt.Println(v)
		}
		fmt.Println("> Параметры указываются через пробел. \n>>> Например: 20 5.5 0.5 2.1")
		fmt.Println(">>> 20 этажей, 5.5 метров этаж, скорость лифта 0.5 метров в секунду, 2.1 сек время между открытием и закрытием дверей")
		os.Exit(0)
	}

	// Создаем дом
	home := coreLift.NewHome(homeConfig.FnHomeOptions...)
	isErrorsInHome, errors := home.Validate()
	if isErrorsInHome {
		for _, v := range errors {
			fmt.Println(v)
		}
		os.Exit(0)
	}

	// Создаем лифт
	lift := coreLift.NewLift(home, homeConfig.FnLiftOptions...)
	return lift
}
