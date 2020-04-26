package utils

import uuid "github.com/satori/go.uuid"

func GenerateUUid() string {
	uid := uuid.NewV4()
	sessionId := uid.String()
	return sessionId
}
