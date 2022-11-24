package log
import (
	"testing"
)

func TestLog(t *testing.T) {
	DebugPrint("%d nihao", 2)
	DebugPrint("%d nihao", 3)
}