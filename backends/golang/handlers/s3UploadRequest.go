package handlers

import (
  "net/http"
  "net/url"
  "strconv"
  "strings"
  "time"
)

type RemainingSignature struct {
  ListSignature     *S3UploadResponse            `json:"list_signature,omitempty"`
  CompleteSignature *S3UploadResponse            `json:"complete_signature,omitempty"`
  ChunkSignatures   map[string]*S3UploadResponse `json:"chunk_signatures,omitempty"`
}

type S3UploadResponse struct {
  Date       string `json:"date,omitempty"`
  UploadId   string `json:"upload_id,omitempty"`
  Key        string `json:"-"`
  Chunk      string `json:"chunk,omitempty"`
  PartNumber string `json:"part_number,omitempty"`
  MimeType   string `json:"mime_type,omitempty"`
  Bucket     string `json:"bucket,omitempty"`
  Signature  string `json:"signature,omitempty"`
}

type S3UploadRequest struct {
  Auth      Auth
  Method    string
  Params    url.Values
  Headers   http.Header
  Acl       string
  Encrypted bool
  Signature string
}

func NewS3UploadRequest(auth Auth, params url.Values, headers http.Header) (*S3UploadRequest, error) {
  encrypted, _ := strconv.ParseBool(valueOf("encrypted", params))
  acl := valueOf("acl", params)
  headers.Set("Date", string(time.Now().In(time.UTC).Format(time.RFC1123)))
  // if valueOf("date", headers) == "" {
  // headers["Date"] = []
  // }
  params["Expires"] = headers["Date"]
  params.Set("Date", headers.Get("Date"))

  params = mapValue("chunk", "partNumber", params)
  params = mapValue("upload_id", "uploadId", params)

  uploadRequest := S3UploadRequest{
    Auth:      auth,
    Params:    params,
    Headers:   headers,
    Acl:       acl,
    Encrypted: encrypted,
  }
  return &uploadRequest, nil
}

func (this *S3UploadRequest) UploadInitSignature() *S3UploadResponse {
  this.Method = "GET"
  return this.signRequest()
}

func (this *S3UploadRequest) UploadPartSignature() *S3UploadResponse {
  this.Method = "PUT"
  return this.signRequest()
}

func (this *S3UploadRequest) UploadCompleteSignature() *S3UploadResponse {
  this.Method = "POST"
  return this.signRequest()
}

func (this *S3UploadRequest) UploadListSignature() *S3UploadResponse {
  this.Method = "GET"
  return this.signRequest()
}

func (this *S3UploadRequest) signRequest() *S3UploadResponse {
  sign(this.Auth, this.Method, this.valueOf("key"), this.Params, this.Headers)
  this.Signature = this.valueOf("Signature")
  return this.newResponse()
}

func (this *S3UploadRequest) newResponse() *S3UploadResponse {
  return &S3UploadResponse{
    Date:       this.valueOf("date"),
    Bucket:     this.valueOf("bucket"),
    UploadId:   this.valueOf("upload_id"),
    Key:        this.valueOf("key"),
    Chunk:      this.valueOf("chunk"),
    PartNumber: this.valueOf("chunk"),
    MimeType:   this.valueOf("mime_type"),
    Signature:  this.Signature,
  }
}

func (this *S3UploadRequest) valueOf(paramName string) string {
  return valueOf(paramName, this.Params)
}

func valueOf(paramName string, params map[string][]string) string {
  for k, v := range params {
    k = strings.ToLower(k)
    if k == strings.ToLower(paramName) {
      return v[0]
    }
  }
  return ""
}

func mapValue(from string, to string, params map[string][]string) map[string][]string {
  fromValue := valueOf(from, params)
  if fromValue != "" {
    params[to] = []string{fromValue}
  }
  return params
}
