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
	require.EqualValues(t, "200 OK", status.Status)

	blockHeight := status.Payload["BlockStorage.BlockHeight"].(map[string]interface{})["Value"]

	require.EqualValues(t, 3715964, blockHeight)

	status, err = Check("http://localhost:6666/500")
	require.EqualError(t, err, "unexpected response code 500")
	require.EqualValues(t, "500 Internal Server Error", status.Status)

	require.Empty(t, status.Payload)
}
