package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/post/min1", testKline)
	r.POST("/post/all_min1", testNewKline)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:6382") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func testKline(c *gin.Context) {
	b, _ := c.GetRawData()
	inString := string(b)

	// 解析参数
	params := strings.Split(inString, "&")
	codes := params[1][5:]
	// try to split codes with ','
	codeArray := strings.Split(codes, ",")

	var result string
	for _, code := range codeArray {
		if code == "" {
			continue
		}
		now := time.Now()
		result += code + ","
		hourMinute := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
		result += hourMinute + ","
		seed := float64(rand.Intn(22) + 3)
		high := fmt.Sprintf("%.2f", seed*0.34)
		open := fmt.Sprintf("%.2f", seed*0.26)
		low := fmt.Sprintf("%.2f", seed*0.39)
		close := fmt.Sprintf("%.2f", seed*0.24)
		volume := fmt.Sprintf("%.0f", seed*12000)
		turover := fmt.Sprintf("%.0f", seed*9000)
		result += high + "," + open + "," + low + "," + close + "," + volume + "," + turover + "#"
	}
	// When codeCount equal 1 must return code fileld
	c.String(http.StatusOK, fmt.Sprint(result))
}

func testNewKline(c *gin.Context) {
	b, _ := c.GetRawData()
	inString := string(b)

	// 解析参数
	params := strings.Split(inString, "&")
	codes := params[1][5:]
	num, _ := strconv.Atoi(params[2][4:])
	// try to split codes with ','
	codeArray := strings.Split(codes, ",")
	codeCount := len(codeArray)

	var result string
	for _, code := range codeArray {
		if code == "" {
			continue
		}
		now := time.Now().Add(time.Minute)
		for i := 0; i < num; i++ {
			now = now.Add(-time.Minute)
		}
		for i := 0; i < num; i++ {
			if codeCount > 1 {
				result += code + ","
			}
			hourMinute := fmt.Sprintf("%02d:%02d", now.Hour(), now.Minute())
			result += hourMinute + ","
			now = now.Add(time.Minute)
			seed := float64(rand.Intn(22) + 3)
			high := fmt.Sprintf("%.2f", seed*0.34)
			open := fmt.Sprintf("%.2f", seed*0.26)
			low := fmt.Sprintf("%.2f", seed*0.39)
			close := fmt.Sprintf("%.2f", seed*0.24)
			volume := fmt.Sprintf("%.0f", seed*12000)
			turover := fmt.Sprintf("%.0f", seed*9000)
			result += high + "," + open + "," + low + "," + close + "," + volume + "," + turover + "\r\n"
		}
	}
	result = "日期\r\nRadomDate\r\n时间,开盘价...\r\n" + result
	// When codeCount equal 1 must return code fileld
	c.String(http.StatusOK, fmt.Sprint(result))
}
