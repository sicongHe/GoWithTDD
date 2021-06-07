package integers

func Repeat(tar string, repeatedTimes int) (ret string) {
	for i := 0; i < repeatedTimes; i++ {
		ret += tar
	}
	return
}
