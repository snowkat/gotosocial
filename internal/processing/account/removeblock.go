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

package account

import (
	"fmt"

	apimodel "github.com/superseriousbusiness/gotosocial/internal/api/model"
	"github.com/superseriousbusiness/gotosocial/internal/db"
	"github.com/superseriousbusiness/gotosocial/internal/gtserror"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
)

func (p *processor) BlockRemove(requestingAccount *gtsmodel.Account, targetAccountID string) (*apimodel.Relationship, gtserror.WithCode) {
	// make sure the target account actually exists in our db
	targetAccount, err := p.db.GetAccountByID(targetAccountID)
	if err != nil {
		return nil, gtserror.NewErrorNotFound(fmt.Errorf("BlockCreate: error getting account %s from the db: %s", targetAccountID, err))
	}

	// check if a block exists, and remove it if it does (storing the URI for later)
	var blockChanged bool
	block := &gtsmodel.Block{}
	if err := p.db.GetWhere([]db.Where{
		{Key: "account_id", Value: requestingAccount.ID},
		{Key: "target_account_id", Value: targetAccountID},
	}, block); err == nil {
		block.Account = requestingAccount
		block.TargetAccount = targetAccount
		if err := p.db.DeleteByID(block.ID, &gtsmodel.Block{}); err != nil {
			return nil, gtserror.NewErrorInternalError(fmt.Errorf("BlockRemove: error removing block from db: %s", err))
		}
		blockChanged = true
	}

	// block status changed so send the UNDO activity to the channel for async processing
	if blockChanged {
		p.fromClientAPI <- gtsmodel.FromClientAPI{
			APObjectType:   gtsmodel.ActivityStreamsBlock,
			APActivityType: gtsmodel.ActivityStreamsUndo,
			GTSModel:       block,
			OriginAccount:  requestingAccount,
			TargetAccount:  targetAccount,
		}
	}

	// return whatever relationship results from all this
	return p.RelationshipGet(requestingAccount, targetAccountID)
}
