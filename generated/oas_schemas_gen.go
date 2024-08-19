// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

type Address string

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

type Date time.Time

type Developer string

type DummyLoginGetOK struct {
	Token OptToken `json:"token"`
}

// GetToken returns the value of Token.
func (s *DummyLoginGetOK) GetToken() OptToken {
	return s.Token
}

// SetToken sets the value of Token.
func (s *DummyLoginGetOK) SetToken(val OptToken) {
	s.Token = val
}

func (*DummyLoginGetOK) dummyLoginGetRes() {}

type Email string

// Квартира.
// Ref: #/components/schemas/Flat
type Flat struct {
	ID      FlatId  `json:"id"`
	HouseID HouseId `json:"house_id"`
	Price   Price   `json:"price"`
	Rooms   Rooms   `json:"rooms"`
	Status  Status  `json:"status"`
}

// GetID returns the value of ID.
func (s *Flat) GetID() FlatId {
	return s.ID
}

// GetHouseID returns the value of HouseID.
func (s *Flat) GetHouseID() HouseId {
	return s.HouseID
}

// GetPrice returns the value of Price.
func (s *Flat) GetPrice() Price {
	return s.Price
}

// GetRooms returns the value of Rooms.
func (s *Flat) GetRooms() Rooms {
	return s.Rooms
}

// GetStatus returns the value of Status.
func (s *Flat) GetStatus() Status {
	return s.Status
}

// SetID sets the value of ID.
func (s *Flat) SetID(val FlatId) {
	s.ID = val
}

// SetHouseID sets the value of HouseID.
func (s *Flat) SetHouseID(val HouseId) {
	s.HouseID = val
}

// SetPrice sets the value of Price.
func (s *Flat) SetPrice(val Price) {
	s.Price = val
}

// SetRooms sets the value of Rooms.
func (s *Flat) SetRooms(val Rooms) {
	s.Rooms = val
}

// SetStatus sets the value of Status.
func (s *Flat) SetStatus(val Status) {
	s.Status = val
}

func (*Flat) flatCreatePostRes() {}
func (*Flat) flatUpdatePostRes() {}

type FlatCreatePostReq struct {
	Number  FlatId  `json:"number"`
	HouseID HouseId `json:"house_id"`
	Price   Price   `json:"price"`
	Rooms   Rooms   `json:"rooms"`
}

// GetNumber returns the value of Number.
func (s *FlatCreatePostReq) GetNumber() FlatId {
	return s.Number
}

// GetHouseID returns the value of HouseID.
func (s *FlatCreatePostReq) GetHouseID() HouseId {
	return s.HouseID
}

// GetPrice returns the value of Price.
func (s *FlatCreatePostReq) GetPrice() Price {
	return s.Price
}

// GetRooms returns the value of Rooms.
func (s *FlatCreatePostReq) GetRooms() Rooms {
	return s.Rooms
}

// SetNumber sets the value of Number.
func (s *FlatCreatePostReq) SetNumber(val FlatId) {
	s.Number = val
}

// SetHouseID sets the value of HouseID.
func (s *FlatCreatePostReq) SetHouseID(val HouseId) {
	s.HouseID = val
}

// SetPrice sets the value of Price.
func (s *FlatCreatePostReq) SetPrice(val Price) {
	s.Price = val
}

// SetRooms sets the value of Rooms.
func (s *FlatCreatePostReq) SetRooms(val Rooms) {
	s.Rooms = val
}

type FlatId int

type FlatUpdatePostReq struct {
	ID      FlatId  `json:"id"`
	HouseID HouseId `json:"house_id"`
	Status  Status  `json:"status"`
}

// GetID returns the value of ID.
func (s *FlatUpdatePostReq) GetID() FlatId {
	return s.ID
}

// GetHouseID returns the value of HouseID.
func (s *FlatUpdatePostReq) GetHouseID() HouseId {
	return s.HouseID
}

// GetStatus returns the value of Status.
func (s *FlatUpdatePostReq) GetStatus() Status {
	return s.Status
}

// SetID sets the value of ID.
func (s *FlatUpdatePostReq) SetID(val FlatId) {
	s.ID = val
}

// SetHouseID sets the value of HouseID.
func (s *FlatUpdatePostReq) SetHouseID(val HouseId) {
	s.HouseID = val
}

// SetStatus sets the value of Status.
func (s *FlatUpdatePostReq) SetStatus(val Status) {
	s.Status = val
}

