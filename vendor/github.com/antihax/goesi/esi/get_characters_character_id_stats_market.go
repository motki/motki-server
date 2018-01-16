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

/* A list of GetCharactersCharacterIdStatsMarket. */
//easyjson:json
type GetCharactersCharacterIdStatsMarketList []GetCharactersCharacterIdStatsMarket

/* market object */
//easyjson:json
type GetCharactersCharacterIdStatsMarket struct {
	AcceptContractsCourier      int64 `json:"accept_contracts_courier,omitempty"`       /* accept_contracts_courier integer */
	AcceptContractsItemExchange int64 `json:"accept_contracts_item_exchange,omitempty"` /* accept_contracts_item_exchange integer */
	BuyOrdersPlaced             int64 `json:"buy_orders_placed,omitempty"`              /* buy_orders_placed integer */
	CancelMarketOrder           int64 `json:"cancel_market_order,omitempty"`            /* cancel_market_order integer */
	CreateContractsAuction      int64 `json:"create_contracts_auction,omitempty"`       /* create_contracts_auction integer */
	CreateContractsCourier      int64 `json:"create_contracts_courier,omitempty"`       /* create_contracts_courier integer */
	CreateContractsItemExchange int64 `json:"create_contracts_item_exchange,omitempty"` /* create_contracts_item_exchange integer */
	DeliverCourierContract      int64 `json:"deliver_courier_contract,omitempty"`       /* deliver_courier_contract integer */
	IskGained                   int64 `json:"isk_gained,omitempty"`                     /* isk_gained integer */
	IskSpent                    int64 `json:"isk_spent,omitempty"`                      /* isk_spent integer */
	ModifyMarketOrder           int64 `json:"modify_market_order,omitempty"`            /* modify_market_order integer */
	SearchContracts             int64 `json:"search_contracts,omitempty"`               /* search_contracts integer */
	SellOrdersPlaced            int64 `json:"sell_orders_placed,omitempty"`             /* sell_orders_placed integer */
}
