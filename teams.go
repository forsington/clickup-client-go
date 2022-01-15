// Copyright (c) 2022, John Moore
// All rights reserved.

// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clickup

import (
	"context"
	"fmt"
	"net/http"
)

type Team struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Color   string       `json:"color"`
	Avatar  string       `json:"avatar"`
	Members []TeamMember `json:"members"`
}

type TeamsResponse struct {
	Teams []Team `json:"teams"`
}

type TeamUser struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profilePicture"`
	Initials       string `json:"initials"`
	Role           int    `json:"role"`
	CustomRole     string `json:"custom_role"`
	LastActive     string `json:"last_active"`
	DateJoined     string `json:"date_joined"`
	DateInvited    string `json:"date_invited"`
}
type TeamMember struct {
	User      TeamUser `json:"user"`
	InvitedBy TeamUser `json:"invited_by"`
}

// Teams returns a listing of teams for the authenticated user (the Client).
func (c *Client) Teams(ctx context.Context) (*TeamsResponse, error) {
	endpoint := fmt.Sprintf("/team")

	var teamsResponse TeamsResponse

	if err := c.call(ctx, http.MethodGet, endpoint, nil, &teamsResponse); err != nil {
		return nil, fmt.Errorf("failed to make clickup request: %w", err)
	}

	return &teamsResponse, nil
}
