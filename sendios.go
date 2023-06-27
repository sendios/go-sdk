package sendios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sendios/go-sdk/internal"
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
	ApiV3 = "https://api.sendios.io/v3/"
	ApiV1 = "https://api.sendios.io/v1/"
)

var m = map[int]string{System: "push/system", Trigger: "push/trigger"}

type Client struct {
	Request    *internal.Request
	EncryptKey []byte
}

func NewClient(clientID, authKey, encryptKey string) *Client {
	auth := &internal.Auth{ClientID: clientID, AuthKey: authKey}
	client := &http.Client{Timeout: time.Second * 10}

	r := &internal.Request{Client: client, Auth: auth}
	sdk := Client{
		Request:    r,
		EncryptKey: []byte(encryptKey),
	}

	return &sdk
}

func (sdk *Client) GetBuyingDecisions(email string) ([]byte, error) {
	params := internal.BuyingDecisionData{Email: email}

	return sdk.Request.Post(ApiV1, "buying/email", params)
}

func (sdk *Client) CreateClientUser(email, clientUserID string, projectID int) ([]byte, error) {
	params := internal.ClientUser{Email: email, ClientUserID: clientUserID, ProjectID: projectID}

	return sdk.Request.Post(ApiV1, "clientuser/create", params)
}

func (sdk *Client) CheckEmail(email string, sanitize bool) ([]byte, error) {
	params := internal.CheckEmail{Email: email, Sanitize: sanitize}

	return sdk.Request.Post(ApiV1, "email/check", params)
}

func (sdk *Client) ValidateEmail(email string, projectID int) ([]byte, error) {
	params := internal.ValidateEmail{Email: email, ProjectID: projectID}

	return sdk.Request.Post(ApiV1, "email/check/send", params)
}

func (sdk *Client) TrackClickByMailID(mailID int) ([]byte, error) {
	return sdk.Request.Post(ApiV1, fmt.Sprintf("trackemail/click/%d", mailID), nil)
}

func (sdk *Client) ProdEventSend(data interface{}) ([]byte, error) {
	return sdk.Request.Post(ApiV1, "product-event/create", data)
}

func (sdk *Client) SendEmail(clientID, typeID, categoryID, projectID int, email string, user, data, meta map[string]string) ([]byte, error) {
	user["email"] = email
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error while json marshaling: %s", err)
	}

	encrypter, err := internal.MakeEncrypt(sdk.EncryptKey)
	if err != nil {
		return nil, fmt.Errorf("error while encrypting: %s", err)
	}

	encrypt, err := encrypter.EncryptData(jsonString)
	if err != nil {
		return nil, fmt.Errorf("error data encrypting: %s", err)
	}

	params := internal.EmailSend{
		TypeID:       typeID,
		Category:     categoryID,
		ProjectID:    projectID,
		ClientID:     clientID,
		User:         user,
		Meta:         meta,
		ValueEncrypt: internal.ValueEncrypt{TemplateData: encrypt},
	}

	route, err := getRoute(categoryID)
	if err != nil {
		return nil, fmt.Errorf("error while getting route: %s", err)
	}

	return sdk.Request.Post(ApiV1, route, params)
}

func (sdk *Client) GetUnsubListByEmailUserID(userID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsubtypes/%d", userID))
}

func (sdk *Client) UnsubEmailUserByTypes(userID int, typeIDs []int) ([]byte, error) {
	params := internal.TypeIDs{TypeIDs: typeIDs}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("%s/%d", "unsubtypes", userID), params)
}

func (sdk *Client) AddTypesToUnsubByEmailUser(userID int, typeIDs []int) ([]byte, error) {
	params := internal.TypeIDs{TypeIDs: typeIDs}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("%s/%d", "unsubtypes/nodiff", userID), params)
}

func (sdk *Client) RemoveUnsubTypesByEmailUser(userID int, typeIDs []int) ([]byte, error) {
	params := internal.TypeIDs{TypeIDs: typeIDs}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsubtypes/nodiff/%d", userID), params)
}

func (sdk *Client) RemoveAllUnsubTypesByEmailUser(userID int) ([]byte, error) {
	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsubtypes/all/%d", userID), nil)
}

func (sdk *Client) UnsubEmailUserClient(userID int) ([]byte, error) {
	return sdk.addEmailUserToUnsubList(userID, SourceClient)
}

