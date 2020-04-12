package extraction
import "testing"

func TestShow(t *testing.T) {
    want := "Hello, world."
	if got := Show(); got != want {
	    t.Errorf("Hello() = %q, want %q", got, want)
    }
}
