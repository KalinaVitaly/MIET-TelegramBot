CREATE TABLE IF NOT EXISTS statistic
(
    Id SERIAL,
    member_id integer NOT NULL,
    request_command CHARACTER VARYING(30) NOT NULL, 
    request_time timestamp,
    FOREIGN KEY (member_id) REFERENCES member (id) ON DELETE CASCADE
);