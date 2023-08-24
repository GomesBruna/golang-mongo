package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"golang-mongo/models"
	"github.com/qiniu/qmgo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct{
	Session *qmgo.Client
	Context context.Context
}

func NewUserController(s *qmgo.Client, c context.Context ) *UserController{
	return &UserController{s, c}
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	u := models.User{}

	if err := uc.Session.Database("mongo-golang").Collection("users").Find(uc.Context,bson.M{"name":id}).One(&u); err != nil{
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf( w, "%s\n", uj)
}

func (uc UserController) DeleteUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	
	if err := uc.Session.Database("mongo-golang").Collection("users").Remove(uc.Context,bson.M{"name":id}); err != nil{
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint( w, "Deleted user ", id, "\n")
}

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	_, err:= uc.Session.Database("mongo-golang").Collection("users").InsertOne(uc.Context,u)
	if err != nil {
		fmt.Println(err)
	}

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf( w, "%s\n", uj)
}

