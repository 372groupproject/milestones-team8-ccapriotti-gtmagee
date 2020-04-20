package main

import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
   Name string
   Time string
}

func main() {
   welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
   
   templates := template.Must(template.ParseFiles("../webpageUtility/entryPage.html"))

  
   
   http.Handle("/static/", 
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static"))))

   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
      
      if err := templates.ExecuteTemplate(w, "entryPage.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   fmt.Println(http.ListenAndServe(":3000", nil));
}