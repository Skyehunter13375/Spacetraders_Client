PRAGMA foreign_keys = ON;

CREATE TABLE server (
    server_up    INTEGER, -- boolean (0/1)
    game_version TEXT,
    agents       INTEGER,
    ships        INTEGER,
    systems      INTEGER,
    waypoints    INTEGER,
    accounts     INTEGER,
    reset_freq   TEXT,
    reset_date   TEXT,
    next_reset   TEXT,
    last_updated TEXT
);

CREATE TABLE leaderboard_creds (
    agent   TEXT PRIMARY KEY,
    credits INTEGER
);

CREATE TABLE leaderboard_charts (
    agent  TEXT PRIMARY KEY,
    charts INTEGER
);

CREATE TABLE agents (
    account_id   TEXT,
    symbol       TEXT PRIMARY KEY,
    faction      TEXT,
    hq           TEXT,
    ships        INTEGER,
    credits      INTEGER,
    last_updated TEXT
);

CREATE TABLE ships (
    symbol       TEXT PRIMARY KEY,
    name         TEXT,
    role         TEXT,
    faction      TEXT,
    last_updated TEXT
);

CREATE TABLE ship_nav (
    ship             TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    system           TEXT,
    waypoint         TEXT,
    status           TEXT,
    flight_mode      TEXT,
    origin           TEXT,
    origin_type      TEXT,
    origin_x         INTEGER,
    origin_y         INTEGER,
    destination      TEXT,
    destination_type TEXT,
    destination_x    INTEGER,
    destination_y    INTEGER,
    arrival          TEXT,
    departure        TEXT
);

CREATE TABLE ship_crew (
    ship     TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    current  INTEGER,
    required INTEGER,
    capacity INTEGER,
    rotation TEXT,
    morale   INTEGER,
    wages    INTEGER
);

CREATE TABLE ship_fuel (
    ship     TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    current  INTEGER,
    capacity INTEGER
);

CREATE TABLE ship_frame (
    ship           TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    symbol         TEXT,
    name           TEXT,
    description    TEXT,
    module_slots   INTEGER,
    mount_points   INTEGER,
    fuel_capacity  INTEGER,
    condition      INTEGER,
    integrity      INTEGER,
    quality        INTEGER,
    power_required INTEGER,
    crew_required  INTEGER
);

CREATE TABLE ship_reactor (
    ship          TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    symbol        TEXT,
    name          TEXT,
    description   TEXT,
    condition     INTEGER,
    integrity     INTEGER,
    power_output  INTEGER,
    quality       INTEGER,
    crew_required INTEGER
);

CREATE TABLE ship_engine (
    ship           TEXT PRIMARY KEY REFERENCES ships(symbol) ON DELETE CASCADE,
    symbol         TEXT,
    name           TEXT,
    description    TEXT,
    condition      INTEGER,
    integrity      INTEGER,
    speed          INTEGER,
    quality        INTEGER,
    power_required INTEGER,
    crew_required  INTEGER
);

CREATE TABLE contracts (
    id                 TEXT PRIMARY KEY,
    faction            TEXT,
    type               TEXT,
    pay_on_accept      INTEGER,
    pay_on_complete    INTEGER,
    accepted           INTEGER, -- boolean 0/1
    fulfilled          INTEGER, -- boolean 0/1
    deadline           TEXT,
    expiration         TEXT,
    deadline_to_accept TEXT,
    last_updated       TEXT
);

CREATE TABLE contract_materials (
    id              TEXT REFERENCES contracts(id) ON DELETE CASCADE,
    material        TEXT,
    destination     TEXT,
    units_required  INTEGER,
    units_fulfilled INTEGER,
    UNIQUE(id, material, destination)
);

CREATE TABLE systems (
    symbol        TEXT PRIMARY KEY,
    sector        TEXT,
    constellation TEXT,
    name          TEXT,
    type          TEXT,
    x_coord       INTEGER,
    y_coord       INTEGER,
    factions      TEXT -- JSON array
);

CREATE TABLE waypoints (
    system       TEXT REFERENCES systems(symbol) ON DELETE CASCADE,
    symbol       TEXT PRIMARY KEY,
    type         TEXT,
    x_coord      INTEGER,
    y_coord      INTEGER,
    orbits       TEXT,
    construction INTEGER,-- boolean 0/1
    traits       TEXT,   -- JSON array
    modifiers    TEXT    -- JSON array
);

CREATE TABLE orbitals (
    waypoint TEXT REFERENCES waypoints(symbol) ON DELETE CASCADE,
    symbol   TEXT PRIMARY KEY
);
