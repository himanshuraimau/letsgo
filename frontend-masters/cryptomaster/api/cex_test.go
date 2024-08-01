package api_test

import (
	"testing"

	"enghimanshu.tech/go/cryptomaster/api"
)

func TestXxx(t *testing.T) {
	_,err:=api.GetRate("")
     
	if err ==nil{
		t.Errorf("Expected error, got nil")
	}
}