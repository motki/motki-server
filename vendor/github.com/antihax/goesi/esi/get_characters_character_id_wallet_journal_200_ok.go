/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.5
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdWalletJournal200Ok. */
//easyjson:json
type GetCharactersCharacterIdWalletJournal200OkList []GetCharactersCharacterIdWalletJournal200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdWalletJournal200Ok struct {
	Date            time.Time                                      `json:"date,omitempty"`              /* Date and time of transaction */
	RefId           int64                                          `json:"ref_id,omitempty"`            /* Unique journal reference ID */
	RefType         string                                         `json:"ref_type,omitempty"`          /* Transaction type, different type of transaction will populate different fields in `extra_info` Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can use the following gist to see string->int mappings: https://gist.github.com/ccp-zoetrope/c03db66d90c2148724c06171bc52e0ec */
	FirstPartyId    int32                                          `json:"first_party_id,omitempty"`    /* first_party_id integer */
	FirstPartyType  string                                         `json:"first_party_type,omitempty"`  /* first_party_type string */
	SecondPartyId   int32                                          `json:"second_party_id,omitempty"`   /* second_party_id integer */
	SecondPartyType string                                         `json:"second_party_type,omitempty"` /* second_party_type string */
	Amount          float64                                        `json:"amount,omitempty"`            /* Transaction amount. Positive when value transferred to the first party. Negative otherwise */
	Balance         float64                                        `json:"balance,omitempty"`           /* Wallet balance after transaction occurred */
	Reason          string                                         `json:"reason,omitempty"`            /* reason string */
	TaxReceiverId   int32                                          `json:"tax_receiver_id,omitempty"`   /* the corporation ID receiving any tax paid */
	Tax             float64                                        `json:"tax,omitempty"`               /* Tax amount received for tax related transactions */
	ExtraInfo       GetCharactersCharacterIdWalletJournalExtraInfo `json:"extra_info,omitempty"`
}
