package internal

import (
	"time"
)

type BuyingDecisionData struct {
	Email string `json:"email"`
}

type ClientUser struct {
	Email        string `json:"email"`
	ClientUserId string `json:"client_user_id"`
	ProjectId    int    `json:"project_id"`
}

type CheckEmail struct {
	Email    string `json:"email"`
	Sanitize bool   `json:"sanitize"`
}

type ValidateEmail struct {
	Email     string `json:"email"`
	ProjectId int    `json:"project_id"`
}

type EmailSend struct {
	TypeId       int               `json:"type_id"`
	Category     int               `json:"category"`
	ClientId     int               `json:"client_id"`
	ProjectId    int               `json:"project_id"`
	User         map[string]string `json:"user"`
	Meta         map[string]string `json:"meta"`
	ValueEncrypt ValueEncrypt      `json:"value_encrypt"`
}

type ValueEncrypt struct {
	TemplateData string `json:"template_data"`
}

type TypeIds struct {
	TypeIds []int `json:"type_ids"`
}

type OnlineByProjectAndEmailUpdating struct {
	Timestamp    time.Time `json:"timestamp"`
	ProjectId    int       `json:"project_id"`
	EncodedEmail string    `json:"encoded_email"`
}

type OnlineByUser struct {
	Timestamp time.Time `json:"timestamp"`
	UserId    int       `json:"user_id"`
}

type Payment struct {
	UserId      int   `json:"user_id"`
	StartDate   int64 `json:"start_date"`
	ExpireDate  int64 `json:"expire_date"`
	TotalCount  int   `json:"total_count"`
	PaymentType int   `json:"payment_type"`
	Amount      int   `json:"amount"`
}

type ForceConfirm struct {
	LastReaction int64  `json:"last_reaction"`
	ProjectId    int    `json:"project_id"`
	EncodedEmail string `json:"encoded_email"`
}

type WebpushSend struct {
	PushUserId int               `json:"push_user_id"`
	Title      string            `json:"title"`
	Url        string            `json:"url"`
	Icon       string            `json:"icon"`
	TypeId     int               `json:"type_id"`
	Meta       map[string]string `json:"meta"`
	Text       string            `json:"text"`
	ImageUrl   string            `json:"image_url"`
	ProjectId  int               `json:"project_id"`
}

type WebpushUserCreate struct {
	UserId int               `json:"user_id"`
	Meta   map[string]string `json:"meta"`
}

type ResponseEmailData struct {
	Data EmailData `json:"data"`
}

type ResponsePushData struct {
	Data PushData `json:"data"`
}

type EmailData struct {
	User EmailUser `json:"user"`
}

type PushData struct {
	PushUser PushUser `json:"result"`
}

type EmailUser struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	ProjectId int    `json:"project_id"`
	Name      string `json:"name"`
}

type PushUser struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	ProjectId int `json:"project_id"`
}

type Auth struct {
	ClientId string
	AuthKey  string
}
