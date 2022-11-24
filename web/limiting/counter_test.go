package limiting
import(
	"testing"
	"fmt"
	"time"
)

func TestCounter (t *testing.T){
	k:=NewLeakyBucket(2, 0.01)
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	time.Sleep(1000*time.Millisecond)
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	/*
	k:=NewFlowRollingCounter(2,2, 1000)
	//fmt.Println(k)
	time.Sleep(100*time.Millisecond)
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	time.Sleep(1000*time.Millisecond)
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	time.Sleep(1000*time.Millisecond)
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k.Pass())
	fmt.Println(k)
	*/
	/*
	f:=NewFlowCounter(1, 10000)
	fmt.Println(f.Pass())
	fmt.Println(f.Pass())
	time.Sleep(10*time.Second)
	fmt.Println(f.Pass())
	fmt.Println(f.Pass())
	fmt.Println(f.Pass())
	fmt.Println(f.Pass())
	*/
}