package routes

import (
	"news-app-be23/internal/features/users"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc users.Handler) {
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	})

	e.POST("/register", uc.SignUp())
	e.POST("/login", uc.Login())
	// todoRoute(e, tc)
}

// func todoRoute(e *echo.Echo, tc todos.Handler) {
// 	jwtKey := os.Getenv("JWT_SECRET")
// 	if jwtKey == "" {
// 		fmt.Println("JWT secret key not found in environment variables")
// 	}

// 	t := e.Group("/todos")
// 	t.Use(echojwt.WithConfig(
// 		echojwt.Config{
// 			SigningKey:    []byte(jwtKey),
// 			SigningMethod: jwt.SigningMethodHS256.Name,
// 		},
// 	))
// 	t.POST("", tc.AddTodo())
// 	t.PUT("/:id", tc.UpdateTodo())
// 	t.GET("", tc.FindTodo())
// 	t.DELETE("/:id", tc.DeleteTodo())
// }
