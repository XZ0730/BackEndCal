package main

import (
	"testing"

	"github.com/XZ0730/tireCV/utils"
)

func TestDetect(t *testing.T) {
	s, err := utils.Calculate("1+(asin(0.5))+2^2+ln(10)")
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(s)

}
