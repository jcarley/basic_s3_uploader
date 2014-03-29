package handlers

import (
  "strconv"
)

type S3UploadRequest struct {
  Auth      Auth
  Params    map[string][]string
  Headers   map[string][]string
  Acl       string
  Encrypted bool
  Signature string
}

type S3UploadResponse struct {
  Date      string `json:"date,omitempty"`
  UploadId  string `json:"upload_id,omitempty"`
  Key       string `json:"-"`
  Chunk     string `json:"chunk,omitempty"`
  MimeType  string `json:"mime_type,omitempty"`
  Bucket    string `json:"bucket,omitempty"`
  Signature string `json:"signature,omitempty"`
}

func NewS3UploadRequest(auth Auth, params, headers map[string][]string) (*S3UploadRequest, error) {
  encrypted, _ := strconv.ParseBool(valueOf("encrypted", params))
  acl := valueOf("acl", params)
  params["Expires"] = append(params["Expires"], "true")

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
  method := "GET"
  sign(this.Auth, method, this.valueOf("key"), this.Params, this.Headers)
  this.Signature = this.valueOf("Signature")
  return this.newResponse()
}

func (this *S3UploadRequest) newResponse() *S3UploadResponse {
  return &S3UploadResponse{
    Date:      this.valueOf("date"),
    Bucket:    this.valueOf("bucket"),
    UploadId:  this.valueOf("upload_id"),
    Key:       this.valueOf("key"),
    Chunk:     this.valueOf("chunk"),
    MimeType:  this.valueOf("mime_type"),
    Signature: this.Signature,
  }
}

func (this *S3UploadRequest) valueOf(paramName string) string {
  return valueOf(paramName, this.Params)
}

func valueOf(paramName string, params map[string][]string) string {
  if value := params[paramName]; value != nil {
    return value[0]
  }
  return ""
}
