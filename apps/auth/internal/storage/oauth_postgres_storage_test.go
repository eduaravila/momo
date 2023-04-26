package storage_test

import (
	"context"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/packages/postgres/queries"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",
	"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
	"Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
	"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.91 Mobile Safari/537.36 OPR/42.9.2246.119956",
	"Opera/9.80 (Android; Opera Mini/28.0.2254/66.318; U; en) Presto/2.12.423 Version/12.16",
}

func TestSessionPosgresStorage_AddSession(t *testing.T) {
	repo := newSessionPostgresStorage(t)

	t.Parallel()
	testCases := []struct {
		name               string
		SessionConstructor func(t *testing.T) *session.Session
	}{
		{
			name:               "standard_session",
			SessionConstructor: newExampleSession,
		},
	}

	for _, c := range testCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			session := c.SessionConstructor(t)
			err := repo.AddSession(ctx, session)
			require.NoError(t, err)
			assertPersistedSession(t, repo, session)
		})
	}
}

func assertPersistedSession(t *testing.T, repo session.Storage, se *session.Session) {
	t.Helper()

	ctx := context.Background()

	persistedSession, err := repo.GetSession(ctx, se.ID)
	require.NoError(t, err)

	assertSession(t, se, persistedSession)
}

var cmpRoundTimeOpt = cmp.Comparer(func(x, y time.Time) bool {
	return x.Round(time.Second).Equal(y.Round(time.Second))
})

func assertSession(t *testing.T, expected, actual *session.Session) {
	t.Helper()

	cmpOpts := cmp.Options{
		cmpRoundTimeOpt,
		cmp.AllowUnexported(session.Session{}),
	}
	assert.True(t, cmp.Equal(expected, actual, cmpOpts))
}

func newExampleClaims(userID string) *session.Claims {
	return session.NewClaims(
		os.Getenv("JWT_ISSUER"),
		userID,
		time.Now().Add(time.Hour*24*30),
		time.Now(),
		time.Now())
}

func newExampleSession(t *testing.T) *session.Session {
	userID := uuid.NewString()

	token, err := session.NewSessionToken(
		uuid.NewString(),
		true,
		newExampleClaims(userID),
	)

	require.NoError(t, err)
	session, err := session.NewSession(
		uuid.NewString(),
		userID,
		time.Now(),
		token,
		userAgents[rand.Intn(len(userAgents))],
		newExampleIP(),
	)
	require.NoError(t, err)

	return session
}

func newExampleIP() string {
	rand.Seed(time.Now().UnixNano())
	ip := net.IPv4(
		byte(rand.Intn(255)),
		byte(rand.Intn(255)),
		byte(rand.Intn(255)),
		byte(rand.Intn(255)),
	)

	return ip.String()
}

func newSessionPostgresStorage(t *testing.T) *storage.OauthPostgresStorage {
	t.Helper()

	db, err := storage.InitPostgresDB()
	require.NoError(t, err)

	return storage.NewSessionPostgresStorage(queries.New(db))
}
