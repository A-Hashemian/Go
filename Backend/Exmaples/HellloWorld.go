
package main 

import(
  "html/template"
  "net/http"

)

type Welcome struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		welcome := Welcome{Name: "Guest"}
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, welcome)
	})

	http.ListenAndServe(":8080", nil)
}