// Дом.
// Ref: #/components/schemas/House
type House struct {
	ID        HouseId         `json:"id"`
	Address   Address         `json:"address"`
	Year      Year            `json:"year"`
	Developer OptNilDeveloper `json:"developer"`
	CreatedAt OptDate         `json:"created_at"`
	UpdateAt  OptDate         `json:"update_at"`
}

// GetID returns the value of ID.
func (s *House) GetID() HouseId {
	return s.ID
}

// GetAddress returns the value of Address.
func (s *House) GetAddress() Address {
	return s.Address
}

// GetYear returns the value of Year.
func (s *House) GetYear() Year {
	return s.Year
}

// GetDeveloper returns the value of Developer.
func (s *House) GetDeveloper() OptNilDeveloper {
	return s.Developer
}

// GetCreatedAt returns the value of CreatedAt.
func (s *House) GetCreatedAt() OptDate {
	return s.CreatedAt
}

// GetUpdateAt returns the value of UpdateAt.
func (s *House) GetUpdateAt() OptDate {
	return s.UpdateAt
}

// SetID sets the value of ID.
func (s *House) SetID(val HouseId) {
	s.ID = val
}

// SetAddress sets the value of Address.
func (s *House) SetAddress(val Address) {
	s.Address = val
}

// SetYear sets the value of Year.
func (s *House) SetYear(val Year) {
	s.Year = val
}

// SetDeveloper sets the value of Developer.
func (s *House) SetDeveloper(val OptNilDeveloper) {
	s.Developer = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *House) SetCreatedAt(val OptDate) {
	s.CreatedAt = val
}

// SetUpdateAt sets the value of UpdateAt.
func (s *House) SetUpdateAt(val OptDate) {
	s.UpdateAt = val
}

func (*House) houseCreatePostRes() {}

type HouseCreatePostReq struct {
	Address   Address         `json:"address"`
	Year      Year            `json:"year"`
	Developer OptNilDeveloper `json:"developer"`
}

// GetAddress returns the value of Address.
func (s *HouseCreatePostReq) GetAddress() Address {
	return s.Address
}

// GetYear returns the value of Year.
func (s *HouseCreatePostReq) GetYear() Year {
	return s.Year
}

// GetDeveloper returns the value of Developer.
func (s *HouseCreatePostReq) GetDeveloper() OptNilDeveloper {
	return s.Developer
}

// SetAddress sets the value of Address.
func (s *HouseCreatePostReq) SetAddress(val Address) {
	s.Address = val
}

// SetYear sets the value of Year.
func (s *HouseCreatePostReq) SetYear(val Year) {
	s.Year = val
}

// SetDeveloper sets the value of Developer.
func (s *HouseCreatePostReq) SetDeveloper(val OptNilDeveloper) {
	s.Developer = val
}

type HouseIDGetOK struct {
	Flats []Flat `json:"flats"`
}

// GetFlats returns the value of Flats.
func (s *HouseIDGetOK) GetFlats() []Flat {
	return s.Flats
}

// SetFlats sets the value of Flats.
func (s *HouseIDGetOK) SetFlats(val []Flat) {
	s.Flats = val
}

func (*HouseIDGetOK) houseIDGetRes() {}

// HouseIDSubscribePostOK is response for HouseIDSubscribePost operation.
type HouseIDSubscribePostOK struct{}

func (*HouseIDSubscribePostOK) houseIDSubscribePostRes() {}

type HouseIDSubscribePostReq struct {
	Email Email `json:"email"`
}

// GetEmail returns the value of Email.
func (s *HouseIDSubscribePostReq) GetEmail() Email {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *HouseIDSubscribePostReq) SetEmail(val Email) {
	s.Email = val
}

type HouseId int

// LoginPostBadRequest is response for LoginPost operation.
type LoginPostBadRequest struct{}

func (*LoginPostBadRequest) loginPostRes() {}

// LoginPostNotFound is response for LoginPost operation.
type LoginPostNotFound struct{}

func (*LoginPostNotFound) loginPostRes() {}

type LoginPostOK struct {
	Token OptToken `json:"token"`
}

// GetToken returns the value of Token.
func (s *LoginPostOK) GetToken() OptToken {
	return s.Token
}

// SetToken sets the value of Token.
func (s *LoginPostOK) SetToken(val OptToken) {
	s.Token = val
}

func (*LoginPostOK) loginPostRes() {}

type LoginPostReq struct {
	ID       OptUserId   `json:"id"`
	Password OptPassword `json:"password"`
}

