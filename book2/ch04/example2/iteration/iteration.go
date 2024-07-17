package iteration

func Repeat(character string, nTimes uint8) string {
	var repeated string = ""
	for i := uint8(0); i < nTimes; i++ {
		repeated += character
	}
	return repeated
}
