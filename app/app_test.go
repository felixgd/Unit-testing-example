package app_test

import (
	"testing"

	. "github.com/felixgd/unitTest-example/app"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app := GetApp()

	assert.NotEmpty(t, app)
}
