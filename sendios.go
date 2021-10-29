package go_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/sendios/go-sdk/internal"
	"net/http"
	"time"
)

const (
	System = iota
	Trigger
)

const (
	UnknownPlatform = iota
	DesktopPlatform
	MobilePlatform
	AndroidPlatform
	IOSPlatform
)

const (
	SourceFbl      = 2
	SourceLink     = 4
	SourceClient   = 8
	SourceSettings = 9
)

const (
	ApiV3 = "http://127.0.0.1:8081/v3/"
	ApiV1 = "http://127.0.0.1:8081/v1/"
)

var m = map[int]string{System: "push/system", Trigger: "push/trigger"}

type SendiosSdk struct {
	Request *internal.Request
}

func NewSendiosSdk(clientId string, authKey string) *SendiosSdk {
	auth := &internal.Auth{ClientId: clientId, AuthKey: authKey}
	client := &http.Client{Timeout: time.Second * 10}

	r := &internal.Request{Client: client, Auth: auth}
	sdk := SendiosSdk{
		Request: r,
	}

	return &sdk
}

func (sdk *SendiosSdk) GetBuyingDecisions(email string) ([]byte, error) {
	params := internal.BuyingDecisionData{Email: email}

	return sdk.Request.Post(ApiV1, "buying/email", params)
}

func (sdk *SendiosSdk) CreateClientUser(email string, clientUserId string, projectId int) ([]byte, error) {
	params := internal.ClientUser{Email: email, ClientUserId: clientUserId, ProjectId: projectId}

	return sdk.Request.Post(ApiV1, "clientuser/create", params)
}

func (sdk *SendiosSdk) CheckEmail(email string, sanitize bool) ([]byte, error) {
	params := internal.CheckEmail{Email: email, Sanitize: sanitize}

	return sdk.Request.Post(ApiV1, "email/check", params)
}

func (sdk *SendiosSdk) ValidateEmail(email string, projectId int) ([]byte, error) {
	params := internal.ValidateEmail{Email: email, ProjectId: projectId}

	return sdk.Request.Post(ApiV1, "email/check/send", params)
}

func (sdk *SendiosSdk) TrackClickByMailId(mailId int) ([]byte, error) {

	return sdk.Request.Post(ApiV1, fmt.Sprintf("trackemail/click/%d", mailId), nil)
}

func (sdk *SendiosSdk) ProdEventSend(data interface{}) ([]byte, error) {

	return sdk.Request.Post(ApiV1, "product-event/create", data)
}

func (sdk *SendiosSdk) SendEmail(clientId int, typeId int, categoryId int, projectId int, email string, user map[string]string, data map[string]string, meta map[string]string) ([]byte, error) {
	user["email"] = email
	jsonString, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("error while json marshaling: %s", err)
	}

	encrypter, err := internal.MakeEncrypt()
	if err != nil {
		return nil, fmt.Errorf("error while encrypting: %s", err)
	}

	encrypt, err := encrypter.EncryptData(jsonString)
	if err != nil {
		return nil, fmt.Errorf("error data encrypting: %s", err)
	}

	params := internal.EmailSend{
		TypeId:       typeId,
		Category:     categoryId,
		ProjectId:    projectId,
		ClientId:     clientId,
		User:         user,
		Meta:         meta,
		ValueEncrypt: internal.ValueEncrypt{TemplateData: encrypt},
	}

	route, err := getRoute(categoryId)
	if err != nil {
		return nil, fmt.Errorf("error while getting route: %s", err)
	}

	return sdk.Request.Post(ApiV1, route, params)
}

func (sdk *SendiosSdk) GetUnsubListByEmailUserId(userId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsubtypes/%d", userId))
}

func (sdk *SendiosSdk) UnsubEmailUserByTypes(userId int, typeIds []int) ([]byte, error) {
	params := internal.TypeIds{TypeIds: typeIds}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("%s/%d", "unsubtypes", userId), params)
}

func (sdk *SendiosSdk) AddTypesToUnsubByEmailUser(userId int, typeIds []int) ([]byte, error) {
	params := internal.TypeIds{TypeIds: typeIds}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("%s/%d", "unsubtypes/nodiff", userId), params)
}

func (sdk *SendiosSdk) RemoveUnsubTypesByEmailUser(userId int, typeIds []int) ([]byte, error) {
	params := internal.TypeIds{TypeIds: typeIds}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsubtypes/nodiff/%d", userId), params)
}

