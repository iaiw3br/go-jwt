package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-jwt/pkg/models/pg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := getDotEnvVariable("PORT")
	pgName := getDotEnvVariable("PG_NAME")
	pgPassword := getDotEnvVariable("PG_PASSWORD")
	pgDatabase := getDotEnvVariable("PG_DATABASE")
	pgMode := getDotEnvVariable("PG_MODE")
	connectionToDB := "postgres://" + pgName + ":" + pgPassword + "@localhost/" + pgDatabase + "?" + pgMode
	db, err := openDB(connectionToDB)

	if err != nil {
		log.Fatal(err)
		return
	}
	app := application{User: &pg.UserModel{DB: db}}
	srv := http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	defer db.Close()
	fmt.Println("Сервер запущен")

	err = srv.ListenAndServe()
	log.Fatal(err)
}

type application struct {
	User *pg.UserModel
}

func getDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

func openDB(connectionToDB string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionToDB)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(getDotEnvVariable("SECRET_JWT")))

	if err != nil {
		return "", err
	}

	return token, nil
}
