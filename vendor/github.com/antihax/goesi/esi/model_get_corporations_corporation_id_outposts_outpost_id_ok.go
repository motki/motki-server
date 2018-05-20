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

/* A list of GetCorporationsCorporationIdOutpostsOutpostIdOk. */
//easyjson:json
type GetCorporationsCorporationIdOutpostsOutpostIdOkList []GetCorporationsCorporationIdOutpostsOutpostIdOk

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdOutpostsOutpostIdOk struct {
	Coordinates              GetCorporationsCorporationIdOutpostsOutpostIdCoordinates `json:"coordinates,omitempty"`
	DockingCostPerShipVolume float32                                                  `json:"docking_cost_per_ship_volume,omitempty"` /* docking_cost_per_ship_volume number */
	OfficeRentalCost         int64                                                    `json:"office_rental_cost,omitempty"`           /* office_rental_cost integer */
	OwnerId                  int32                                                    `json:"owner_id,omitempty"`                     /* The entity that owns the station (e.g. the entity whose logo is on the station services bar) */
	ReprocessingEfficiency   float32                                                  `json:"reprocessing_efficiency,omitempty"`      /* reprocessing_efficiency number */
	ReprocessingStationTake  float32                                                  `json:"reprocessing_station_take,omitempty"`    /* reprocessing_station_take number */
	Services                 []GetCorporationsCorporationIdOutpostsOutpostIdService   `json:"services,omitempty"`                     /* A list of services the given outpost provides */
	StandingOwnerId          int32                                                    `json:"standing_owner_id,omitempty"`            /* The owner ID that sets the ability for someone to dock based on standings. */
	SystemId                 int32                                                    `json:"system_id,omitempty"`                    /* The ID of the solar system the outpost rests in */
	TypeId                   int32                                                    `json:"type_id,omitempty"`                      /* The type ID of the given outpost */
}
