CREATE TABLE user_saved_homes
(
    user_id     INTEGER   NOT NULL
        CONSTRAINT user_saved_homes_users_id_fk
            REFERENCES users,
    property_id INTEGER   NOT NULL
        CONSTRAINT user_saved_homes_properties_id_fk
            REFERENCES properties
            ON UPDATE CASCADE ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE user_saved_homes
    ADD CONSTRAINT user_saved_homes_pk
        PRIMARY KEY (user_id, property_id);

