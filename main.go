package main 

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log" 
    "events/models"



)


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func getEvents (c *gin.Context){

    events, err := models.GetEvents(40)

    checkErr(err)

    if events == nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": "no records found"})

        return 
} else {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{"data": events})
}
}
 
    
// func addEvent(c *gin.Context){
//
//     var json.models.Event 
//
//     if err:= c.ShouldBindJSON(&json); err != nil{
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
//     }
//
// 	success, err := models.AddPerson(json)
//
// 	if success {
// 		c.JSON(http.StatusOK, gin.H{"message": "Success"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }
//


func main(){
err := models.ConnectDatabase()
checkErr(err)

    r := gin.Default()
r.LoadHTMLGlob("templates/*")
    r.GET("/events", getEvents)
   // r.POST("events", addEvents) 
r.GET("/login", func(c *gin.Context) {
    c.HTML(http.StatusOK, "test.tmpl", gin.H{})
})

r.Run()



}



