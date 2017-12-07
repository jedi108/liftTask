package coreLift

import (
	"testing"
)

func TestGetLiftInit(t *testing.T) {
	myHome := NewHome(SetHomeCountFlors(10), SetHomeHeightFloorInMetr(20))
	testLift := NewLift(myHome, SetLiftSpeedMetrSec(2), SetLiftTimeOpenCloseDoors(3))
	if testLift.CurrentFloor != 1 {
		t.Fatal("failed init lift not in floor 1")
	}
	if testLift.SpeedTimeMetrSec != 2 {
		t.Fatal("failed init SetLiftSpeedMetrSec", testLift.SpeedTimeMetrSec)
	}
	if testLift.OpenCloseTimeSec != 3 {
		t.Fatal("failed init SetLiftTimeOpenCloseDoors", testLift.OpenCloseTimeSec)
	}
}

