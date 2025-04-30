package repositories

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

type SessionRepository struct {
    rdb *redis.Client
}

func NewSessionRepository(rdb *redis.Client) *SessionRepository {
    return &SessionRepository{rdb: rdb}
}

func (s *SessionRepository) SaveSession(ctx context.Context, sessionID string, userID string, expiration time.Duration) error {
    return s.rdb.Set(ctx, fmt.Sprintf("session:%s", sessionID), userID, expiration).Err()
}

func (s *SessionRepository) GetUserIDBySession(ctx context.Context, sessionID string) (string, error) {
    return s.rdb.Get(ctx, fmt.Sprintf("session:%s", sessionID)).Result()
}

func (s *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
    return s.rdb.Del(ctx, fmt.Sprintf("session:%s", sessionID)).Err()
}

func (s *SessionRepository) CacheUser(ctx context.Context, userID string, userData string, expiration time.Duration) error {
    return s.rdb.Set(ctx, fmt.Sprintf("user_cache:%s", userID), userData, expiration).Err()
}

func (s *SessionRepository) GetCachedUser(ctx context.Context, userID string) (string, error) {
    return s.rdb.Get(ctx, fmt.Sprintf("user_cache:%s", userID)).Result()
}
