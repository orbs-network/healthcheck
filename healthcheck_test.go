package healthcheck

import (
	"github.com/orbs-network/healthcheck/service"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestCheck(t *testing.T) {
	service.Dummy()

	go http.ListenAndServe("0.0.0.0:6666", nil)

	status, err := Check("http://localhost:6666/metrics")
	require.NoError(t, err)
	require.EqualValues(t, "200 OK", status)

	status, err = Check("http://localhost:6666/500")
	require.EqualError(t, err, "wrong response")
	require.EqualValues(t, "500 Internal Server Error", status)

}
