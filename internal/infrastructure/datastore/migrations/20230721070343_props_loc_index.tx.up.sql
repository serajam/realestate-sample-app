CREATE INDEX properties_location_gist_index ON properties USING gist (location);
CREATE INDEX properties_city_index ON properties (city);