// GetID returns the value of ID.
func (s *LoginPostReq) GetID() OptUserId {
	return s.ID
}

// GetPassword returns the value of Password.
func (s *LoginPostReq) GetPassword() OptPassword {
	return s.Password
}

// SetID sets the value of ID.
func (s *LoginPostReq) SetID(val OptUserId) {
	s.ID = val
}

// SetPassword sets the value of Password.
func (s *LoginPostReq) SetPassword(val OptPassword) {
	s.Password = val
}

// NewOptDate returns new OptDate with value set to v.
func NewOptDate(v Date) OptDate {
	return OptDate{
		Value: v,
		Set:   true,
	}
}

// OptDate is optional Date.
type OptDate struct {
	Value Date
	Set   bool
}

// IsSet returns true if OptDate was set.
func (o OptDate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDate) Reset() {
	var v Date
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDate) SetTo(v Date) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDate) Get() (v Date, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDate) Or(d Date) Date {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFlatCreatePostReq returns new OptFlatCreatePostReq with value set to v.
func NewOptFlatCreatePostReq(v FlatCreatePostReq) OptFlatCreatePostReq {
	return OptFlatCreatePostReq{
		Value: v,
		Set:   true,
	}
}

// OptFlatCreatePostReq is optional FlatCreatePostReq.
type OptFlatCreatePostReq struct {
	Value FlatCreatePostReq
	Set   bool
}

// IsSet returns true if OptFlatCreatePostReq was set.
func (o OptFlatCreatePostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFlatCreatePostReq) Reset() {
	var v FlatCreatePostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFlatCreatePostReq) SetTo(v FlatCreatePostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFlatCreatePostReq) Get() (v FlatCreatePostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFlatCreatePostReq) Or(d FlatCreatePostReq) FlatCreatePostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFlatUpdatePostReq returns new OptFlatUpdatePostReq with value set to v.
func NewOptFlatUpdatePostReq(v FlatUpdatePostReq) OptFlatUpdatePostReq {
	return OptFlatUpdatePostReq{
		Value: v,
		Set:   true,
	}
}

// OptFlatUpdatePostReq is optional FlatUpdatePostReq.
type OptFlatUpdatePostReq struct {
	Value FlatUpdatePostReq
	Set   bool
}

// IsSet returns true if OptFlatUpdatePostReq was set.
func (o OptFlatUpdatePostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFlatUpdatePostReq) Reset() {
	var v FlatUpdatePostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFlatUpdatePostReq) SetTo(v FlatUpdatePostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFlatUpdatePostReq) Get() (v FlatUpdatePostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFlatUpdatePostReq) Or(d FlatUpdatePostReq) FlatUpdatePostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptHouseCreatePostReq returns new OptHouseCreatePostReq with value set to v.
func NewOptHouseCreatePostReq(v HouseCreatePostReq) OptHouseCreatePostReq {
	return OptHouseCreatePostReq{
		Value: v,
		Set:   true,
	}
}

// OptHouseCreatePostReq is optional HouseCreatePostReq.
type OptHouseCreatePostReq struct {
	Value HouseCreatePostReq
	Set   bool
}

// IsSet returns true if OptHouseCreatePostReq was set.
func (o OptHouseCreatePostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptHouseCreatePostReq) Reset() {
	var v HouseCreatePostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptHouseCreatePostReq) SetTo(v HouseCreatePostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptHouseCreatePostReq) Get() (v HouseCreatePostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptHouseCreatePostReq) Or(d HouseCreatePostReq) HouseCreatePostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptHouseIDSubscribePostReq returns new OptHouseIDSubscribePostReq with value set to v.
func NewOptHouseIDSubscribePostReq(v HouseIDSubscribePostReq) OptHouseIDSubscribePostReq {
	return OptHouseIDSubscribePostReq{
		Value: v,
		Set:   true,
	}
}

// OptHouseIDSubscribePostReq is optional HouseIDSubscribePostReq.
type OptHouseIDSubscribePostReq struct {
	Value HouseIDSubscribePostReq
	Set   bool
}

// IsSet returns true if OptHouseIDSubscribePostReq was set.
func (o OptHouseIDSubscribePostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptHouseIDSubscribePostReq) Reset() {
	var v HouseIDSubscribePostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptHouseIDSubscribePostReq) SetTo(v HouseIDSubscribePostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptHouseIDSubscribePostReq) Get() (v HouseIDSubscribePostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptHouseIDSubscribePostReq) Or(d HouseIDSubscribePostReq) HouseIDSubscribePostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptLoginPostReq returns new OptLoginPostReq with value set to v.
func NewOptLoginPostReq(v LoginPostReq) OptLoginPostReq {
	return OptLoginPostReq{
		Value: v,
		Set:   true,
	}
}

// OptLoginPostReq is optional LoginPostReq.
type OptLoginPostReq struct {
	Value LoginPostReq
	Set   bool
}

// IsSet returns true if OptLoginPostReq was set.
func (o OptLoginPostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptLoginPostReq) Reset() {
	var v LoginPostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptLoginPostReq) SetTo(v LoginPostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptLoginPostReq) Get() (v LoginPostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptLoginPostReq) Or(d LoginPostReq) LoginPostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilDeveloper returns new OptNilDeveloper with value set to v.
func NewOptNilDeveloper(v Developer) OptNilDeveloper {
	return OptNilDeveloper{
		Value: v,
		Set:   true,
	}
}

// OptNilDeveloper is optional nullable Developer.
type OptNilDeveloper struct {
	Value Developer
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilDeveloper was set.
func (o OptNilDeveloper) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilDeveloper) Reset() {
	var v Developer
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilDeveloper) SetTo(v Developer) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilDeveloper) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *OptNilDeveloper) SetToNull() {
	o.Set = true
	o.Null = true
	var v Developer
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilDeveloper) Get() (v Developer, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilDeveloper) Or(d Developer) Developer {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPassword returns new OptPassword with value set to v.
func NewOptPassword(v Password) OptPassword {
	return OptPassword{
		Value: v,
		Set:   true,
	}
}

// OptPassword is optional Password.
type OptPassword struct {
	Value Password
	Set   bool
}

// IsSet returns true if OptPassword was set.
func (o OptPassword) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPassword) Reset() {
	var v Password
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPassword) SetTo(v Password) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPassword) Get() (v Password, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPassword) Or(d Password) Password {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptRegisterPostReq returns new OptRegisterPostReq with value set to v.
func NewOptRegisterPostReq(v RegisterPostReq) OptRegisterPostReq {
	return OptRegisterPostReq{
		Value: v,
		Set:   true,
	}
}

// OptRegisterPostReq is optional RegisterPostReq.
type OptRegisterPostReq struct {
	Value RegisterPostReq
	Set   bool
}

// IsSet returns true if OptRegisterPostReq was set.
func (o OptRegisterPostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptRegisterPostReq) Reset() {
	var v RegisterPostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptRegisterPostReq) SetTo(v RegisterPostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptRegisterPostReq) Get() (v RegisterPostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptRegisterPostReq) Or(d RegisterPostReq) RegisterPostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptToken returns new OptToken with value set to v.
func NewOptToken(v Token) OptToken {
	return OptToken{
		Value: v,
		Set:   true,
	}
}

// OptToken is optional Token.
type OptToken struct {
	Value Token
	Set   bool
}

// IsSet returns true if OptToken was set.
func (o OptToken) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptToken) Reset() {
	var v Token
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptToken) SetTo(v Token) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptToken) Get() (v Token, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptToken) Or(d Token) Token {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserId returns new OptUserId with value set to v.
func NewOptUserId(v UserId) OptUserId {
	return OptUserId{
		Value: v,
		Set:   true,
	}
}

// OptUserId is optional UserId.
type OptUserId struct {
	Value UserId
	Set   bool
}

// IsSet returns true if OptUserId was set.
func (o OptUserId) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserId) Reset() {
	var v UserId
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserId) SetTo(v UserId) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserId) Get() (v UserId, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserId) Or(d UserId) UserId {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type Password string

type Price int

type R400 struct {
	// Текст бизнесовой ошибки.
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *R400) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *R400) SetMessage(val string) {
	s.Message = val
}

