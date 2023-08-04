package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type bindingType int64

const (
	BindType_JSON  = 1
	BindType_URI   = 2
	BindType_Query = 3
)

func BindData(c *gin.Context, req interface{}, bindingTypes ...bindingType) error {
	for _, v := range bindingTypes {
		switch v {
		case BindType_JSON:
			if err := c.BindJSON(req); err != nil {
				return err
			}
		case BindType_URI:
			if err := c.BindUri(req); err != nil {
				return err
			}
		case BindType_Query:
			if err := c.BindQuery(req); err != nil {
				return err
			}
		}
	}
	return nil
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ResponseNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func ResponseCreated(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func ResponseBadRequest(c *gin.Context, err error) {
	ResponseErrorWithCode(c, http.StatusBadRequest, err)
}

func ResponseInternalServerError(c *gin.Context, err error) {
	if IsNotFoundDBError(err) {
		ResponseNotFound(c, err)
		return
	}
	ResponseErrorWithCode(c, http.StatusInternalServerError, err)
}

func ResponseUnauthorized(c *gin.Context, err error) {
	ResponseErrorWithCode(c, http.StatusUnauthorized, err)
}

func ResponseNotFound(c *gin.Context, err error) {
	if err == nil {
		err = errors.New("not found")
	}
	ResponseErrorWithCode(c, http.StatusNotFound, err)
}

func ResponseErrorWithCode(c *gin.Context, code int, err error) {
	c.JSON(code, NewError(err))
}

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func IsNotFoundDBError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, gorm.ErrRecordNotFound)
}

