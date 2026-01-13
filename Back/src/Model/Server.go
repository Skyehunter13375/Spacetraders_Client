package Model

type ServerState struct {
	Status       string       `json:"status"`
	Version      string       `json:"version"`
	LastReset    string       `json:"resetDate"`
	ServerResets ServResets   `json:"serverResets"`
	Stats        ServStats    `json:"stats"`
	Leaderboards Leaderboards `json:"leaderboards"`
	LastCheckIn  string
}

type ServResets struct {
	NextReset string `json:"next"`
	ResetFreq string `json:"frequency"`
}

type ServStats struct {
	Accounts  int `json:"accounts"`
	Agents    int `json:"agents"`
	Ships     int `json:"ships"`
	Systems   int `json:"systems"`
	Waypoints int `json:"waypoints"`
}

type Leaderboards struct {
	MostCredits []LeaderCredits `json:"mostCredits"`
	MostCharted []LeaderCharts  `json:"mostSubmittedCharts"`
}

type LeaderCredits struct {
	Agent string `json:"agentSymbol"`
	Creds int64  `json:"credits"`
}

type LeaderCharts struct {
	Agent  string `json:"agentSymbol"`
	Charts int64  `json:"chartCount"`
}

// {:
//   "status": "string",
//   "version": "string",
//   "resetDate": "string",
//   "description": "string",
//   "stats": {
//     "accounts": 0,
//     "agents": 0,
//     "ships": 0,
//     "systems": 0,
//     "waypoints": 0
//   },
//   "leaderboards": {
//     "mostCredits": [
//       {
//         "agentSymbol": "string",
//         "credits": -9007199254740991
//       }
//     ],
//     "mostSubmittedCharts": [
//      {
//         "agentSymbol": "string",
//         "chartCount": 0
//       }
//     ]
//   },
//   "serverResets": {
//     "next": "string",
//     "frequency": "string"
//   },
//   "announcements": [
//     {
//       "title": "string",
//       "body": "string"
//     }
//   ],
//   "links": [
//     {
//       "name": "string",
//       "url": "string"
//     }
//   ]
// }
