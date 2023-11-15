package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/metrics"
	"golang.org/x/crypto/bcrypt"
)

// Authentication Key Generation Config
const (
	charset     string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	authKeySize int    = 5
)

var (
	generatedAuthKeys = metrics.NewCounter(
		"generatedAuthKeys",
		"total number of generated authentication keys",
	)

	validAuthKeys = metrics.NewCounter(
		"validAuthKeys",
		"totoal number of successfully matching keys",
	)
)

// Receiver Config
//type ReceiverType int
//
//const (
//	Email ReceiverType = iota + 1
//	Mobile
//)

// GenerateAuthKey2FA generates a 2FA authentication key.
// The generated key will be hashed and stored via localStorage in JavaScript
// and should be removed from localStorage cache once verified.

type Authenticator interface {
	GenerateAuthKey2FA(context.Context) (string, error)
	VerifyAuthKey2FA(context.Context, string, string) (bool, error)
}

type impl struct {
	weaver.Implements[Authenticator]
}

func (t *impl) GenerateAuthKey2FA(ctx context.Context) (string, error) {
	generatedAuthKeys.Inc()
	log := t.Logger(ctx)

	source := rand.New(rand.NewSource(time.Now().Unix()))
	var authKey strings.Builder

	for i := 0; i < authKeySize; i++ {
		sample := source.Intn(len(charset))
		authKey.WriteString(string(charset[sample]))
	}
	fmt.Println(authKey.String())

	hashChannel := make(chan string, 1)
	go func() {
		authKeyHash, err := HashPassword(ctx, authKey.String())
		if err != nil {
			log.Error("vivian: [error]", "err", "failure hashing the authentication key")
			hashChannel <- ""
			return
		}
		hashChannel <- authKeyHash
	}()
	hash := <-hashChannel

	if hash == "" {
		log.Error("vivian: [error]", "err", "failure hashing the authentication key")
		return "", nil
	}

	log.Debug("vivian: [ok]", "authentication key generated", http.StatusOK)
	return hash, nil
}

func (t *impl) VerifyAuthKey2FA(ctx context.Context, authkey_hash, input string) (bool, error) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	if SanitizeCheck(input) {
		status := bcrypt.CompareHashAndPassword([]byte(authkey_hash), []byte(input))
		if status != nil {
			t.Logger(ctx).Debug("vivian: [warning]", "key invalid", http.StatusNotAcceptable)
			return status == nil, status
		} else {
			validAuthKeys.Inc()
			t.Logger(ctx).Debug("vivian: [ok]", "key verified", status == nil, "status", http.StatusOK)
			return status == nil, status
		}
	}

	return false, nil
}
