package comparisons

func SignInt(val int32) int32 {
	return SignLong(int64(val))
}

func SignLong(val int64) int32 {
	if val > 0 {
		return 1
	} else if val < 0 {
		return -1
	} else {
		return 0
	}
}

func SignFloat(val float32, gFlag bool) int32 {
	return SignDouble(float64(val), gFlag)
}

func SignDouble(val float64, gFlag bool) int32 {
	if val > 0 {
		return 1
	} else if val < 0 {
		return -1
	} else if val == 0 {
		return 0
	} else if gFlag {
		return 1
	} else {
		return -1
	}
}
