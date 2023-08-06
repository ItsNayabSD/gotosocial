// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gtsmodel

import (
	"context"
	"time"
)

// Mention refers to the 'tagging' or 'mention' of a user within a status.
type Mention struct {
	ID               string    `bun:"type:CHAR(26),pk,nullzero,notnull,unique"`                    // id of this item in the database
	CreatedAt        time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item created
	UpdatedAt        time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item last updated
	StatusID         string    `bun:"type:CHAR(26),nullzero,notnull"`                              // ID of the status this mention originates from
	Status           *Status   `bun:"rel:belongs-to"`                                              // status referred to by statusID
	OriginAccountID  string    `bun:"type:CHAR(26),nullzero,notnull"`                              // ID of the mention creator account
	OriginAccountURI string    `bun:",nullzero,notnull"`                                           // ActivityPub URI of the originator/creator of the mention
	OriginAccount    *Account  `bun:"rel:belongs-to"`                                              // account referred to by originAccountID
	TargetAccountID  string    `bun:"type:CHAR(26),nullzero,notnull"`                              // Mention target/receiver account ID
	TargetAccount    *Account  `bun:"rel:belongs-to"`                                              // account referred to by targetAccountID
	Silent           *bool     `bun:",nullzero,notnull,default:false"`                             // Prevent this mention from generating a notification?

	/*
		NON-DATABASE CONVENIENCE FIELDS
		These fields are just for convenience while passing the mention
		around internally, to make fewer database calls and whatnot. They're
		not meant to be put in the database!
	*/

	// NameString is for putting in the namestring of the mentioned user
	// before the mention is dereferenced. Should be in a form along the lines of:
	// @whatever_username@example.org
	//
	// This will not be put in the database, it's just for convenience.
	NameString string `bun:"-"`
	// TargetAccountURI is the AP ID (uri) of the user mentioned.
	//
	// This will not be put in the database, it's just for convenience.
	TargetAccountURI string `bun:"-"`
	// TargetAccountURL is the web url of the user mentioned.
	//
	// This will not be put in the database, it's just for convenience.
	TargetAccountURL string `bun:"-"`
	// A pointer to the gtsmodel account of the mentioned account.
}

// ParseMentionFunc describes a function that takes a lowercase account name
// in the form "@test@whatever.example.org" for a remote account, or "@test"
// for a local account, and returns a fully populated mention for that account,
// with the given origin status ID and origin account ID.
//
// If the account is remote and not yet found in the database, then ParseMentionFunc
// will try to webfinger the remote account and put it in the database before returning.
//
// Mentions generated by this function are not put in the database, that's still up to
// the caller to do.
type ParseMentionFunc func(ctx context.Context, targetAccount string, originAccountID string, statusID string) (*Mention, error)
