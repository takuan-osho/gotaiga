package cmd

import "time"

type NormalLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UserAuthDetail struct {
	ID                            int         `json:"id"`
	Username                      string      `json:"username"`
	FullName                      string      `json:"full_name"`
	FullNameDisplay               string      `json:"full_name_display"`
	Color                         string      `json:"color"`
	Bio                           string      `json:"bio"`
	Lang                          string      `json:"lang"`
	Theme                         string      `json:"theme"`
	Timezone                      string      `json:"timezone"`
	IsActive                      bool        `json:"is_active"`
	Photo                         string      `json:"photo"`
	BigPhoto                      string      `json:"big_photo"`
	GravatarID                    string      `json:"gravatar_id"`
	Roles                         []string    `json:"roles"`
	TotalPrivateProjects          int         `json:"total_private_projects"`
	TotalPublicProjects           int         `json:"total_public_projects"`
	Email                         string      `json:"email"`
	UUID                          string      `json:"uuid"`
	DateJoined                    time.Time   `json:"date_joined"`
	ReadNewTerms                  bool        `json:"read_new_terms"`
	AcceptedTerms                 bool        `json:"accepted_terms"`
	MaxPrivateProjects            interface{} `json:"max_private_projects"`
	MaxPublicProjects             interface{} `json:"max_public_projects"`
	MaxMembershipsPrivateProjects interface{} `json:"max_memberships_private_projects"`
	MaxMembershipsPublicProjects  interface{} `json:"max_memberships_public_projects"`
	AuthToken                     string      `json:"auth_token"`
}
