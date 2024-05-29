package app

import (
	"context"
	"fmt"

	"github.com/siampanya/capcha-with-gemini/internal/service"
)

func App(ctx context.Context) interface{} {

	var url string
	fmt.Print("Enter captcha url: ")
	fmt.Scanf("%s", &url)

	service.Gemini(ctx, url)

	return App(ctx)
}
