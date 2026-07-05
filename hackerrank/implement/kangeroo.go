package implement


func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}

	distanceGap := x2 - x1
	speedGap := v1 - v2

	if distanceGap%speedGap == 0 {
		return "YES"
	}

	return "NO"
}