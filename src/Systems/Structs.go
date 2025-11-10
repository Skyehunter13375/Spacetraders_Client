package Waypoints

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

type Waypoint struct {
	System       string        `json:"systemSymbol"`
	Symbol       string        `json:"symbol"`
	Type         string        `json:"type"`
	X            int           `json:"x"`
	Y            int           `json:"Y"`
	Orbits       string        `json:"orbits"`
	Factions     []WayFactions `json:"faction"`
	Traits       []WayTraits   `json:"traits"`
	Modifiers    []WayMods     `json:"modifiers"`
	Chart        WayChart      `json:"chart"`
	Construction bool          `json:"isUnderConstruction"`
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
