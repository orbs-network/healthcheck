package healthcheck

import (
	"context"
	"github.com/orbs-network/healthcheck/service"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	service.Dummy()

	go http.ListenAndServe("0.0.0.0:6666", nil)

	status, err := Check("http://localhost:6666/status")
	require.NoError(t, err)
	require.EqualValues(t, "Last Successful Committed Block was too long ago", status.Status)
	require.NotEmpty(t, status.Payload)
	require.Empty(t, status.Error)

	status500, err500 := Check("http://localhost:6666/status.500")
	require.EqualError(t, err500, "unexpected response code 500")
	require.EqualValues(t, "Last Successful Committed Block was too long ago", status500.Status)
	require.NotEmpty(t, status500.Payload)
	require.EqualValues(t, "unexpected response code 500", status500.Error)

	statusFailed, errFailed := Check("http://localhost:6666/status.failed")
	require.EqualValues(t, "Not provisioned", statusFailed.Status)
	require.EqualError(t, errFailed, "failed to parse status JSON: invalid character '\\x01' looking for beginning of value")
	require.EqualValues(t, "failed to parse status JSON: invalid character '\\x01' looking for beginning of value", statusFailed.Error)
}

func TestDumpToDisk(t *testing.T) {
	successfulOutput := "./tmp/success.txt"
	os.Remove(successfulOutput)

	ctx, cancelSuccess := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelSuccess()

	DumpToDisk(ctx, successfulOutput, []byte("success"), WRITE_MODE)

	data, err := ioutil.ReadFile(successfulOutput)
	require.NoError(t, err)
	require.Equal(t, data, []byte("success"))

	timedOutOutput := "./tmp/timed-out.txt"
	os.Remove(timedOutOutput)

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancelTimeout()

	DumpToDisk(ctx, timedOutOutput, []byte("lost data"), WRITE_MODE)

	_, err = ioutil.ReadFile(timedOutOutput)
	require.EqualError(t, err, "open ./tmp/timed-out.txt: no such file or directory")
}

func TestTimeoutValues(t *testing.T) {
	// the sum of timeout deadlines should be less than 30s

	require.Less(t, (CHECK_TIMEOUT + DISK_TIMEOUT*2).Nanoseconds(), 30*time.Second.Nanoseconds())
}
