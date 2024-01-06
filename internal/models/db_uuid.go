package models

import (
	"database/sql/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DbUUID uuid.UUID

func StringToDbUUID(s string) (DbUUID, error) {
	id, err := uuid.Parse(s)
	return DbUUID(id), err
}

func (id DbUUID) String() string {
	return uuid.UUID(id).String()
}

func (id DbUUID) GormDataType() string {
	return "binary(16)"
}

func (id DbUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(id)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (id *DbUUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*id = DbUUID(s)
	return err
}

func (id *DbUUID) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*id = DbUUID(parseByte)
	return err
}

func (id DbUUID) Value() (driver.Value, error) {
	return uuid.UUID(id).MarshalBinary()
}

func GetUserId(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.GetString("x-user-id"))
}

func NewDbUUID() DbUUID {
	return DbUUID(uuid.New())
}
