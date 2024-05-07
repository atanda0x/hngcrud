package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os/exec"

// 	"github.com/julienschmidt/httprouter"
// )

// func getCommand(command string, argument ...string) string {
// 	out, _ := exec.Command(command, argument...).Output()
// 	return string(out)
// }

// func goVersion(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
// 	res := getCommand("/usr/local/go/bin/go", "version")
// 	io.WriteString(w, res)
// 	return
// }

// func getFile(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
// 	fmt.Fprintf(w, getCommand("/bin/cat", param.ByName("name")))
// }

// func main() {
// 	router := httprouter.New()
// 	router.Get("/api/v1/go-version", goVersion)
// 	router.Get("/api/v1/show-file/:name", getFile)
// 	log.Fatal(http.ListenAndServe(":9000", router))
// }
