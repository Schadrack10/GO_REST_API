# Restful application programing interface with Golang

# Introduction
>> restful api allows us to serve data to the client by ways of modifying existing or creating new data
>> communication between the client and the server , whereby the client sends a request and the server sends a response , the      response which contains data that the client will render

# Set-up dependency tracking: go mod

Creating a new module (or mod) in Go is a common practice to manage dependencies and versioning within a Go project. The Go programming language introduced a module system to address some of the challenges associated with dependency management in previous versions.

in the terminal we add the below command:
<b>replace REST_APi with my directory </b>

    go mod init example/REST_API

# Set-up upstream branch [git]
Pushing to Git command when getting error : no upstream branch

    git push --set-upstream https://github.com/Schadrack10/GO_REST_API.git master


# gin
>> we will use the library gin to create restfull api 
>> this is a github package

in the terminal we add

    go get github.com/gin-gonic/gin

# Begin to create or define the  structure for our data , we use a struct
   
    type todo struct {
        ID string 
        item string
        completed bool
    }

    var todos = []todo{
        {ID:"1",item:"",completed:""},
        {ID:"2",item:"",completed:""},
    }

>> now we need to convert the structure to json, we do this by adding json literal next to the struct
>> now we can communicate with client side with the json
   
       type todo struct {
        ID string `json:"id"`
        item string `json:"title"`
        completed bool `json:"completed"`
    }

# adding imports , Creating the api port and initialising gin

    import (
	"net/htt" 
	"github.com/gin-gonic/gin"
    )


>>  creating the api port and rendering the todo data 
   
    func getTodos(context *gin.Context) {
        context.IndentedJSON(http.StatusOK, todos)
    }


    func main() {

        router := gin.Default()  

        router.GET("/todos", getTodos)

        router.Run("localhost:9090")
    }
