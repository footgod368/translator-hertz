package youdao

import (
	"context"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		resp, err := Query(context.Background(), "simple")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	})
}
