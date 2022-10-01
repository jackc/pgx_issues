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

PREPARE s AS INSERT INTO promos
    (promo_id, user_id, reason, type)
VALUES
    ($1, $2, $3, $4)
ON CONFLICT (user_id) WHERE type = $4 AND reason = $3
DO NOTHING;

EXECUTE s ('00ebd890-f5ac-47c7-9365-4ce9875c04a1', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
EXECUTE s ('8e3775d9-af90-472f-9720-d0341ff7bba7', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
EXECUTE s ('f983db27-62d9-4ef1-bc67-e1e492eee48e', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
EXECUTE s ('10532081-851c-4dc4-9d83-8750bd4cf78d', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
EXECUTE s ('98e049e4-8762-4c95-be6c-31f4d8f9b04e', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
EXECUTE s ('355972ac-20a4-4c17-b28d-70d8cb2dd5b8', '132b64e6-1dc9-46cb-9349-5a0d7469622b', 'promo_reason', 'promo_type');
