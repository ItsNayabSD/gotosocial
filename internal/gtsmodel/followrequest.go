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

import "time"

// FollowRequest represents one account requesting to follow another, and the metadata around that request.
type FollowRequest struct {
	ID              string    `bun:"type:CHAR(26),pk,nullzero,notnull,unique"`                    // id of this item in the database
	CreatedAt       time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item created
	UpdatedAt       time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item last updated
	URI             string    `bun:",notnull,nullzero,unique"`                                    // ActivityPub uri of this follow (request).
	AccountID       string    `bun:"type:CHAR(26),unique:frsrctarget,notnull,nullzero"`           // Who does this follow request originate from?
	Account         *Account  `bun:"rel:belongs-to"`                                              // Account corresponding to accountID
	TargetAccountID string    `bun:"type:CHAR(26),unique:frsrctarget,notnull,nullzero"`           // Who is the target of this follow request?
	TargetAccount   *Account  `bun:"rel:belongs-to"`                                              // Account corresponding to targetAccountID
	ShowReblogs     *bool     `bun:",nullzero,notnull,default:true"`                              // Does this follow also want to see reblogs and not just posts?
	Notify          *bool     `bun:",nullzero,notnull,default:false"`                             // does the following account want to be notified when the followed account posts?
}
