package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	// 创建CSRF保护中间件
	csrfMiddleware := csrf.Protect(
		[]byte("your-secret-key"), // 替换为你的密钥
		csrf.Secure(false),        // 如果你的应用程序不使用HTTPS，将Secure设置为false
	)

	// 使用CSRF保护中间件
	r.Use(func(c *gin.Context) {
		csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 让 Gorilla CSRF 中间件处理请求
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
			w.Header().Set("Vary", "Cookie") // 通常需要设置 Vary 标头

			// 继续处理 Gin 请求
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	})

	// 创建GET路由，用于生成包含CSRF令牌的表单
	r.GET("/form", func(c *gin.Context) {
		tmpl, err := template.ParseFiles("./form.temp")
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		// 利用给定数据渲染模板，并将结果写入w
		err = tmpl.Execute(c.Writer, gin.H{"csrfToken": c.Writer.Header().Get("X-CSRF-Token")})
		if err != nil {
			fmt.Println(err)
			return
		}

	})

	// 创建POST路由，用于处理提交的表单数据
	r.POST("/submit", func(c *gin.Context) {
		// 验证CSRF令牌
		// csrf.Token(r) 和 csrf.HeaderField(r) 在内部验证令牌
		// 如果验证失败，中间件会自动返回错误响应
		// 你无需额外调用 VerifyToken
		// 在这里处理表单提交
		// ...

		// 提取表单数据示例：
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println(username, password)
		// 处理表单提交的逻辑
		// 例如，这里可以验证用户名和密码等
		// ...

		c.String(http.StatusOK, "Form submitted successfully")
	})

	r.Run(":8080")
}
