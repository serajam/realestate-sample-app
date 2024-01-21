CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS users
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100) NULL     DEFAULT '',
    surname            VARCHAR(100) NULL     DEFAULT '',
    email              VARCHAR(150) NOT NULL UNIQUE,
    password           bytea        NOT NULL,
    is_verified        BOOLEAN      NOT NULL DEFAULT FALSE,
    reset_token        VARCHAR(50),
    reset_token_expiry TIMESTAMP,
    created_at         TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TYPE token_action_type AS ENUM ('email_verification', 'password_reset');

CREATE TABLE IF NOT EXISTS user_token_actions
(
    id           SERIAL PRIMARY KEY,
    token        VARCHAR(32)       NOT NULL,
    user_id      INTEGER REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE DEFAULT NULL,
    action       token_action_type NOT NULL,
    token_expiry TIMESTAMP         NOT NULL,
    payload      jsonb             NOT NULL DEFAULT '[]',
    created_at   TIMESTAMP         NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP         NOT NULL DEFAULT NOW()
);

/*1
enum PropertyHomeTypeEnum {
  CONDO = '1',
  HOUSE = '2',
  COMMERCIAL = '3',
  SMALL_HOUSE = '4',
  TOWN_HOUSE = '5',
  MULTI_FAMILY = '6',
  LAND = '7',
  MANUFACTURED = '8',
  OBJECT = '9',
}

enum PropertyConditionEnum {
  UNDER_CONSTRUCTION = '1',
  PREPARATION = '2',
  FINISHED = '3'
}

enum PropertyTypeEnum {
  SOLD = '1',
  SALE = '2',
  RENT = '3',
}

*/

CREATE SEQUENCE properties_sequence
    START 10001
    INCREMENT 1;

CREATE TABLE IF NOT EXISTS search_filters
(
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    name            VARCHAR(100)            NOT NULL,
    email_frequency SMALLINT                NOT NULL DEFAULT 0,
    filters         jsonb                   NOT NULL DEFAULT '[]',
    polygon         jsonb                   NOT NULL DEFAULT '[]',
    sort            SMALLINT                NOT NULL DEFAULT 0,
    created_at      TIMESTAMP               NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP               NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS properties
(
    id               INTEGER PRIMARY KEY             DEFAULT NEXTVAL('properties_sequence'),
    user_id          INTEGER REFERENCES users (id)
        ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    location         geometry(POINT, 4326)  NOT NULL,
    price            NUMERIC                NOT NULL,
    price_currency   VARCHAR(3),
    address          TEXT                   NOT NULL,
    country          VARCHAR(50),
    city             VARCHAR(50),
    state            VARCHAR(50),
    street           VARCHAR(100),
    house_number     VARCHAR(50),
    neighbourhood    VARCHAR(100),
    zip_code         VARCHAR(20),
    home_size        NUMERIC(15, 2),
    lot_size         NUMERIC(15, 2),
    year_build       SMALLINT CHECK (year_build BETWEEN 0 AND 9999),
    bedroom          SMALLINT                        DEFAULT 0,
    bathroom         SMALLINT                        DEFAULT 0,
    broker_name      VARCHAR(200),
    home_type        SMALLINT CHECK (home_type BETWEEN 1 AND 20),
    property_type    SMALLINT CHECK (property_type BETWEEN 1 AND 20),
    condition        SMALLINT CHECK (condition BETWEEN 1 AND 20),
    description      TEXT,
    floor            SMALLINT                        DEFAULT 0,
    total_floors     SMALLINT                        DEFAULT 0,
    is_active_status bool                            DEFAULT FALSE,
    has_images       bool                            DEFAULT FALSE,
    has_garage       bool                            DEFAULT FALSE,
    has_video        bool                            DEFAULT FALSE,
    has_3d_tour      bool                            DEFAULT FALSE,
    total_parking    SMALLINT                        DEFAULT 0,
    has_ac           bool                            DEFAULT FALSE,
    created_at       TIMESTAMP              NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP              NOT NULL DEFAULT NOW()
)
