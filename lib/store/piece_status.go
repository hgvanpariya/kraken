package store

import (
	"regexp"

	"code.uber.internal/infra/kraken/lib/store/base"
)

func init() {
	base.RegisterMetadata(regexp.MustCompile("_status"), &pieceStatusFactory{})
}

type pieceStatusFactory struct{}

func (f pieceStatusFactory) Create(suffix string) base.MetadataType {
	return NewPieceStatus()
}

type pieceStatus struct{}

// NewPieceStatus initializes and returns an new MetadataType obj.
func NewPieceStatus() base.MetadataType {
	return pieceStatus{}
}

func (p pieceStatus) GetSuffix() string {
	return "_status"
}

func (p pieceStatus) Movable() bool {
	return true
}
