
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

/*
In this code a struct named Welcome is defined and inside this struct there is a variable named Name  This variable will be used to hold the user's name.
In the main function a function is defined with http.HandleFunc that will run when a request comes to the "/" path. 
In this function an instance is created from the Welcome struct and the Name variable is assigned the value "Guest" by default. 
Then, using the r.FormValue function, the "name" value from the form is retrieved and if this value is not empty, this value is assigned to the variable Name.


The index.html file is as follows:

<!DOCTYPE html>
<html>
<head>
	<title>Welcome</title>
</head>
<body>
	<h1>Hello, {{.Name}}!</h1>
	<form action="/" method="POST">
		<label for="name">Enter your name:</label>
		<input type="text" name="name" id="name" />
		<button type="submit">Submit</button>
	</form>
</body>
</html>


In the HTML code The {{.Name}} placeholder will replace the value of the Name variable in the Welcome struct also a form was created and this form


*/
