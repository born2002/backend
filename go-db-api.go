package main

import (
	"example/go-db-api/model"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/nextjs_employee?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	r.GET("/employees", func(c *gin.Context) {
		var employees []model.Employee
		db.Find(&employees)
		c.JSON(200, employees)
	})

	r.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		var employees model.Employee
		db.First(&employees, id)
		c.JSON(200, employees)
	})

	r.POST("/employees", func(c *gin.Context) {
		var employees model.Employee
		if err := c.ShouldBindJSON(&employees); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Bind(&employees)
		result := db.Create(&employees)
		c.JSON(200, gin.H{"RowsAffected": result.RowsAffected})
	})

	r.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		var employee model.Employee
		db.First(&employee, id)
		db.Delete(&employee)
		c.JSON(200, employee)
	})

	//Put Data
	//  r.PUT("/employees", func(c *gin.Context) {
	//   var employee model.Employee
	//   var updatedEmployee model.Employee
	//   if err := c.ShouldBindJSON(&employee); err != nil {
	//    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//    return
	//   }
	//   db.First(&updatedEmployee, employee.ID)
	//   updatedEmployee.Employeename = employee.Employeename
	//   updatedEmployee.Employeeusername = employee.Employeeusername
	//   updatedEmployee.Employeepassword = employee.Employeepassword
	//   db.Save(updatedEmployee)
	//   c.JSON(200, updatedEmployee)
	//  })

	//Put Data
	r.PUT("/employees/:id", func(c *gin.Context) {

		c.Header("Content-Type", "application/json")

		id := c.Param("id")

		var employee model.Employee
		var updatedemployee model.Employee

		c.BindJSON(&employee)
		c.BindJSON(&updatedemployee)

		db.First(&updatedemployee, id)

		updatedemployee.Employeeid = employee.Employeeid
		updatedemployee.Employeename = employee.Employeename
		updatedemployee.Employeeusername = employee.Employeeusername
		updatedemployee.Employeepassword = employee.Employeepassword

		db.Save(updatedemployee)
		c.JSON(200, updatedemployee)

		//c.JSON(200, gin.H{"id : " + id: "is puted"});
	})

	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
