package main

import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-mongo/controllers"
	"fmt"
	"github.com/qiniu/qmgo"
	"context"
)

func main(){

	r := httprouter.New()
	client, context := getSession()
	defer client.Close(context)
	uc := controllers.NewUserController(client,context)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("Service running in port 9000")
	http.ListenAndServe("localhost:9000", r)
		
}

func getSession() (*qmgo.Client, context.Context){

	//arrumar o fato de estar usando o admin para logar, podemos criar um usuario para nao utilizar o admnin
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://root:example@mongo/"})
	if err != nil{
		panic(err)
	}
	return client, ctx
}