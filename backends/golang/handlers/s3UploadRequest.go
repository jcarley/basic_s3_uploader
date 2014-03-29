package handlers

type S3UploadRequest struct {
  Date      string `json:"date"`
  UploadId  string `json:"upload_id"`
  Key       string `json:"-"`
  Chunk     string `json:"chunk"`
  MimeType  string `json:"mime_type"`
  Bucket    string `json:"bucket"`
  Signature string `json:"signature"`
  Acl       string `json:"-"`
  Encrypted bool   `json:"-"`
}

func (this *S3UploadRequest) NewS3UploadRequest(date string, bucket string, uploadId string, key string, chunk string, mimeType string, acl string, encrypted bool) (*S3UploadRequest, error) {
  uploadRequest := S3UploadRequest{
    Date:      date,
    Bucket:    bucket,
    UploadId:  uploadId,
    Key:       key,
    Chunk:     chunk,
    MimeType:  mimeType,
    Acl:       acl,
    Encrypted: encrypted,
  }
  return &uploadRequest, nil
}

func (this *S3UploadRequest) uploadInitSignature() string {

  // func sign(auth aws.Auth, method, canonicalPath string, params, headers map[string][]string) {
  return ""
}

// def upload_init_signature
// if encrypted_upload?
// encode("POST\n\n\n\nx-amz-acl:#{@acl}\nx-amz-date:#{@date}\nx-amz-server-side-encryption:AES256\n/#{@bucket}/#{@key}?uploads")
// else
// encode("POST\n\n\n\nx-amz-acl:#{@acl}\nx-amz-date:#{@date}\n/#{@bucket}/#{@key}?uploads")
// end
// end
