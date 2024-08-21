package firebaseadmin

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseAdmin struct {
	client *auth.Client
}

func NewFirebaseAdmin(serviceAccountBase64 string) (*FirebaseAdmin, error) {
	serviceAccountJSON, err := base64.StdEncoding.DecodeString(serviceAccountBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 service account string: %v", err)
	}

	opt := option.WithCredentialsJSON(serviceAccountJSON)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &FirebaseAdmin{client: client}, nil
}

func (firebaseAdmin *FirebaseAdmin) VerifyToken(token string) (*auth.Token, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}

	tokenData, err := firebaseAdmin.client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("error verifying token: %v", err)
	}

	return tokenData, nil
}

func (firebaseAdmin *FirebaseAdmin) GetUser(uid string) (*auth.UserRecord, error) {
	user, err := firebaseAdmin.client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (firebaseAdmin *FirebaseAdmin) CreateCustomToken(
	uid string,
	customClaims map[string]interface{},
) (string, error) {
	customToken, err := firebaseAdmin.client.
		CustomTokenWithClaims(context.Background(), uid, customClaims)
	if err != nil {
		return "", fmt.Errorf("error generating custom token: %v", err)
	}

	return customToken, nil
}
