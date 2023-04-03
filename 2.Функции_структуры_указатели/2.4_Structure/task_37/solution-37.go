package main

type Ammunition struct {
	On    bool
	Ammo  int
	Power int
}

func (am *Ammunition) Shoot() bool {
	if am.On == true && am.Ammo > 0 {
		am.Ammo--
		return true
	} else {
		return false
	}
}

func (am *Ammunition) RideBike() bool {
	if am.On == true && am.Power > 0 {
		am.Power--
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := new(Ammunition)
}
