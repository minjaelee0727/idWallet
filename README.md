# idWallet

This README is written by gpt-4o. 

## Overview

**idWallet** is a decentralized identity wallet system enabling users to store, present, and verify digital credentials securely. It uses **blockchain anchoring** to ensure the integrity and traceability of credentials, all without relying on zero-knowledge proof schemes.

This monorepo includes:

- 🧠 A **Go backend** that issues, anchors, and verifies credentials on-chain
- 📱 A **Swift-based iOS app** that acts as a mobile identity wallet for users

The architecture is built for **transparency, verifiability**, and **developer-friendliness**, balancing privacy with simplicity.

---

## ✨ Features

### ✅ Cross-Platform Identity Flow

- Users manage verifiable credentials on iOS.
- Backend handles credential issuance, blockchain anchoring, and verification.
- Supports third-party verifiers without requiring zero-knowledge proof infrastructure.

### 🔐 Blockchain Anchored Credentials (Backend)
- Credentials are hashed and the digest is stored on-chain.
- Fast and transparent verification by hash comparison.
- No reliance on central authorities.

### 📱 Secure iOS Wallet App
- Built in Swift with intuitive UI for managing identities.
- Local secure storage using Apple’s keychain and biometric protection.
- Sends and receives credentials via the REST backend.

### 🧩 Modular, Lightweight Architecture
- Go REST API handles core business logic and persists to SQLite.
- Pluggable blockchain layer (abstracted to support different RPCs).
- Designed for future compatibility with DID, smart contracts, and revocation lists.

---

## 📁 Project Structure

```
idWallet/
├── backend/         # Go-based credential service and blockchain anchor
│   ├── main.go
│   ├── blockchain.go
│   ├── rest.go
│   ├── db.go
│   ├── wallet.go
│   └── utils.go
├── iOS/             # Swift iOS Wallet App (Xcode project)
│   ├── idWallet.xcodeproj
│   ├── Models/
│   ├── Views/
│   ├── Controllers/
│   └── Network/
```

---


### 🧠 Backend (Go)

#### Prerequisites

- Go 1.16+
- SQLite3
- Access to a blockchain RPC endpoint (e.g., Infura, Ganache)

### 📱 iOS App (Swift)

#### Prerequisites

- Xcode 12+
- iOS 14+ device or simulator
---

## 🔌 REST API

### `POST /issue`

Issues and anchors a credential.

```json
{
  "holder": "did:example:1234",
  "data": {
    "name": "Alice",
    "org": "Waterloo University",
    "issuedAt": "2025-05-08"
  }
}
```

### `POST /verify`

Verifies credential by recomputing and comparing the hash.

```json
{
  "credential": {
    "holder": "did:example:1234",
    "data": {
      "name": "Alice",
      "org": "Waterloo University",
      "issuedAt": "2025-05-08"
    }
  }
}
```

---

## 🧪 Example Flow

1. User creates a credential on iOS → Sends to backend
2. Backend:
   - Hashes the credential
   - Stores it locally
   - Anchors hash to blockchain
3. Later, anyone can verify:
   - Client sends credential
   - Backend rehashes and checks blockchain state
   - Responds with `valid: true/false`
