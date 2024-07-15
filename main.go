package main

import (
	"github.com/mahdi-cpp/go-english/repository"
)

func main() {
	repository.ConnectDatabase()
	repository.CreatNewUser()
	Run()
}
