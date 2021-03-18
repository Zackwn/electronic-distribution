package core

type Electrons = int

func ValidateElectrons(electrons int) (Electrons, bool) {
	// max number of electrons: 118
	if electrons > 118 {
		return 0, false
	}

	return Electrons(electrons), true
}
