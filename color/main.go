package color

import "fmt"

const (
	RED   = "\033[31m"
	GREEN = "\033[32m"
	RESET = "\033[0m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", RED, s, RESET)
}

func Redf(s string, args ...interface{}) string {
	s = fmt.Sprintf(s, args...)
	return fmt.Sprintf("%s%s%s", RED, s, RESET)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", GREEN, s, RESET)
}

func Greenf(s string, args ...interface{}) string {
	s = fmt.Sprintf(s, args...)
	return fmt.Sprintf("%s%s%s", GREEN, s, RESET)
}