func (*R400) flatCreatePostRes()       {}
func (*R400) flatUpdatePostRes()       {}
func (*R400) houseCreatePostRes()      {}
func (*R400) houseIDGetRes()           {}
func (*R400) houseIDSubscribePostRes() {}

// Ref: #/components/responses/401
type R401 struct{}

func (*R401) flatCreatePostRes()       {}
func (*R401) flatUpdatePostRes()       {}
func (*R401) houseCreatePostRes()      {}
func (*R401) houseIDGetRes()           {}
func (*R401) houseIDSubscribePostRes() {}

type R5xx struct {
	// Описание ошибки.
	Message string `json:"message"`
	// Идентификатор запроса. Предназначен для более
	// быстрого поиска проблем.
	RequestID OptString `json:"request_id"`
	// Код ошибки. Предназначен для классификации проблем и
	// более быстрого решения проблем.
	Code OptInt `json:"code"`
}

// GetMessage returns the value of Message.
func (s *R5xx) GetMessage() string {
	return s.Message
}

// GetRequestID returns the value of RequestID.
func (s *R5xx) GetRequestID() OptString {
	return s.RequestID
}

// GetCode returns the value of Code.
func (s *R5xx) GetCode() OptInt {
	return s.Code
}

// SetMessage sets the value of Message.
func (s *R5xx) SetMessage(val string) {
	s.Message = val
}

