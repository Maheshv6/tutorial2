package main 

import (
       	"fmt"
        "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type  Foodies  struct { 
       Breakfast  string  
       Lunch    string
       Dinner   string
}

       type Hotel struct {
         gorm.Model
	 Name string
         Items Itemsy   `json:"items" gorm:"foreignkey:Refer"`
	                      
         }

         type Itemsy struct {
          gorm.Model
          Lunch     string
          Dinner   string
         }


type   Review struct {
      Rating   string      `json:"laugh"`
}


type  Breakfast   interface {

      Fooodies(c *gin.Context)

}

type Items struct {
         
     db *gorm.DB

}

       

 func Customer()  Breakfast  {

	db, err := gorm.Open("mysql", "root:Mahe@786@tcp(127.0.0.1:3306)/mahedb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&Hotel{} ,&Itemsy{} )

	return &Items{ db }

}

 func Inputvalidation (  req   *gin.Context )  *Foodies{

       var  inputs Foodies

        req.ShouldBind(&inputs)

  
     return  &inputs

 }

func    APIvalidation (inputs * Foodies , req   *Items   ) bool   {



	    fmt.Println( inputs)

           req.db.Create( &Hotel{Name: "KFC", Items: Itemsy{Lunch: "egg" , Dinner:"chicken"}})

       return true

}

func  (req   *Items )  Fooodies(c *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:=  Inputvalidation(c   )

	if inputs == nil {
	     c.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

 
        resp := APIvalidation(  inputs ,req  )

	if resp == false {
	   c.JSON(201,"invalid email or password ")
              return 
	}
  
     var Response  Review = Review{ "ssuuppeerr" }

      c.AsciiJSON(200 ,Response)
}



func main(){

	
	 
           r := gin.Default()

	   obj:= Customer()

         v1 := r.Group("/bull")
	 {
	  v1.GET("/login",obj.Fooodies )
	  }
	
	 r.Run(":9090")


}