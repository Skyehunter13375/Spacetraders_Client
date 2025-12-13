package Model

// ┣━━━━━━━━━━━━━━━━━━━━━┫ System (All at once - no detail) ┣━━━━━━━━━━━━━━━━━━━━━┫
type System struct {
	Symbol        string         `json:"symbol"`
	Sector        string         `json:"sectorSymbol"`
	Constellation string         `json:"constellation"`
	Name          string         `json:"name"`
	Type          string         `json:"type"`
	Xcoord        int            `json:"x"`
	Ycoord        int            `json:"y"`
	Waypoints     []SysWaypoints `json:"waypoints"`
}

type SysWaypoints struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Xcoord int    `json:"x"`
	Ycoord int    `json:"y"`
	Orbits string `json:"orbits"`
}

type SysWayOrbitals struct {
	Symbol string `json:"symbol"`
}

type SysFactions struct {
	Symbol string `json:"symbol"`
}

// {
//     "data": {
//         "symbol": "X1-XQ13",
//         "sectorSymbol": "X1",
//         "type": "BLUE_STAR",
//         "x": 2216,
//         "y": -10966,
//         "waypoints": [
//             {
//                 "symbol": "X1-XQ13-A1",
//                 "type": "PLANET",
//                 "x": 8,
//                 "y": 26,
//                 "orbitals": [
//                     {"symbol": "X1-XQ13-A2"},
//                     {"symbol": "X1-XQ13-A3"}
//                 ]
//             },
//             {
//                 "symbol": "X1-XQ13-XB5E",
//                 "type": "ENGINEERED_ASTEROID",
//                 "x": -28,
//                 "y": -5,
//                 "orbitals": []
//             }
//         ],
//         "factions": [],
//         "constellation": "Purva Bhadrapada",
//         "name": "Purva Bhadrapada XLIV"
//     }
// }

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Individual Waypoints ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
type Waypoint struct {
	System       string      `json:"systemSymbol"`
	Symbol       string      `json:"symbol"`
	Type         string      `json:"type"`
	X            int         `json:"x"`
	Y            int         `json:"Y"`
	Orbits       string      `json:"orbits"`
	Traits       []WayTraits `json:"traits"`
	Modifiers    []WayMods   `json:"modifiers"`
	Chart        WayChart    `json:"chart"`
	Construction bool        `json:"isUnderConstruction"`
	// Factions     []WayFactions `json:"faction"`
}

type WayFactions struct {
	Symbol string `json:"symbol"`
}

type WayTraits struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Descr  string `json:"description"`
}

type WayMods struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Descr  string `json:"description"`
}

type WayChart struct {
	Symbol string `json:"symbol"`
	Agent  string `json:"submittedBy"`
	Date   string `json:"submittedOn"`
}

// {
//     "data": {
//         "symbol": "string",
//         "type": "PLANET",
//         "systemSymbol": "string",
//         "x": 0,
//         "y": 0,
//         "orbitals": [
//         {
//             "symbol": "string"
//         }
//         ],
//         "orbits": "string",
//         "faction": {
//             "symbol": "COSMIC"
//         },
//         "traits": [
//         {
//             "symbol": "UNCHARTED",
//             "name": "string",
//             "description": "string"
//         }
//         ],
//         "modifiers": [
//         {
//             "symbol": "STRIPPED",
//             "name": "string",
//             "description": "string"
//         }
//         ],
//         "chart": {
//         "waypointSymbol": "string",
//         "submittedBy": "string",
//         "submittedOn": "2019-08-24T14:15:22Z"
//         },
//         "isUnderConstruction": true
//     }
// }

// ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫ Shipyard ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
type Shipyard struct {
	Symbol       string            `json:"symbol"`
	Types        []SY_Types        `json:"shipTypes"`
	Transactions []SY_Transactions `json:"transactions"`
	Ships        []SY_Ship         `json:"ships"`
}

type SY_Types struct {
	Type string `json:"type"`
}

type SY_Transactions struct {
	Waypoint  string `json:"waypointSymbol"`
	Ship      string `json:"shipSymbol"`
	Type      string `json:"shipType"`
	Price     int64  `json:"price"`
	Agent     string `json:"agentSymbol"`
	Timestamp string `json:"timestamp"`
}

type SY_Ship struct {
	Type        string              `json:"type"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Supply      string              `json:"supply"`
	Activity    string              `json:"activity"`
	Price       int64               `json:"shipPrice"`
	ModFee      int64               `json:"modificationsFee"`
	Frame       ShipFrame     `json:"frame"`
	Reactor     ShipReactor   `json:"reactor"`
	Engine      ShipEngine    `json:"engine"`
	Modules     []ShipModules `json:"modules"`
	Mounts      []ShipMounts  `json:"mounts"`
	Crew        ShipCrew      `json:"crew"`
}

// {
//   "data": {
//     "symbol": "string",
//     "shipTypes": [
//       {
//         "type": "SHIP_PROBE"
//       }
//     ],
//     "transactions": [
//       {
//         "waypointSymbol": "string",
//         "shipSymbol": "string",
//         "shipType": "string",
//         "price": 0,
//         "agentSymbol": "string",
//         "timestamp": "2019-08-24T14:15:22Z"
//       }
//     ],
//     "ships": [
//       {
//         "type": "SHIP_PROBE",
//         "name": "string",
//         "description": "string",
//         "supply": "SCARCE",
//         "activity": "WEAK",
//         "purchasePrice": 0,
//         "frame": {
//           "symbol": "FRAME_PROBE",
//           "name": "string",
//           "description": "string",
//           "condition": 0,
//           "integrity": 0,
//           "moduleSlots": 0,
//           "mountingPoints": 0,
//           "fuelCapacity": 0,
//           "requirements": {
//             "power": 0,
//             "crew": 0,
//             "slots": 0
//           },
//           "quality": 0
//         },
//         "reactor": {
//           "symbol": "REACTOR_SOLAR_I",
//           "name": "string",
//           "description": "string",
//           "condition": 0,
//           "integrity": 0,
//           "powerOutput": 1,
//           "requirements": {
//             "power": 0,
//             "crew": 0,
//             "slots": 0
//           },
//           "quality": 0
//         },
//         "engine": {
//           "symbol": "ENGINE_IMPULSE_DRIVE_I",
//           "name": "string",
//           "description": "string",
//           "condition": 0,
//           "integrity": 0,
//           "speed": 1,
//           "requirements": {
//             "power": 0,
//             "crew": 0,
//             "slots": 0
//           },
//           "quality": 0
//         },
//         "modules": [
//           {
//             "symbol": "MODULE_MINERAL_PROCESSOR_I",
//             "capacity": 0,
//             "range": 0,
//             "name": "string",
//             "description": "string",
//             "requirements": {
//               "power": 0,
//               "crew": 0,
//               "slots": 0
//             }
//           }
//         ],
//         "mounts": [
//           {
//             "symbol": "MOUNT_GAS_SIPHON_I",
//             "name": "string",
//             "description": "string",
//             "strength": 0,
//             "deposits": [
//               "QUARTZ_SAND"
//             ],
//             "requirements": {
//               "power": 0,
//               "crew": 0,
//               "slots": 0
//             }
//           }
//         ],
//         "crew": {
//           "required": 0,
//           "capacity": 0
//         }
//       }
//     ],
//     "modificationsFee": 0
//   }
// }
