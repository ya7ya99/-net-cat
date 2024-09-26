package main



func NameExists(name string) bool {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for _, clientName := range clients {
		if clientName == name {
			return false
		}
	}
	for _, v := range name {
		if v == ' ' || v < 32 || v > 126 {
			return false
		}
	}
	return true
}
