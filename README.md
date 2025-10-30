# vigilant-happiness

An endpoint that processes a response to return a Go struct.

## https://vigilant-happiness-df7653c99ec9.herokuapp.com/webhook 

Send a webhook payload to convert to a Go type struct.

Payload:

{
  "event": "taskCreated",
  "history_items": [
    {
      "id": "2800763136717140857",
      "type": 1,
      "date": "1642734631523",
      "field": "status",
      "parent_id": "162641062",
      "data": {
        "status_type": "open"
      },
      "source": null,
      "user": {
        "id": 183,
        "username": "John",
        "email": "john@company.com",
        "color": "#7b68ee",
        "initials": "J",
        "profilePicture": null
      },
      "before": {
        "status": null,
        "color": "#000000",
        "type": "removed",
        "orderindex": -1
      },
      "after": {
        "status": "to do",
        "color": "#f9d900",
        "orderindex": 0,
        "type": "open"
      }
    },
    {
      "id": "2800763136700363640",
      "type": 1,
      "date": "1642734631523",
      "field": "task_creation",
      "parent_id": "162641062",
      "data": {},
      "source": null,
      "user": {
        "id": 183,
        "username": "John",
        "email": "john@company.com",
        "color": "#7b68ee",
        "initials": "J",
        "profilePicture": null
      },
      "before": null,
      "after": null
    }
  ],
  "task_id": "1vj37mc",
  "webhook_id": "7fa3ec74-69a8-4530-a251-8a13730bd204"
}

Response: 

type Data struct {
    StatusType string `json:"status_type" `
}

type Before struct {
    Status any `json:"status" `
    Color string `json:"color" `
    Type string `json:"type" `
    Orderindex int `json:"orderindex" `
}

type After struct {
    Type string `json:"type" `
    Status string `json:"status" `
    Color string `json:"color" `
    Orderindex int `json:"orderindex" `
}

type User struct {
    Id int `json:"id" `
    Username string `json:"username" `
    Email string `json:"email" `
    Color string `json:"color" `
    Initials string `json:"initials" `
    ProfilePicture any `json:"profilePicture" `
}

type HistoryItem struct {
    Date string `json:"date" `
    User User `json:"user" `
    Data Data `json:"data" `
    Source any `json:"source" `
    Before Before `json:"before" `
    After After `json:"after" `
    Id string `json:"id" `
    Type int `json:"type" `
    Field string `json:"field" `
    ParentId string `json:"parent_id" `
}

type taskCreated struct {
    Event string `json:"event" `
    HistoryItems []HistoryItem `json:"history_items" `
    TaskId string `json:"task_id" `
    WebhookId string `json:"webhook_id" `
}

## https://vigilant-happiness-df7653c99ec9.herokuapp.com/type?name={type_name}

Send JSON to convert to a Go type struct.

Request Body:

{
    "username": "PeterJBishop",
    "email": "pjb.den@gmail.com",
    "password": "password1",
    "address": {
        "street": "2059 Albion Street",
        "city": "Denver",
        "state": "CO",
        "zip": 80207
    }
}

Response:

type user struct {
    Username string `json:"username" `
    Email string `json:"email" `
    Password string `json:"password" `
    Address Address `json:"address" `
}

type Address struct {
    Zip int `json:"zip" `
    Street string `json:"street" `
    City string `json:"city" `
    State string `json:"state" `
}