func (sdk *SendiosSdk) RemoveAllUnsubTypesByEmailUser(userId int) ([]byte, error) {

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsubtypes/all/%d", userId), nil)
}

func (sdk *SendiosSdk) UnsubEmailUserClient(userId int) ([]byte, error) {

	return sdk.addEmailUserToUnsubList(userId, SourceClient)
}

func (sdk *SendiosSdk) UnsubEmailUserBySettings(userId int) ([]byte, error) {

	return sdk.addEmailUserToUnsubList(userId, SourceSettings)
}

func (sdk *SendiosSdk) UnsubEmailUserByAdmin(email string, projectId int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	return sdk.Request.Post(ApiV1, fmt.Sprintf("unsub/admin/%d/email/%s", projectId, encodedEmail), nil)
}

func (sdk *SendiosSdk) SubscribeEmailUser(userId int) ([]byte, error) {

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsub/%d", userId), nil)
}

func (sdk *SendiosSdk) IsUnsubUser(userId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/isunsub/%d", userId))
}

func (sdk *SendiosSdk) IsUnsubByEmailAndProjectId(email string, projectId int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectId(email, projectId)
	if err != nil {
		return nil, fmt.Errorf("can not get email user by project %d and email %s. Error: %s", projectId, email, err)
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while email user parsing: %s", err)
	}

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/isunsub/%d", user.Id))
}

func (sdk *SendiosSdk) GetUnsubscribeReason(email string, projectId int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectId(email, projectId)
	if err != nil {
		return nil, fmt.Errorf("can not get email user by project %d and email %s. Error: %s", projectId, email, err)
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while email user parsing: %s", err)
	}

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/unsubreason/%d", user.Id))
}

func (sdk *SendiosSdk) GetUnsubscribesByDate(time int64) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/list/%d", time))
}

func (sdk *SendiosSdk) GetEmailUserByEmailAndProjectId(email string, projectId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("user/project/%d/email/%s", projectId, email))
}

func (sdk *SendiosSdk) GetEmailUserById(id int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("user/id/%d", id))
}

func (sdk *SendiosSdk) SetUserFieldsByEmailAndProjectId(email string, projectId int, data map[string]string) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	return sdk.Request.Put(ApiV1, fmt.Sprintf("userfields/project/%d/emailhash/%s", projectId, encodedEmail), data)
}

func (sdk *SendiosSdk) SetUserFieldsByUserId(userId int, data map[string]string) ([]byte, error) {

	return sdk.Request.Put(ApiV1, fmt.Sprintf("userfields/user/%d", userId), data)
}

func (sdk *SendiosSdk) GetUserFieldsByEmailAndProjectId(email string, projectId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("userfields/project/%d/email/%s", projectId, email))
}

func (sdk *SendiosSdk) GetUserFieldsByUserId(userId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("userfields/user/%d", userId))
}

func (sdk *SendiosSdk) SetOnlineByEmailAndProjectId(email string, projectId int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)
	params := internal.OnlineByProjectAndEmailUpdating{
		ProjectId:    projectId,
		EncodedEmail: encodedEmail,
		Timestamp:    time.Now(),
	}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/project/%d/email/%s/online", projectId, encodedEmail), params)
}

func (sdk *SendiosSdk) SetOnlineByUser(userId int) ([]byte, error) {
	params := internal.OnlineByUser{UserId: userId, Timestamp: time.Now()}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/%d/online", userId), params)
}

func (sdk *SendiosSdk) AddPaymentByEmailAndProjectId(email string, projectId int, startDate, expireDate int64, totalCount, paymentType, amount int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectId(email, projectId)
	if err != nil {
		return nil, err
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.Payment{
		UserId:      user.Id,
		StartDate:   startDate,
		ExpireDate:  expireDate,
		TotalCount:  totalCount,
		PaymentType: paymentType,
		Amount:      amount,
	}

	return sdk.Request.Post(ApiV1, "lastpayment", params)
}

func (sdk *SendiosSdk) AddPaymentByUserId(userId int, startDate, expireDate int64, totalCount, paymentType, amount int) ([]byte, error) {
	params := internal.Payment{
		UserId:      userId,
		StartDate:   startDate,
		ExpireDate:  expireDate,
		TotalCount:  totalCount,
		PaymentType: paymentType,
		Amount:      amount,
	}

	return sdk.Request.Post(ApiV1, "lastpayment", params)
}

func (sdk *SendiosSdk) ForceConfirmByEmailAndProject(email string, projectId int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	params := internal.ForceConfirm{
		EncodedEmail: encodedEmail,
		ProjectId:    projectId,
		LastReaction: time.Now().Unix(),
	}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/project/%d/email/%s/confirm", projectId, encodedEmail), params)
}

func (sdk *SendiosSdk) UnsubscribePushUserByEmailUserId(userId int) ([]byte, error) {
	res, err := sdk.GetPushUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUser.Id), nil)
}

