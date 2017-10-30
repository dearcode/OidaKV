package proto

import (
	"fmt"
)

const (
	EcOk            = 0  //success
	EcTimeout       = -1 // Request timeout
	EcInvPkgHeadLen = -2 // Pkg length should be less than 2MB
	EcInvHeadCrc    = -3 //Invalid Head CRC
)

var (
	ErrTimeout       = initErr(EcTimeout, "timeout")
	ErrInvPkgHeadLen = initErr(EcInvPkgHeadLen, "pkg head length out of range")
	ErrInvHeadCrc    = initErr(EcInvHeadCrc, "invalid head crc")
)

type TableError struct {
	ErrorCode int8   `json:"err_code"`
	ErrorMsg  string `json:"err_msg"`
}

func (e TableError) Error() string {
	if len(e.ErrorMsg) > 0 {
		return fmt.Sprintf("%s(%d)", e.ErrorMsg, e.ErrorCode)
	} else {
		return fmt.Sprintf("error code %d", e.ErrorCode)
	}
}

var tableErrors = make([]TableError, 256)

func initErr(code int8, msg string) TableError {
	tableErrors[int(code)+128] = TableError{code, msg}
	return tableErrors[int(code)+128]
}

func GetErr(code int8) TableError {
	if tableErrors[int(code)+128].ErrorCode != 0 {
		return tableErrors[int(code)+128]
	} else {
		return TableError{code, ""}
	}
}
