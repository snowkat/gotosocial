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

import "context"

// Basic wraps basic database functionality.
type Basic interface {
	// CreateTable creates a table for the given interface.
	// For implementations that don't use tables, this can just return nil.
	CreateTable(i interface{}) Error

	// DropTable drops the table for the given interface.
	// For implementations that don't use tables, this can just return nil.
	DropTable(i interface{}) Error

	// RegisterTable registers a table for use in many2many relations.
	// For implementations that don't use tables, or many2many relations, this can just return nil.
	RegisterTable(i interface{}) Error

	// Stop should stop and close the database connection cleanly, returning an error if this is not possible.
	// If the database implementation doesn't need to be stopped, this can just return nil.
	Stop(ctx context.Context) Error

	// IsHealthy should return nil if the database connection is healthy, or an error if not.
	IsHealthy(ctx context.Context) Error

	// GetByID gets one entry by its id. In a database like postgres, this might be the 'id' field of the entry,
	// for other implementations (for example, in-memory) it might just be the key of a map.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	// In case of no entries, a 'no entries' error will be returned
	GetByID(id string, i interface{}) Error

	// GetWhere gets one entry where key = value. This is similar to GetByID but allows the caller to specify the
	// name of the key to select from.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	// In case of no entries, a 'no entries' error will be returned
	GetWhere(where []Where, i interface{}) Error

	// GetAll will try to get all entries of type i.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	// In case of no entries, a 'no entries' error will be returned
	GetAll(i interface{}) Error

	// Put simply stores i. It is up to the implementation to figure out how to store it, and using what key.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	Put(i interface{}) Error

	// Upsert stores or updates i based on the given conflict column, as in https://www.postgresqltutorial.com/postgresql-upsert/
	// It is up to the implementation to figure out how to store it, and using what key.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	Upsert(i interface{}, conflictColumn string) Error

	// UpdateByID updates i with id id.
	// The given interface i will be set to the result of the query, whatever it is. Use a pointer or a slice.
	UpdateByID(id string, i interface{}) Error

	// UpdateOneByID updates interface i with database the given database id. It will update one field of key key and value value.
	UpdateOneByID(id string, key string, value interface{}, i interface{}) Error

	// UpdateWhere updates column key of interface i with the given value, where the given parameters apply.
	UpdateWhere(where []Where, key string, value interface{}, i interface{}) Error

	// DeleteByID removes i with id id.
	// If i didn't exist anyway, then no error should be returned.
	DeleteByID(id string, i interface{}) Error

	// DeleteWhere deletes i where key = value
	// If i didn't exist anyway, then no error should be returned.
	DeleteWhere(where []Where, i interface{}) Error
}
