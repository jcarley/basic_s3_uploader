package handlers

import (
  // "encoding/json"
  "github.com/gorilla/mux"
  // "io"
  // "io/ioutil"
  "net/http"
)

func App(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "web/public/views/index.html")
}

func getInitSignature(w http.ResponseWriter, r *http.Request) {
  // content_type :json
  // hash = S3UploadRequest.new(:type => :init, :params => params).to_h
  // hash.to_json
}

func getRemainingSignatures(w http.ResponseWriter, r *http.Request) {
}

func Install(r *mux.Router) {
  r.HandleFunc("/", App)
  // r.HandleFunc("/ws", WsHandler)

  r.HandleFunc("/get_init_signature", getInitSignature)
  r.HandleFunc("/get_remaining_signatures", getRemainingSignatures)
  r.PathPrefix("/public/").Handler(http.FileServer(http.Dir("web/")))
}
