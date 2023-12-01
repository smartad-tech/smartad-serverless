package utils

func CategoryUuidToString(categoryUuid string) string {
	switch categoryUuid {
	case "c44453a1-8184-4905-a81e-04372f9b76e6":
		return "Adult Men"
	case "7186a646-17ec-4afb-a18f-104c34830eac":
		return "Adult Women"
	default:
		return "Invalid category"
	}
}
