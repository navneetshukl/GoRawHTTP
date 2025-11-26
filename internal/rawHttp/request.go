package rawHttp

import (
	"fmt"
	"strings"
)

func ParseRequest(reqBody string) {
	headers := make(map[string]interface{})
	requestHeaders := strings.Split(reqBody, "\r\n")
	// for idx := range requestHeaders {
	// 	fmt.Println(idx,"   ",requestHeaders[idx])
	// }
	n := len(requestHeaders)
	for i := 2; i < n; i++ {
		data:=strings.Split(requestHeaders[i], ":")
		if len(data)>1{
		key:=data[0]
		value:=data[1]
		headers[key]=value
		}

	}
	for k,v:=range headers{
		fmt.Println("key is ",k)
		fmt.Println("value is ",v)
	}

	fmt.Println("auth header is ",headers["Authorization"])

}
