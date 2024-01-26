package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	} else {
		return true
	}
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	} else {
		return false
	}
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
		return true
	} else {
		return false
	}
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

	/*
	   If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep. The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.
	   If Annalyn does not have her dog then she and the prisoner must be very sneaky! Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.

	*/
	if !archerIsAwake && petDogIsPresent {
		return true
	} else if !archerIsAwake && !knightIsAwake && prisonerIsAwake {
		return true
	} else {
		return false
	}
	panic("Please implement the CanFreePrisoner() function")
}
