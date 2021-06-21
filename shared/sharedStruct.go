package shared

type ConfigItem struct {
	Name     string
	Value    string
	Children []*ConfigItem
}
