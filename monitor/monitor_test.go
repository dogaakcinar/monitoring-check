package monitor

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStartMonitor(t *testing.T) {
	// Test case 1: ALERT_INTERVAL_SECOND is not set
	os.Setenv("ALERT_INTERVAL_SECOND", "")
	var buf bytes.Buffer
	log.SetOutput(&buf)
	StartMonitor()
	expectedLog := "ALERT_INTERVAL_SECOND is not set or is an empty string"
	if !strings.Contains(buf.String(), expectedLog) {
		t.Errorf("Expected log message '%s', got '%s'", expectedLog, buf.String())
	}

	buf.Reset()

	// Test case 2: Failed to parse ALERT_INTERVAL_SECOND
	os.Setenv("ALERT_INTERVAL_SECOND", "abc")
	StartMonitor()
	expectedLog = "Failed to parse ALERT_INTERVAL_SECOND"
	if !strings.Contains(buf.String(), expectedLog) {
		t.Errorf("Expected log message '%s', got '%s'", expectedLog, buf.String())
	}

	buf.Reset()

	// Test case 3: No requests in the last X seconds
	os.Setenv("ALERT_INTERVAL_SECOND", "60")
	mu.Lock()
	lastRequestTime = time.Now().Add(-61 * time.Second)
	mu.Unlock()
	StartMonitor()
	expectedLog = "No requests in the last"
	if !strings.Contains(buf.String(), expectedLog) {
		t.Errorf("Expected log message containing '%s', got '%s'", expectedLog, buf.String())
	}

	// Reset environment variables after testing
	os.Setenv("ALERT_INTERVAL_SECOND", "")
	os.Unsetenv("ALERT_INTERVAL_SECOND")
}
