package resttk

import (
	"testing"
)

func TestCookieCreation(t *testing.T) {

	cookie := Cookie("testname", "testvalue", "/localhost", "localhost.localdomain", 300, true, true)
	if cookie == nil {
		t.Errorf("Cookie should not be nil")
	}
}
