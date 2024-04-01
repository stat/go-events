package validator

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	InitializeError = errors.New("Validator could not initialize")
)

func Initialize(engine *gin.Engine) error {
	v := binding.Validator.Engine()

	validator, ok := v.(*validator.Validate)

	if !ok {
		return InitializeError
	}

	validator.SetTagName("validate")

	return nil
}
