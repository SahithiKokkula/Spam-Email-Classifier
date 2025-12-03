package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type PredictionRequest struct {
	EmailText string `json:"email_text"`
}

type PredictionResponse struct {
	IsSpam    bool    `json:"is_spam"`
	Confidence float64 `json:"confidence"`
	Label     string  `json:"label"`
	Error     string  `json:"error,omitempty"`
}

func predictSpam(emailText string) (*PredictionResponse, error) {
	var cmd *exec.Cmd
	
	workDir := "."
	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", "predict.py", emailText)
	} else {
		cmd = exec.Command("python3", "predict.py", emailText)
	}
	
	cmd.Dir = workDir
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("prediction failed: %v", err)
	}

	var response PredictionResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("failed to parse prediction: %v", err)
	}

	return &response, nil
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PredictionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := PredictionResponse{Error: "Invalid request body"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if strings.TrimSpace(req.EmailText) == "" {
		response := PredictionResponse{Error: "Email text cannot be empty"}
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := predictSpam(req.EmailText)
	if err != nil {
		response := PredictionResponse{Error: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	r := mux.NewRouter()
	
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/predict", predictHandler).Methods("POST", "OPTIONS")
	api.HandleFunc("/health", healthHandler).Methods("GET")
	
	frontendDir := "./frontend"
	if _, err := os.Stat(frontendDir); os.IsNotExist(err) {
		frontendDir = "../frontend"
	}
	
	fs := http.FileServer(http.Dir(frontendDir))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	log.Fatal(http.ListenAndServe(port, corsHandler))
}

