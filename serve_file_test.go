package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {

	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources") 

	// * untuk membaca file pathnya wajib menggunakan /*filepath, karena sudah bawaan httprouter
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:5000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HttpRouter", string(body), "Data Tidak Cocok")
}