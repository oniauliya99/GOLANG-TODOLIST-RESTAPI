package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"codebrains.io/todolist/models"
	"codebrains.io/todolist/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func helloWorld(c *fiber.Ctx)error  {
	return c.SendString("Hello World")
}

func initDatabase()  {
	var err error
	dsn := "root:tanyaoni@tcp(127.0.0.1:3306)/db_todolist?charset=utf8mb4&parseTime=True&loc=Local"
  	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic("Failed to connect to database")
	}
	fmt.Println("Database connect")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrate DB")
}

func setupRoutes(app *fiber.App){
	app.Get("/todos",models.GetTodos)
	app.Get("/todos/:id",models.GetTodoById)
	app.Post("/todos",models.CreateTodo)
	app.Put("/todos/:id",models.UpdateTodo)
	app.Delete("/todos/:id",models.DeleteTodo)
}  

func main()  {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	app.Get("/",helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}