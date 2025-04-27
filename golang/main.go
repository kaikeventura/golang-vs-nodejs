package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input map[string]int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	n, exists := input["number"]
	if !exists {
		http.Error(w, "Campo 'number' não encontrado", http.StatusBadRequest)
		return
	}

	results := []int{}
	for i := 0; i < 1000; i++ {
		results = append(results, fibonacci(n))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/fibonacci", fibonacciHandler)

	port := 8080
	fmt.Printf("Iniciando servidor na porta %d...\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

