package models

import (
	"encoding/json"
)

type Stream struct {
	Message string `json:"message"`
}
