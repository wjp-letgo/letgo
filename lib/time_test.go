package lib
import(
	"testing"
	"fmt"
)
func TestTime(t *testing.T){
	fmt.Println(GetMinuteByTimeLong(1677287479000))
	fmt.Println(HumpName(UnderLineName("wjpNameGee")))
	fmt.Println(Time(),1625189194-1625189188)
	fmt.Println(WordRank([]string{"施华洛世奇手链","项链施华洛世奇手链","施华洛世奇 手链","手链施华洛世奇","ddd"}))
}
