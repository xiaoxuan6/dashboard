package tests

import (
    "dashboard/pkg/email"
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestEmail(t *testing.T) {
    _ = os.Setenv("MAIL_HOST", "smtp.qq.com")
    _ = os.Setenv("MAIL_PORT", "587")
    _ = os.Setenv("MAIL_USERNAME", "672224976@qq.com")
    _ = os.Setenv("MAIL_PASSWORD", "xxx")
    _ = os.Setenv("MAIL_TO", "1527736751@qq.com")
    _ = os.Setenv("MAIL_SUBJECT", "大富人家一天一颗预估很发达代发")

    opt := email.Option{
        Text: []byte("test"),
    }

    err := email.Send(opt)
    t.Log(err)
    assert.Nil(t, err)
}
