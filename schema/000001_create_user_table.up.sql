CREATE TABLE IF NOT EXISTS member
(
    id SERIAL PRIMARY KEY,
    user_tg_id integer NOT NULL,
    first_name  CHARACTER VARYING(30) COLLATE pg_catalog."default",
    last_name  CHARACTER VARYING(30) COLLATE pg_catalog."default",
    username  CHARACTER VARYING(30) COLLATE pg_catalog."default" NOT NULL,
    group_name  CHARACTER VARYING(30) COLLATE pg_catalog."default",
    user_deauth BOOLEAN default FALSE
);