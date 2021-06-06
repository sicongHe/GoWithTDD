package greeting

const chinese = "Chinese"
const english = "English"
const englishHiPrefix = "hi "
const chineseHiPrefix = "你好 "
const englishHelloPrefix = "hello "

func Hi(user string, lang string) string {
	if user == "" {
		user = "gopher"
	}
	return greetingPrefix(lang) + user
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case chinese:
		prefix = chineseHiPrefix
	default:
		prefix = englishHiPrefix
	}
	return
}

func Hello() string {
	return englishHelloPrefix
}
