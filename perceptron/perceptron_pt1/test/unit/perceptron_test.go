package perceptrontest

import (
	"testing"

	"../../perceptron"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	testing := perceptron.New(3, .0003)

	if result == nil {
		t.Errorf("Generated QRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}

}
