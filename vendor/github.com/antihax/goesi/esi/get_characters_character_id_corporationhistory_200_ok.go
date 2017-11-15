/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.2
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

/* A list of GetCharactersCharacterIdCorporationhistory200Ok. */
//easyjson:json
type GetCharactersCharacterIdCorporationhistory200OkList []GetCharactersCharacterIdCorporationhistory200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdCorporationhistory200Ok struct {
	StartDate     time.Time `json:"start_date,omitempty"`     /* start_date string */
	CorporationId int32     `json:"corporation_id,omitempty"` /* corporation_id integer */
	IsDeleted     bool      `json:"is_deleted,omitempty"`     /* True if the corporation has been deleted */
	RecordId      int32     `json:"record_id,omitempty"`      /* An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous */
}
