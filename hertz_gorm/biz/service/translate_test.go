package service

import (
	"context"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
	"testing"
)

func TestTranslate(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		resp, err := Translate(context.Background(), translator.TranslateRequest{Text: "insist"})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	})
}
