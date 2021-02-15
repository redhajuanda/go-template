-- +migrate Up
CREATE TABLE users (
    id uuid NOT NULL PRIMARY KEY,
    username varchar NOT NULL,
    first_name varchar(32) NOT NULL,
    last_name varchar(32),
    email varchar(100) NOT NULL,
    email_verified_at timestamp NULL,
    password varchar NOT NULL DEFAULT '',
    sso_source varchar NOT NULL DEFAULT '',
    profile_pic varchar NOT NULL DEFAULT '',
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- UNIQUE (username),
    UNIQUE (email)
);

INSERT INTO public.users (id, username, first_name, last_name, email, email_verified_at, "password", sso_source, profile_pic, is_active, created_at, updated_at)
    VALUES ('290faed6-d5fe-465c-9329-811b3bcf4d9e', 'redhajuanda', 'Redha', 'Juanda', 'redhajuanda@gmail.com', '2020-12-16 07:22:02.508', '$2a$04$1lOwfT6Bt78FlICF5p9oh.iiH5RqoakLzfFyhHHvAzEuBK9PpCpmu', '', 'https://lh3.googleusercontent.com/a-/AOh14GgT67aRIIqXSOsZApBzn85ag0v03kPKaFRTxWaxEw=s96-c', TRUE, '2020-12-16 07:22:02.508', '2020-12-16 07:22:02.508');

-- +migrate Down
DROP TABLE users;

