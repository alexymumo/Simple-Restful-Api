package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type students struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Department string `json:"department"`
    Faculty string `json:"faculty"`
}

var student = []students{
	{
        ID : "C025-01-2270/2014",
        Name: "Alex Mumo", 
        Department: "Information Technology" ,
        Faculty: "Technology"},

	{
        ID : "C025-01-2260/2014",
        Name: "John Patel",
        Department: "Computer Science",
        Faculty: "Technology"},

	{
        ID : "C025-01-7889/2020",
        Name: "Philip Luck", 
        Department: "Food Science",
        Faculty: "Food Science"},

    {
        ID : "C025-01-7839/4589",
        Name: "John Hill",
        Department: "Electrical Engineering",
        Faculty: "Engineering"},
}


//hangler to return all the students 
func getStudents(c *gin.Context){
    c.IndentedJSON(http.StatusOK, student)
}

//handler to add new student

func addStudents(c *gin.Context){
    var newStudent students
    if err := c.BindJSON(&newStudent); err != nil{
        return
    }
    student = append(student, newStudent)
    c.IndentedJSON(http.StatusCreated, newStudent)
}

// main function
func main(){
    router := gin.Default()

    router.GET("/student", getStudents)
    router.POST("/student", addStudents)    
    router.Run("localhost:8080")

}
