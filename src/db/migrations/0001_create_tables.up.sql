CREATE TABLE IF NOT EXISTS objects(
        id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
        name text NULL ,
        timezone int NULL
);

CREATE TABLE IF NOT EXISTS technological_windows(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    object_id uuid NOT NULL ,
    date_start timestamp NOT NULL ,
    date_end timestamp NOT NULL ,
    UNIQUE (object_id, date_start, date_end),
    CONSTRAINT fk_objects
        FOREIGN KEY (object_id)
            REFERENCES objects(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE

);
