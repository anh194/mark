package entity

import (
	"fmt"
)


type Essay struct {
	ID string `json:"id"`
	Header string `json:"header"`
	Body string `json:"body"`
}


func (essay Essay) ToString() string{
	return fmt.Sprintf("ID: %s\nHeader: %s\nBody: %s\n", essay.ID, essay.Header, essay.Body)
}