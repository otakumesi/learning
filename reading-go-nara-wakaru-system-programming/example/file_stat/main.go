package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]\n", os.Args[0])
		os.Exit(1)
	}

	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("  FileName: %v\n", info.Name())
	fmt.Printf("  FileSize: %v\n", info.Size())
	fmt.Printf("  ModTime: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  IsDir: %v\n", info.Mode().IsDir())
	fmt.Printf("  IsRegular: %v\n", info.Mode().IsRegular())
	fmt.Printf("  Permission: %o\n", info.Mode().Perm())
	fmt.Printf("  ModeText: %v\n", info.Mode().String())
}
