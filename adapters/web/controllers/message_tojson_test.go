package controllers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := messageToJson(msg)
	require.Equal(t, `{"message":"Hello Json"}`, string(result))
}