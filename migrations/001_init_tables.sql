-- +goose Up

CREATE TABLE users (
                       id            SERIAL PRIMARY KEY,
                       username      TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL
);

CREATE TABLE devices (
                         id         SERIAL PRIMARY KEY,
                         name       TEXT UNIQUE NOT NULL,
                         type       TEXT NOT NULL,
                         location   TEXT,
                         created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE light_readings (
                                id          SERIAL PRIMARY KEY,
                                device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                is_on       BOOLEAN NOT NULL,
                                brightness  INTEGER NOT NULL CHECK (brightness >= 0 AND brightness <= 100),
                                updated_at  TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE dht22_readings (
                                id          SERIAL PRIMARY KEY,
                                device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                temperature DOUBLE PRECISION NOT NULL,
                                humidity    DOUBLE PRECISION NOT NULL,
                                recorded_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE water_leak_readings (
                                     id          SERIAL PRIMARY KEY,
                                     device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                     is_leaking  BOOLEAN NOT NULL,
                                     recorded_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE climate_control_states (
                                        id          SERIAL PRIMARY KEY,
                                        device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                        is_on       BOOLEAN NOT NULL,
                                        target_temp DOUBLE PRECISION,
                                        recorded_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE ds18b20_readings (
                                  id          SERIAL PRIMARY KEY,
                                  device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                  temperature DOUBLE PRECISION NOT NULL,
                                  recorded_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE motion_readings (
                                 id          SERIAL PRIMARY KEY,
                                 device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                                 motion      BOOLEAN NOT NULL,
                                 recorded_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE door_states (
                             id          SERIAL PRIMARY KEY,
                             device_name TEXT REFERENCES devices(name) ON DELETE CASCADE,
                             is_open     BOOLEAN NOT NULL,
                             recorded_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS door_states;
DROP TABLE IF EXISTS motion_readings;
DROP TABLE IF EXISTS ds18b20_readings;
DROP TABLE IF EXISTS climate_control_states;
DROP TABLE IF EXISTS water_leak_readings;
DROP TABLE IF EXISTS dht22_readings;
DROP TABLE IF EXISTS light_readings;
DROP TABLE IF EXISTS devices;
