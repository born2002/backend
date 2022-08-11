package model
import  "gorm.io/gorm"
type Employee struct {
 gorm.Model
 Employeeid  string
 Employeename  string
 Employeeusername    string
 Employeepassword  string
}

