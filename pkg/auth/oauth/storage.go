// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package oauth

import (
	"fmt"
	"time"

	"github.com/RangelReale/osin"
	"github.com/TheThingsNetwork/ttn/pkg/auth"
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/store"
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/types"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
)

type storage struct {
	keys  *auth.Keys
	store store.Store
}

// UserData is the userdata that gets carried around with authorization requests.
type UserData struct {
	UserID string
}

// Clone the store if needed.
func (s *storage) Clone() osin.Storage {
	return s
}

// Close the store, releasing resources it might be holding.
func (s *storage) Close() {}

// GetClient loads the OAuth Client by client_id.
func (s *storage) GetClient(clientID string) (osin.Client, error) {
	client, err := s.store.Clients.GetByID(clientID)
	if err != nil {
		return nil, err
	}

	if client.GetClient().State != ttnpb.STATE_APPROVED {
		return nil, nil
	}

	return client, nil
}

// SaveAuthorize saves authorization data.
func (s *storage) SaveAuthorize(data *osin.AuthorizeData) error {
	userID := ""
	udata, ok := data.UserData.(*UserData)
	if ok && udata != nil {
		userID = udata.UserID
	}

	return s.store.OAuth.SaveAuthorizationCode(&types.AuthorizationData{
		AuthorizationCode: data.Code,
		ClientID:          data.Client.GetId(),
		CreatedAt:         data.CreatedAt,
		ExpiresIn:         time.Duration(data.ExpiresIn) * time.Second,
		Scope:             data.Scope,
		RedirectURI:       data.RedirectUri,
		State:             data.State,
		UserID:            userID,
	})
}

// LoadAuthorize loads the client and authorization data for the authorization code.
func (s *storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	data, err := s.store.OAuth.GetAuthorizationCode(code)
	if err != nil {
		return nil, err
	}

	client, err := s.store.Clients.GetByID(data.ClientID)
	if err != nil {
		return nil, err
	}

	// ensure the expiration
	if exp := data.CreatedAt.Add(data.ExpiresIn); exp.Before(time.Now()) {
		return nil, fmt.Errorf("Authorization code is expired by %v", time.Now().Sub(exp))
	}

	return &osin.AuthorizeData{
		Code:        data.AuthorizationCode,
		Client:      client,
		ExpiresIn:   int32(data.ExpiresIn.Seconds()),
		Scope:       data.Scope,
		RedirectUri: data.RedirectURI,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData: &UserData{
			UserID: data.UserID,
		},
	}, nil
}

// RemoveAuthorize deletes the authorization code.
func (s *storage) RemoveAuthorize(code string) error {
	return s.store.OAuth.DeleteAuthorizationCode(code)
}

// SaveAccess saves the access data for later use.
func (s *storage) SaveAccess(data *osin.AccessData) error {
	if data.RefreshToken != "" {
		err := s.store.OAuth.SaveRefreshToken(&types.RefreshData{
			RefreshToken: data.RefreshToken,
			ClientID:     data.Client.GetId(),
			Scope:        data.Scope,
			CreatedAt:    data.CreatedAt,
			RedirectURI:  data.RedirectUri,
		})

		return err
	}

	return nil
}

// LoadAccess loads the access data based on the access token.
func (s *storage) LoadAccess(accessToken string) (*osin.AccessData, error) {
	claims, err := auth.ClaimsFromToken(s.keys, accessToken)
	if err != nil {
		return nil, err
	}

	client, err := s.store.Clients.GetByID(claims.Client)
	if err != nil {
		return nil, err
	}

	return &osin.AccessData{
		Client:    client,
		ExpiresIn: int32(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()).Seconds()),
		Scope:     Scope(claims.Rights),
		CreatedAt: time.Unix(claims.IssuedAt, 0),
		UserData: &UserData{
			UserID: claims.Creator,
		},
	}, nil
}

// RemoveAccess revokes access data.
func (s *storage) RemoveAccess(accessToken string) error {
	// Access tokens aren't saved in the database, so do nothing.
	return nil
}

// LoadRefresh loads the access data based on the refresh token.
func (s *storage) LoadRefresh(refreshToken string) (*osin.AccessData, error) {
	data, err := s.store.OAuth.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	client, err := s.store.Clients.GetByID(data.ClientID)
	if err != nil {
		return nil, err
	}

	return &osin.AccessData{
		RefreshToken: data.RefreshToken,
		Client:       client,
		Scope:        data.Scope,
		CreatedAt:    data.CreatedAt,
		RedirectUri:  data.RedirectURI,
	}, nil
}

// RemoveRefresh deletes the refresh token.
func (s *storage) RemoveRefresh(refreshToken string) error {
	return s.store.OAuth.DeleteRefreshToken(refreshToken)
}
