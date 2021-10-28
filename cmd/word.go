package cmd

import (
	"github.com/GoProject/GoToolSet/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper                      = iota + 1 // 转大写
	ModeLower                                 // 转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var wordDesc = strings.Join([]string{
	"该命令支持各种单词格式转换，有如下模式：",
	"1：转大写",
	"2：转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  wordDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var result string
		switch mode {
		case ModeUpper:
			result = word.ToUpper(str)
		case ModeLower:
			result = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			result = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			result = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			result = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该模式，执行 help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", result)
	},
}

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入需要转换的单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入转换模式")
}
