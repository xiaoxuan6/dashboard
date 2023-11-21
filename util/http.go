package util

import (
    "errors"
    "fmt"
    "github.com/avast/retry-go"
    "io/ioutil"
    "net/http"
)

func Get(url string) (string, error) {
    var response string
    err := retry.Do(
        func() error {
            res, err := http.Get(url)
            defer func() {
                if res != nil {
                    _ = res.Body.Close()
                }
            }()

            if err != nil {
                return errors.New(fmt.Sprintf("http get error: %s", err.Error()))
            }

            b, err := ioutil.ReadAll(res.Body)
            if err != nil {
                return errors.New(fmt.Sprintf("read body error: %s", err.Error()))
            }

            response = string(b)
            return nil
        },
        retry.Attempts(3),
        retry.LastErrorOnly(true),
    )

    return response, err
}
