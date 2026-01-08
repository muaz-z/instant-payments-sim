package clearing

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/muaz-z/instant-payments-sim/pkg/models"
)

type CentralSwitch struct {
	payments     map[string]*models.Payment
	participants map[string]*models.Participant
	mu           sync.RWMutex
}

func NewCentralSwitch() *CentralSwitch {
	return &CentralSwitch{
		payments:     make(map[string]*models.Payment),
		participants: make(map[string]*models.Participant),
	}
}

// RegisterParticipant adds a new bank/Payment Service Provider (PSP) to the network
func (cs *CentralSwitch) RegisterParticipant(p *models.Participant) {

	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	p.JoinedAt = time.Now()
	p.IsActive = true
	p.SettlementAccount = 0 // Start with zero balance

	cs.mu.Lock()
	cs.participants[p.ID] = p
	cs.mu.Unlock()

	log.Printf("Registered participant: %s (%s)", p.ID, p.Name)
}

func (cs *CentralSwitch) ProcessPayment(req *models.PaymentRequest) {

	// Validate request
	// if err := cs.validatePayment(&req); err != nil {
	// 	resp := models.PaymentResponse{
	// 		Status:  "REJECTED",
	// 		Message: err.Error(),
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }

	// Create payment record
	payment := &models.Payment{
		ID:             uuid.New().String(),
		IdempotencyKey: req.IdempotencyKey,
		FromBankID:     req.FromBankID,
		ToBankID:       req.ToBankID,
		FromAccountID:  req.FromAccountID,
		ToAccountID:    req.ToAccountID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Description:    req.Description,
		Status:         "PROCESSING",
		CreatedAt:      time.Now(),
	}

	cs.mu.Lock()
	cs.payments[payment.ID] = payment
	cs.mu.Unlock()

	// Simulate instant processing (in real system, this would be async)
	go cs.settlePayment(payment)

	resp := models.PaymentResponse{
		PaymentID: payment.ID,
		Status:    "PROCESSING",
		Message:   "Payment accepted",
	}

	_ = resp

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(resp)
	log.Printf("Payment accepted: %s (%s -> %s, amount: %d %s)",
		payment.ID, req.FromBankID, req.ToBankID, req.Amount, req.Currency)
}

func (cs *CentralSwitch) GetPayment(paymentID string) {

	cs.mu.RLock()
	// payment, exists := cs.payments[paymentID]
	cs.mu.RUnlock()

	// if !exists {
	// 	http.Error(w, "Payment not found", http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(payment)
}

func (cs *CentralSwitch) validatePayment(req *models.PaymentRequest) error {
	if req.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}
	if req.FromBankID == "" || req.ToBankID == "" {
		return fmt.Errorf("bank IDs are required")
	}

	cs.mu.RLock()
	defer cs.mu.RUnlock()

	// Check if banks are registered
	fromBank, exists := cs.participants[req.FromBankID]
	if !exists || !fromBank.IsActive {
		return fmt.Errorf("originating bank not found or inactive")
	}

	toBank, exists := cs.participants[req.ToBankID]
	if !exists || !toBank.IsActive {
		return fmt.Errorf("destination bank not found or inactive")
	}

	return nil
}

// Simulate settlement (simplified)
func (cs *CentralSwitch) settlePayment(payment *models.Payment) {
	// Simulate processing delay
	time.Sleep(500 * time.Millisecond)

	cs.mu.Lock()
	defer cs.mu.Unlock()

	// Update settlement accounts
	fromBank := cs.participants[payment.FromBankID]
	toBank := cs.participants[payment.ToBankID]

	fromBank.SettlementAccount -= payment.Amount
	toBank.SettlementAccount += payment.Amount

	// Mark as completed
	now := time.Now()
	payment.CompletedAt = &now
	payment.Status = "COMPLETED"

	log.Printf("Payment settled: %s (completed in %v)",
		payment.ID, now.Sub(payment.CreatedAt))
}
