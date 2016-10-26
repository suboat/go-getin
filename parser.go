package getin

// 解析getin格式的url
// https://github.com/suboat/getin

import (
	"strconv"
	"strings"
)

const (
	tagEquRoot = "="
	tagEquDeep = "~"
	tagSepRoot = "&"
	tagSepDeep = "+"
)

const (
	GetinKeyTypeString = ""     // 字符串
	GetinKeyTypeInt    = "_int" // 整形
	GetinKeyTypeFloat  = "_flt" // 浮点
	GetinKeyTypeList   = "_lis" // 数组
	GetinKeyTypeMap    = "_obj" // map
)

// 解析key类型
func GetinKeyType(s string) (k string, t string) {
	if len(s) <= 4 {
		// 没有_xyz空间,直接判定为string格式
		k = s
		t = GetinKeyTypeString
		return
	}
	k = s[0 : len(s)-4]
	t = s[len(s)-4 : len(s)]
	switch t {
	case GetinKeyTypeInt, GetinKeyTypeFloat, GetinKeyTypeList, GetinKeyTypeMap:
		break
	default:
		// 无法识别,暂定为字符串
		k = s
		t = GetinKeyTypeString
		break
	}
	return
}

// 解析字符串
func GetGetinMap(s string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	for _, _s := range strings.Split(s, tagSepRoot) {
		if err = parserGetinKv(m, _s, 0); err != nil {
			return
		}
	}
	return
}

// 取key value值
func parserGetinKv(m map[string]interface{}, s string, deep int) (err error) {
	var (
		key    string
		valStr string
	)

	// 解析key
	var i int
	if deep == 0 {
		i = strings.Index(s, tagEquRoot)
	} else {
		i = strings.Index(s, tagEquDeep)
	}
	if i == -1 {
		valStr = s
	} else {
		key = s[0:i]
		valStr = s[i+1 : len(s)]
	}

	// 解析val
	if len(key) > 0 {
		_key, _typ := GetinKeyType(key)
		var (
			deepNow = deep + 1
			vLis    = strings.Split(valStr, tagSepDeep)
		)

		for _, val := range vLis {
			switch _typ {
			case GetinKeyTypeString:
				// 字符串,直接赋值
				//println("-key:", _key, " -val:", val)
				if _, _ok := m[_key]; _ok == false {
					m[_key] = val
				} else {
					// _key已经被定义过,不覆盖
				}
				break
			case GetinKeyTypeInt:
				// 解析成整型
				if _, _ok := m[_key]; _ok == false {
					if _n, _err := strconv.Atoi(val); _err == nil {
						m[_key] = _n
					}
				} else {
					// _key已经被定义过,不覆盖
				}
				break
			case GetinKeyTypeFloat:
				// 解析成浮点
				if _, _ok := m[_key]; _ok == false {
					if _f, _err := strconv.ParseFloat(val, 64); _err == nil {
						m[_key] = _f
					}
				} else {
					// _key已经被定义过,不覆盖
				}
				break
			case GetinKeyTypeList:
				// 解析成数组
				//println("-lis:", _key, " -val:", val)
				var (
					lisSub []string // 现支持字符串数组
				)
				if _v, _ok := m[_key]; _ok == true {
					if _lis, _ok := _v.([]string); _ok == true {
						lisSub = _lis
					} else {
						// _key已被事先定义
					}
				} else {
					lisSub = []string{}
				}
				m[_key] = append(lisSub, val)
				break
			case GetinKeyTypeMap:
				// 解析成map
				var (
					mapSub map[string]interface{}
				)
				if _v, _ok := m[_key]; _ok == true {
					if _m, _ok := _v.(map[string]interface{}); _ok == true {
						mapSub = _m
					} else {
						// _key已被事先定义
					}
				} else {
					mapSub = make(map[string]interface{})
					m[_key] = mapSub
				}
				if mapSub != nil {
					parserGetinKv(mapSub, val, deepNow)
				}
				//println("new obj", _key, val)
				break
			default:
				break
			}
		}
	} else {
		// not key
	}

	return
}
