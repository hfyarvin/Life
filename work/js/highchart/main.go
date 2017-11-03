package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Date struct {
	Temp []float64
	Hum  []float64
	Time []string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	checkError(err)
	d := new(Date)
	d.Temp = []float64{7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6}
	d.Hum = []float64{70, 69, 45, 50, 50, 60, 50, 60, 42, 43, 42, 43}
	err = t.Execute(w, d)
	checkError(err)
}

func main() {
	http.Handle("/highcharts/", http.StripPrefix("/highcharts/", http.FileServer(http.Dir("highcharts"))))
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":9333", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
