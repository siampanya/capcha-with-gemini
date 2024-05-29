package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
)

func Gemini(ctx context.Context, url string) {

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("APIKEY")))

	if err != nil {
		panic(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")

	capcha, err := GetFilesFromhttpImages(url)
	if err != nil {
		panic(err)
	}

	prompt := []genai.Part{
		genai.ImageData("jpeg", capcha),
		genai.Text("What is number, read captcha, show JSON pattern {code:xxxx}"),
	}

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Candidates[0].Content.Parts[0])

}
