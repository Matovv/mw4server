package errors

import "errors"

var ErrSessionFull = errors.New("session is full")
var ErrSessionNotFound = errors.New("session not found")
var ErrSendQueueFull = errors.New("send queue is full")
var ErrDisconnected = errors.New("client is disconnected")
var ErrPrefabUnitNotFound = errors.New("prefab unit was not found")