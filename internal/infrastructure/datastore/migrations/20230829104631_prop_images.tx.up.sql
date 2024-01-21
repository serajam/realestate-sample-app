CREATE TABLE IF NOT EXISTS property_images (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    property_id INTEGER      NOT NULL,
    user_id INTEGER      NOT NULL,
    description TEXT DEFAULT '' NOT NULL,
    position SMALLINT DEFAULT 0 NOT NULL,
    is_main BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    FOREIGN KEY (property_id) REFERENCES properties (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
