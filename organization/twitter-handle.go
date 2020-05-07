package organization

import (
	"fmt"
	"strings"
)

type TwitterHandle string

func (handle TwitterHandle) Valid() bool {
	return strings.HasPrefix(string(handle), "@")
}

func (handle TwitterHandle) RedirectUrl() string {
	cleanHandle := strings.TrimPrefix(string(handle), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandle)
}
