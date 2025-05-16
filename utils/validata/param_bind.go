package validata

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
	})
}

type ValidataResponse struct {
	Field map[string]any `json:"field"`
	Msg   string         `json:"msg"`
}

func ValidateErr(err error) (resp ValidataResponse) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		resp.Msg = err.Error()
		return
	}
	m := make(map[string]any)
	var msgs []string

	for _, e := range errs {
		msg := e.Translate(trans)
		_list := strings.Split(msg, "---")
		m[_list[0]] = _list[1]
		msgs = append(msgs, _list[1])
	}
	resp.Field = m
	resp.Msg = strings.Join(msgs, ";")
	return
}
