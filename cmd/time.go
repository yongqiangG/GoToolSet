package cmd

import (
	"github.com/GoProject/GoToolSet/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "当前时间",
	Long:  "当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := time.LoadLocation("Asia/Shanghai")
		now := time.Now().In(location)

		log.Printf("输出结果：%s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "时间加减计算",
	Long:  "时间加减计算",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calculateTime == "" {
			currentTimer = time.Now().In(location)
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				layout = "2006-01-02 15:04:05"
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0).In(location)
			}
		}
		t, err := timer.GetCalculateTine(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTine err: %v", err)
		}
		log.Printf("输出结果：%s，%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，可以是时间戳或者格式化的时间（2006-01-02 15:04:05）`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
