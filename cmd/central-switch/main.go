package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muaz-z/instant-payments-sim/internal/clearing"
	"github.com/muaz-z/instant-payments-sim/pkg/models"
)

func main() {
	// Create the central switch
	cs := clearing.NewCentralSwitch()

	// Setup HTTP routes
	http.HandleFunc("/register", handleRegister(cs))
	http.HandleFunc("/payments", handlePayment(cs))
	http.HandleFunc("/payments/status", handlePaymentStatus(cs))

	// Start server
	port := ":8080"
	log.Printf("Central Switch starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// handleRegister returns a handler for participant registration
func handleRegister(cs *clearing.CentralSwitch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var p models.Participant
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// if err := cs.RegisterParticipant(&p); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		cs.RegisterParticipant(&p)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

func handlePayment(cs *clearing.CentralSwitch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.PaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// resp, err := cs.ProcessPayment(&req)
		// if err != nil {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(resp)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(resp)
		cs.ProcessPayment(&req)
	}
}

func handlePaymentStatus(cs *clearing.CentralSwitch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		paymentID := r.URL.Query().Get("id")
		if paymentID == "" {
			http.Error(w, "Missing payment ID", http.StatusBadRequest)
			return
		}

		// payment, err := cs.GetPayment(paymentID)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusNotFound)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(payment)
		cs.GetPayment(paymentID)

	}
}
