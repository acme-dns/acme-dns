package api

import (
	"context"
	"testing"

	"github.com/joohoi/acme-dns/pkg/acmedns"
	"github.com/mholt/acmez/v3/acme"
)

type mockNameserver struct {
	acmedns.AcmednsNS
	authKey string
}

func (m *mockNameserver) SetOwnAuthKey(key string) {
	m.authKey = key
}

func TestChallengeProvider(t *testing.T) {
	mock := &mockNameserver{}
	servers := []acmedns.AcmednsNS{mock}
	cp := NewChallengeProvider(servers)

	ctx := context.Background()
	challenge := acme.Challenge{
		Type:             "dns-01",
		Token:            "test-token",
		KeyAuthorization: "test-key-auth",
	}

	// Test Present
	err := cp.Present(ctx, challenge)
	if err != nil {
		t.Errorf("Present failed: %v", err)
	}
	expectedKey := challenge.DNS01KeyAuthorization()
	if mock.authKey != expectedKey {
		t.Errorf("Expected auth key %s, got %s", expectedKey, mock.authKey)
	}

	// Test CleanUp
	err = cp.CleanUp(ctx, challenge)
	if err != nil {
		t.Errorf("CleanUp failed: %v", err)
	}
	if mock.authKey != "" {
		t.Errorf("Expected empty auth key after CleanUp, got %s", mock.authKey)
	}

	// Test Wait
	err = cp.Wait(ctx, challenge)
	if err != nil {
		t.Errorf("Wait failed: %v", err)
	}
}
