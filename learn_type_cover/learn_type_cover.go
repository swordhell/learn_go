package main

func main() {
	var v1 int32 = 100
	var v2 int32 = 200

	var rdata interface{} = v1
	var value interface{} = v2
	if rdata.(uint32) == uint32(value.(float64)) {
		return
	}

}
