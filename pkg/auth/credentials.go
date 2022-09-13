package auth

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	log "github.com/sirupsen/logrus"
)

type CredentialsCache struct {
	AccessKeyId     string    `json:"accessKeyId"`
	Expiration      time.Time `json:"expiration"`
	SecretAccessKey string    `json:"secretAccessKey"`
	SessionToken    string    `json:"sessionToken"`
	// don't (un)marshal this
	cfgPath string `json:"-"`
}

func (cc *CredentialsCache) Load() (*CredentialsCache, error) {
	var (
		loadedCredentials CredentialsCache
	)
	data, err := os.ReadFile(cc.cfgPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &loadedCredentials)
	if err != nil {
		return nil, err
	}
	return &loadedCredentials, nil
}

func (cc *CredentialsCache) Save(credentials *CredentialsCache) error {
	data, err := json.Marshal(credentials)
	if err != nil {
		return err
	}
	return os.WriteFile(cc.cfgPath, data, 0644)
}

func (cc *CredentialsCache) Valid() bool {
	loaded, err := cc.Load()
	if err != nil {
		log.Debugf("Could not load config: %+v\n", err)
		return false
	}
	if time.Now().UTC().Before(loaded.Expiration) {
		return true
	}
	return false
}

func (cc *CredentialsCache) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     cc.AccessKeyId,
		SecretAccessKey: cc.SecretAccessKey,
		SessionToken:    cc.SessionToken,
		Source:          "AWS K8S Cache",
		CanExpire:       true,
		Expires:         cc.Expiration,
	}, nil
}

func NewCredentials(cfgPath string) *CredentialsCache {
	return &CredentialsCache{cfgPath: cfgPath}
}
