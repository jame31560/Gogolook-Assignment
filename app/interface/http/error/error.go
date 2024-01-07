package http_error

import "task/app/pkg/status"

type ErrRes struct {
	ErrMsg string `json:"error_message"`
}

func NewErrRes(st status.Status) ErrRes {
	return ErrRes{
		ErrMsg: st.ErrMsg(),
	}
}
