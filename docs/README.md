
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# api

```go
import "github.com/rs401/lg/api"
```

Package api provides HTTP endpoints


# authclient

```go
import "github.com/rs401/lg/api/authclient"
```

Package authclient provides a client to the auth service


- [type AuthSvcClient](<#type-authsvcclient>)
  - [func NewAuthServiceClient(client *rpc.Client) AuthSvcClient](<#func-newauthserviceclient>)


## type AuthSvcClient

AuthSvcClient interface defining client methods

```go
type AuthSvcClient interface {
    SignUp(req *models.SignUpRequest, res *models.User) error
    SignIn(req *models.SignInRequest, res *models.User) error
    GetUser(req *models.GetUserRequest, res *models.User) error
    ListUsers(req string, res *models.GetUsersResponse) error
    UpdateUser(req *models.User, res *models.User) error
    DeleteUser(req *models.GetUserRequest, res *models.GetUserRequest) error
}
```

### func NewAuthServiceClient

```go
func NewAuthServiceClient(client *rpc.Client) AuthSvcClient
```

NewAuthServiceClient takes a pointer to an rpc\.Client and returns an AuthSvcClient\.

# handlers

```go
import "github.com/rs401/lg/api/handlers"
```

Package handlers provides handlerfuncs


- [type AuthHandlers](<#type-authhandlers>)
  - [func NewAuthHandlers(authSvcClient authclient.AuthSvcClient) AuthHandlers](<#func-newauthhandlers>)


## type AuthHandlers

AuthHandlers interface defining HandlerFuncs

```go
type AuthHandlers interface {
    SignUp(w http.ResponseWriter, r *http.Request)
    SignIn(w http.ResponseWriter, r *http.Request)
    UpdateUser(w http.ResponseWriter, r *http.Request)
    GetUser(w http.ResponseWriter, r *http.Request)
    GetUsers(w http.ResponseWriter, r *http.Request)
    DeleteUser(w http.ResponseWriter, r *http.Request)
}
```

### func NewAuthHandlers

```go
func NewAuthHandlers(authSvcClient authclient.AuthSvcClient) AuthHandlers
```

NewAuthHandlers takes an authclient\.AuthSvcClient and returns an AuthHandlers

# middlewares

```go
import "github.com/rs401/lg/api/middlewares"
```

### Package middlewares provides middlewares

Package middlewares provides middlewares


- [func HeadersMiddleware(next http.Handler) http.Handler](<#func-headersmiddleware>)
- [func LogMiddleware(next http.Handler) http.Handler](<#func-logmiddleware>)
- [func SetupMiddleWares(r *mux.Router)](<#func-setupmiddlewares>)


## func HeadersMiddleware

```go
func HeadersMiddleware(next http.Handler) http.Handler
```

HeadersMiddleware sets "Content\-Type" header to "application/json"

## func LogMiddleware

```go
func LogMiddleware(next http.Handler) http.Handler
```

LogMiddleware logs information about the request

## func SetupMiddleWares

```go
func SetupMiddleWares(r *mux.Router)
```

SetupMiddleWares takes a \*mux\.Router and sets it to use the middlewares

# routes

```go
import "github.com/rs401/lg/api/routes"
```

Package routes provides utility to setup routes


- [func SetupRoutes(r *mux.Router, hndlrs handlers.AuthHandlers)](<#func-setuproutes>)


## func SetupRoutes

```go
func SetupRoutes(r *mux.Router, hndlrs handlers.AuthHandlers)
```

SetupRoutes takes a \*mux\.Router and a AuthHandlers to configure \*mux\.Routes

# tokenutils

```go
import "github.com/rs401/lg/api/tokenutils"
```

Package tokenutils provides utilities for JWTs


- [func ExtractToken(r *http.Request) string](<#func-extracttoken>)
- [func VerifyRefreshToken(r *http.Request) (*jwt.Token, error)](<#func-verifyrefreshtoken>)
- [func VerifyToken(r *http.Request) (*jwt.Token, error)](<#func-verifytoken>)
- [type Claims](<#type-claims>)
- [type Tokens](<#type-tokens>)
  - [func CreateToken(userid uint) (*Tokens, error)](<#func-createtoken>)


## func ExtractToken

```go
func ExtractToken(r *http.Request) string
```

ExtractToken extracts an access token from request header

## func VerifyRefreshToken

```go
func VerifyRefreshToken(r *http.Request) (*jwt.Token, error)
```

VerifyRefreshToken verifies a refresh token is legit

## func VerifyToken

```go
func VerifyToken(r *http.Request) (*jwt.Token, error)
```

VerifyToken verifies a token is legit

## type Claims

Claims holds a UserId and standard claims

```go
type Claims struct {
    UserId uint `json:"userid"` // User ID
    jwt.StandardClaims
}
```

## type Tokens

Tokens holds an access token and a refresh token

```go
type Tokens struct {
    AccessToken  string
    RefreshToken string
}
```

### func CreateToken

```go
func CreateToken(userid uint) (*Tokens, error)
```

CreateToken takes an ID and generates a Tokens


# auth

```go
import "github.com/rs401/lg/auth"
```

Package auth provides RPC methods to be called




# models

```go
import "github.com/rs401/lg/auth/models"
```

### Package models provides data structures

Package models provides data structures


- [type GetUserRequest](<#type-getuserrequest>)
- [type GetUsersResponse](<#type-getusersresponse>)
- [type SignInRequest](<#type-signinrequest>)
- [type SignInResponse](<#type-signinresponse>)
- [type SignUpRequest](<#type-signuprequest>)
- [type User](<#type-user>)


## type GetUserRequest

GetUserRequest holds an Id

```go
type GetUserRequest struct {
    Id uint `json:"id"`
}
```

## type GetUsersResponse

GetUsersResponse holds a slice of \*Users

```go
type GetUsersResponse struct {
    Users []*User
}
```

## type SignInRequest

SignInRequest holds an Email and Password

```go
type SignInRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

## type SignInResponse

SignInResponse holds a token

```go
type SignInResponse struct {
    Token string `json:"token"`
}
```

## type SignUpRequest

SignUpRequest holds a Name\, Email and Password

```go
type SignUpRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
```

## type User

User model for user

```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"<-;unique;not null"`
    Email     string    `json:"email" gorm:"<-;unique;not null"`
    Password  []byte    `json:"-"`
    CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
```

# repository

```go
import "github.com/rs401/lg/auth/repository"
```

Package repository provides methods to interact with the database


- [Variables](<#variables>)
- [type UsersRepository](<#type-usersrepository>)
  - [func NewUsersRepository(conn db.Connection) UsersRepository](<#func-newusersrepository>)


## Variables

ErrorBadID custom error

```go
var ErrorBadID error = errors.New("bad id")
```

## type UsersRepository

UsersRepository interface defines methods for interacting with the database

```go
type UsersRepository interface {
    Save(user *models.User) error
    GetById(id uint) (*models.User, error)
    GetByEmail(email string) (*models.User, error)
    GetAll() ([]*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
}
```

### func NewUsersRepository

```go
func NewUsersRepository(conn db.Connection) UsersRepository
```

NewUsersRepository takes a db\.Connection and returns a UsersRepository

# service

```go
import "github.com/rs401/lg/auth/service"
```

Package service provides RPC methods to call repository actions


- [type AuthSvc](<#type-authsvc>)
  - [func NewAuthService(usersRepository repository.UsersRepository) AuthSvc](<#func-newauthservice>)


## type AuthSvc

AuthSvc interface defines the RPC methods to call repository actions

```go
type AuthSvc interface {
    SignUp(*models.SignUpRequest, *models.User) error
    SignIn(*models.SignInRequest, *models.User) error
    GetUser(*models.GetUserRequest, *models.User) error
    ListUsers(string, *models.GetUsersResponse) error
    UpdateUser(*models.User, *models.User) error
    DeleteUser(*models.GetUserRequest, *models.GetUserRequest) error
}
```

### func NewAuthService

```go
func NewAuthService(usersRepository repository.UsersRepository) AuthSvc
```

NewAuthService takes a users repository and returns an AuthSvc



# db

```go
import "github.com/rs401/lg/db"
```

### Package db provides utilities to configure and connect to a database

Package db provides utilities to configure and connect to a database


- [type Config](<#type-config>)
  - [func NewConfig() Config](<#func-newconfig>)
- [type Connection](<#type-connection>)
  - [func NewConnection(cfg Config) (Connection, error)](<#func-newconnection>)
- [type ConnectionError](<#type-connectionerror>)
  - [func (ce *ConnectionError) Error() string](<#func-connectionerror-error>)


## type Config

Config interface defines methods for retreiving config information

```go
type Config interface {
    ConnStr() string
    DbName() string
}
```

### func NewConfig

```go
func NewConfig() Config
```

NewConfig constructs and returns a new Config

## type Connection

Connection interface defines a method for retrieving the \*gorm\.DB

```go
type Connection interface {
    DB() *gorm.DB
}
```

### func NewConnection

```go
func NewConnection(cfg Config) (Connection, error)
```

NewConnection takes a db\.Config and returns a Connection and an error

## type ConnectionError

ConnectionError custom error type

```go
type ConnectionError struct{}
```

### func \(\*ConnectionError\) Error

```go
func (ce *ConnectionError) Error() string
```

Error implements the error interface




Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)