package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

func main() {
	format := render.New()
	router := chi.NewRouter()

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(router, "", "/public", http.Dir(filesDir))

	handlersForPractice(router, format)

	router.Route("/test", func(methodRouter chi.Router) {
		handlersForMethods(methodRouter, format)
	})

	handlersForURLShortening(router, format)

	handlersForURLTime(router, format)

	router.Route("/user", func(methodRouter chi.Router) {
		handlersForURLFile(methodRouter, format)
	})

	router.Route("/expire", func(methodRouter chi.Router) {
		handlersForURLExpire(methodRouter, format)
	})

	handlersForImageSharing(router, format)

	router.Route("/struct", func(structRouter chi.Router) {
		handlersForStruct(structRouter, format)
	})

	router.Route("/array", func(arrayRouter chi.Router) {
		handlersForArray(arrayRouter, format)
	})

	router.Route("/map", func(mapRouter chi.Router) {
		handlersForMap(mapRouter, format)
	})

	fmt.Println("Server is running!")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}

//общая структура, в которой хранятся данные пользователя
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
}

//структура, в которой хранятся данные пользователя с дополнительным полем Friend
type UserWithFriends struct {
	ID      int
	Name    string
	Surname string
	Age     int
	Sex     string
	Friend  []int //поле является списком идентbфикаторов других пользователей
}

//метод String(), форматирующий структуру User в тип string
func (arg User) String() string {
	result := fmt.Sprintf("ID: %d\nName: %s\nSurname: %s\nAge: %d\nSex: %s\n", arg.ID, arg.Name, arg.Surname, arg.Age, arg.Sex)
	return result
}

//метод ToString(), форматирующий структуру UserWithFriends в тип string
func (arg UserWithFriends) ToString() string {
	result := fmt.Sprintf("ID: %d\nName: %s\nSurname: %s\nAge: %d\nSex: %s\nFriends: %d\n", arg.ID, arg.Name, arg.Surname, arg.Age, arg.Sex, arg.Friend)
	return result
}
