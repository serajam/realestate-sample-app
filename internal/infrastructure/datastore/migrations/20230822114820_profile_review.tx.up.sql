UPDATE users
SET is_account_locked      = FALSE,
    is_account_deactivated = TRUE;

ALTER TABLE users
    RENAME COLUMN is_account_locked TO locked;

ALTER TABLE users
    ALTER COLUMN locked SET NOT NULL;

ALTER TABLE users
    ALTER COLUMN locked SET DEFAULT FALSE;

ALTER TABLE users
    RENAME COLUMN is_account_deactivated TO active;

ALTER TABLE users
    ALTER COLUMN active SET NOT NULL;

ALTER TABLE users
    ALTER COLUMN active SET DEFAULT TRUE;

ALTER TABLE properties
    ADD active BOOLEAN DEFAULT TRUE NOT NULL;

ALTER TABLE properties
    DROP COLUMN is_active_status;
