/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package db

import "github.com/superseriousbusiness/gotosocial/internal/gtsmodel"

// Status contains functions for getting statuses, creating statuses, and checking various other fields on statuses.
type Status interface {
	// GetStatusByID returns one status from the database, with all rel fields populated (if possible).
	GetStatusByID(id string) (*gtsmodel.Status, Error)

	// GetStatusByURI returns one status from the database, with all rel fields populated (if possible).
	GetStatusByURI(uri string) (*gtsmodel.Status, Error)

	// GetStatusByURL returns one status from the database, with all rel fields populated (if possible).
	GetStatusByURL(uri string) (*gtsmodel.Status, Error)

	// PutStatus stores one status in the database.
	PutStatus(status *gtsmodel.Status) Error

	// CountStatusReplies returns the amount of replies recorded for a status, or an error if something goes wrong
	CountStatusReplies(status *gtsmodel.Status) (int, Error)

	// CountStatusReblogs returns the amount of reblogs/boosts recorded for a status, or an error if something goes wrong
	CountStatusReblogs(status *gtsmodel.Status) (int, Error)

	// CountStatusFaves returns the amount of faves/likes recorded for a status, or an error if something goes wrong
	CountStatusFaves(status *gtsmodel.Status) (int, Error)

	// GetStatusParents gets the parent statuses of a given status.
	//
	// If onlyDirect is true, only the immediate parent will be returned.
	GetStatusParents(status *gtsmodel.Status, onlyDirect bool) ([]*gtsmodel.Status, Error)

	// GetStatusChildren gets the child statuses of a given status.
	//
	// If onlyDirect is true, only the immediate children will be returned.
	GetStatusChildren(status *gtsmodel.Status, onlyDirect bool, minID string) ([]*gtsmodel.Status, Error)

	// IsStatusFavedBy checks if a given status has been faved by a given account ID
	IsStatusFavedBy(status *gtsmodel.Status, accountID string) (bool, Error)

	// IsStatusRebloggedBy checks if a given status has been reblogged/boosted by a given account ID
	IsStatusRebloggedBy(status *gtsmodel.Status, accountID string) (bool, Error)

	// IsStatusMutedBy checks if a given status has been muted by a given account ID
	IsStatusMutedBy(status *gtsmodel.Status, accountID string) (bool, Error)

	// IsStatusBookmarkedBy checks if a given status has been bookmarked by a given account ID
	IsStatusBookmarkedBy(status *gtsmodel.Status, accountID string) (bool, Error)

	// GetStatusFaves returns a slice of faves/likes of the given status.
	// This slice will be unfiltered, not taking account of blocks and whatnot, so filter it before serving it back to a user.
	GetStatusFaves(status *gtsmodel.Status) ([]*gtsmodel.StatusFave, Error)

	// GetStatusReblogs returns a slice of statuses that are a boost/reblog of the given status.
	// This slice will be unfiltered, not taking account of blocks and whatnot, so filter it before serving it back to a user.
	GetStatusReblogs(status *gtsmodel.Status) ([]*gtsmodel.Status, Error)
}
