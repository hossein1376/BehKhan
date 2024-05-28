CREATE TABLE IF NOT EXISTS books
(
    id    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title TEXT            NOT NULL,
    CONSTRAINT authors_pk
        PRIMARY KEY (id)
);
