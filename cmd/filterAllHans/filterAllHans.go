package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

func main() {
	// 要扫描的目录路径
	flagDirPath := flag.String("src-path", "./src/", "source code path")
	flagOutoutFile := flag.String("output-file", "./allHansUnicode.txt", "all hans unicode output file")
	flag.Parse()

	outputFile, err := os.OpenFile(*flagOutoutFile, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		log.Fatalf("打开输出文件 %s 失败: %v\n", *flagOutoutFile, err)
		return
	}
	defer outputFile.Close()

	uniqueChars := make(map[rune]bool)
	err = filepath.WalkDir(*flagDirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("读取文件 %s 失败: %v\n", path, err)
				return nil
			}
			log.Printf("读取文件 %s", path)
			for _, char := range string(content) {
				if unicode.Is(unicode.Han, char) || char > 127 {
					uniqueChars[char] = true
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("扫描目录时发生错误: %v\n", err)
		return
	}
	for char := range uniqueChars {
		outputFile.WriteString(fmt.Sprintf("%c", char))
		fmt.Printf("%c", char)
	}

	log.Print("完成")
}
