package models

import (
	"time"
)

type PaymentStatus string

const (
	StatusPending    PaymentStatus = "PENDING"
	StatusProcessing PaymentStatus = "PROCESSING"
	StatusCompleted  PaymentStatus = "COMPLETED"
	StatusFailed     PaymentStatus = "FAILED"
	StatusRejected   PaymentStatus = "REJECTED"
)

type Payment struct {
	ID             string
	IdempotencyKey string
	FromBankID     string
	ToBankID       string
	FromAccountID  string
	ToAccountID    string
	Amount         int64
	Currency       string
	Description    string
	Status         PaymentStatus
	CreatedAt      time.Time
	CompletedAt    *time.Time
	ErrorMessage   string
}

type Participant struct {
	ID                string
	Name              string
	SettlementAccount int64
	IsActive          bool
	JoinedAt          time.Time
}

type PaymentRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	FromBankID     string `json:"from_bank_id"`
	ToBankID       string `json:"to_bank_id"`
	FromAccountID  string `json:"from_account_id"`
	ToAccountID    string `json:"to_account_id"`
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
}

type PaymentResponse struct {
	PaymentID string        `json:"payment_id"`
	Status    PaymentStatus `json:"status"`
	Message   string        `json:"message,omitempty"`
}
