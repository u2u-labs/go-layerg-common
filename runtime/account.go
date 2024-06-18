package runtime

type AccountUpdate struct {
	UserID      string
	Username    string
	Metadata    map[string]interface{}
	DisplayName string
	Timezone    string
	Location    string
	LangTag     string
	AvatarUrl   string
}