func (sdk *Client) UnsubEmailUserBySettings(userID int) ([]byte, error) {
	return sdk.addEmailUserToUnsubList(userID, SourceSettings)
}

func (sdk *Client) UnsubEmailUserByAdmin(email string, projectID int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	return sdk.Request.Post(ApiV1, fmt.Sprintf("unsub/admin/%d/email/%s", projectID, encodedEmail), nil)
}

func (sdk *Client) SubscribeEmailUser(userID int) ([]byte, error) {
	return sdk.Request.Delete(ApiV1, fmt.Sprintf("unsub/%d", userID), nil)
}

func (sdk *Client) IsUnsubUser(userID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/isunsub/%d", userID))
}

func (sdk *Client) IsUnsubByEmailAndProjectID(email string, projectID int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectID(email, projectID)
	if err != nil {
		return nil, fmt.Errorf("can not get email user by project %d and email %s. Error: %s", projectID, email, err)
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while email user parsing: %s", err)
	}

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/isunsub/%d", user.ID))
}

func (sdk *Client) GetUnsubscribeReason(email string, projectID int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectID(email, projectID)
	if err != nil {
		return nil, fmt.Errorf("can not get email user by project %d and email %s. Error: %s", projectID, email, err)
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while email user parsing: %s", err)
	}

	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/unsubreason/%d", user.ID))
}

func (sdk *Client) GetUnsubscribesByDate(time int64) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("unsub/list/%d", time))
}

func (sdk *Client) GetEmailUserByEmailAndProjectID(email string, projectID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("user/project/%d/email/%s", projectID, email))
}

func (sdk *Client) GetEmailUserByID(id int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("user/id/%d", id))
}

func (sdk *Client) SetUserFieldsByEmailAndProjectID(email string, projectID int, data map[string]string) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	return sdk.Request.Put(ApiV1, fmt.Sprintf("userfields/project/%d/emailhash/%s", projectID, encodedEmail), data)
}

func (sdk *Client) SetUserFieldsByUserID(userID int, data map[string]string) ([]byte, error) {
	return sdk.Request.Put(ApiV1, fmt.Sprintf("userfields/user/%d", userID), data)
}

func (sdk *Client) GetUserFieldsByEmailAndProjectID(email string, projectID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("userfields/project/%d/email/%s", projectID, email))
}

func (sdk *Client) GetUserFieldsByUserID(userID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("userfields/user/%d", userID))
}

func (sdk *Client) SetOnlineByEmailAndProjectID(email string, projectID int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)
	params := internal.OnlineByProjectAndEmailUpdating{
		ProjectID:    projectID,
		EncodedEmail: encodedEmail,
		Timestamp:    time.Now(),
	}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/project/%d/email/%s/online", projectID, encodedEmail), params)
}

func (sdk *Client) SetOnlineByUser(userID int) ([]byte, error) {
	params := internal.OnlineByUser{UserID: userID, Timestamp: time.Now()}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/%d/online", userID), params)
}

func (sdk *Client) AddPaymentByEmailAndProjectID(email string, projectID int, startDate, expireDate int64, totalCount, paymentType, amount int) ([]byte, error) {
	res, err := sdk.GetEmailUserByEmailAndProjectID(email, projectID)
	if err != nil {
		return nil, err
	}

	user, err := parseUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.Payment{
		UserID:      user.ID,
		StartDate:   startDate,
		ExpireDate:  expireDate,
		TotalCount:  totalCount,
		PaymentType: paymentType,
		Amount:      amount,
	}

	return sdk.Request.Post(ApiV1, "lastpayment", params)
}

func (sdk *Client) AddPaymentByUserID(userID int, startDate, expireDate int64, totalCount, paymentType, amount int) ([]byte, error) {
	params := internal.Payment{
		UserID:      userID,
		StartDate:   startDate,
		ExpireDate:  expireDate,
		TotalCount:  totalCount,
		PaymentType: paymentType,
		Amount:      amount,
	}

	return sdk.Request.Post(ApiV1, "lastpayment", params)
}

