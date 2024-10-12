package main

import (
	"flag"
	"fmt"
	"github.com/dextercai/MomentumTranslateUtils/pkg/bdfConv"
	"github.com/samber/lo"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	flagBdfFontPath := flag.String("bdf-font", "default_font.bdf", "file path of bdf font")
	flagUnicodeListFilePath := flag.String("unicode-list-file", "default_unicode.txt", "file contained unicode list")
	flagExportFontName := flag.String("export-font-name", "dextercai_momentum_utils_custom_font", "font name in export c file")
	flagExportCFile := flag.String("export-c-file", "export.c", "export c file")

	flag.Parse()

	unicodeListFile, err := os.Open(*flagUnicodeListFilePath)
	if err != nil {
		log.Fatalf("error occoured when open unicode list err: %s", err)
		return
	}
	allUnicode, err := io.ReadAll(unicodeListFile)
	if err != nil {
		log.Fatalf("error occoured when read unicode list err: %s", err)
		return
	}
	mapStrList := []string{"32-128"}

	for _, char := range allUnicode {
		if unicode.Is(unicode.Han, rune(char)) || char > 127 {
			mapStrList = append(mapStrList, fmt.Sprintf("$%X", char))
		}
	}

	mapStrList = lo.Union(mapStrList)

	ret := bdfConv.CreateTargetCFont(*flagBdfFontPath, strings.Join(mapStrList, ","), *flagExportFontName, *flagExportCFile)

	if ret != 0 {
		log.Printf("some error occoured")
		return
	}

	log.Printf("conv done")
}
