package lib

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

//SubString 字符串截取
//参数str：输入字符串
//参数start：起始位置
//参数end：结束位置
//返回值：截完后的字符串
func SubString(str string, start int, end int) string {
	arr := []rune(str)
	if end == -1 {
		return string(arr[start:])
	} else {
		return string(arr[start:end])
	}

}

//FirstToUpper 将首字母转化为大写
//参数str：输入字符串
//返回值：首字母大写字符串
func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//FirstToLower 首字母小写
//参数str：输入字符串
//返回值：首字母小写字符串
func FirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 65 && strArry[0] <= 90 {
		strArry[0] += 32
	}
	return string(strArry)
}

//InStringArray 是否包含字符
func InStringArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//InInt64Array 是否包含int64
func InInt64Array(need int64, needArr []int64) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//InIntArray 是否包含int
func InIntArray(need int, needArr []int) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//ResolveAddress 解析地址
func ResolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":1122"
	case 1:
		return fmt.Sprintf("%s:1122", addr[0])
	case 2:
		return fmt.Sprintf("%s:%s", addr[0], addr[1])
	default:
		return fmt.Sprintf("%s:%s", addr[0], addr[1])
	}
}

//ReplaceIndex 替换指定第n个处
func ReplaceIndex(s, old, new string, n int) string {
	arr := strings.Split(s, old)
	r := ""
	for i, v := range arr {
		if v != "" {
			if i == n {
				r += v + new
			} else {
				r += v + old
			}
		}
	}
	return r
}

//IsFloat 判断字符串是否是一个小数
func IsFloat(s string) bool {
	match1, _ := regexp.MatchString(`^[\+-]?\d*\.\d+$`, s)
	match2, _ := regexp.MatchString(`^[\+-]?\d+\.\d*$`, s)
	return match1 || match2
}

//IsInt  判断字符串是否是一个整型
func IsInt(s string) bool {
	match1, _ := regexp.MatchString(`^[\+-]?\d+$`, s)
	return match1
}

//正则表达式相关函数
func Reg(reg string, content string, index int) string {
	r, _ := regexp.Compile(reg)
	if r != nil {
		match := r.FindAllStringSubmatch(content, -1)
		if len(match) > 0 && len(match[0]) > index {
			return match[0][index]
		}
	}
	return ""
}

//正则替换
//参数reg：正则表达式
//参数new_str：替换成字符串
//参数content：待匹配字符串
//返回值：替换后的字符串
func RegReplace(reg, newStr, content string) string {
	r, _ := regexp.Compile(reg)
	if r != nil {
		return r.ReplaceAllString(content, newStr)
	}
	return ""
}

//GetRootDomain 获得根域名
func GetRootDomain(ul string) (root string) {
	pattern := "([a-z0-9--]{1,200})\\.(ac\\.cn|bj\\.cn|sh\\.cn|tj\\.cn|cq\\.cn|he\\.cn|sn\\.cn|sx\\.cn|nm\\.cn|ln\\.cn|jl\\.cn|hl\\.cn|js\\.cn|zj\\.cn|ah\\.cn|fj\\.cn|jx\\.cn|sd\\.cn|ha\\.cn|hb\\.cn|hn\\.cn|gd\\.cn|gx\\.cn|hi\\.cn|sc\\.cn|gz\\.cn|yn\\.cn|gs\\.cn|qh\\.cn|nx\\.cn|xj\\.cn|tw\\.cn|hk\\.cn|mo\\.cn|xz\\.cn" +
		"|com\\.cn|com\\.net|net\\.cn|org\\.cn|gov\\.cn|我爱你|在线|中国|网址|网店|中文网|公司|网络|集团" +
		"|com|cn|cc|org|net|xin|xyz|vip|shop|top|club|wang|fun|info|online|tech|store|site|ltd|ink|biz|group|link|work|pro|mobi|ren|kim|name|tv|red" +
		"|cool|team|live|pub|company|zone|today|video|art|chat|gold|guru|show|life|love|email|fund|city|plus|design|social|center|world|auto):?\\d*$"
	root = Reg(pattern, ul, 0)
	if root != "" {
		arr := strings.Split(root, ":")
		root = arr[0]
	}
	return root
}

//UnderLineName 驼峰转下划线命名
func UnderLineName(name string) string {
	var buf []byte
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(r)))
		} else {
			buf = append(buf, byte(r))
		}
	}
	return string(buf)
}

//HumpName 下划线命名转驼峰
func HumpName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

//WordRankResult 词频返回
type WordRankResult struct {
	Text  string  `json:"text"`
	Count int     `json:"count"`
	Rank  float32 `json:"rank"`
}

//String
func (e WordRankResult) String() string {
	return ObjectToString(e)
}

type WordRankResults []WordRankResult

func (s WordRankResults) Len() int {
	return len(s)
}
func (s WordRankResults) Less(i, j int) bool {
	return s[i].Count < s[j].Count
}
func (s WordRankResults) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//统计词频
func WordRank(arrWords []string) WordRankResults {
	var lst WordRankResults
	var rmap map[string]int = make(map[string]int)
	var sum int = 0
	for _, v := range arrWords {
		v = strings.Trim(v, " ")
		if v != "" && utf8.RuneCountInString(v) > 1 {
			rmap[v] += 1
			sum += 1
		}
	}
	for k2, v2 := range rmap {
		lst = append(lst, WordRankResult{
			Text:  k2,
			Count: v2,
			Rank:  float32(v2) / float32(sum),
		})
	}
	sort.Sort(sort.Reverse(lst))
	return lst
}
