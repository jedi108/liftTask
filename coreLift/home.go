package coreLift

const MinFloor int = -1 //Минимальный этаж - парковка (ps: Все данные, которых не хватает в задаче, можно выбрать на свое усмотрение.)

/*
	Пользовательские параметры дома для подключаемых лифтов
 */
type HomeParams struct {
	CountFloors     int     // кол-во этажей в подъезде — N (от 5 до 20);
	HeightFloorMetr float64 // высота одного этажа;
	MinFloor        int     // минимальный этаж
}

func NewHome(opts ...func(*HomeParams)) *HomeParams {
	home := &HomeParams{MinFloor: MinFloor}
	for _, fn := range opts {
		fn(home)
	}
	return home
}

func (he *HomeParams) Validate() (bool, []string) {
	var (
		hasErrors bool = false
		Errors    []string
	)
	if he.CountFloors < 5 {
		hasErrors = true
		Errors = append(Errors, "Количество этажей меньше 5")
	}
	if he.CountFloors > 20 {
		hasErrors = true
		Errors = append(Errors, "Количество этажей больше 20")
	}
	return hasErrors, Errors
}

func SetHomeCountFlors(floors int) func(*HomeParams) {
	return func(entry *HomeParams) {
		entry.CountFloors = floors
	}
}

func SetHomeHeightFloorInMetr(metr float64) func(*HomeParams) {
	return func(entry *HomeParams) {
		entry.HeightFloorMetr = metr
	}
}

func (he *HomeParams) GetHomeCountFloors() int {
	return he.CountFloors
}

func (he *HomeParams) GetHomeHeightFloor() float64 {
	return he.HeightFloorMetr
}
