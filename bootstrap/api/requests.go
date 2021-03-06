//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package api

import "github.com/mainflux/mainflux/bootstrap"

type apiReq interface {
	validate() error
}

type addReq struct {
	key         string
	ExternalID  string   `json:"external_id"`
	ExternalKey string   `json:"external_key"`
	Channels    []string `json:"channels"`
	Name        string   `json:"name"`
	Content     string   `json:"content"`
}

func (req addReq) validate() error {
	if req.ExternalID == "" || req.ExternalKey == "" {
		return bootstrap.ErrMalformedEntity
	}

	return nil
}

type entityReq struct {
	key string
	id  string
}

func (req entityReq) validate() error {
	if req.key == "" {
		return bootstrap.ErrUnauthorizedAccess
	}

	return nil
}

type updateReq struct {
	key      string
	id       string
	Channels []string        `json:"channels"`
	Name     string          `json:"name"`
	Content  string          `json:"content"`
	State    bootstrap.State `json:"state"`
}

func (req updateReq) validate() error {
	if req.key == "" {
		return bootstrap.ErrUnauthorizedAccess
	}

	// Can't explicitly update state to NewThing or Created.
	if req.State != bootstrap.Inactive &&
		req.State != bootstrap.Active {
		return bootstrap.ErrMalformedEntity
	}

	return nil
}

type listReq struct {
	key    string
	filter bootstrap.Filter
	offset uint64
	limit  uint64
}

func (req listReq) validate() error {
	if req.key == "" {
		return bootstrap.ErrUnauthorizedAccess
	}

	return nil
}

type bootstrapReq struct {
	key string
	id  string
}

func (req bootstrapReq) validate() error {
	if req.key == "" {
		return bootstrap.ErrUnauthorizedAccess
	}

	if req.id == "" {
		return bootstrap.ErrMalformedEntity
	}

	return nil
}

type changeStateReq struct {
	key   string
	id    string
	State bootstrap.State `json:"state"`
}

func (req changeStateReq) validate() error {
	if req.id == "" || req.key == "" {
		return bootstrap.ErrUnauthorizedAccess
	}

	if req.State != bootstrap.Inactive &&
		req.State != bootstrap.Active {
		return bootstrap.ErrMalformedEntity
	}

	return nil
}
