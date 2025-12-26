package Model

type RegPayload struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

type RegResult struct {
	Token    string   `json:"token"`
	Agent    Agent    `json:"agent"`
	Faction  Faction  `json:"faction"`
	Contract Contract `json:"contract"`
	Ships    []Ship   `json:"ships"`
}

/* SAMPLE
{
	"data": {
		"token": "string_here"
		"agent": {
			"accountId": "string_here",
			"symbol": "NULLSKY2",
			"headquarters": "X1-KX27-A1",
			"credits": 175000,
			"startingFaction": "VOID",
			"shipCount": 2
		},
		"faction": {
			"symbol": "VOID",
			"name": "Voidfarers",
			"description": "The Voidfarers are a group of nomadic traders and adventurers who travel the galaxy in search of riches and adventure, willing to take risks and explore the unknown.",
			"traits": [
				{
					"symbol": "DARING",
					"name": "Daring",
					"description": "Willing to take risks and challenges. Sometimes unafraid to explore new and unknown territories, and may be willing to take bold and decisive action in order to achieve their goals. Sometimes able to overcome challenges that would be insurmountable for others."
				},
				{
					"symbol": "EXPLORATORY",
					"name": "Exploratory",
					"description": "Dedicated to exploration and discovery. Sometimes interested in mapping new territories and uncovering the secrets of the universe. Sometimes able to overcome obstacles and challenges in order to advance the boundaries of human knowledge and understanding."
				},
				{
					"symbol": "RESOURCEFUL",
					"name": "Resourceful",
					"description": "Known for their ingenuity and ability to make the most out of limited resources. Able to improvise and adapt to changing circumstances, using whatever is available to them in order to overcome challenges and achieve their goals."
				},
				{
					"symbol": "FLEXIBLE",
					"name": "Flexible",
					"description": "Able to adapt to changing circumstances and environments. Sometimes able to quickly switch between different strategies and tactics in order to respond to new challenges or opportunities. Sometimes able to improvise and think on their feet, making them difficult to predict or outmaneuver."
				}
			],
			"isRecruiting": true,
			"headquarters": "X1-FQ57"
		},
		"contract": {
			"id": "cmj4le1dfb11pri6v7hsu2y3c",
			"factionSymbol": "VOID",
			"type": "PROCUREMENT",
			"terms": {
				"deadline": "2025-12-20T17:51:06.096Z",
				"payment": {
					"onAccepted": 1615,
					"onFulfilled": 9190
				},
				"deliver": [
					{
						"tradeSymbol": "ALUMINUM_ORE",
						"destinationSymbol": "X1-KX27-H52",
						"unitsRequired": 46,
						"unitsFulfilled": 0
					}
				]
			},
			"accepted": false,
			"fulfilled": false,
			"expiration": "2025-12-14T17:51:06.096Z",
			"deadlineToAccept": "2025-12-14T17:51:06.096Z"
		},
		"ships": [
			{
				"symbol": "NULLSKY2-1",
				"registration": {
					"name": "NULLSKY2-1",
					"factionSymbol": "VOID",
					"role": "COMMAND"
				},
				"nav": {
					"systemSymbol": "X1-KX27",
					"waypointSymbol": "X1-KX27-A1",
					"route": {
						"destination": {
							"symbol": "X1-KX27-A1",
							"type": "PLANET",
							"systemSymbol": "X1-KX27",
							"x": 19,
							"y": 16
						},
						"origin": {
							"symbol": "X1-KX27-A1",
							"type": "PLANET",
							"systemSymbol": "X1-KX27",
							"x": 19,
							"y": 16
						},
						"departureTime": "2025-12-13T17:51:06.112Z",
						"arrival": "2025-12-13T17:51:06.112Z"
					},
					"status": "DOCKED",
					"flightMode": "CRUISE"
				},
				"crew": {
					"current": 57,
					"required": 57,
					"capacity": 80,
					"rotation": "STRICT",
					"morale": 100,
					"wages": 0
				},
				"frame": {
					"symbol": "FRAME_FRIGATE",
					"name": "Frigate",
					"condition": 1,
					"integrity": 1,
					"description": "A medium-sized, multi-purpose spacecraft, often used for combat, transport, or support operations.",
					"moduleSlots": 8,
					"mountingPoints": 5,
					"fuelCapacity": 400,
					"requirements": {
						"power": 8,
						"crew": 25
					},
					"quality": 4
				},
				"reactor": {
					"symbol": "REACTOR_FISSION_I",
					"name": "Fission Reactor I",
					"condition": 1,
					"integrity": 1,
					"description": "A basic fission power reactor, used to generate electricity from nuclear fission reactions.",
					"powerOutput": 31,
					"requirements": {
						"crew": 8
					},
					"quality": 5
				},
				"engine": {
					"symbol": "ENGINE_ION_DRIVE_II",
					"name": "Ion Drive II",
					"condition": 1,
					"integrity": 1,
					"description": "An advanced propulsion system that uses ionized particles to generate high-speed, low-thrust acceleration, with improved efficiency and performance.",
					"speed": 36,
					"requirements": {
						"power": 6,
						"crew": 8
					},
					"quality": 4
				},
				"modules": [
					{
						"symbol": "MODULE_CARGO_HOLD_II",
						"name": "Expanded Cargo Hold",
						"description": "An expanded cargo hold module that provides more efficient storage space for a ship's cargo.",
						"requirements": {
							"power": 2,
							"crew": 2,
							"slots": 2
						},
						"capacity": 40
					},
					{
						"symbol": "MODULE_CREW_QUARTERS_I",
						"name": "Crew Quarters",
						"description": "A module that provides living space and amenities for the crew.",
						"requirements": {
							"power": 1,
							"crew": 2,
							"slots": 1
						},
						"capacity": 40
					},
					{
						"symbol": "MODULE_CREW_QUARTERS_I",
						"name": "Crew Quarters",
						"description": "A module that provides living space and amenities for the crew.",
						"requirements": {
							"power": 1,
							"crew": 2,
							"slots": 1
						},
						"capacity": 40
					},
					{
						"symbol": "MODULE_MINERAL_PROCESSOR_I",
						"name": "Mineral Processor",
						"description": "Crushes and processes extracted minerals and ores into their component parts, filters out impurities, and containerizes them into raw storage units.",
						"requirements": {
							"power": 1,
							"crew": 0,
							"slots": 2
						}
					},
					{
						"symbol": "MODULE_GAS_PROCESSOR_I",
						"name": "Gas Processor",
						"description": "Filters and processes extracted gases into their component parts, filters out impurities, and containerizes them into raw storage units.",
						"requirements": {
							"power": 1,
							"crew": 0,
							"slots": 2
						}
					}
				],
				"mounts": [
					{
						"symbol": "MOUNT_SENSOR_ARRAY_II",
						"name": "Sensor Array II",
						"description": "An advanced sensor array that improves a ship's ability to detect and track other objects in space with greater accuracy and range.",
						"requirements": {
							"power": 2,
							"crew": 2
						},
						"strength": 4
					},
					{
						"symbol": "MOUNT_GAS_SIPHON_II",
						"name": "Gas Siphon II",
						"description": "An advanced gas siphon that can extract gas and other resources from gas giants and other gas-rich bodies more efficiently and at a higher rate.",
						"requirements": {
							"power": 2,
							"crew": 2
						},
						"strength": 20
					},
					{
						"symbol": "MOUNT_MINING_LASER_II",
						"name": "Mining Laser II",
						"description": "An advanced mining laser that is more efficient and effective at extracting valuable minerals from asteroids and other space objects.",
						"requirements": {
							"power": 2,
							"crew": 2
						},
						"strength": 5
					},
					{
						"symbol": "MOUNT_SURVEYOR_II",
						"name": "Surveyor II",
						"description": "An advanced survey probe that can be used to gather information about a mineral deposit with greater accuracy.",
						"requirements": {
							"power": 3,
							"crew": 4
						},
						"strength": 2,
						"deposits": [
							"QUARTZ_SAND",
							"SILICON_CRYSTALS",
							"PRECIOUS_STONES",
							"ICE_WATER",
							"AMMONIA_ICE",
							"IRON_ORE",
							"COPPER_ORE",
							"SILVER_ORE",
							"ALUMINUM_ORE",
							"GOLD_ORE",
							"PLATINUM_ORE",
							"DIAMONDS",
							"URANITE_ORE"
						]
					}
				],
				"cargo": {
					"capacity": 40,
					"units": 0,
					"inventory": [

					]
				},
				"fuel": {
					"current": 400,
					"capacity": 400,
					"consumed": {
						"amount": 0,
						"timestamp": "2025-12-13T17:51:06.112Z"
					}
				},
				"cooldown": {
					"shipSymbol": "NULLSKY2-1",
					"totalSeconds": 0,
					"remainingSeconds": 0
				}
			},
			{
				"symbol": "NULLSKY2-2",
				"registration": {
					"name": "NULLSKY2-2",
					"factionSymbol": "VOID",
					"role": "SATELLITE"
				},
				"nav": {
					"systemSymbol": "X1-KX27",
					"waypointSymbol": "X1-KX27-H53",
					"route": {
						"destination": {
							"symbol": "X1-KX27-H53",
							"type": "MOON",
							"systemSymbol": "X1-KX27",
							"x": -18,
							"y": 43
						},
						"origin": {
							"symbol": "X1-KX27-H53",
							"type": "MOON",
							"systemSymbol": "X1-KX27",
							"x": -18,
							"y": 43
						},
						"departureTime": "2025-12-13T17:51:06.137Z",
						"arrival": "2025-12-13T17:51:06.137Z"
					},
					"status": "DOCKED",
					"flightMode": "CRUISE"
				},
				"crew": {
					"current": 0,
					"required": 0,
					"capacity": 0,
					"rotation": "STRICT",
					"morale": 100,
					"wages": 0
				},
				"frame": {
					"symbol": "FRAME_PROBE",
					"name": "Probe",
					"condition": 1,
					"integrity": 1,
					"description": "A small, unmanned spacecraft used for exploration, reconnaissance, and scientific research.",
					"moduleSlots": 0,
					"mountingPoints": 0,
					"fuelCapacity": 0,
					"requirements": {
						"power": 1,
						"crew": 0
					},
					"quality": 1
				},
				"reactor": {
					"symbol": "REACTOR_SOLAR_I",
					"name": "Solar Reactor I",
					"condition": 1,
					"integrity": 1,
					"description": "A basic solar power reactor, used to generate electricity from solar energy.",
					"powerOutput": 3,
					"requirements": {
						"crew": 0
					},
					"quality": 1
				},
				"engine": {
					"symbol": "ENGINE_IMPULSE_DRIVE_I",
					"name": "Impulse Drive I",
					"condition": 1,
					"integrity": 1,
					"description": "A basic low-energy propulsion system that generates thrust for interplanetary travel.",
					"speed": 9,
					"requirements": {
						"power": 1,
						"crew": 0
					},
					"quality": 1
				},
				"modules": [

				],
				"mounts": [

				],
				"cargo": {
					"capacity": 0,
					"units": 0,
					"inventory": [

					]
				},
				"fuel": {
					"current": 0,
					"capacity": 0,
					"consumed": {
						"amount": 0,
						"timestamp": "2025-12-13T17:51:06.137Z"
					}
				},
				"cooldown": {
					"shipSymbol": "NULLSKY2-2",
					"totalSeconds": 0,
					"remainingSeconds": 0
				}
			}
		]
	}
}
*/
