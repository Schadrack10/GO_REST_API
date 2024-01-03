# Restful application programing interface with Golang

# Introduction
>> restful api allows us to serve data to the client by ways of modifying existing or creating new data
>> communication between the client and the server , whereby the client sends a request and the server sends a response , the      response which contains data that the client will render

# Set-up dependency tracking: go mod

Creating a new module (or mod) in Go is a common practice to manage dependencies and versioning within a Go project. The Go programming language introduced a module system to address some of the challenges associated with dependency management in previous versions.

in the terminal we add the below command:
<b>replace REST_APi with my directory </b>

    go mod init example/REST_API

# gin
>> we will use the library gin to create restfull api 
>> this is a github package

in the terminal we add

    go get github.com/gin-gonic/gin