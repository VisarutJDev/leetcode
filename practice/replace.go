package practice

import (
	"errors"
	"strings"
)

// RepalceEmpty function to replace empty between string
func RepalceEmpty(s string) (string, error) {
	l := len(s)
	news := strings.TrimSpace(s)
	news = strings.ReplaceAll(news, " ", "%20")
	if len(news) != l {
		return "", errors.New("not equal true lenght")
	}

	return news, nil
}
