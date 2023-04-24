package organization

import "time"

type OrganizationPolicy struct {
}

type Organization struct {
	name    string
	founder []Member
	members []Member
	policy  OrganizationPolicy
}

type Role struct {
	name string
}

type Member struct {
	name string
	role Role
}

type Thread struct {
	organization Organization
	data_started time.Time
	members      []Member
}

type ChatMember struct {
	sender    Member
	receiver  Member
	time_send time.Time
	time_read time.Time
}

type ChatOrganization struct {
	sender    Member
	receiver  Organization
	time_send time.Time
}

type ChatThread struct {
	sender   Member
	receiver Thread
}

type ChatContent struct {
	message string
	files   []File
	links   []Link
}

type File struct {
	name string
	path string
}

type Link struct {
	name string
	path string
}
