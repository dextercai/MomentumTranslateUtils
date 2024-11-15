# MomentumTranslateUtils

u8g2开发时Unicode字库的适配。也适用于 Momentum 项目汉化时的工具。

## 包含

### BDF字库裁剪与u8g2代码生成工具

```bash
CGO_ENABLED=1
go install github.com/dextercai/MomentumTranslateUtils/cmd/genU8g2FontCode@main
genU8g2FontCode --help
```

```text
Usage of genU8g2FontCode:
  -bdf-font string
        file path of bdf font (default "default_font.bdf")
  -export-c-file string
        export c file (default "export.c")
  -export-font-name string
        font name in export c file (default "dextercai_momentum_utils_custom_font")
  -unicode-list-file string
        file contained unicode list (default "default_unicode.txt")
```

### 代码扫描及字符提取工具
```text
CGO_ENABLED=1
go install github.com/dextercai/MomentumTranslateUtils/cmd/filterAllHans@main
filterAllHans --help
```


```text
Usage of filterAllHans:
  -output-file string
        all hans unicode output file (default "./allHansUnicode.txt")
  -src-path string
        source code path (default "./src/")
```

## Reference
- https://github.com/olikraus/u8g2
