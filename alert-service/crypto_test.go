package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	
	ctx := context.Background()
	cryptoWatcher, err := NewCryptoWatcher(ctx, []currency{BTC, ETH, SOL})
	assert.NoError(t, err)
	

	err = cryptoWatcher.Run(ctx)
	assert.NoError(t, err)
}
