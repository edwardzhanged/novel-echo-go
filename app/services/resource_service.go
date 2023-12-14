package services

import "github.com/mojocn/base64Captcha"

type ResourceService interface {
	GetImgVerifyCode()
}

type ResourceApi struct{}

func (resource *ResourceApi) GetImgVerifyCode() (id string, b64s string, ans string, err error) {
	var DriverMath = &base64Captcha.DriverMath{
		Height:          60,
		Width:           240,
		ShowLineOptions: 3,
		NoiseCount:      4,
	}
	store := base64Captcha.DefaultMemStore

	driver := DriverMath.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err = c.Generate()
	if err != nil {
		return "", "", "", err
	}
	return id, b64s, ans, nil
}
