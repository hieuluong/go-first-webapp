package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/api/k8s/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("C:\\ExtraApps\\kubectl.exe", "get", "pods", "-o", "json")
		test := "test abc"
		_ = test
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Hello Hieu, %q", stdout)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
