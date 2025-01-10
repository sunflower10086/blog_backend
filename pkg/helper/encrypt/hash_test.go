package encrypt_test

import (
	"fmt"
	"sunflower-blog-svc/pkg/helper/encrypt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	fmt.Println(encrypt.PasswordHash("123456"))
}
