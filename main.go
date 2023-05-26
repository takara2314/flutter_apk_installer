package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func detector() {
	out, err := exec.Command("flutter", "devices").Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if !strings.Contains(line, "mobile") {
			continue
		}

		contents := strings.Split(line, "•")
		id := strings.TrimSpace(contents[1])

		fmt.Println("実機デバイスを検出しました:", id)
		fmt.Println("インストールしています…")

		out, err := exec.Command("flutter", "install", "-d", id).CombinedOutput()
		outMsg := string(out)
		if err != nil {
			errorMsgIndex := strings.Index(outMsg, "Error:")
			fmt.Println("エラーが発生しました (´･ω･`)")
			fmt.Println(outMsg[errorMsgIndex:])
			os.Exit(1)
		}
		fmt.Println(outMsg)

		return
	}

	fmt.Println("実機デバイスは見つかりませんでした (´･ω･`)")
}

func main() {
	detector()
}
