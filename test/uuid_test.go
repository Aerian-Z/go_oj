package test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	t.Log(uuid.NewV4().String())
}
