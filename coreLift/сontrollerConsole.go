package coreLift

import (
	"fmt"
	"strings"
	"strconv"
)

type ConsoleController struct{}

func StartConsoleController() *ConsoleController {
	return &ConsoleController{}
}

func (cc *ConsoleController) PrintHomeLiftInfo(lift *RealLift) {
	fmt.Println("____________________________________________")
	fmt.Println("этажей в доме = ", lift.HomeEntry.CountFloors)
	fmt.Println("высота этажа в метрах = ", lift.HomeEntry.HeightFloorMetr)
	fmt.Println("скорость лифта при движении в метрах в секунду = ", lift.SpeedTimeMetrSec)
	fmt.Println("время между открытием и закрытием дверей = ", lift.OpenCloseTimeSec)
}

func (cc *ConsoleController) PrintCommands() {
	fmt.Println("____________________________________________")
	fmt.Println("Введите `q` \t - для выхода")
	fmt.Println("Введите `cN` \t - вызовет лифт из подъезда на ваш этаж N, например c10 - вызов лифта с 10го этажа")
	fmt.Println("Введите `N` \t - внутри лифта выбрать N номер этажа")
}

/*
Возможный ввод пользователя:
	- вызов лифта на этаж из подъезда;
	- нажать на кнопку этажа внутри лифта.
	- ps: выход из программы)))
*/
func (cc *ConsoleController) InputLift(lift *RealLift) {
	c := make(chan bool)
	go func() {
		var input string
		var toFloor int
		for {
			input = ""
			fmt.Scanln(&input)
			switch {
			case strings.ToLower(input) == "q":
				//для выхода из программы (просто закрываем канал)
				close(c)

			case len(input) > 0 && strings.ToLower(input[:1]) == "c":
				//вызовет лифт на этаж N, например c10 - вызов лифта с 10го этажа
				toFloor, err := strconv.Atoi(input[1:])
				if err == nil {
					lift.DriveLiftToFloor(toFloor)
				} else {
					fmt.Println("error enter", input)
				}

			case func(input string) bool {
				// внутри лифта выбрать N номер этажа
				var err error
				toFloor, err = strconv.Atoi(input)
				return err == nil
			}(input):
				lift.DriveLiftToFloor(toFloor)
			default:
				fmt.Print("Ошибка ввода\n")
			}
		}
	}()
	<-c
}
