package main

import (
	"html/template"
	"log"
	"net/http"
)

type Skill struct {
	Name string
	Icon string
}

type PageData struct {
	Title   string
	Stacks  []Skill
	Contact map[string]string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	funcMap := template.FuncMap{
		"mod": func(i, j int) int { return i % j },
	}
	tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Sir Ivan's Portfolio",
			Stacks: []Skill{
				{Name: "Python", Icon: "python.png"},
				{Name: "Java", Icon: "java.png"},
				{Name: "Next", Icon: "next.png"},
				{Name: "React", Icon: "react.png"},
				{Name: "Spring Boot", Icon: "spring-boot.png"},
				{Name: "Jira", Icon: "jira.png"},
				{Name: "Odoo", Icon: "odoo.png"},
				{Name: "Figma", Icon: "figma.png"},
				{Name: "Django", Icon: "django.png"},
				{Name: "Postman", Icon: "postman.png"},
				{Name: "PostgreSQL", Icon: "postgresql.png"},
				{Name: "Docker", Icon: "docker.png"},
				{Name: "Flutter", Icon: "flutter.png"},
				{Name: "Dart", Icon: "dart.png"},
				{Name: "Godot", Icon: "godot.png"},
				{Name: "GO", Icon: "go.png"},
				{Name: "Tableau", Icon: "tableau.png"},
			},
			Contact: map[string]string{
				"Email":    "mailto:ivanrbnc@gmail.com",
				"GitHub":   "https://github.com/ivanrbnc",
				"LinkedIn": "https://linkedin.com/in/ivan-cezeliano",
			},
		}

		tmpl.Execute(w, data)
	})

	log.Println("⚔️  Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
