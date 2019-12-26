/*
@Time : 2019/10/3 22:59
@Author : Hermes
@File : errcheck
@Description:
*/
package netUtil

import (
	"io"
	"net"
	"os"
	"syscall"
)

const TIMEOUT_ERROR = 1
const EOF_ERROR = 2
const Temporary_ERROR = 3
const DNS_ERROR = 4
const SYSCALL_ECONNREFUSED_ERROR = 5
const SYSCALL_ETIMEDOUT_ERROR = 6
const UNKONW_SYSCALL_ERROR = 7
const UNKONW_ERROR = 8

func GetNetErrorType(err error) int {
	if err == io.EOF {
		return EOF_ERROR
	}

	netErr, ok := err.(net.Error)
	if ok {
		if netErr.Timeout() {
			return TIMEOUT_ERROR
		} else if netErr.Temporary() {
			return Temporary_ERROR
		}
	}

	opErr, ok := netErr.(*net.OpError)
	if ok {
		switch t := opErr.Err.(type) {
		case *net.DNSError:
			return DNS_ERROR
		case *os.SyscallError:
			if errno, ok := t.Err.(syscall.Errno); ok {
				switch errno {
				case syscall.ECONNREFUSED:
					return SYSCALL_ECONNREFUSED_ERROR
				case syscall.ETIMEDOUT:
					return SYSCALL_ETIMEDOUT_ERROR
				}
			}
		default:
			return UNKONW_SYSCALL_ERROR
		}
	}

	return UNKONW_ERROR
}
