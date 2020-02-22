package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://raw.githubusercontent.com/ray-bang8/piscine-go/master/atoibase.go")

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
