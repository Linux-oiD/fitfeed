package entity

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type Passkey struct {
	Base
	UserID          uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	CredentialID    []byte    `gorm:"type:bytea;index" json:"credential_id"`
	PublicKey       []byte    `gorm:"type:bytea" json:"public_key"`
	AttestationType string    `gorm:"size:255" json:"attestation_type"`
	AAGUID          []byte    `gorm:"type:bytea" json:"aaguid"`
	SignCount       uint32    `json:"sign_count"`
}

func (p Passkey) WebAuthnCredential() webauthn.Credential {
	return webauthn.Credential{
		ID:              p.CredentialID,
		PublicKey:       p.PublicKey,
		AttestationType: p.AttestationType,
		Authenticator: webauthn.Authenticator{
			AAGUID:    p.AAGUID,
			SignCount: p.SignCount,
		},
	}
}
