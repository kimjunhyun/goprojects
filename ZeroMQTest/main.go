package gozmq

/*
#cgo pkg-config: libzmq
#include <zmq.h>
*/
import "C"

type ZmqOsSocketType C.SOCKET

func (self ZmqOsSocketType) ToRaw() C.SOCKET {
	return C.SOCKET(self)
}