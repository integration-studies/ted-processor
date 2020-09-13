CREATE TABLE pre_transaction
(
    id             varchar(200) PRIMARY KEY,
    from_account   VARCHAR(50)   NOT NULL,
    to_account   VARCHAR(50)   NOT NULL,
    value   numeric(10,4)         NOT NULL,
    time TIMESTAMP           NOT NULL,
    device_type   VARCHAR(50)   NOT NULL,
    status   VARCHAR(50)   NOT NULL,
    started_at TIMESTAMP           NOT NULL,
    metadata json NOT NULL
);

