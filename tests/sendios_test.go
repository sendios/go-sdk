package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sendios"
	"sendios/internal"
	"testing"
	"time"
)

func TestNewSendiosSdk(t *testing.T) {
	type args struct {
		clientId string
		authKey  string
	}
	tests := []struct {
		name string
		args args
		want *sendios.SendiosSdk
	}{
		{
			"new_sendios_object",
			args{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"},
			&sendios.SendiosSdk{
				&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sendios.NewSendiosSdk(tt.args.clientId, tt.args.authKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSendiosSdk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddPaymentByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email       string
		projectId   int
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
		{"add_payment_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 1, 1625479419, 1625479419, 1, 1, 1},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3701},"data":{"date":"2021-07-05 13:03:39.000000","message":"done","status":true}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, `{"_meta":{"count":3,"status":"SUCCESS","time":3701},"data":{"date":"2021-07-05 13:03:39.000000","message":"done","status":true}}`)
			}))
			defer ts.Close()

			params := internal.Payment{
				UserId:      1,
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
				t.Errorf("AddPaymentByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddPaymentByUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId      int
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
		{"add_payment_by_user_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, 1625479419, 1625479419, 1, 1, 1},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3964},"data":{"date":"2021-07-05 13:57:57.000000","message":"done","status":true}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				_, _ = fmt.Fprintln(w, `{"_meta":{"count":3,"status":"SUCCESS","time":3964},"data":{"date":"2021-07-05 13:57:57.000000","message":"done","status":true}}`)
			}))
			defer ts.Close()

			params := internal.Payment{
				UserId:      tt.args.userId,
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
				t.Errorf("AddPaymentByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_AddTypesToUnsubByEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId  int
		typeIds []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"add_types_to_unsub_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, []int{1, 2, 3, 4, 5}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3665},"data":null}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, `{"_meta":{"count":1,"status":"SUCCESS","time":3665},"data":null}`)
			}))
			defer ts.Close()

			params := internal.TypeIds{TypeIds: tt.args.typeIds}
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

		{"check_email_invalid",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", true},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3499},"data":{"domain":"gmail.com","email":"test@gmail.com","orig":"test@gmail.com","reason":"system","trusted":true,"valid":false,"vendor":"Google"}}`),
			true},

		{"check_email_valid",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", true},
			[]byte(`{"_meta":{"count":6,"status":"SUCCESS","time":4449},"data":{"domain":"corp.sendios.io","email":"volodymyr.voloshyn@corp.sendios.io","orig":"volodymyr.voloshyn@corp.sendios.io","trusted":true,"valid":true,"vendor":"Unknown"}}`),
			true},

		{"check_email_invalid_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", false},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3962},"data":{"domain":"gmail.com","email":"test@gmail.com","orig":"test@gmail.com","reason":"system","trusted":true,"valid":false,"vendor":"Google"}}`),
			true},

		{"check_email_valid_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", false},
			[]byte(`{"_meta":{"count":6,"status":"SUCCESS","time":4941},"data":{"domain":"corp.sendios.io","email":"volodymyr.voloshyn@corp.sendios.io","orig":"volodymyr.voloshyn@corp.sendios.io","trusted":true,"valid":true,"vendor":"Unknown"}}`),
			true},

		{"check_email_valid_untrusted_sanitize_false",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test.com", false},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":3741},"data":{"domain":"test.com","email":"test.com","orig":"test.com","reason":"invalid","trusted":false,"valid":false,"vendor":"Unknown"}}`),
			true},

		{"check_email_valid_untrusted_sanitize_true",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test.com", true},
			[]byte(`{"_meta":{"count":7,"status":"SUCCESS","time":4070},"data":{"domain":"test.com","email":"test.com@test.com","orig":"test.com","reason":"mx_record","trusted":false,"valid":false,"vendor":"Unknown"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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
		clientUserId string
		projectId    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"create_client_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", "1", 2},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":4301},"data":{"date":"2021-07-05 15:45:39.000000","message":"done","status":true}}`),
			true},

		{"create_client_user_already_exist",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", "1", 2},
			[]byte(`{"_meta":{"count":3,"status":"SUCCESS","time":3414},"data":{"date":"2021-07-05 15:51:09.000000","message":"done","status":true}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.ClientUser{Email: tt.args.email, ClientUserId: tt.args.clientUserId, ProjectId: tt.args.projectId}
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
//		userId    int
//		projectId int
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
//			sdk := &sendios.SendiosSdk{
//				Request: tt.fields.Request,
//			}
//			got, err := sdk.CreatePushUser(tt.args.userId, tt.args.projectId, tt.args.url, tt.args.publicKey, tt.args.authToken)
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
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"force_confirm_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 2},
			[]byte(`{"status":"Accepted"}`),
			true},

		{"force_confirm_by_invalid_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"testgmail.com", 2},
			[]byte(`{"status":"Accepted"}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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
				ProjectId:    tt.args.projectId,
				LastReaction: time.Now().Unix(),
			}

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/project/%d/email/%s/confirm", tt.args.projectId, encodedEmail), params)
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
		{"buying_decision",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com"},
			[]byte(`{"_meta":{"count":2,"status":"SUCCESS","time":3923},"data":{"decision":false,"email":"test@gmail.com"}}`),
			true},

		{"buying_decision_validation_error",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"testgmail.com"},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3407},"data":{"error":"Request validation error:  email - This value is not a valid email address."}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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

