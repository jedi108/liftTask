package coreLift

import (
	"fmt"
	"time"
	"context"
	"sync"
)

/*
	Пользовательские параметры лифта
 */
type LiftParams struct {
	SpeedTimeMetrSec float64 //скорость лифта при движении в метрах в секунду (ускорением пренебрегаем, считаем, что когда лифт едет — он сразу едет с определенной скоростью);
	OpenCloseTimeSec float64 //время между открытием и закрытием дверей.
}

type LiftDoing struct {
	sync.RWMutex
	isBusy     bool         // false - лифт свободен, true - лифт занят
	byInTheWay map[int]bool // Этажи по пути
	byOutOfWay map[int]bool // Этажи вне пути TODO: можно сделать очередь, если лифт находится в не предпологаемого пути
}

/*
	Шаблон/Структура работающего лифта
 */
type RealLift struct {
	LiftDoing
	HomeEntry     *HomeParams // привязка к дому
	LiftParams                // параметры лифта(можно увеличивать не меняя код клиентский
	CurrentFloor  int         // текущий этаж
	toFloor       int         // Целевой этаж в данном направлении лифта
	liftDirection func()      // действие над лифтом
}

/*
func NewLift - создание нового лифта, в будующем интерфейс позволяет прикреплять к одному дому неск лифтов
	HomeParams - сопоставляем лифт(ы) с домом
	opts - параметры лифта, можно сделать опциональными в будующем или добавить опций не меняя интерфейс
 */
func NewLift(home *HomeParams, opts ...func(*LiftParams)) *RealLift {
	lift := &RealLift{
		CurrentFloor: 1,
		HomeEntry:    home,
		LiftDoing: LiftDoing{
			byInTheWay: make(map[int]bool),
			byOutOfWay: make(map[int]bool),
		},
	}

	for _, fn := range opts {
		fn(&lift.LiftParams)
	}
	return lift
}

func SetLiftSpeedMetrSec(secInMetr float64) func(*LiftParams) {
	return func(lift *LiftParams) {
		lift.SpeedTimeMetrSec = secInMetr
	}
}

func SetLiftTimeOpenCloseDoors(secondsOC float64) func(*LiftParams) {
	return func(lift *LiftParams) {
		lift.OpenCloseTimeSec = secondsOC
	}
}


func (rl *RealLift) addOnTheWay(Userfloor int) {
	rl.LiftDoing.RLock()
	defer rl.LiftDoing.RUnlock()
	if _, ok := rl.byInTheWay[Userfloor]; !ok {
		rl.byInTheWay[Userfloor] = true
		fmt.Println("#: Лифт по пути остановится на этаже", Userfloor)
	} else {
		fmt.Println("<-= Лифт пока занят =->")
	}
}


func (rl *RealLift) isBusyDoing(Userfloor int) {
	switch {
	case rl.toFloor > rl.CurrentFloor: // Путь вверх
		if Userfloor < rl.toFloor && Userfloor > rl.CurrentFloor {
			// По пути вверх
			rl.addOnTheWay(Userfloor)
		}
	case rl.toFloor < rl.CurrentFloor: //Путь вниз
		if Userfloor > rl.toFloor && Userfloor < rl.CurrentFloor {
			// По пути вниз
			rl.addOnTheWay(Userfloor)
		}
	}
}

func (rl *RealLift) DriveLiftToFloor(Userfloor int) {
	rl.LiftDoing.RLock()
	defer rl.LiftDoing.RUnlock()
	if rl.isBusy {
		rl.isBusyDoing(Userfloor)
		return
	}

	if Userfloor > rl.HomeEntry.CountFloors || Userfloor < rl.HomeEntry.MinFloor {
		fmt.Println("Не возможно отправить лифт на несуществующий этаж", Userfloor)
		return
	}

	go func() {
		rl.toFloor = Userfloor
		rl.liftDirection = rl.GetDirectionFunc()
		rl.doLiftDirection()
	}()
}

func (rl *RealLift) GetDirectionFunc() func() {
	switch {
	case rl.CurrentFloor > rl.toFloor:
		return func() {
			rl.CurrentFloor--
		}
	case rl.CurrentFloor < rl.toFloor:
		return func() {
			rl.CurrentFloor++
		}
	}
	return func() {}
}

func (rl *RealLift) doLiftDirection() {
	for {
		if rl.CurrentFloor == rl.toFloor {
			fmt.Println("Приехали на этаж", rl.CurrentFloor)
			fmt.Println("лифт открыл двери")
			rl.liftRunCommand(func() {}, rl.OpenCloseTimeSec)
			fmt.Println("лифт закрыл двери")
			rl.isBusy = false
			return
		} else {
			rl.isBusy = true
			rl.сheckAndGetPeoplesInWay()
			fmt.Printf("лифт проезжает %d этаж\n", rl.CurrentFloor)
			rl.liftRunCommand(rl.liftDirection, rl.SpeedTimeMetrSec*rl.HomeEntry.HeightFloorMetr)
		}
	}
}

func (rl *RealLift) сheckAndGetPeoplesInWay() {
	if _, ok := rl.byInTheWay[rl.CurrentFloor]; ok {
		fmt.Printf("Так как по пути был вызываемый этаж %d, возможно берем попутчиков\n", rl.CurrentFloor)
		fmt.Println("лифт открыл двери")
		rl.liftRunCommand(func() {}, rl.OpenCloseTimeSec)
		fmt.Println("лифт закрыл двери")
		rl.Lock()
		defer rl.Unlock()
		delete(rl.byInTheWay, rl.CurrentFloor)
	}
}

func (rl *RealLift) liftRunCommand(command func(), intTimeMilliseconds float64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(intTimeMilliseconds)*time.Second)
	defer func() {
		cancel()
	}()
L:
	for {
		select {
		//case <-time.After(1 * time.Second):
		//	continue
		case <-ctx.Done():
			command()
			//fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			break L
		}
	}
}
