package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/thegeorgenikhil/go-config-example/config"
	"github.com/thegeorgenikhil/go-config-example/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) UserHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	userCollection := m.App.Client.Database(m.App.DatabaseName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	switch method {
	case "GET":
		var users []models.User
		cur, err := userCollection.Find(ctx, bson.D{})
		if err != nil {
			log.Println(err)
		}

		if err = cur.All(ctx, &users); err != nil {
			log.Println(err)
		}

		if users == nil {
			w.Write([]byte("No users found!"))
			return
		}

		responseJSON, _ := json.Marshal(users)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJSON))
	case "POST":
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)

		user.ID = primitive.NewObjectID()

		if err != nil {
			log.Println(err)
		}

		_, err = userCollection.InsertOne(ctx, user)

		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully!"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Invalid Method!"))
	}
}
