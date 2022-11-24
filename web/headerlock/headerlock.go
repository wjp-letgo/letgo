package headerlock
import (
	"sync"
)
//读写操作锁
var HeaderMapMutex = sync.RWMutex{}