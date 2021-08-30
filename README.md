##How to run 

###Sendios sdk initializing
```
    sdk := sendios.NewSendiosSdk("id", "client_key")
```

##Methods
####Create web push user 
```
    result, err := sdk.CreatePushUser(userId, projectId int, url, publicKey, authToken string)
```

###Create client user
```
    result, err := sdk.CreateClientUser(email string, clientUserId string, projectId int)
```

###Check email
```
    result, err := sdk.CheckEmail(email string, sanitize bool)
```

###Unsub email user by types
```
    result, err := sdk.AddTypesToUnsubByEmailUser(userId int, typeIds []int)
```

####Add payment information by email and project
```
    result, err := sdk.AddPaymentByEmailAndProjectId(email string, projectId int, startDate, expireDate int64, totalCount, paymentType, amount int)
```

####Add payment information by user id
```
    result, err := sdk.AddPaymentByUserId(userId int, startDate, expireDate int64, totalCount, paymentType, amount int)
```

####Create prod event
```
result, err := sdk.ProdEventSend(data interface{})
```

####Set online by email and project 
```
result, err :=  sdk.SetOnlineByEmailAndProjectId(email string, projectId int)
```

####Set online by user id
```
result, err :=  sdk.SetOnlineByUser(userId int)
```

####Force confirm by email and project 
```
result, err := sdk.ForceConfirmByEmailAndProject(email string, projectId int)
```

result, err := sdk.GetBuyingDecisions(email)
result, err := sdk.GetEmailUserByEmailAndProjectId(email, projectId)
result, err:= sdk.GetEmailUserById(1)
result, err := sdk.GetPushUserById(0)
result, err := sdk.GetUnsubListByEmailUserId(0)
result, err := sdk.GetUnsubscribeReason(email, projectId)
result, err := sdk.GetUnsubscribesByDate(time.Now().Unix())
result, err := sdk.GetUserFieldsByEmailAndProjectId(email, projectId)
result, err := sdk.IsUnsubByEmailAndProjectId(email, projectId)
result, err := sdk.IsUnsubUser(userId)
