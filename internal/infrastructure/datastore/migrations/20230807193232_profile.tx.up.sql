ALTER TABLE users
    RENAME COLUMN is_verified TO is_verified_email;

ALTER TABLE users
    ADD account_id INTEGER;

ALTER TABLE users
    ADD screen_name TEXT;

ALTER TABLE users
    ADD photo_url TEXT;

ALTER TABLE users
    ADD phone TEXT;

ALTER TABLE users
    ADD additional_phones TEXT[];

ALTER TABLE users
    ADD country TEXT;

ALTER TABLE users
    ADD city TEXT;

ALTER TABLE users
    ADD zip TEXT;

ALTER TABLE users
    ADD latitude TEXT;

ALTER TABLE users
    ADD longitude TEXT;

ALTER TABLE users
    ADD is_created_password bool;

ALTER TABLE users
    ADD is_verified_phone bool;

ALTER TABLE users
    ADD is_google_linked bool;

ALTER TABLE users
    ADD is_facebook_linked bool;

ALTER TABLE users
    ADD is_apple_linked bool;

ALTER TABLE users
    ADD is_account_locked bool;

ALTER TABLE users
    ADD is_account_deactivated bool;

