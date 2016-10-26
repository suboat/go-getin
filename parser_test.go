package getin

import (
	"encoding/json"
	"testing"
)

// 测试用例来自https://github.com/suboat/getin
var testUrl1 = "limit_int=10&skip_int=20&plat=china&sort_lis=active&sort_lis=create&rate_flt=10.5"
var testUrl2 = "plat=china&key_obj=cate~person&key_obj=group~social"
var testUrl3 = "plat=china&key_obj=class_lis~music+class_lis~drawing"
var testUrl3_1 = "plat=china&key_obj=class_lis~music&key_obj=class_lis~drawing"
var testUrl4 = "key_obj=magic_obj~name~jack+magic_obj~method_obj~want~jump"

// 测试解析
func Test_Parser(t *testing.T) {
	for _, s := range []string{testUrl1, testUrl2, testUrl3, testUrl3_1, testUrl4} {
		if m, err := GetGetinMap(s); err != nil {
			t.Fatal(err)
		} else {
			if b, _err := json.Marshal(m); _err == nil {
				t.Logf("%s <- %s", string(b), s)
			}
		}
	}

}
