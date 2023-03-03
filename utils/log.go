package utils

import (
	"fmt"
	"os"
	"strconv"
)

func LogPort() {
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("http://localhost:%d\n", portInt)
}
