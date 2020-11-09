package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func landingpage(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() 
    fmt.Println(r.Form) 
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "go to /aritcles for posting or viewing the aritcles") // write data to response
}

func article(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of posting article
        fmt.Println("ID:", r.Form["id"])
        fmt.Println("Title:", r.Form["title"])
        fmt.Println("SubTitle:", r.Form["stitle"])
        fmt.Println("Content:", r.Form["content"])
        fmt.Println("TimeStamp:", r.Form["timestamp"])
    }
}

func main() {
    http.HandleFunc("/", landingpage) // setting router rule
    http.HandleFunc("/articles", article)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

