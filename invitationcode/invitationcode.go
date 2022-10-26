package invitationcode

import "strings"

const (
	BASE    = "SE8D9ZGW2YLT6NBQ7FCP5IKM3JUA4RHV"
	DECIMAL = 32
	PAD     = "X"
	LEN     = 6
)

// EncodeInviteCode 根据id生产invite code
func EncodeInviteCode(inviteId uint64) string {
	id := inviteId
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(inviteId)+i)%DECIMAL])
		}
	}
	return res
}

func DecodeInviteCode(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)
	baseArr := []byte(BASE)       // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += uint64(index) * b
		r++
	}
	return res
}
