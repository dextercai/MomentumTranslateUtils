# MomentumTranslateUtils

适用于 Momentum项目 汉化时的小工具

## 包含

### BDF字库裁剪与u8g2代码生成工具

```bash
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