package services

import (
	"github.com/edwardzhanged/novel-go/app/utils"
	"github.com/mojocn/base64Captcha"
)

type ResourceService interface {
	GetImgVerifyCode()
	VerifyImgAnswer()
}

type ResourceApi struct{}

func (resource *ResourceApi) GetImgVerifyCode() (id string, b64s string, err error) {
	var DriverMath = &base64Captcha.DriverMath{
		Height:          60,
		Width:           240,
		ShowLineOptions: 3,
		NoiseCount:      4,
	}

	driver := DriverMath.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, utils.Store)
	id, b64s, _, err = c.Generate()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil
}

func (resource *ResourceApi) VerifyImgAnswer(id string, answer string) bool {
	return utils.Store.Verify(id, answer, true)
}
