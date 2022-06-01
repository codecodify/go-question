package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUuid(t *testing.T) {
	u := uuid.NewV4()
	fmt.Println(u.String())
}
