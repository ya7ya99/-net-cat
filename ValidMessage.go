package main

func ValidMessage(message string) bool {

	if message == "" {
		return false
	}
	for _, v := range message {
		if  v < 32 || v > 126 {
			return false
		}
	}
	return true
}
