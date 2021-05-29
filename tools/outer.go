package tools

const englishHiPrefix = "hi"
const englishHelloPrefix = "hello"

func Hi(user string) string {
	return englishHiPrefix + " " + user + "!\n"
}

func Hello() string {
	return englishHelloPrefix + "!\n"
}
