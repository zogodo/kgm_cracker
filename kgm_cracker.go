package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"io/ioutil"
	"net/http"
)

// CGO_ENABLED: 0
// GOOS: darwin、freebsd、linux、windows
// GOARCH: 386、amd64、arm
//
//-go:generate bash -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kgm_cracker kgm_cracker.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/kgm_cracker.exe kgm_cracker.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/kgm_cracker_x86.exe kgm_cracker.go"

type Handler struct {
}

func (h Handler) ServeHTTP(rep http.ResponseWriter, req *http.Request) {
	fmt.Printf("收到验证请求，已返回成功激活数据\n")
	rep.Write([]byte("{\"code\":0,\"message\":\"破解成功\"}"))
}

func main() {

	defer fmt.Scan("%c")

	var err error
	var file *os.File
	var handler Handler
	fmt.Printf("Please add\n")
	fmt.Printf("127.0.0.1 yinyuezhushou.com\n")
	fmt.Printf("to your hosts file (C:\\Windows\\System32\\drivers\\etc\\hosts)\n\n")

	switch runtime.GOOS {
	case "windows":

		//exec.Command("explorer.exe", "C:\\Windows\\System32\\drivers\\etc\\").Run()
		filePath := "C:\\Windows\\System32\\drivers\\etc\\hosts"
		fileAdd := "\n127.0.0.1 yinyuezhushou.com"

		file, err = os.OpenFile(filePath, os.O_RDONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Print(err)
			goto end
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Print(err)
			goto end
		}

		fileContentStr := string(data)
		fmt.Sprintf("hosts content:\n%s\n\n", fileContentStr)
		if strings.Contains(fileContentStr, fileAdd) {
			fmt.Printf("hosts was OK\n")
			goto end
		}

		n, err := file.Write([]byte(fileContentStr+fileAdd))
		if err != nil || n < 1 {
			fmt.Print(err)
			goto end
		}
		fmt.Printf("hosts is OK\n")

	default:
		fmt.Printf("Maybe you should run on Windows\n")
	}

end:
	if file != nil {
		file.Close()
	}

	err = http.ListenAndServe("127.0.0.1:8008", handler)
	if err != nil {
		panic(err)
	}
}
