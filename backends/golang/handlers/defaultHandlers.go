package handlers

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
)

func marshal(item interface{}, w http.ResponseWriter) {
  bytes, err := json.Marshal(item)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  w.Write(bytes)
}

func App(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "web/public/views/index.html")
}

func getInitSignature(w http.ResponseWriter, r *http.Request) {
  auth, _ := EnvAuth()

  values := r.URL.Query()
  headers := r.Header
  uploadRequest, _ := NewS3UploadRequest(auth, values, headers)
  response := uploadRequest.UploadInitSignature()

  marshal(response, w)
}

func getRemainingSignatures(w http.ResponseWriter, r *http.Request) {
}

func Install(r *mux.Router) {
  r.HandleFunc("/", App)

  r.HandleFunc("/get_init_signature", getInitSignature)
  r.HandleFunc("/get_remaining_signatures", getRemainingSignatures)
  r.PathPrefix("/public/").Handler(http.FileServer(http.Dir("web/")))
}
