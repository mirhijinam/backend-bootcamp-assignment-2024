CREATE TABLE IF NOT EXISTS houses (
    id SERIAL PRIMARY KEY,
    address TEXT NOT NULL,
    year INTEGER NOT NULL,
    developer TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP
);

CREATE TYPE flat_status AS enum ('created', 'approved', 'declined', 'on_moderation');

CREATE TABLE IF NOT EXISTS flats (
    flat_number INTEGER NOT NULL,
    house_id INTEGER NOT NULL REFERENCES houses(id),    
    price INTEGER NOT NULL,
    rooms INTEGER NOT NULL,
    status flat_status NOT NULL DEFAULT 'created',

    CONSTRAINT house_flat_number_pk PRIMARY KEY (flat_number, house_id)
);

CREATE TYPE user_role AS enum ('moderator', 'client');

CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL
);


CREATE TABLE IF NOT EXISTS subscriptions (
    user_id UUID NOT NULL,
    house_id INTEGER NOT NULL REFERENCES houses(id)
);
CREATE TYPE notification_status AS enum ('pending', 'sended');

CREATE TABLE IF NOT EXISTS notifications (
    user_id UUID NOT NULL,
    flat_number INTEGER NOT NULL,
    house_id INTEGER NOT NULL,
    status notification_status NOT NULL DEFAULT 'pending',

    FOREIGN KEY (flat_number, house_id) REFERENCES flats(flat_number,house_id)
);
