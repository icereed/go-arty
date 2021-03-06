// Copyright (c) 2016 John E. Vincent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright (c) 2018 Target Brands, Inc.

package artifactory

import (
	"fmt"
)

// UsersService handles communication with the user related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SECURITY
type UsersService service

// SecurityUser represents a security user in Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Security+Configuration+JSON#SecurityConfigurationJSON-application/vnd.org.jfrog.artifactory.security.User+json
type SecurityUser struct {
	Name                     *string   `json:"name,omitempty"`
	Email                    *string   `json:"email,omitempty"`
	Password                 *string   `json:"password,omitempty"`
	Admin                    *bool     `json:"admin,omitempty"`
	ProfileUpdatable         *bool     `json:"profileUpdatable,omitempty"`
	DisableUIAccess          *bool     `json:"disableUIAccess,omitempty"`
	InternalPasswordDisabled *bool     `json:"internalPasswordDisabled,omitempty"`
	LastLoggedIn             *string   `json:"lastLoggedIn,omitempty"`
	Realm                    *string   `json:"realm,omitempty"`
	Groups                   *[]string `json:"groups,omitempty"`
}

func (u SecurityUser) String() string {
	return Stringify(u)
}

// User represents a user in Artifactory.
//
// Docs: This struct is currently undocumented by JFrog
type User struct {
	Name                     *string   `json:"name,omitempty"`
	Email                    *string   `json:"email,omitempty"`
	Admin                    *bool     `json:"admin,omitempty"`
	GroupAdmin               *bool     `json:"groupAdmin,omitempty"`
	ProfileUpdatable         *bool     `json:"profileUpdatable,omitempty"`
	InternalPasswordDisabled *bool     `json:"internalPasswordDisabled,omitempty"`
	Groups                   *[]string `json:"groups,omitempty"`
	LastLoggedIn             *string   `json:"lastLoggedIn,omitempty"`
	LastLoggedInMillis       *int64    `json:"lastLoggedInMillis,omitempty"`
	Realm                    *string   `json:"realm,omitempty"`
	OfflineMode              *bool     `json:"offlineMode,omitempty"`
	DisableUIAccess          *bool     `json:"disableUIAccess,omitempty"`
	ProWithoutLicense        *bool     `json:"proWithoutLicense,omitempty"`
	ExternalRealmLink        *string   `json:"externalRealmLink,omitempty"`
	ExistsInDB               *bool     `json:"existsInDB,omitempty"`
	HideUploads              *bool     `json:"hideUploads,omitempty"`
	RequireProfileUnlock     *bool     `json:"requireProfileUnlock,omitempty"`
	RequireProfilePassword   *bool     `json:"requireProfilePassword,omitempty"`
	Locked                   *bool     `json:"locked,omitempty"`
	CredentialsExpired       *bool     `json:"credentialsExpired,omitempty"`
	NumberOfGroups           *int      `json:"numberOfGroups,omitempty"`
	NumberOfPermissions      *int      `json:"numberOfPermissions,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// APIKey represents an api key in Artifactory.
type APIKey struct {
	APIKey *string `json:"apiKey,omitempty"`
}

func (a APIKey) String() string {
	return Stringify(a)
}

// DeleteAPIKey represents a response from deleting an API key in Artifactory
type DeleteAPIKey struct {
	Info *string `json:"info,omitempty"`
}

func (d DeleteAPIKey) String() string {
	return Stringify(d)
}

// GetAll returns a list of all users.
//
// Docs: This endpoint is currently undocumented by JFrog
func (s *UsersService) GetAll() (*[]User, *Response, error) {
	u := fmt.Sprintf("/api/users")
	v := new([]User)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetAllSecurity returns a list of all users.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetUsers
func (s *UsersService) GetAllSecurity() (*[]SecurityUser, *Response, error) {
	u := fmt.Sprintf("/api/security/users")
	v := new([]SecurityUser)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetSecurity returns the provided user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetUserDetails
func (s *UsersService) GetSecurity(user string) (*SecurityUser, *Response, error) {
	u := fmt.Sprintf("/api/security/users/%s", user)
	v := new(SecurityUser)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// CreateSecurity constructs a user with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceUser
func (s *UsersService) CreateSecurity(user *SecurityUser) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/users/%s", *user.Name)
	v := new(string)

	resp, err := s.client.Call("PUT", u, user, v)
	return v, resp, err
}

// UpdateSecurity modifies a user with the provided details.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateUser
func (s *UsersService) UpdateSecurity(user *SecurityUser) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/users/%s", *user.Name)
	v := new(string)

	resp, err := s.client.Call("POST", u, user, v)
	return v, resp, err
}

// DeleteSecurity removes the provided user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-DeleteUser
func (s *UsersService) DeleteSecurity(user string) (*string, *Response, error) {
	u := fmt.Sprintf("/api/security/users/%s", user)
	v := new(string)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// GetAPIKey returns the api key of the authenticated user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetAPIKey
func (s *UsersService) GetAPIKey() (*APIKey, *Response, error) {
	u := "/api/security/apiKey"
	v := new(APIKey)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// CreateAPIKey constructs an api key for the authenticated user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-CreateAPIKey
func (s *UsersService) CreateAPIKey() (*APIKey, *Response, error) {
	u := "/api/security/apiKey"
	v := new(APIKey)

	resp, err := s.client.Call("POST", u, nil, v)
	return v, resp, err
}

// RegenerateAPIKey recreates an api key for the authenticated user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RegenerateAPIKey
func (s *UsersService) RegenerateAPIKey() (*APIKey, *Response, error) {
	u := "/api/security/apiKey"
	v := new(APIKey)

	resp, err := s.client.Call("PUT", u, nil, v)
	return v, resp, err
}

// DeleteAPIKey removes an api key for the authenticated user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeAPIKey
func (s *UsersService) DeleteAPIKey() (*DeleteAPIKey, *Response, error) {
	u := "/api/security/apiKey"
	v := new(DeleteAPIKey)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// DeleteUserAPIKey removes an api key for the provided user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeUserAPIKey
func (s *UsersService) DeleteUserAPIKey(user string) (*DeleteAPIKey, *Response, error) {
	u := fmt.Sprintf("/api/security/apiKey/%s", user)
	v := new(DeleteAPIKey)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// DeleteAllAPIKeys removes all api keys.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeAllAPIKeys
func (s *UsersService) DeleteAllAPIKeys() (*DeleteAPIKey, *Response, error) {
	u := "/api/security/apiKey?deleteAll=1"
	v := new(DeleteAPIKey)

	resp, err := s.client.Call("DELETE", u, nil, v)
	return v, resp, err
}

// GetEncryptedPassword returns the encrypted password of the authenticated user.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetUserEncryptedPassword
func (s *UsersService) GetEncryptedPassword() (*string, *Response, error) {
	u := "/api/security/encryptedPassword"
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}
