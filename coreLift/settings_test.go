package coreLift

import "testing"

func TestSetHomeOptionsValid(t *testing.T) {
	var args []string
	args = append(args, "exe")
	args = append(args, "1")
	args = append(args, "10.1")

	homeConfig := NewHomeConfig().SetHomeOptions(args)

	if homeConfig.HasErrors == true {
		t.Fatal("we have errors")
	}

	home := NewHome(homeConfig.FnHomeOptions...)

	if home.CountFloors != 1 {
		t.Fatal("not 1")
	}

	if home.HeightFloorMetr != 10.1 {
		t.Fatal("not 10.1")
	}
}

func TestSetHomeOptionsNoValid(t *testing.T) {
	var args []string
	args = append(args, "exe")

	homeConfig := NewHomeConfig().SetHomeOptions(args)
	if homeConfig.HasErrors == false {
		t.Fatal("we have no errors")
	}
}
