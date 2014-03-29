package handlers

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

var testAuth = Auth{"0PN5J17HBGZHT7JJ3X82", "uV3F3YluFJax1cknvbcGwgjvx4QpvB+leU8dUj2o", ""}

func TestSigningUrl(t *testing.T) {

  method := "GET"
  path := "/johnsmith/photos/puppy.jpg"
  headers := map[string][]string{
    "Host": {"johnsmith.s3.amazonaws.com"},
    "Date": {"Tue, 27 Mar 2007 19:36:42 +0000"},
  }

  sign(testAuth, method, path, nil, headers)

  expected := "AWS 0PN5J17HBGZHT7JJ3X82:xXjDGYUmKxnwqr5KXNPGldn5LbA="
  assert.Equal(t, headers["Authorization"], []string{expected})
}

func TestSignExampleList(t *testing.T) {
  method := "GET"
  path := "/johnsmith/"
  params := map[string][]string{
    "prefix":   {"photos"},
    "max-keys": {"50"},
    "marker":   {"puppy"},
    "Expires":  {"true"},
  }
  headers := map[string][]string{
    "Host":       {"johnsmith.s3.amazonaws.com"},
    "Date":       {"Tue, 27 Mar 2007 19:42:41 +0000"},
    "User-Agent": {"Mozilla/5.0"},
  }
  sign(testAuth, method, path, params, headers)
  expected := "g2t9J4eOm7fqoSi67Bbyn1A5k74="
  assert.Equal(t, params["Signature"], []string{expected})
}
