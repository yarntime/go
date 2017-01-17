package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "YXBpVmVyc2lvbjogdjEKa2luZDogU2VydmljZQptZXRhZGF0YToKICBuYW1lOiB7eyB0ZW1wbGF0ZSAiZnVsbG5hbWUiIC4gfX0KICBsYWJlbHM6CiAgICBjaGFydDogInt7IC5DaGFydC5OYW1lIH19LXt7IC5DaGFydC5WZXJzaW9uIH19IgpzcGVjOgogIHR5cGU6IHt7IC5WYWx1ZXMuc2VydmljZS50eXBlIH19CiAgcG9ydHM6CiAgLSBwb3J0OiB7eyAuVmFsdWVzLnNlcnZpY2UuZXh0ZXJuYWxQb3J0IH19CiAgICB0YXJnZXRQb3J0OiB7eyAuVmFsdWVzLnNlcnZpY2UuaW50ZXJuYWxQb3J0IH19CiAgICBwcm90b2NvbDogVENQCiAgICBuYW1lOiB7eyAuVmFsdWVzLnNlcnZpY2UubmFtZSB9fQogIHNlbGVjdG9yOgogICAgYXBwOiB7eyB0ZW1wbGF0ZSAiZnVsbG5hbWUiIC4gfX0K"
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	l, _ := base64.StdEncoding.Decode(base64Text, []byte(str))
	fmt.Printf("base64: %s\n", base64Text[:l])
}
