package main

import (
	"ai/hd2-killer/internal/os"
	"log"
	"runtime"
)

func main() {
	var processName string

	if runtime.GOOS == "windows" {
		processName = "helldivers2.exe"
	} else {
		// TODO idk, unix not tested
		processName = "helldivers2.exe"
	}

	if err := os.KillProcessByName(processName); err != nil {
		log.Fatalf("kill process by name: %s", err)
	}
}
