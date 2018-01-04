package main
 
var testMap = make(map[string]string)
 
func main() {
	for i := 0; i < 33; i++ {
		go func() {
			testMap["aa"] = "bb"
		}()
	}
}