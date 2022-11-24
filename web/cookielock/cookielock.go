package cookielock
import (
	"sync"
)
//读写操作锁
var CookieMapMutex = sync.RWMutex{}