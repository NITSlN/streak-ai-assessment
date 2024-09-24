package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NITSlN/streak-ai-assessment/handlers"
)

type Request struct {
	Nums   []int `json:"numbers"`
	Target int   `json:"target"`
}

type Response struct {
	Solution [][]int `json:"solution"`
}

func main() {
	http.HandleFunc("/find-pairs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		result := handlers.FindPairs(req.Nums, req.Target)
		json.NewEncoder(w).Encode(Response{Solution: result})
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
