package jwt

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	token, time, err := GenerateToken("admin", "admin321", 3600)
	fmt.Println("token:",
		token, "time:", time, "err:", err)
}
