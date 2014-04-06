package handlers

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
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
  auth, _ := EnvAuth()

  values := r.URL.Query()
  headers := r.Header
  uploadRequest, _ := NewS3UploadRequest(auth, values, headers)
  listResponse := uploadRequest.UploadListSignature()
  completeResponse := uploadRequest.UploadCompleteSignature()

  totalParts, err := strconv.Atoi(values.Get("total_chunks"))
  if err != nil {
    panic(err)
  }

  chunksSignatures := make(map[string]*S3UploadResponse)
  for index := 0; index < totalParts; index++ {
    fmt.Println(index)

    chunkNumber := strconv.Itoa(index + 1)
    values.Set("chunk", chunkNumber)
    values.Set("part_number", chunkNumber)

    request, _ := NewS3UploadRequest(auth, values, headers)
    chunksSignatures[chunkNumber] = request.UploadPartSignature()
  }

  remainingSignature := RemainingSignature{
    ListSignature:     listResponse,
    CompleteSignature: completeResponse,
    ChunkSignatures:   chunksSignatures,
  }

  marshal(&remainingSignature, w)
}

func Install(r *mux.Router) {
  r.HandleFunc("/", App)

  r.HandleFunc("/get_init_signature", getInitSignature)
  r.HandleFunc("/get_remaining_signatures", getRemainingSignatures)
  r.PathPrefix("/public/").Handler(http.FileServer(http.Dir("web/")))
}
