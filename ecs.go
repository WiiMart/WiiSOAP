//	Copyright (C) 2018-2020 CornierKhan1
//
//	WiiSOAP is SOAP Server Software, designed specifically to handle Wii Shop Channel SOAP.
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU Affero General Public License as published
//    by the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU Affero General Public License for more details.
//
//    You should have received a copy of the GNU Affero General Public License
//    along with this program.  If not, see http://www.gnu.org/licenses/.

package main

import (
	"errors"
	"fmt"
	"github.com/wii-tools/wadlib"
	"log"
)

const (
	QueryOwnedTitles = `SELECT owned_titles.title_id, tickets.version
		FROM owned_titles, tickets
		WHERE owned_titles.title_id = tickets.title_id
		AND owned_titles.account_id = $1`

	QueryTicketStatement = `SELECT ticket, version FROM tickets WHERE title_id = $1`

	AssociateTicketStatement = `INSERT INTO owned_titles (account_id, title_id, version)
		VALUES ($1, $2, $3)`

	// SharedBalanceAmount describes the maximum signed 32-bit integer value.
	// It is not an actual tracked points value, but exists to permit reuse.
	SharedBalanceAmount = 2147483647
)

func getBalance() Balance {
	return Balance{
		Amount:   SharedBalanceAmount,
		Currency: "POINTS",
	}
}

func checkDeviceStatus(e *Envelope) {
	e.AddCustomType(getBalance())
	e.AddKVNode("ForceSyncTime", "0")
	e.AddKVNode("ExtTicketTime", e.Timestamp())
	e.AddKVNode("SyncTime", e.Timestamp())
}

func notifyETicketsSynced(e *Envelope) {
	// TODO: Implement handling of synchronization timing
}

func listETickets(e *Envelope) {
	accountId, err := e.AccountId()
	if err != nil {
		e.Error(2, "missing account ID", err)
		return
	}

	rows, err := pool.Query(ctx, QueryOwnedTitles, accountId)
	if err != nil {
		log.Printf("error executing statement: %v\n", err)
		e.Error(2, "database error", errors.New("failed to execute db operation"))
		return
	}

	// Add all available titles for this account.
	defer rows.Close()
	for rows.Next() {
		var titleId string
		var version int
		err = rows.Scan(&titleId, &version)
		if err != nil {
			log.Printf("error executing statement: %v\n", err)
			e.Error(2, "database error", errors.New("failed to execute db operation"))
			return
		}

		e.AddCustomType(Tickets{
			TitleId: titleId,
			Version: version,

			// We do not support migration, ticket IDs, or revocation.
			TicketId:     "0",
			RevokeDate:   0,
			MigrateCount: 0,
			MigrateLimit: 0,
		})
	}

	e.AddKVNode("ForceSyncTime", "0")
	e.AddKVNode("ExtTicketTime", e.Timestamp())
	e.AddKVNode("SyncTime", e.Timestamp())
}

func getETickets(e *Envelope) {
	e.AddKVNode("ForceSyncTime", "0")
	e.AddKVNode("ExtTicketTime", e.Timestamp())
	e.AddKVNode("SyncTime", e.Timestamp())
}

func purchaseTitle(e *Envelope) {
	//accountId, err := e.AccountId()
	//if err != nil {
	//	e.Error(2, "missing account ID", err)
	//	return
	//}

	// Determine the title ID we're going to purchase.
	titleId, err := e.getKey("TitleId")
	if err != nil {
		e.Error(2, "missing account ID", err)
		return
	}

	// Query the ticket and current version for this title.
	var ticket []byte
	var version int
	row := pool.QueryRow(ctx, QueryTicketStatement, titleId)

	err = row.Scan(&ticket, &version)
	if err != nil {
		log.Printf("unexpected error purchasing: %v", err)
		// TODO(spotlightishere): Can we more elegantly return an error when a title may not exist here?
		e.Error(2, "error purchasing", nil)
	}

	// Associate the given title ID with the user.
	//_, err = pool.Exec(ctx, AssociateTicketStatement, accountId, titleId, version)
	//if err != nil {
	//	log.Printf("unexpected error purchasing: %v", err)
	//	e.Error(2, "error purchasing", nil)
	//}

	// The returned ticket is expected to have two other certificates associated.
	ticketString := b64(append(ticket, wadlib.CertChainTemplate...))

	e.AddCustomType(getBalance())
	e.AddCustomType(Transactions{
		TransactionId: "00000000",
		Date:          e.Timestamp(),
		Type:          "PURCHGAME",
		TotalPaid:     0,
		Currency:      "POINTS",
		ItemId:        0,
	})
	e.AddKVNode("SyncTime", e.Timestamp())

	//// Two cert types must be present.
	//type Certs struct {
	//	XMLName xml.Name `xml:"Certs"`
	//	Value   string   `xml:",chardata"`
	//}

	e.AddKVNode("ETickets", ticketString)
	e.AddKVNode("Certs", b64(wadlib.CertChainTemplate))
	e.AddKVNode("Certs", b64(wadlib.CertChainTemplate))
	e.AddKVNode("TitleId", titleId)
}

func listPurchaseHistory(e *Envelope) {
	e.AddCustomType([]Transactions{
		{
			TransactionId: "12345678",
			Date:          e.Timestamp(),
			Type:          "SERVICE",
			TotalPaid:     7,
			Currency:      "POINTS",
			ItemId:        0,
			TitleId:       "000100014843494A",
			ItemPricing: []Limits{
				LimitStruct(DR),
			},
			ReferenceId:    1,
			ReferenceValue: 19224,
		},
	})

	e.AddKVNode("ListResultTotalSize", "1")
}

// genServiceUrl returns a URL with the given service against a configured URL.
// Given a baseUrl of example.com and genServiceUrl("ias", "IdentityAuthenticationSOAP"),
// it would return http://ias.example.com/ias/services/ias/IdentityAuthenticationSOAP.
func genServiceUrl(service string, path string) string {
	return fmt.Sprintf("http://%s.%s/%s/services/%s", service, baseUrl, service, path)
}

func getECConfig(e *Envelope) {
	contentUrl := fmt.Sprintf("http://ccs.%s/ccs/download", baseUrl)
	e.AddKVNode("ContentPrefixURL", contentUrl)
	e.AddKVNode("UncachedContentPrefixURL", contentUrl)
	e.AddKVNode("SystemContentPrefixURL", contentUrl)
	e.AddKVNode("SystemUncachedContentPrefixURL", contentUrl)

	e.AddKVNode("EcsURL", genServiceUrl("ecs", "ECommerceSOAP"))
	e.AddKVNode("IasURL", genServiceUrl("ias", "IdentityAuthenticationSOAP"))
	e.AddKVNode("CasURL", genServiceUrl("cas", "CatalogingSOAP"))
	e.AddKVNode("NusURL", genServiceUrl("nus", "NetUpdateSOAP"))
}
