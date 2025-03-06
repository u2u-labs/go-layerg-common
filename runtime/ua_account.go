package runtime

import "time"

type UADirectLoginResponse struct {
	Success bool              `json:"success"`
	Data    UADirectLoginData `json:"data"`
	Message string            `json:"message"`
}

type UADirectLoginData struct {
	User               AccountUA `json:"user"`
	W                  W         `json:"w"`
	RefreshToken       string    `json:"refreshToken"`
	RefreshTokenExpire int64     `json:"refreshTokenExpire"`
	AccessToken        string    `json:"accessToken"`
	AccessTokenExpire  int64     `json:"accessTokenExpire"`
	UserID             string    `json:"userId"`
	APIKey             string    `json:"apiKey"`
}

type AccountUA struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	Type                string    `json:"type"`
	EoaWallet           string    `json:"eoaWallet"`
	EncryptedPrivateKey string    `json:"encryptedPrivateKey"`
	Username            *string   `json:"username"`
	Firstname           *string   `json:"firstname"`
	Lastname            *string   `json:"lastname"`
	Avatar              *string   `json:"avatar"`
	Signature           *string   `json:"signature"`
	PrimaryAAWallet     string    `json:"primaryAAWallet"`
	TelegramID          *string   `json:"telegramId"`
	TelegramUsername    *string   `json:"telegramUsername"`
	TelegramFirstName   *string   `json:"telegramFirstName"`
	TelegramLastName    *string   `json:"telegramLastName"`
	TelegramAvatarUrl   *string   `json:"telegramAvatarUrl"`
	Email               *string   `json:"email"`
	FacebookID          *string   `json:"facebookId"`
	FacebookEmail       *string   `json:"facebookEmail"`
	FacebookFirstName   *string   `json:"facebookFirstName"`
	FacebookLastName    *string   `json:"facebookLastName"`
	FacebookAvartaUrl   *string   `json:"facebookAvartaUrl"`
	GoogleID            *string   `json:"googleId"`
	GoogleEmail         *string   `json:"googleEmail"`
	GoogleFirstName     *string   `json:"googleFirstName"`
	GoogleLastName      *string   `json:"googleLastName"`
	GoogleAvatarUrl     *string   `json:"googleAvatarUrl"`
	TwitterID           *string   `json:"twitterId"`
	TwitterEmail        *string   `json:"twitterEmail"`
	TwitterFirstName    *string   `json:"twitterFirstName"`
	TwitterLastName     *string   `json:"twitterLastName"`
	TwitterAvatarUrl    *string   `json:"twitterAvatarUrl"`
}

type W struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	AAAddress      string    `json:"aaAddress"`
	FactoryAddress string    `json:"factoryAddress"`
	UserID         string    `json:"userId"`
}
