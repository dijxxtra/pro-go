package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Данные заполняемые пользователем
type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// Данные передаются с использованием POST запроса
type formData struct {
	*Rsvp
	Errors []string
}

// Все данные пользователей
var responses = make([]*Rsvp, 0, 10)

// Все используемые шаблоны
var templates = make(map[string]*template.Template, 3)

// Создание Приветственной страницы
func welcomeHundler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

// Создание страницы со списком участников
func listHundler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

// Создание страницы с заполняемой формой
func formHandler(writer http.ResponseWriter, request *http.Request) {
	// Если запрос на получение страницы (GET)
	if request.Method == http.MethodGet {
		// Создай страницу с пустыми данными
		templates["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
		// Если POST запрос
	} else if request.Method == http.MethodPost {
		// Взять данные с формы (в виде карты, где ключ это значение name (в форме), а значение - список со строковыми данными)
		request.ParseForm()
		responseData := Rsvp{
			Name:  request.Form["name"][0],
			Email: request.Form["email"][0],
			Phone: request.Form["phone"][0],
			// т.к передается строка, а значение WillAttend - bool
			WillAttend: request.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{&responseData, errors})
		} else {
			// Записываем запрос в общий список запросов
			responses = append(responses, &responseData)
			if responseData.WillAttend {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}
	}
}

func loadTemplate() {
	// Имена файлов
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		// t - *template.Template объединяющий файлы layout с другими шаблонами
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			// welcome: *template.Template > layout.html + welcome.html
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	// Загрузка шаблонов
	loadTemplate()

	// Создание обработчика
	http.HandleFunc("/", welcomeHundler)
	http.HandleFunc("/list", listHundler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
