package piscine

func FindNextPrime(nb int) int {

	if nb <= 1 {
		return 2
	} else {
		for i := 1; i <= nb; i++ {
			if KB(nb) == false {
				nb++
			} else {
				break
			}
		}
		return nb
	}
}
func KB(nb int) bool {

	if nb <= 1 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}

}
