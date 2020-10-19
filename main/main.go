package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	filename := string(fdTime) + ".txt"
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()

	c := cron.New()
	c.AddFunc("*/1 * * * * *", func() {
		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		timeString := string(timeStamp)
		v, _ := mem.VirtualMemory()
		c, _ := cpu.Percent(time.Second, false)

		memUsedPercentFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.UsedPercent), 64)
		cpuUsageFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", c[0]), 64)
		memUsedPercent := strconv.FormatFloat(memUsedPercentFloat, 'f', -1, 64)
		cpuUsage := strconv.FormatFloat(cpuUsageFloat, 'f', -1, 64)

		fdContent := strings.Join([]string{timeString, ",", "Mem UsedPercent:", memUsedPercent, ",", "CPU Usage:", cpuUsage, "\n"}, "")
		buf := []byte(fdContent)
		fd.Write(buf)
	})
	c.Start()
	defer c.Stop()
	select {}
}