func (sdk *SendiosSdk) UnsubscribePushUserById(pushUserId int) ([]byte, error) {

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUserId), nil)
}

func (sdk *SendiosSdk) UnsubscribePushUserByProjectIdAndHash(projectId int, hash string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIdAndHash(projectId, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUser.Id), nil)
}

func (sdk *SendiosSdk) SubscribePushUserByEmailUserId(userId int) ([]byte, error) {
	res, err := sdk.GetPushUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("webpush/subscribe/%d", pushUser.Id), nil)
}

func (sdk *SendiosSdk) SubscribePushUserByProjectIdAndHash(projectId int, hash string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIdAndHash(projectId, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("webpush/subscribe/%d", pushUser.Id), nil)
}

func (sdk *SendiosSdk) SendPushByEmailUserId(userId int, title, text, url, iconUrl string, typeId int, meta map[string]string, imageUrl string) ([]byte, error) {
	res, err := sdk.GetPushUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.WebpushSend{
		PushUserId: pushUser.Id,
		Title:      title,
		Text:       text,
		Url:        url,
		Icon:       iconUrl,
		TypeId:     typeId,
		Meta:       meta,
		ImageUrl:   imageUrl,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)
}

func (sdk *SendiosSdk) SendPushByProjectIdAndHash(projectId int, hash, title, text, url, iconUrl string, typeId int, meta map[string]string, imageUrl string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIdAndHash(projectId, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.WebpushSend{
		PushUserId: pushUser.Id,
		Title:      title,
		Text:       text,
		Icon:       iconUrl,
		TypeId:     typeId,
		Meta:       meta,
		ImageUrl:   imageUrl,
		Url:        url,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)

}

func (sdk *SendiosSdk) SendPushByProject(projectId int, title, text, url, iconUrl string, typeId int, meta map[string]string, imageUrl string) ([]byte, error) {
	params := internal.WebpushSend{
		Title:     title,
		Text:      text,
		Icon:      iconUrl,
		TypeId:    typeId,
		Meta:      meta,
		ImageUrl:  imageUrl,
		ProjectId: projectId,
		Url:       url,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)
}

func (sdk *SendiosSdk) CreatePushUser(userId, projectId int, url, publicKey, authToken string) ([]byte, error) {
	meta := map[string]string{"url": url, "public_key": publicKey, "auth_token": authToken}
	params := internal.WebpushUserCreate{
		UserId: userId,
		Meta:   meta,
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/project/%d", projectId), params)
}

func (sdk *SendiosSdk) GetPushUserById(userId int) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("webpush/user/get/%d", userId))
}

func (sdk *SendiosSdk) GetPushUserByProjectIdAndHash(projectId int, hash string) ([]byte, error) {

	return sdk.Request.Get(ApiV1, fmt.Sprintf("webpush/project/get/%d/hash/%s", projectId, hash))
}

func (sdk *SendiosSdk) addEmailUserToUnsubList(userId int, sourceId int) ([]byte, error) {

	return sdk.Request.Post(ApiV1, fmt.Sprintf("unsub/%d/source/%d", userId, sourceId), nil)
}

func getRoute(categoryId int) (string, error) {
	elem, ok := m[categoryId]

	if ok != true {
		return "", fmt.Errorf("not found route by category %d", categoryId)
	}

	return elem, nil
}

func parseUserFromResponseData(res []byte) (internal.EmailUser, error) {
	var responseData *internal.ResponseEmailData
	err := json.Unmarshal(res, &responseData)

	if err != nil {
		return internal.EmailUser{}, fmt.Errorf("error while unmarshling email user data: %s", err)
	}

	return responseData.Data.User, nil
}

func parsePushUserFromResponseData(res []byte) (internal.PushUser, error) {
	var responseData *internal.ResponsePushData
	err := json.Unmarshal(res, &responseData)

	if err != nil {
		return internal.PushUser{}, fmt.Errorf("error while unmarshling push user data: %s", err)
	}

	return responseData.Data.PushUser, nil
}
