package tests

import (
    "dashboard/lib"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLogin(t *testing.T) {
    response := login()
    t.Log(response)
    assert.Equal(t, 200, response.Status)
    data := response.Data.(*lib.LoginResponse)
    assert.NotEmpty(t, data.Token)
    t.Log(data.Token)
}
