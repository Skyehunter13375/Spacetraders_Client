CREATE TABLE tokens (
    type  TEXT PRIMARY KEY,
    token TEXT
);
CREATE TABLE server (
    server_up    BOOLEAN,
    game_version TEXT,
    agents       BIGINT,
    ships        BIGINT,
    systems      BIGINT,
    waypoints    BIGINT,
    accounts     BIGINT,
    reset_freq   TEXT,
    reset_date   TIMESTAMP WITHOUT TIME ZONE,
    next_reset   TIMESTAMP WITHOUT TIME ZONE,
    last_updated TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE agents (
    account_id   TEXT,
    symbol       TEXT PRIMARY KEY,
    faction      TEXT,
    hq           TEXT,
    ships        INT,
    credits      INT,
    last_updated TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE ship (
    symbol       TEXT PRIMARY KEY,
    name         TEXT,
    role         TEXT,
    faction      TEXT,
    last_updated TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE ship_nav (
    ship             TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    system           TEXT,
    waypoint         TEXT,
    status           TEXT,
    flight_mode      TEXT,
    origin           TEXT,
    origin_type      TEXT,
    origin_x         INT,
    origin_y         INT,
    destination      TEXT,
    destination_type TEXT,
    destination_x    INT,
    destination_y    INT,
    arrival          TIMESTAMPTZ,
    departure        TIMESTAMPTZ
);
CREATE TABLE ship_crew (
    ship     TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    current  INT,
    required INT,
    capacity INT,
    rotation TEXT,
    morale   INT,
    wages    INT
);
CREATE TABLE ship_fuel (
    ship     TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    current  INT,
    capacity INT
);
CREATE TABLE ship_frame (
    ship           TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    symbol         TEXT,
    name           TEXT,
    description    TEXT,
    module_slots   INT,
    mount_points   INT,
    fuel_capacity  INT,
    condition      INT,
    integrity      INT,
    quality        INT,
    power_required INT,
    crew_required  INT
);
CREATE TABLE ship_reactor (
    ship          TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    symbol        TEXT,
    name          TEXT,
    description   TEXT,
    condition     INT,
    integrity     INT,
    power_output  INT,
    quality       INT,
    crew_required INT
);
CREATE TABLE ship_engine (
    ship           TEXT PRIMARY KEY REFERENCES ship(symbol) ON DELETE CASCADE,
    symbol         TEXT,
    name           TEXT,
    description    TEXT,
    condition      INT,
    integrity      INT,
    speed          INT,
    quality        INT,
    power_required INT,
    crew_required  INT
);
CREATE TABLE contract (
    id                 TEXT PRIMARY KEY,
    faction            TEXT,
    type               TEXT,
    pay_on_accept      BIGINT,
    pay_on_complete    BIGINT,
    accepted           BOOLEAN,
    fulfilled          BOOLEAN,
    deadline           TIMESTAMP WITH TIME ZONE,
    expiration         TIMESTAMP WITH TIME ZONE,
    deadline_to_accept TIMESTAMP WITH TIME ZONE,
    last_updated       TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE contract_materials (
    id              TEXT REFERENCES contract(id) ON DELETE CASCADE,
    material        TEXT,
    destination     TEXT,
    units_required  BIGINT,
    units_fulfilled BIGINT,
    UNIQUE(id,material,destination)
);
CREATE TABLE systems (
    symbol        TEXT PRIMARY KEY,
    sector        TEXT,
    constellation TEXT,
    name          TEXT,
    type          TEXT,
    x_coord       INT,
    y_coord       INT,
    factions      TEXT
);
CREATE TABLE waypoints (
    system  TEXT REFERENCES systems(symbol) ON DELETE CASCADE,
    symbol  TEXT PRIMARY KEY,
    type    TEXT,
    x_coord INT,
    y_coord INT,
    orbits  TEXT
);
CREATE TABLE orbitals (
    waypoint TEXT REFERENCES waypoints(symbol) ON DELETE CASCADE,
    symbol   TEXT PRIMARY KEY
);