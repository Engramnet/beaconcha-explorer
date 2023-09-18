package models

import (
	"time"
)

// SlotsPageData is a struct to hold info for the slots page
type ValidatorSlotsPageData struct {
	Index uint64 `json:"index"`
	Name  string `json:"name"`

	Slots          []*ValidatorSlotsPageDataSlot `json:"slots"`
	SlotCount      uint64                        `json:"slot_count"`
	FirstSlot      uint64                        `json:"first_slot"`
	LastSlot       uint64                        `json:"last_slot"`
	GraffitiFilter string                        `json:"graffiti_filter"`

	IsDefaultPage    bool   `json:"default_page"`
	TotalPages       uint64 `json:"total_pages"`
	PageSize         uint64 `json:"page_size"`
	CurrentPageIndex uint64 `json:"page_index"`
	CurrentPageSlot  uint64 `json:"page_slot"`
	PrevPageIndex    uint64 `json:"prev_page_index"`
	PrevPageSlot     uint64 `json:"prev_page_slot"`
	NextPageIndex    uint64 `json:"next_page_index"`
	NextPageSlot     uint64 `json:"next_page_slot"`
	LastPageSlot     uint64 `json:"last_page_slot"`
}

type ValidatorSlotsPageDataSlot struct {
	Slot                  uint64    `json:"slot"`
	Epoch                 uint64    `json:"epoch"`
	Ts                    time.Time `json:"ts"`
	Finalized             bool      `json:"scheduled"`
	Scheduled             bool      `json:"finalized"`
	Status                uint8     `json:"status"`
	Proposer              uint64    `json:"proposer"`
	ProposerName          string    `json:"proposer_name"`
	AttestationCount      uint64    `json:"attestation_count"`
	DepositCount          uint64    `json:"deposit_count"`
	ExitCount             uint64    `json:"exit_count"`
	ProposerSlashingCount uint64    `json:"proposer_slashing_count"`
	AttesterSlashingCount uint64    `json:"attester_slashing_count"`
	SyncParticipation     float64   `json:"sync_participation"`
	EthTransactionCount   uint64    `json:"eth_transaction_count"`
	WithEthBlock          bool      `json:"with_eth_block"`
	EthBlockNumber        uint64    `json:"eth_block_number"`
	Graffiti              []byte    `json:"graffiti"`
	BlockRoot             []byte    `json:"block_root"`
}
