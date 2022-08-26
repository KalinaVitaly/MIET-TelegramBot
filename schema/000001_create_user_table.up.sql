CREATE TABLE IF NOT EXISTS member
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    user_tg_id integer NOT NULL,
    first_name character varying(30) COLLATE pg_catalog."default",
    last_name character varying(30) COLLATE pg_catalog."default",
    username character varying(30) COLLATE pg_catalog."default" NOT NULL,
    group_name character varying(30) COLLATE pg_catalog."default",
    CONSTRAINT member_pkey PRIMARY KEY (id)
);