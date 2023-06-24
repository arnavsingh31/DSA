package main

func distanceTraveled(mainTank int, additionalTank int) int {
	mileage := 10

	if mainTank < 5 {
		return mileage * mainTank
	}
	distance := 0
	for mainTank >= 5 {
		distance += 5 * mileage

		if additionalTank >= 1 {
			mainTank = mainTank - 5 + 1
			additionalTank--
		} else {
			mainTank = mainTank - 5
		}

	}

	if mainTank > 0 {
		distance += mainTank * mileage
	}

	return distance
}
