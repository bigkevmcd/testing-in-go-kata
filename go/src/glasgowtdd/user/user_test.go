package user

import "testing"

type User struct {
	Name string
	Age  int
}

func TestUser(t *testing.T) {
	u := User{Name: "testing", Age: 20}

}