func TestSendiosSdk_GetEmailUserByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"email_user_by_email_and_project",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", 2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3668},"data":{"user":{"activation":null,"channel_id":null,"clicks":0,"country":false,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":null,"last_online":null,"last_payment":{"active":1,"amount":1,"expires_at":1625479419,"id":1,"payment_count":1,"payment_type":1,"project_id":2,"started_at":1625479419,"user_id":5005},"last_reaction":null,"last_request":null,"meta":{"profile":{"age":null,"ak":null,"partner_id":null,"photo":null}},"name":"Volodymyr","project_id":2,"project_title":"Test project 1","sends":0,"sent_mails":[],"subchannel_id":null,"unsub_promo":[],"unsubscribe":[],"unsubscribe_types":[{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":1,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":2,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":3,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":4,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":5,"type_sig":"Undefined"}],"webpush":{"last_click":null,"last_push":"2020-10-26 14:47:25","reg_date":"2017-02-13 16:55:25"}}}}`),
			true},

		{"email_user_by_email_and_project_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"test@gmail.com", 2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3600},"data":{"error":"Not found  user for project 2 and email test@gmail.com"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/user/project/%d/email/%s", tt.args.projectId, tt.args.email))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmailUserByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetEmailUserById(t *testing.T) {
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
		{"email_user_by_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4507},"data":{"user":{"activation":null,"channel_id":null,"clicks":0,"country":false,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":null,"last_online":null,"last_payment":{"active":1,"amount":1,"expires_at":1625479419,"id":1,"payment_count":1,"payment_type":1,"project_id":2,"started_at":1625479419,"user_id":5005},"last_reaction":null,"last_request":null,"meta":{"profile":{"age":null,"ak":null,"partner_id":null,"photo":null}},"name":"Volodymyr","project_id":2,"project_title":"Test project 1","sends":0,"sent_mails":[],"subchannel_id":null,"unsub_promo":[],"unsubscribe":[],"unsubscribe_types":[{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":1,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":2,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":3,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":4,"type_sig":"Undefined"},{"created_at":"2021-07-05 15:06:24","sharded":0,"type_id":5,"type_sig":"Undefined"}],"webpush":{"last_click":null,"last_push":"2020-10-26 14:47:25","reg_date":"2017-02-13 16:55:25"}}}}`),
			true},

		{"email_user_by_id_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4477},"data":{"error":"User not found"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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
				t.Errorf("GetEmailUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetPushUserById(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"push_user_by_user_id",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3630},"data":{"result":{"hash":"NULL","id":717067,"invalid":null,"last_click":null,"last_online":null,"last_push":"2020-10-26 14:47:25","last_show":null,"meta":{"auth_token":"FpOCjghkdFV1JEUKOgJ-cw==","public_key":"BHv8Z_EomoVe33d2oGdwdbmT9Jd3rJd4VPZRXjzNPgtJzeHOu-hERyap53cV74nKVPQQ7BlNVwAKMGZaZsSr16A=","url":"https://android.googleapis.com/gcm/send/dtnHwCXnsBw:APA91bFPw56tcPNcVdWdOCzqhIdnPf4pyBIaTXg_tjMDDlLHlk4zo5BwpOTYJGbG_ZpMCuBZg_R3LIJaMalUHCQExEWD7__CNpYRcR-UHhkoHa_r9p3sD00kYr9qy9iBCztSTGWgZHzy"},"platform_id":null,"project_id":2,"reg_date":"2017-02-13 16:55:25","send_platform_id":null,"type":null,"user_id":5005}}}`),
			true},

		{"push_user_by_user_id_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3464},"data":{"error":"Not found user by id: 0"}}`),
			true},

		{"push_user_by_user_id_forbidden",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4494},"data":{"error":"Forbidden"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/webpush/user/get/%d", tt.args.userId))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPushUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUnsubListByEmailUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"unsub_list_by_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":5,"status":"SUCCESS","time":4030},"data":[{"created_at":"2021-07-05 15:06:24","name":"SystemMail07082","type_id":1},{"created_at":"2021-07-05 15:06:24","name":"Test2","type_id":2},{"created_at":"2021-07-05 15:06:24","name":"Qwerty","type_id":3},{"created_at":"2021-07-05 15:06:24","name":"Foobar","type_id":4},{"created_at":"2021-07-05 15:06:24","name":"TriggerMail04082","type_id":5}]}`),
			true},

		{"unsub_list_by_email_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4240},"data":{"error":"User not found"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsubtypes/%d", tt.args.userId))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnsubListByEmailUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUnsubscribeReason(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		UserId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"unsub_reason_by_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3537},"data":{"error":"User not found"}}`),
			true},
		{"unsub_reason_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3421},"data":{"result":false}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/unsubreason/%d", tt.args.UserId))
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
		{"unsub_by_date",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{time.Now().Unix()},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":4477},"data":{"error":"User not found"}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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

func TestSendiosSdk_GetUserFieldsByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"user_fields_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendios.io", 2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":5623},"data":{"result":{"custom_fields":[],"user":{"city_id":null,"confirm":1,"country_id":null,"created_at":"2021-06-17 10:29:27","email":"volodymyr.voloshyn@corp.sendios.io","err_response":0,"gender":"m","id":5005,"language":"en","last_mailed":0,"last_online":0,"last_reaction":0,"list_id":0,"meta":"[]","name":"Volodymyr","platform_id":1,"project_id":2,"status":1,"valid_id":null,"vendor_id":3,"vip":1}}}}`),
			true},

		{"user_fields_by_email_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"example@gmail.com", 2},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3726},"data":{"error":"User not found."}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/userfields/project/%d/email/%s", tt.args.projectId, tt.args.email))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFieldsByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_GetUserFieldsByUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.GetUserFieldsByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFieldsByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFieldsByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_IsUnsubByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"is_unsub_by_email_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3758},"data":{"error":"User not found"}}`),
			true},
		{"is_unsub_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3387},"data":{"result":false}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/isunsub/%d", tt.args.userId))
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsUnsubByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_IsUnsubUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"is_unsub_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3758},"data":{"error":"User not found"}}`),
			true},
		{"is_unsub",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3387},"data":{"result":false}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Get(ts.URL, fmt.Sprintf("/unsub/isunsub/%d", tt.args.userId))
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
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"remove_all_unsub_types",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4969},"data":null}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsubtypes/all/%d", tt.args.userId), nil)
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
		userId  int
		typeIds []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"remove_unsub_types_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, []int{1, 2, 3}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4275},"data":null}`),
			true},

		{"remove_unsub_types_email_user_not_found",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{0, []int{1, 2, 3}},
			[]byte(`{"_meta":{"count":1,"status":"ERROR","time":3615},"data":{"error":"User not found"}}`),
			true},

		{"remove_unsub_types_email_user_empty_types",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1, []int{}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3460},"data":null}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.TypeIds{TypeIds: tt.args.typeIds}

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsubtypes/nodiff/%d", tt.args.userId), params)
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
		clientId   int
		typeId     int
		categoryId int
		projectId  int
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
		{"send_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{3, 1, 1, 2, "test@gmail.com", map[string]string{"id": "1"}, map[string]string{"data": "test"}, map[string]string{"meta": "email"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4275},"data":null}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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
				t.Errorf("IsUnsubByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// later
func TestSendiosSdk_SendPushByEmailUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId   int
		title    string
		text     string
		url      string
		iconUrl  string
		typeId   int
		meta     map[string]string
		imageUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByEmailUserId(tt.args.userId, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeId, tt.args.meta, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPushByEmailUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendPushByEmailUserId() got = %v, want %v", got, tt.want)
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
		projectId int
		title     string
		text      string
		url       string
		iconUrl   string
		typeId    int
		meta      map[string]string
		imageUrl  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByProject(tt.args.projectId, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeId, tt.args.meta, tt.args.imageUrl)
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

//later
func TestSendiosSdk_SendPushByProjectIdAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectId int
		hash      string
		title     string
		text      string
		url       string
		iconUrl   string
		typeId    int
		meta      map[string]string
		imageUrl  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.SendPushByProjectIdAndHash(tt.args.projectId, tt.args.hash, tt.args.title, tt.args.text, tt.args.url, tt.args.iconUrl, tt.args.typeId, tt.args.meta, tt.args.imageUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendPushByProjectIdAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendPushByProjectIdAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetOnlineByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"set_online_by_email_and_project",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendio.io", 2},
			[]byte(`{"status":"Accepted"}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
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
				ProjectId:    tt.args.projectId,
				EncodedEmail: encodedEmail,
				Timestamp:    time.Now(),
			}
			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/project/%d/email/%s/online", tt.args.projectId, encodedEmail), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlineByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetOnlineByUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"set_online_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{1},
			[]byte(`{"status":"Accepted"}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			params := internal.OnlineByUser{UserId: tt.args.userId, Timestamp: time.Now()}

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/users/%d/online", tt.args.userId), params)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOnlineByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetUserFieldsByEmailAndProjectId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		email     string
		projectId int
		data      map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"set_user_fields_by_email",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{"volodymyr.voloshyn@corp.sendio.io", 2, map[string]string{"test": "test"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3850},"data":{"result":true}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			encodedEmail := internal.Base64Encoder(tt.args.email)
			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/userfields/project/%d/emailhash/%s", tt.args.projectId, encodedEmail), tt.args.data)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserFieldsByEmailAndProjectId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SetUserFieldsByUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
		data   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"set_user_fields_by_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2, map[string]string{"test": "test"}},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":4112},"data":{"result":true}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Put(ts.URL, fmt.Sprintf("/userfields/user/%d", tt.args.userId), tt.args.data)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserFieldsByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribeEmailUser(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{"subscribe_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3716},"data":{"subscribe":{"rowCount":1}}}`),
			true},
		{"subscribe_email_user",
			fields{&internal.Request{Client: &http.Client{Timeout: time.Second * 10}, Auth: &internal.Auth{"3", "VeaGGspBXpGQeZGbfegEeq5PPJ2CsjQ6"}}},
			args{2},
			[]byte(`{"_meta":{"count":1,"status":"SUCCESS","time":3616},"data":{"subscribe":{"rowCount":0}}}`),
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				fmt.Fprintln(w, string(tt.want))
			}))
			defer ts.Close()

			got, err := sdk.Request.Delete(ts.URL, fmt.Sprintf("/unsub/%d", tt.args.userId), nil)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribeEmailUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribePushUserByEmailUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.SubscribePushUserByEmailUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribePushUserByEmailUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribePushUserByEmailUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_SubscribePushUserByProjectIdAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectId int
		hash      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.SubscribePushUserByProjectIdAndHash(tt.args.projectId, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscribePushUserByProjectIdAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscribePushUserByProjectIdAndHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_TrackClickByMailId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		mailId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.TrackClickByMailId(tt.args.mailId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrackClickByMailId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackClickByMailId() got = %v, want %v", got, tt.want)
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
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserByAdmin(tt.args.email, tt.args.projectId)
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
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserBySettings(tt.args.userId)
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
		userId  int
		typeIds []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserByTypes(tt.args.userId, tt.args.typeIds)
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
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubEmailUserClient(tt.args.userId)
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

func TestSendiosSdk_UnsubscribePushUserByEmailUserId(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserByEmailUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserByEmailUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserByEmailUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubscribePushUserById(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		pushUserId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserById(tt.args.pushUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendiosSdk_UnsubscribePushUserByProjectIdAndHash(t *testing.T) {
	type fields struct {
		Request *internal.Request
	}
	type args struct {
		projectId int
		hash      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.UnsubscribePushUserByProjectIdAndHash(tt.args.projectId, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsubscribePushUserByProjectIdAndHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsubscribePushUserByProjectIdAndHash() got = %v, want %v", got, tt.want)
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
		projectId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &sendios.SendiosSdk{
				Request: tt.fields.Request,
			}
			got, err := sdk.ValidateEmail(tt.args.email, tt.args.projectId)
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