func (sdk *Client) ForceConfirmByEmailAndProject(email string, projectID int) ([]byte, error) {
	encodedEmail := internal.Base64Encoder(email)

	params := internal.ForceConfirm{
		EncodedEmail: encodedEmail,
		ProjectID:    projectID,
		LastReaction: time.Now().Unix(),
	}

	return sdk.Request.Put(ApiV3, fmt.Sprintf("users/project/%d/email/%s/confirm", projectID, encodedEmail), params)
}

func (sdk *Client) UnsubscribePushUserByEmailUserID(userID int) ([]byte, error) {
	res, err := sdk.GetPushUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUser.ID), nil)
}

func (sdk *Client) UnsubscribePushUserByID(pushUserID int) ([]byte, error) {
	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUserID), nil)
}

func (sdk *Client) UnsubscribePushUserByProjectIDAndHash(projectID int, hash string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIDAndHash(projectID, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/unsubscribe/%d", pushUser.ID), nil)
}

func (sdk *Client) SubscribePushUserByEmailUserID(userID int) ([]byte, error) {
	res, err := sdk.GetPushUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("webpush/subscribe/%d", pushUser.ID), nil)
}

func (sdk *Client) SubscribePushUserByProjectIDAndHash(projectID int, hash string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIDAndHash(projectID, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	return sdk.Request.Delete(ApiV1, fmt.Sprintf("webpush/subscribe/%d", pushUser.ID), nil)
}

func (sdk *Client) SendPushByEmailUserID(userID int, title, text, url, iconUrl string, typeID int, meta map[string]string, imageUrl string) ([]byte, error) {
	res, err := sdk.GetPushUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.WebpushSend{
		PushUserID: pushUser.ID,
		Title:      title,
		Text:       text,
		Url:        url,
		Icon:       iconUrl,
		TypeID:     typeID,
		Meta:       meta,
		ImageUrl:   imageUrl,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)
}

func (sdk *Client) SendPushByProjectIDAndHash(projectID int, hash, title, text, url, iconUrl string, typeID int, meta map[string]string, imageUrl string) ([]byte, error) {
	res, err := sdk.GetPushUserByProjectIDAndHash(projectID, hash)
	if err != nil {
		return nil, fmt.Errorf("error while getting push user: %s", err)
	}

	pushUser, err := parsePushUserFromResponseData(res)
	if err != nil {
		return nil, fmt.Errorf("error while parsing push user: %s", err)
	}

	params := internal.WebpushSend{
		PushUserID: pushUser.ID,
		Title:      title,
		Text:       text,
		Icon:       iconUrl,
		TypeID:     typeID,
		Meta:       meta,
		ImageUrl:   imageUrl,
		Url:        url,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)
}

func (sdk *Client) SendPushByProject(projectID int, title, text, url, iconUrl string, typeID int, meta map[string]string, imageUrl string) ([]byte, error) {
	params := internal.WebpushSend{
		Title:     title,
		Text:      text,
		Icon:      iconUrl,
		TypeID:    typeID,
		Meta:      meta,
		ImageUrl:  imageUrl,
		ProjectID: projectID,
		Url:       url,
	}

	return sdk.Request.Post(ApiV1, "webpush/send", params)
}

func (sdk *Client) CreatePushUser(userID, projectID int, url, publicKey, authToken string) ([]byte, error) {
	meta := map[string]string{"url": url, "public_key": publicKey, "auth_token": authToken}
	params := internal.WebpushUserCreate{
		UserID: userID,
		Meta:   meta,
	}

	return sdk.Request.Post(ApiV1, fmt.Sprintf("webpush/project/%d", projectID), params)
}

func (sdk *Client) GetPushUserByID(userID int) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("webpush/user/get/%d", userID))
}

func (sdk *Client) GetPushUserByProjectIDAndHash(projectID int, hash string) ([]byte, error) {
	return sdk.Request.Get(ApiV1, fmt.Sprintf("webpush/project/get/%d/hash/%s", projectID, hash))
}

func (sdk *Client) addEmailUserToUnsubList(userID, sourceID int) ([]byte, error) {
	return sdk.Request.Post(ApiV1, fmt.Sprintf("unsub/%d/source/%d", userID, sourceID), nil)
}

func getRoute(categoryID int) (string, error) {
	elem, ok := m[categoryID]

	if !ok {
		return "", fmt.Errorf("not found route by category %d", categoryID)
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
