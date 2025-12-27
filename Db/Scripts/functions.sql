CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE OR REPLACE FUNCTION prevent_admin_insert()
    RETURNS TRIGGER AS
$$
BEGIN
    RAISE EXCEPTION 'Insertions are not allowed in admins table';
END;
$$ LANGUAGE plpgsql;