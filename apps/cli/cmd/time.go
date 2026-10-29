package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/luobochuanqi/superset/apps/cli/internal/timer"
	"github.com/spf13/cobra"
)

var (
	calculateTime string
	duration      string
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间相关",
	Long:  "时间格式处理",
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("当前时间：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculatedTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算时间",
	Long:  "在基准时间（默认当前时间）上加上时间间隔，得到结果时间",
	Run: func(cmd *cobra.Command, args []string) {
		layout := "2006-01-02 15:04:05"
		currentTime := timer.GetNowTime()
		if calculateTime != "" {
			switch strings.Count(calculateTime, " ") {
			case 0:
				layout = "2006-01-02"
			case 1:
				layout = "2006-01-02 15:04"
			case 2:
				layout = "2006-01-02 15:04:05"
			default:
				log.Fatalf("时间格式不正确")
			}
			var err error
			currentTime, err = time.Parse(layout, calculateTime)
			if err != nil {
				log.Fatalf("时间格式不正确：%v", err)
			}
		}
		result, err := timer.GetCalculatedTime(currentTime, duration)
		if err != nil {
			log.Fatalf("时间间隔格式不正确：%v", err)
		}
		log.Printf("计算结果：%s, %d", result.Format(layout), result.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculatedTimeCmd)
	calculatedTimeCmd.Flags().StringVarP(&calculateTime, "time", "t", "", "需要计算的时间，格式为：2006-01-02 15:04:05")
	calculatedTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "需要计算的时间间隔，格式为：1h30m")
}
