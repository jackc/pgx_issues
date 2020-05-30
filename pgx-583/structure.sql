CREATE TYPE color AS ENUM (
    'red',
    'green',
    'blue'
    );

CREATE TABLE blah (
                      id int,
                      single_color color
);

INSERT INTO blah(id,single_color) VALUES(1, 'blue');
