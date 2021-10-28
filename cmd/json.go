package cmd

import (
	"fmt"
	"github.com/GoProject/GoToolSet/internal/json"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	JSONPretty = iota + 1
)

var jsonDesc = strings.Join([]string{
	"该命令支持JSON相关转换，有如下模式：",
	"1：JSON美化",
}, "\n")

var jsonPrettyCmd = &cobra.Command{
	Use:   "pretty",
	Short: "格式化JSON",
	Long:  "格式化JSON",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonStr=" + jsonStr)
		log.Printf("输出结果：\n%s", json.JSONPretty(jsonStr, false))
	},
}

var jsonColorPrettyCmd = &cobra.Command{
	Use:   "cpretty",
	Short: "格式化JSON（彩色）",
	Long:  "格式化JSON（彩色）",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("输出结果：\n%s", json.JSONPretty(jsonStr, true))
	},
}

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "JSON转换",
	Long:  "JSON转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var jsonStr string

func init() {
	jsonCmd.AddCommand(jsonPrettyCmd)
	jsonCmd.AddCommand(jsonColorPrettyCmd)
	jsonPrettyCmd.Flags().StringVarP(&jsonStr, "jsonStr", "j", "", "请输入需要转换的JSON字符串")
	jsonColorPrettyCmd.Flags().StringVarP(&jsonStr, "jsonStr", "j", "", "请输入需要转换的JSON字符串")
}
