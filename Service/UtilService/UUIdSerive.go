package UtilService

import uuid "github.com/satori/go.uuid"

func UUId() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}
