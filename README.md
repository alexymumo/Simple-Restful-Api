# Simple-Restful-Api
A simple restful api written using go and gin

First run the the command go mod init command, giving it the path of the module your code will be in.
This command creates a go.mod file in which dependencies you add will be listed for tracking.
At the command line, use go get to add the github.com/gin-gonic/gin module as a dependency for your module. Use a dot argument to mean “get dependencies for code in the current directory.”![go_get]
(https://user-images.githubusercontent.com/56880898/133260512-a1d39c95-b1f8-4932-a09d-c08a5bbce124.png)

From the command line in the directory containing main.go, run the code. Use a dot argument to mean “run code in the current directory.”![go_get](https://user-images.githubusercontent.com/56880898/133260806-e772ad41-af66-4782-86f3-bfde8b247e79.png)
Once the code is running, you have a running HTTP server to which you can send requests.
From a new command line window, use curl tomake a request to your running web service.
curl http://localhost:8080/student
The command should display the data you seeded the service with.!
[outout](https://user-images.githubusercontent.com/56880898/133261649-a8d6d747-9fad-4ed6-a001-fcb0d3b76371.png)




