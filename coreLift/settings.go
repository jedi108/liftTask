/**
	Консольный конфигуратор дома и лифта
 */
package coreLift

import (
	"strconv"
)

type HomeLiftSettings struct {
	FnHomeOptions []func(entry *HomeParams)
	FnLiftOptions []func(entry *LiftParams)
	Errors        []string
	HasErrors     bool
}

func NewHomeConfig() *HomeLiftSettings {
	return &HomeLiftSettings{}
}

func (gs *HomeLiftSettings) SetHomeOptions(args []string) *HomeLiftSettings {
	var (
		valInt  int
		valFl64 float64
	)

	if len(args) > 1 && gs.validateSetInt(args[1], &valInt) == nil {
		gs.FnHomeOptions = append(gs.FnHomeOptions, SetHomeCountFlors(valInt))
	} else {
		gs.HasErrors = true
		gs.Errors = append(gs.Errors, "не указано кол-во этажей в подъезде")
	}

	if len(args) > 2 && gs.validateSetFloat64(args[2], &valFl64) == nil {
		gs.FnHomeOptions = append(gs.FnHomeOptions, SetHomeHeightFloorInMetr(valFl64))
	} else {
		gs.HasErrors = true
		gs.Errors = append(gs.Errors, "не указана высота одного этажа")
	}

	return gs
}

func (gs *HomeLiftSettings) SetLiftOptions(args []string) *HomeLiftSettings {
	var valFloat64 float64
	if len(args) > 3 && gs.validateSetFloat64(args[3], &valFloat64) == nil {
		gs.FnLiftOptions = append(gs.FnLiftOptions, SetLiftSpeedMetrSec(valFloat64))
	} else {
		gs.HasErrors = true
		gs.Errors = append(gs.Errors, "не указана скорость лифта при движении в метрах в секунду")
	}

	if len(args) > 4 && gs.validateSetFloat64(args[4], &valFloat64) == nil {
		gs.FnLiftOptions = append(gs.FnLiftOptions, SetLiftTimeOpenCloseDoors(valFloat64))
	} else {
		gs.HasErrors = true
		gs.Errors = append(gs.Errors, "не указана время между открытием и закрытием дверей")
	}
	return gs
}

func (gs *HomeLiftSettings) validateSetInt(arg string, valStructInt *int) (error) {
	val, err := strconv.Atoi(arg)
	if err == nil {
		*(valStructInt) = val
	}
	return err
}

func (gs *HomeLiftSettings) validateSetFloat64(arg string, valFloat64 *float64) (error) {
	float64val, err := strconv.ParseFloat(arg, 64)
	if err == nil {
		*(valFloat64) = float64val
	}
	return err
}
