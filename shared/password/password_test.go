package password

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	hashedPwd, _ := HashAndSalt([]byte("qwerty"))
	fmt.Println(hashedPwd)
	fmt.Println(ComparePasswords(hashedPwd, []byte("qwerty")))
}
