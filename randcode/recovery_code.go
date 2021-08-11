package randcode

import "fmt"


func GenerateRecoveryCode(strSideLen, strs int) []string {
	m := make(map[string]struct{})
	codes := make([]string, 0, strs)
	i := 0
	for i < strs {
		code := fmt.Sprintf("%s-%s", RandStringBytesMaskImprSrcUnsafe(strSideLen), RandStringBytesMaskImprSrcUnsafe(strSideLen))
		if _, ok := m[code]; !ok {
			m[code] = struct{}{}
			codes = append(codes, code)
			i++
		}
	}
	return codes
}
