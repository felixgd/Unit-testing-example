package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/felixgd/unitTest-example/service"
)

func TestReturnText(t *testing.T) {
	mockInternal := NewService("test")

	td := []struct {
		name        string
		expected    string
		expectedErr error
	}{
		{
			name:        "Andres",
			expected:    "Hello! Andres",
			expectedErr: nil,
		},
		{
			name:        "Felix",
			expected:    "Hello! Felix",
			expectedErr: nil,
		},
		{
			name:        "testestetestetstett",
			expected:    "",
			expectedErr: errors.New("Name is too long."),
		},
		{
			name:        "Agustiniano",
			expected:    "",
			expectedErr: errors.New("Name is too long."),
		},
	}
	for _, d := range td {
		s, err := mockInternal.ReturnText(d.name)

		if d.expectedErr != nil {
			assert.Equal(t, d.expectedErr, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, d.expected, s)
	}

}
