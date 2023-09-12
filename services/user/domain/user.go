package domain

import (
	m "gokit-basic/common/model"
	"math/rand"
	"time"

	guuid "github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserDomain struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Age       string    `json:"age" gorm:"column:age"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*UserDomain) TableName() string {
	return "user"
}

func ToUserDomainMapper(v *m.SingleUser) *UserDomain {
	return &UserDomain{
		ID:        RandomUUID(),
		Name:      v.Name,
		Age:       v.Age,
		Phone:     v.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *UserDomain) ToUserProtoMappter() *m.SingleUserResponse {
	return &m.SingleUserResponse{
		Id:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Age:       u.Age,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func RandomUUID() string {
	return guuid.NewString()
}

func NameGenerator() string {
	rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixMicro()
	name := namegenerator.NewNameGenerator(seed)

	return name.Generate()
}
