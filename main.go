package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/api/k8s/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("C:\\Apps\\kubectl.exe", "get", "pods", "-o", "json")

		stdout, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		test := string(stdout)
		_ = test
		fmt.Fprintf(w, "%q", stdout)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
