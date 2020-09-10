package log

import (

)

func main() {
    // 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色
    gin.DisableConsoleColor()
​
    // 写入日志文件
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)
​
    // 如果你需要同时写入日志文件和控制台上显示，使用下面代码
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
​
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
​
    router.Run(":8080")
}

// 自定义日志格式
// func main() {
    router := gin.New()
​
    // LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
    // 默认 gin.DefaultWriter = os.Stdout
    router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
​
        // 你的自定义格式
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
                param.ClientIP,
                param.TimeStamp.Format(time.RFC1123),
                param.Method,
                param.Path,
                param.Request.Proto,
                param.StatusCode,
                param.Latency,
                param.Request.UserAgent(),
                param.ErrorMessage,
        )
    }))
    router.Use(gin.Recovery())
​
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
​
    router.Run(":8080")
}