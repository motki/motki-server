/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.2
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

/* A list of GetCorporationsCorporationIdStarbases200Ok. */
//easyjson:json
type GetCorporationsCorporationIdStarbases200OkList []GetCorporationsCorporationIdStarbases200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdStarbases200Ok struct {
	MoonId          int32     `json:"moon_id,omitempty"`          /* The moon this starbase (POS) is anchored on, unanchored POSes do not have this information */
	OnlinedSince    time.Time `json:"onlined_since,omitempty"`    /* When the POS onlined, for starbases (POSes) in online state */
	ReinforcedUntil time.Time `json:"reinforced_until,omitempty"` /* When the POS will be out of reinforcement, for starbases (POSes) in reinforced state */
	StarbaseId      int64     `json:"starbase_id,omitempty"`      /* Unique ID for this starbase (POS) */
	State           string    `json:"state,omitempty"`            /* state string */
	SystemId        int32     `json:"system_id,omitempty"`        /* The solar system this starbase (POS) is in, unanchored POSes have this information */
	TypeId          int32     `json:"type_id,omitempty"`          /* Starbase (POS) type */
	UnanchorAt      time.Time `json:"unanchor_at,omitempty"`      /* When the POS started unanchoring, for starbases (POSes) in unanchoring state */
}
