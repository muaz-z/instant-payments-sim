# Instant Payments Simulator (`instant-payments-sim`)

A simulation of an **instant account-to-account payment network** inspired by real-world systems like **DuitNow (Malaysia)**, **Pix (Brazil)**, and **Wero (Europe)**.  
This project demonstrates how **QR-based, bank-to-bank instant payments** work, including ledger management, transaction flow, and merchant confirmation — all in a **safe, simulated environment**.

---

## 🚀 Features

- **QR Code Payments**

  - Merchants generate QR codes for payment requests
  - Users scan QR codes and approve transactions in-app

- **Instant Account-to-Account Transfers**

  - Simulated real-time transfer between user and merchant accounts
  - Ledger ensures **transaction integrity** and prevents double-spending

- **Transaction Ledger**

  - Tracks balances and transaction history for users and merchants
  - Handles reversals / refunds (simulated)

- **Idempotency & Atomicity**

  - Ensures each payment is processed exactly once
  - Avoids race conditions in concurrent payments

- **Simulated Merchant Dashboard**
  - Shows received payments
  - Displays QR codes for testing

---

## 🧰 Tech Stack

- **Backend:** Go (Golang) – REST API / Payment logic
- **Database:** PostgreSQL (ledger, users, merchants)
- **Cache / Concurrency:** Redis (idempotency / locks)
- **Mobile / Frontend:** React Native
- **QR Scanning:** In-app camera or QR code reader library
- **Authentication:** JWT / OAuth (simulated)

---

## 🔄 Payment Flow

1. Merchant generates a payment QR code with:
   - Merchant ID
   - Payment amount
   - Reference ID
2. User scans QR code via mobile app
3. App displays payment details and asks for confirmation
4. Backend (Go):
   - Validates balance
   - Updates ledger atomically
   - Marks transaction as completed
5. Merchant receives confirmation instantly

---

## 💡 Learning Goals

This project is designed to **teach and showcase**:

- How **instant payments** work under the hood
- Account-to-account transaction modeling
- Ledger design and consistency
- QR-based payment flows and merchant interactions
- Idempotency, atomicity, and basic concurrency control

---

## ⚠️ Disclaimer

This is a **simulation only**.

- No real money is transferred
- Not connected to any banks or financial networks
- Safe for learning and demonstration purposes
