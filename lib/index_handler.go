package lib

import (
    "dashboard/constants"
)

type IndexHandler struct {
}

func (i IndexHandler) Run() *Response {
    return successWithData(constants.Settings)
}
