package utils_test

import (
	"github.com/smartad-tech/smartad-serverless/internal/utils"
	"testing"
)

const AdultWomenUuid = "7186a646-17ec-4afb-a18f-104c34830eac"

func TestCategoryUuidToString(t *testing.T) {
	categoryName := utils.CategoryUuidToString(AdultWomenUuid)
	if categoryName != "Adult Women" {
		t.Fail()
	}
}
