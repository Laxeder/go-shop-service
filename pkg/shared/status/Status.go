package status

type Status string

const (
	Undefined Status = ""
	Enabled   Status = "enabled"
	Disabled  Status = "disabled"
)
