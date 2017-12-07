package coreLift

import "testing"

func TestNewHome(t *testing.T) {
	myHome := NewHome(SetHomeCountFlors(10), SetHomeHeightFloorInMetr(20))
	if myHome.GetHomeCountFloors() != 10 && myHome.GetHomeHeightFloor() != 20 {
		t.Fatal("Error")
	}
}
