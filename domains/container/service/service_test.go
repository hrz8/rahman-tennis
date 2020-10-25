package service

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/kerti/evm/02-kitara-store/model"
	"github.com/stretchr/testify/assert"
)

const (
	url = "http://localhost:8080/orders/process"
)

func TestOrderConcurrency(t *testing.T) {
	payload1, _ := json.Marshal(models..OrderProcessInput{
		OrderID: uuid.FromStringOrNil("5b27773a-efce-4e21-8474-2694ebdaa084"),
	})

	// Run both requests concurrently
	r1, err := http.Post(url, "application/json", body1)
	assert.Nil(t, err)
	response1 := r1

	// If one of the responses returns OK, the other one should not return OK
	if response1.StatusCode == http.StatusOK {
		assert.NotEqual(t, response1.StatusCode, http.StatusOK)
	}
}
