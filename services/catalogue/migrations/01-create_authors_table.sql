CREATE TABLE IF NOT EXISTS authors
(
    id   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name TEXT            NOT NULL,
    CONSTRAINT authors_pk
        PRIMARY KEY (id)
);
