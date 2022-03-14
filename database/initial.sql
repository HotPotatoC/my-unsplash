CREATE TABLE IF NOT EXISTS "photos" (
    "id" bigint GENERATED ALWAYS AS IDENTITY,
    "label" varchar(255) NOT NULL,
    "label_tsvector" tsvector NOT NULL,
    "url" varchar(255) NOT NULL,
    "password" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL
);

ALTER TABLE "photos"
    ADD PRIMARY KEY ("id");

CREATE INDEX "photos_label_tsvector_idx" ON "photos" USING gin (to_tsvector('english', "label"));

CREATE OR REPLACE FUNCTION photo_label_tsvector_update ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.label_tsvector := to_tsvector('english', NEW.label);
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER "photos_label_tsvector_update"
    BEFORE INSERT OR UPDATE ON "photos"
    FOR EACH ROW
    EXECUTE PROCEDURE photo_label_tsvector_update ();

