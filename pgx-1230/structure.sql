CREATE TYPE promo_reason AS ENUM ('promo_reason');
CREATE TYPE promo_type AS ENUM ('promo_type');

CREATE TABLE promos
(
promo_id    UUID         NOT NULL PRIMARY KEY,
user_id     UUID         NOT NULL,
reason      promo_reason NOT NULL,
type        promo_type   NOT NULL
);

CREATE UNIQUE INDEX promotions_user_id_first_transfer_key ON promos (user_id)
WHERE type = 'promo_type' AND reason = 'promo_reason';
