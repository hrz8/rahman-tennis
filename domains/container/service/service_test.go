package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

const (
	// change the HOST and PORT first if you want to run the test locally
	// and you had modify the app configuration
	url          = "http://127.0.0.1:8099/api/v1/players/"
	urlContainer = "http://127.0.0.1:8099/api/v1/containers/player/"
	rahmanID     = "4db77dd4-09b0-4633-aed2-a8382e17a748"
	verifiedID   = "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755"
)

func TestNotReadyAndVerified(t *testing.T) {
	var rahmanCanAddBall bool
	r1, err := http.Get(fmt.Sprintf("%s%s", urlContainer, rahmanID))
	assert.Nil(t, err)
	rd1, err := ioutil.ReadAll(r1.Body)
	assert.Nil(t, err)
	iv1 := gjson.Get(string(rd1), "data.isVerified")
	rahmanCanAddBall = iv1.Bool()

	// test for rahman, expected response is OK and still can PUT ball
	// or if test is run multiple times until container full,
	// response will going to be BadRequest and cannot PUT ball
	reqAddBallRahman, err := http.NewRequest("PUT", fmt.Sprintf("%s%s", url, rahmanID), nil)
	assert.Nil(t, err)
	reqAddBallRahman.Header.Set("Content-Type", "application/json")

	rahmanResponse, err := http.DefaultClient.Do(reqAddBallRahman)
	assert.Nil(t, err)
	rahmanResponseData, err := ioutil.ReadAll(rahmanResponse.Body)
	assert.Nil(t, err)

	rahmanStatus := gjson.Get(string(rahmanResponseData), "status")
	if rahmanCanAddBall {
		assert.Equal(t, rahmanStatus.Int(), int64(http.StatusBadRequest))
	} else {
		assert.Equal(t, rahmanStatus.Int(), int64(http.StatusOK))
	}

	// test for verrified user, expected response is BadRequest and cannot PUT ball again
	reqAddBallVerified, err := http.NewRequest("PUT", fmt.Sprintf("%s%s", url, verifiedID), nil)
	assert.Nil(t, err)
	reqAddBallVerified.Header.Set("Content-Type", "application/json")

	verifiedResponse, err := http.DefaultClient.Do(reqAddBallVerified)
	assert.Nil(t, err)
	verifiedResponseData, err := ioutil.ReadAll(verifiedResponse.Body)
	assert.Nil(t, err)

	verifiedStatus := gjson.Get(string(verifiedResponseData), "status")

	assert.Equal(t, verifiedStatus.Int(), int64(http.StatusBadRequest))
}
