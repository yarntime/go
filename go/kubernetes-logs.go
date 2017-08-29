package main

import (
	"async/client"
	"bufio"
	"fmt"
	"k8s.io/client-go/pkg/api/v1"
)

func main() {

	fmt.Println("starting")

	c := client.NewK8sClint()

	req := c.GetLogs("NAMESPACE", "POD_NAME", &v1.PodLogOptions{
		Follow:     true,
		Timestamps: false,
		Container:  "CONTAINER_NAME",
		TailLines:  newInt64(100),
	})

	stream, err := req.Stream()
	if err != nil {
		fmt.Sprintf("Error opening stream to %s/%s: %s\n", "NAMESPACE", "POD_NAME", "CONTAINER_NAME")
		return
	}
	defer stream.Close()

	reader := bufio.NewReader(stream)

	for {
		line, _ := reader.ReadBytes('\n')

		str := string(line)

		fmt.Print(str)
	}
}

func newInt64(val int64) *int64 {
	p := new(int64)
	*p = val
	return p
}
