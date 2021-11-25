package uuid

import (
	"github.com/google/uuid"
)

func GenUuid() string {
	v := uuid.New()
	return v.String()
}
