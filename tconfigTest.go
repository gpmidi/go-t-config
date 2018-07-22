package go_t_config

import "testing"

func getTestingLogger(t *testing.T) (*TLogI, error) {

}

func TestNewTConfig(t *testing.T) {
	l, err := getTestingLogger(t)
	if err != nil {
		t.Error("Got error", err)
	}
	var tc *TConfigI
	tc, err = NewTConfig(l)

}