// SetRequestID sets the value of RequestID.
func (s *R5xx) SetRequestID(val OptString) {
	s.RequestID = val
}

// SetCode sets the value of Code.
func (s *R5xx) SetCode(val OptInt) {
	s.Code = val
}

func (*R5xx) dummyLoginGetRes()        {}
func (*R5xx) flatCreatePostRes()       {}
func (*R5xx) flatUpdatePostRes()       {}
func (*R5xx) houseCreatePostRes()      {}
func (*R5xx) houseIDGetRes()           {}
func (*R5xx) houseIDSubscribePostRes() {}
func (*R5xx) loginPostRes()            {}
func (*R5xx) registerPostRes()         {}

// RegisterPostBadRequest is response for RegisterPost operation.
type RegisterPostBadRequest struct{}

func (*RegisterPostBadRequest) registerPostRes() {}

type RegisterPostOK struct {
	UserID OptUserId `json:"user_id"`
}

// GetUserID returns the value of UserID.
func (s *RegisterPostOK) GetUserID() OptUserId {
	return s.UserID
}

// SetUserID sets the value of UserID.
func (s *RegisterPostOK) SetUserID(val OptUserId) {
	s.UserID = val
}

func (*RegisterPostOK) registerPostRes() {}

type RegisterPostReq struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
	UserType UserType `json:"user_type"`
}

// GetEmail returns the value of Email.
func (s *RegisterPostReq) GetEmail() Email {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *RegisterPostReq) GetPassword() Password {
	return s.Password
}

// GetUserType returns the value of UserType.
func (s *RegisterPostReq) GetUserType() UserType {
	return s.UserType
}

// SetEmail sets the value of Email.
func (s *RegisterPostReq) SetEmail(val Email) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *RegisterPostReq) SetPassword(val Password) {
	s.Password = val
}

// SetUserType sets the value of UserType.
func (s *RegisterPostReq) SetUserType(val UserType) {
	s.UserType = val
}

type Rooms int

// Статус квартиры.
// Ref: #/components/schemas/Status
type Status string

const (
	StatusCreated      Status = "created"
	StatusApproved     Status = "approved"
	StatusDeclined     Status = "declined"
	StatusOnModeration Status = "on moderation"
)

// AllValues returns all Status values.
func (Status) AllValues() []Status {
	return []Status{
		StatusCreated,
		StatusApproved,
		StatusDeclined,
		StatusOnModeration,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Status) MarshalText() ([]byte, error) {
	switch s {
	case StatusCreated:
		return []byte(s), nil
	case StatusApproved:
		return []byte(s), nil
	case StatusDeclined:
		return []byte(s), nil
	case StatusOnModeration:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Status) UnmarshalText(data []byte) error {
	switch Status(data) {
	case StatusCreated:
		*s = StatusCreated
		return nil
	case StatusApproved:
		*s = StatusApproved
		return nil
	case StatusDeclined:
		*s = StatusDeclined
		return nil
	case StatusOnModeration:
		*s = StatusOnModeration
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type Token string

type UserId uuid.UUID

// Тип пользователя.
// Ref: #/components/schemas/UserType
type UserType string

const (
	UserTypeClient    UserType = "client"
	UserTypeModerator UserType = "moderator"
)

// AllValues returns all UserType values.
func (UserType) AllValues() []UserType {
	return []UserType{
		UserTypeClient,
		UserTypeModerator,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UserType) MarshalText() ([]byte, error) {
	switch s {
	case UserTypeClient:
		return []byte(s), nil
	case UserTypeModerator:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserType) UnmarshalText(data []byte) error {
	switch UserType(data) {
	case UserTypeClient:
		*s = UserTypeClient
		return nil
	case UserTypeModerator:
		*s = UserTypeModerator
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type Year int
