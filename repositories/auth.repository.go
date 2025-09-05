package repositories

import (
	"authentication/requests"
	"authentication/utils"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	SetOTP(ctx context.Context, phone string, code int, ttl time.Duration)
	GetOTP(ctx context.Context, phone string) string
	UserExists(ctx context.Context, phone string) bool
	CreateUser(ctx context.Context, phone string) map[string]string
	GetUser(ctx context.Context, phone string) map[string]string
	SetRefreshToken(ctx context.Context, phone, refreshToken string, ttl time.Duration) error
	GetRefreshToken(ctx context.Context, phone string) (string, error)
	ListUsers(ctx context.Context, request requests.UsersList) []map[string]string
}

type authRepository struct {
	redisConnection *redis.Client
}

func NewAuthRepository(redisConnection *redis.Client) AuthRepository {
	return &authRepository{
		redisConnection: redisConnection,
	}
}

func (r *authRepository) SetOTP(ctx context.Context, phone string, code int, ttl time.Duration) {
	key := "otp:" + phone
	set, err := r.redisConnection.SetNX(ctx, key, code, ttl).Result()
	if err != nil {
		panic(err)
	}
	if !set {
		panic(utils.PanicMessage{MessageKey: 5})
	}
}

func (r *authRepository) GetOTP(ctx context.Context, phone string) string {
	key := "otp:" + phone
	res, err := r.redisConnection.Get(ctx, key).Result()
	if err == redis.Nil {
		panic(utils.PanicMessage{MessageKey: 3})
	} else if err != nil {
		panic(err)
	}
	return res
}

func (r *authRepository) UserExists(ctx context.Context, phone string) bool {
	key := "user:" + phone
	exists, err := r.redisConnection.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return exists > 0
}

func (r *authRepository) CreateUser(ctx context.Context, phone string) map[string]string {
	user := map[string]string{
		"id":    fmt.Sprintf("user-%d", time.Now().UnixNano()),
		"phone": phone,
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	key := "user:" + phone
	if err := r.redisConnection.Set(ctx, key, data, 0).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Saved user:", key)

	if err := r.redisConnection.LPush(ctx, "users", phone).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pushed phone into users list:", phone)

	return user
}

func (r *authRepository) GetUser(ctx context.Context, phone string) map[string]string {
	key := "user:" + phone
	data, err := r.redisConnection.Get(ctx, key).Result()
	if err == redis.Nil {
		panic(utils.PanicMessage{MessageKey: 4})
	} else if err != nil {
		panic(err)
	}

	var user map[string]string
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		panic(err)
	}

	return user
}

func (r *authRepository) SetRefreshToken(ctx context.Context, phone, refreshToken string, ttl time.Duration) error {
	key := "refresh:" + phone
	return r.redisConnection.Set(ctx, key, refreshToken, ttl).Err()
}

func (r *authRepository) GetRefreshToken(ctx context.Context, phone string) (string, error) {
	key := "refresh:" + phone
	token, err := r.redisConnection.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("refresh token not found")
	}
	return token, err
}

func (r *authRepository) ListUsers(ctx context.Context, request requests.UsersList) []map[string]string {
	phones, err := r.redisConnection.LRange(ctx, "users", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	filteredPhones := make([]string, 0)
	if request.PhoneLike != "" {
		for _, phone := range phones {
			if strings.Contains(phone, request.PhoneLike) || phone == request.PhoneLike {
				filteredPhones = append(filteredPhones, phone)
			}
		}
	} else {
		filteredPhones = phones
	}

	start := (request.Page - 1) * request.PageSize
	end := start + request.PageSize
	if start >= int64(len(filteredPhones)) {
		return []map[string]string{}
	}
	if end > int64(len(filteredPhones)) {
		end = int64(len(filteredPhones))
	}

	users := make([]map[string]string, 0, end-start)
	for _, phone := range filteredPhones[start:end] {
		data, err := r.redisConnection.Get(ctx, "user:"+phone).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			panic(err)
		}

		var user map[string]string
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users
}
