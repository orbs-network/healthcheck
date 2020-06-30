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

	status, err := Check("http://localhost:6666/status")
	require.NoError(t, err)
	require.EqualValues(t, "Last Successful Committed Block was too long ago", status.Status)
	require.NotEmpty(t, status.Payload)

	status500, err500 := Check("http://localhost:6666/status.500")
	require.EqualError(t, err500, "unexpected response code 500")
	require.EqualValues(t, "Last Successful Committed Block was too long ago", status500.Status)
	require.NotEmpty(t, status500.Payload)

	statusFailed, errFailed := Check("http://localhost:6666/status.failed")
	require.EqualValues(t, "Not provisioned", statusFailed.Status)
	require.EqualError(t, errFailed, "failed to parse status JSON: invalid character '\\x01' looking for beginning of value")
	require.EqualValues(t, "failed to parse status JSON: invalid character '\\x01' looking for beginning of value", statusFailed.Error)
}
