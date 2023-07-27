package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	sendios "github.com/sendios/go-sdk"
	"github.com/sendios/go-sdk/internal"
)

func TestNewSendiosSdk(t *testing.T) {
	type args struct {
		clientID   string
		authKey    string
		encryptKey string
	}
	tests := []struct {
		name string
		args args
		want *sendios.Client
	}{
		{
			name: "new_sendios_object",
			args: args{clientID: "3", authKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6", encryptKey: "asldfajskdfha"},
			want: &sendios.Client{
				Request:    &internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}},
				EncryptKey: []byte("asldfajskdfha"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sendios.NewClient(tt.args.clientID, tt.args.authKey, tt.args.encryptKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddPaymentByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email       string
		projectID   int
		startDate   int64
		expireDate  int64
		totalCount  int
		paymentType int
		amount      int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"add_payment_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 1, 1625479419, 1625479419, 1, 1, 1},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3701},"data":{"date":"2021-07-05 13:03:39.000000","message":"done","status":true}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, `{"_meta":{"count":3,"status":"SUCCESS","time":3701},"data":{"date":"2021-07-05 13:03:39.000000","message":"done","status":true}}`)
			}))
			defer ts.Close()

			params := internal.Payment{
				UserID:      1,
				StartDate:   tt.args.startDate,
				ExpireDate:  tt.args.expireDate,
				TotalCount:  tt.args.totalCount,
				PaymentType: tt.args.paymentType,
				Amount:      tt.args.amount,
			}

			got, err := sdk.Request.Post(ts.URL, "/lastpayment", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPaymentByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddPaymentByUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID      int
		startDate   int64
		expireDate  int64
		totalCount  int
		paymentType int
		amount      int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"add_payment_by_user_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, 1625479419, 1625479419, 1, 1, 1},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3964},"data":{"date":"2021-07-05 13:57:57.000000","message":"done","status":true}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				_, _ = fmt.Fprintln(w, `{"_meta":{"count":3,"status":"SUCCESS","time":3964},"data":{"date":"2021-07-05 13:57:57.000000","message":"done","status":true}}`)
			}))
			defer ts.Close()

			params := internal.Payment{
				UserID:      tt.args.userID,
				StartDate:   tt.args.startDate,
				ExpireDate:  tt.args.expireDate,
				TotalCount:  tt.args.totalCount,
				PaymentType: tt.args.paymentType,
				Amount:      tt.args.amount,
			}

			got, err := sdk.Request.Post(ts.URL, "/lastpayment", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPaymentByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddTypesToUnsubByEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID  int
		typeIDs []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"add_types_to_unsub_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, []int{1, 2, 3, 4, 5}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3665},"data":null}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, `{"_meta":{"count":1,"status":"SUCCESS","time":3665},"data":null}`)
			}))
			defer ts.Close()

			params := internal.TypeIDs{TypeIDs: tt.args.typeIDs}
			got, err := sdk.Request.Post(ts.URL, "/unsubtypes/nodiff", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTypesToUnsubByEmailUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_CheckEmail(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email    string
		sanitize bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"check_email_invalid",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", true},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3499},"data":{"domain":"gmail.com","email":"test@gmail.com","orig":"test@gmail.com","reason":"system","trusted":true,"valid":false,"vendor":"Google"}}`),
			true,
		},

		{
			"check_email_valid",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", true},
			[]byte(`{"_meta":{"count":6,"status":"SUCCESS","time":4449},"data":{"domain":"corp.sendios.io","email":"volodymyr.voloshyn@corp.sendios.io","orig":"volodymyr.voloshyn@corp.sendios.io","trusted":true,"valid":true,"vendor":"Unknown"}}`),
			true,
		},

		{
			"check_email_invalid_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", false},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3962},"data":{"domain":"gmail.com","email":"test@gmail.com","orig":"test@gmail.com","reason":"system","trusted":true,"valid":false,"vendor":"Google"}}`),
			true,
		},

		{
			"check_email_valid_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", false},
			[]byte(`{"_meta":{"count":6,"status":"SUCCESS","time":4941},"data":{"domain":"corp.sendios.io","email":"volodymyr.voloshyn@corp.sendios.io","orig":"volodymyr.voloshyn@corp.sendios.io","trusted":true,"valid":true,"vendor":"Unknown"}}`),
			true,
		},

		{
			"check_email_valid_untrusted_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test.com", false},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3741},"data":{"domain":"test.com","email":"test.com","orig":"test.com","reason":"invalid","trusted":false,"valid":false,"vendor":"Unknown"}}`),
			true,
		},

		{
			"check_email_valid_untrusted_sanitize_true",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test.com", true},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":4070},"data":{"domain":"test.com","email":"test.com@test.com","orig":"test.com","reason":"mx_record","trusted":false,"valid":false,"vendor":"Unknown"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.CheckEmail{Email: tt.args.email, Sanitize: tt.args.sanitize}
			got, err := sdk.Request.Post(ts.URL, "/email/check", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_CreateClientUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email        string
		clientUserID string
		projectID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"create_client_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", "1", 2},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":4301},"data":{"date":"2021-07-05 15:45:39.000000","message":"done","status":true}}`),
			true,
		},

		{
			"create_client_user_already_exist",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", "1", 2},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3414},"data":{"date":"2021-07-05 15:51:09.000000","message":"done","status":true}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.ClientUser{Email: tt.args.email, ClientUserID: tt.args.clientUserID, ProjectID: tt.args.projectID}
			got, err := sdk.Request.Post(ts.URL, "/clientuser/create", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateClientUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// trouble with origin header
//func TestSendiosSdk_CreatePushUser(t *testing.T) {
//	type fields struct {
//		Request *internal.Request
//	}
//	type args struct {
//		userID    int
//		projectID int
//		url       string
//		publicKey string
//		authToken string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			sdk := &sendios.Client{
//				Request: tt.fields.Request,
//			}
//			got, err := sdk.CreatePushUser(tt.args.userID, tt.args.projectID, tt.args.url, tt.args.publicKey, tt.args.authToken)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreatePushUser() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreatePushUser() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestSendiosSdk_ForceConfirmByEmailAndProject(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"force_confirm_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 2},
			[]byte(`{"status":"Accepted"}`),
			true,
		},

		{
			"force_confirm_by_invalid_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"testgmail.com", 2},
			[]byte(`{"status":"Accepted"}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			encodedEmail := internal.Base64Encoder(tt.args.email)

			params := internal.ForceConfirm{
				EncodedEmail: encodedEmail,
				ProjectID:    tt.args.projectID,
				LastReaction: time.Now().Unix(),
			}

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/project/%d/email/%s/confirm", tt.args.projectID, encodedEmail), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceConfirmByEmailAndProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetBuyingDecisions(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"buying_decision",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com"},
			[]byte(`{"_meta":{"count":2,"status":"SUCCESS","time":3923},"data":{"decision":false,"email":"test@gmail.com"}}`),
			true,
		},

		{
			"buying_decision_validation_error",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"testgmail.com"},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3407},"data":{"error":"Request validation error:  email - This value is not a valid email address."}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.BuyingDecisionData{Email: tt.args.email}
			got, err := sdk.Request.Post(ts.URL, "/buying/email", params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBuyingDecisions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetEmailUserByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"email_user_by_email_and_project",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", 2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3668},"data":{"user":{"activation":null,"channel_id":null,"clicks":0,"country":false,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":null,"last_online":null,"last_payment":{"active":1,"amount":1,"expires_at":1625479419,"id":1,"payment_count":1,"payment_type":1,"project_id":2,"started_at":1625479419,"user_id":5005},"last_reaction":null,"last_request":null,"meta":{"profile":{"age":null,"ak":null,"partner_id":null,"photo":null}},"name":"Volodymyr","project_id":2,"project_title":"Test project 1","sends":0,"sent_mails":[],"subchannel_id":null,"unsub_promo":[],"unsubscribe":[],"unsubscribe_types":[{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":1,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":2,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":3,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":4,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":5,"type_sig":"Undefined"}],"webpush":{"last_click":null,"last_push":"2020-10-26 14:47:25","reg_date":"2017-02-13 16:55:25"}}}}`),
			true,
		},

		{
			"email_user_by_email_and_project_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3600},"data":{"error":"Not found  user for project 2 and email test@gmail.com"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/user/project/%d/email/%s", tt.args.projectID, tt.args.email))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmailUserByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetEmailUserByID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"email_user_by_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4507},"data":{"user":{"activation":null,"channel_id":null,"clicks":0,"country":false,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":null,"last_online":null,"last_payment":{"active":1,"amount":1,"expires_at":1625479419,"id":1,"payment_count":1,"payment_type":1,"project_id":2,"started_at":1625479419,"user_id":5005},"last_reaction":null,"last_request":null,"meta":{"profile":{"age":null,"ak":null,"partner_id":null,"photo":null}},"name":"Volodymyr","project_id":2,"project_title":"Test project 1","sends":0,"sent_mails":[],"subchannel_id":null,"unsub_promo":[],"unsubscribe":[],"unsubscribe_types":[{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":1,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":2,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":3,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":4,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":5,"type_sig":"Undefined"}],"webpush":{"last_click":null,"last_push":"2020-10-26 14:47:25","reg_date":"2017-02-13 16:55:25"}}}}`),
			true,
		},

		{
			"email_user_by_id_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4477},"data":{"error":"User not found"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/user/id/%d", tt.args.id))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmailUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetPushUserByID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"push_user_by_user_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3630},"data":{"result":{"hash":"NULL","id":717067,"invalid":null,"last_click":null,"last_online":null,"last_push":"2020-10-26 14:47:25","last_show":null,"meta":{"auth_token":"FpOCjghkdFV1JEUKOgJ-cw==","public_key":"BHv8Z_EomoVe33d2oGdwdbmT9Jd3rJd4VPZRXjzNPgtJzeHOu-hERyap53cV74nKVPQQ7BlNVwAKMGZaZsSr16A=","url":"https://android.googleapis.com/gcm/send/dtnHwCXnsBw:APA91bFPw56tcPNcVdWdOCzqhIDnPf4pyBIaTXg_tjMDDlLHlk4zo5BwpOTYJGbG_ZpMCuBZg_R3LIJaMalUHCQExEWD7__CNpYRcR-UHhkoHa_r9p3sD00kYr9qy9iBCztSTGWgZHzy"},"platform_id":null,"project_id":2,"reg_date":"2017-02-13 16:55:25","send_platform_id":null,"type":null,"user_id":5005}}}`),
			true,
		},

		{
			"push_user_by_user_id_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3464},"data":{"error":"Not found user by id: 0"}}`),
			true,
		},

		{
			"push_user_by_user_id_forbidden",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4494},"data":{"error":"Forbidden"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/webpush/user/get/%d", tt.args.userID))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPushUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUnsubListByEmailUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"unsub_list_by_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":5,"status":"SUCCESS","time":4030},"data":[{"created_at":"2021-07-05 15:06:24","name":"SystemMail07082","type_id":1},{"created_at":"2021-07-05 15:06:24","name":"Test2","type_id":2},{"created_at":"2021-07-05 15:06:24","name":"Qwerty","type_id":3},{"created_at":"2021-07-05 15:06:24","name":"Foobar","type_id":4},{"created_at":"2021-07-05 15:06:24","name":"TriggerMail04082","type_id":5}]}`),
			true,
		},

		{
			"unsub_list_by_email_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4240},"data":{"error":"User not found"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsubtypes/%d", tt.args.userID))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnsubListByEmailUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUnsubscribeReason(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		UserID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"unsub_reason_by_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3537},"data":{"error":"User not found"}}`),
			true,
		},
		{
			"unsub_reason_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3421},"data":{"result":false}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/unsubreason/%d", tt.args.UserID))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnsubscribeReason() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUnsubscribesByDate(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		time int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"unsub_by_date",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{time.Now().Unix()},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4477},"data":{"error":"User not found"}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/list/%d", tt.args.time))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnsubscribesByDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUserFieldsByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"user_fields_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", 2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":5623},"data":{"result":{"custom_fields":[],"user":{"city_id":null,"confirm":1,"country_id":null,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":0,"last_online":0,"last_reaction":0,"list_id":0,"meta":"[]","name":"Volodymyr","platform_id":1,"project_id":2,"status":1,"valid_id":null,"vendor_id":3,"vip":1}}}}`),
			true,
		},

		{
			"user_fields_by_email_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"example@gmail.com", 2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3726},"data":{"error":"User not found."}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/userfields/project/%d/email/%s", tt.args.projectID, tt.args.email))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFieldsByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUserFieldsByUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.GetUserFieldsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFieldsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFieldsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_IsUnsubByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"is_unsub_by_email_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3758},"data":{"error":"User not found"}}`),
			true,
		},
		{
			"is_unsub_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3387},"data":{"result":false}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/isunsub/%d", tt.args.userID))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsUnsubByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_IsUnsubUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"is_unsub_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3758},"data":{"error":"User not found"}}`),
			true,
		},
		{
			"is_unsub",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3387},"data":{"result":false}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/isunsub/%d", tt.args.userID))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsUnsubUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_RemoveAllUnsubTypesByEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"remove_all_unsub_types",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4969},"data":null}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsubtypes/all/%d", tt.args.userID), nil)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAllUnsubTypesByEmailUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_RemoveUnsubTypesByEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID  int
		typeIDs []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"remove_unsub_types_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, []int{1, 2, 3}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4275},"data":null}`),
			true,
		},

		{
			"remove_unsub_types_email_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0, []int{1, 2, 3}},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3615},"data":{"error":"User not found"}}`),
			true,
		},

		{
			name:    "remove_unsub_types_email_user_empty_types",
			fields:  fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args:    args{1, []int{}},
			want:    []byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3460},"data":null}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.TypeIDs{TypeIDs: tt.args.typeIDs}

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsubtypes/nodiff/%d", tt.args.userID), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveUnsubTypesByEmailUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// later
func TestSendiosSdk_SendEmail(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		clientID   int
		typeID     int
		categoryID int
		projectID  int
		email      string
		user       map[string]string
		data       map[string]string
		meta       map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"send_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{3, 1, 1, 2, "test@gmail.com", map[string]string{"id": "1"}, map[string]string{"data": "test"}, map[string]string{"meta": "email"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4275},"data":null}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, "/push/system")
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsUnsubByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// later
func TestSendiosSdk_SendPushByEmailUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID   int
		title    string
		text     string
		url      string
		iconUrl  string
		typeID   int
		meta     map[string]string
		imageUrl string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByEmailUserID(tt.args.userID, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeID, tt.args.meta, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPushByEmailUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendPushByEmailUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// later
func TestSendiosSdk_SendPushByProject(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectID int
		title     string
		text      string
		url       string
		iconUrl   string
		typeID    int
		meta      map[string]string
		imageUrl  string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByProject(tt.args.projectID, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeID, tt.args.meta, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPushByProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendPushByProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// later
func TestSendiosSdk_SendPushByProjectIDAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectID int
		hash      string
		title     string
		text      string
		url       string
		iconUrl   string
		typeID    int
		meta      map[string]string
		imageUrl  string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByProjectIDAndHash(tt.args.projectID, tt.args.hash, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeID, tt.args.meta, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPushByProjectIDAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendPushByProjectIDAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetOnlineByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"set_online_by_email_and_project",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendio.io", 2},
			[]byte(`{"status":"Accepted"}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			encodedEmail := internal.Base64Encoder(tt.args.email)
			params := internal.OnlineByProjectAndEmailUpdating{
				ProjectID:    tt.args.projectID,
				EncodedEmail: encodedEmail,
				Timestamp:    time.Now(),
			}
			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/project/%d/email/%s/online", tt.args.projectID, encodedEmail), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlineByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetOnlineByUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"set_online_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"status":"Accepted"}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.OnlineByUser{UserID: tt.args.userID, Timestamp: time.Now()}

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/%d/online", tt.args.userID), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlineByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetUserFieldsByEmailAndProjectID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
		data      map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"set_user_fields_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendio.io", 2, map[string]string{"test": "test"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3850},"data":{"result":true}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			encodedEmail := internal.Base64Encoder(tt.args.email)
			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/userfields/project/%d/emailhash/%s", tt.args.projectID, encodedEmail), tt.args.data)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserFieldsByEmailAndProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetUserFieldsByUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
		data   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"set_user_fields_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2, map[string]string{"test": "test"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4112},"data":{"result":true}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/userfields/user/%d", tt.args.userID), tt.args.data)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserFieldsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribeEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"subscribe_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3716},"data":{"subscribe":{"rowCount":1}}}`),
			true,
		},
		{
			"subscribe_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{ClientID: "3", AuthKey: "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3616},"data":{"subscribe":{"rowCount":0}}}`),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsub/%d", tt.args.userID), nil)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribeEmailUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribePushUserByEmailUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.SubscribePushUserByEmailUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribePushUserByEmailUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribePushUserByEmailUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribePushUserByProjectIDAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectID int
		hash      string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.SubscribePushUserByProjectIDAndHash(tt.args.projectID, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribePushUserByProjectIDAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribePushUserByProjectIDAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_TrackClickByMailID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		mailID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.TrackClickByMailID(tt.args.mailID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrackClickByMailID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackClickByMailID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubEmailUserByAdmin(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserByAdmin(tt.args.email, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubEmailUserByAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubEmailUserByAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubEmailUserBySettings(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserBySettings(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubEmailUserBySettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubEmailUserBySettings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubEmailUserByTypes(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID  int
		typeIDs []int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserByTypes(tt.args.userID, tt.args.typeIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubEmailUserByTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubEmailUserByTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubEmailUserClient(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserClient(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubEmailUserClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubEmailUserClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubscribePushUserByEmailUserID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserByEmailUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserByEmailUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserByEmailUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubscribePushUserByID(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		pushUserID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserByID(tt.args.pushUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubscribePushUserByProjectIDAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectID int
		hash      string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserByProjectIDAndHash(tt.args.projectID, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserByProjectIDAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserByProjectIDAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_ValidateEmail(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectID int
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.Client{
				Request: tt.fields.Request,
			}
			got, err := sdk.ValidateEmail(tt.args.email, tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
