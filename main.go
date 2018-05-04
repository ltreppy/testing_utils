package main

import (
  "net/http"
  "flag"
  "html/template"
  "go.uber.org/zap"
)

var logger zap.Logger

func main(){
  var port int
  logger, _ := zap.NewDevelopment()
  logger.Sugar()
  defer logger.Sync()
  flag.IntVar(&port, "p", 8080, "Port number for proxy utility")
  flag.Parse()
  http.HandleFunc("/", adminHandler)
  http.HandleFunc("/api/", apiHandler)
  if err := http.ListenAndServe("localhost:8080", nil); err != http.ErrServerClosed {
      logger.Fatal("Error in ListenAndServer %v", zap.Error(err))
  }
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("static/main.html")
  if err != nil {
    logger.Fatal("Error in template parsing %v", zap.Error(err))
  }
  err = t.Execute(w, nil)
  if err != nil {
    logger.Fatal("Error in template execution", zap.Error(err))
  }
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

}
