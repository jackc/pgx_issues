CREATE TABLE test_text (
	id INT8 NOT NULL DEFAULT unique_rowid(),
	name STRING NULL,
	num INT8 NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, name, num)
);

INSERT INTO
    test_text (id, name, num)
VALUES
    (1, 'one', 1001),
    (2, 'two', 1002);
