package handlers

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

// encode("PUT\n\n#{@mime_type}\n\nx-amz-date:#{@date}\n/#{@bucket}/#{@key}?partNumber=#{@chunk}&uploadId=#{@upload_id}")
func TestUploadPartSignature(t *testing.T) {
  params := map[string][]string{
    "bucket":    {"foo_bar_bucket"},
    "upload_id": {"50abcdefg999"},
    "key":       {"quax/bar/wow.mov"},
    "chunk":     {"4"},
    "mime_type": {"application/mov"},
    "Expires":   {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "Date":      {"Tue, 27 Mar 2007 19:42:41 +0000"},
  }
  headers := map[string][]string{
    "Host":       {"johnsmith.s3.amazonaws.com"},
    "Date":       {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "User-Agent": {"Mozilla/5.0"},
  }
  request, _ := NewS3UploadRequest(testAuth, params, headers)
  response := request.UploadPartSignature()

  assert.Equal(t, "foo_bar_bucket", response.Bucket)
  assert.Equal(t, "50abcdefg999", response.UploadId)
  assert.Equal(t, "quax/bar/wow.mov", response.Key)
  assert.Equal(t, "4", response.Chunk)
  assert.Equal(t, "4", response.PartNumber)
  assert.Equal(t, "application/mov", response.MimeType)
  assert.Equal(t, "Tue, 27 Mar 2007 19:42:41 +0000", response.Date)
  assert.Equal(t, "+l3Ry9oeeIQrRzZecQ8hlMAasvk=", response.Signature)
}

// encode("POST\n\n#{@mime_type}\n\nx-amz-date:#{@date}\n/#{@bucket}/#{@key}?uploadId=#{@upload_id}")
func TestUploadCompleteSignature(t *testing.T) {
  params := map[string][]string{
    "bucket":    {"foo_bar_bucket"},
    "upload_id": {"50abcdefg999"},
    "key":       {"quax/bar/wow.mov"},
    "mime_type": {"application/mov"},
    "Expires":   {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "Date":      {"Tue, 27 Mar 2007 19:42:41 +0000"},
  }
  headers := map[string][]string{
    "Host":       {"johnsmith.s3.amazonaws.com"},
    "Date":       {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "User-Agent": {"Mozilla/5.0"},
  }
  request, _ := NewS3UploadRequest(testAuth, params, headers)
  response := request.UploadCompleteSignature()

  assert.Equal(t, "foo_bar_bucket", response.Bucket)
  assert.Equal(t, "50abcdefg999", response.UploadId)
  assert.Equal(t, "quax/bar/wow.mov", response.Key)
  assert.Equal(t, "application/mov", response.MimeType)
  assert.Equal(t, "Tue, 27 Mar 2007 19:42:41 +0000", response.Date)
  assert.Equal(t, "1l2y7FFwE+HTnFSg2gFM2Tp9Eik=", response.Signature)
}

// encode("GET\n\n\n\nx-amz-date:#{@date}\n/#{@bucket}/#{@key}?uploadId=#{@upload_id}")
func TestUploadListSignature(t *testing.T) {
  params := map[string][]string{
    "bucket":    {"foo_bar_bucket"},
    "upload_id": {"50abcdefg999"},
    "key":       {"quax/bar/wow.mov"},
    "Expires":   {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "Date":      {"Tue, 27 Mar 2007 19:42:41 +0000"},
  }
  headers := map[string][]string{
    "Host":       {"johnsmith.s3.amazonaws.com"},
    "Date":       {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "User-Agent": {"Mozilla/5.0"},
  }
  request, _ := NewS3UploadRequest(testAuth, params, headers)
  response := request.UploadListSignature()

  assert.Equal(t, "foo_bar_bucket", response.Bucket)
  assert.Equal(t, "50abcdefg999", response.UploadId)
  assert.Equal(t, "quax/bar/wow.mov", response.Key)
  assert.Equal(t, "Tue, 27 Mar 2007 19:42:41 +0000", response.Date)
  assert.Equal(t, "tHzdxjC269mA/YaZZihInEy0H/I=", response.Signature)
}
