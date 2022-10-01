package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

const insert = `
INSERT INTO promos
    (promo_id, user_id, reason, type)
VALUES
    ($1, $2, $3, $4)
ON CONFLICT (user_id) WHERE type = $4 AND reason = $3
DO NOTHING
`

type PromoReason string

const PromoReasonS PromoReason = "promo_reason"

type PromoType string

const PromoTypeS PromoType = "promo_type"

type Promo struct {
	PromoID     uuid.UUID
	UserID      uuid.UUID
	PromoReason PromoReason
	PromoType   PromoType
}

func requireNoError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	requireNoError(err)
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, `set log_statement = 'all'`)
	requireNoError(err)

	p := Promo{
		UserID:      uuid.New(),
		PromoReason: PromoReasonS,
		PromoType:   PromoTypeS,
	}

	for i := 0; i < 200; i++ {
		promo := p

		promo.PromoID = uuid.New()
		ct, insertErr := conn.Exec(ctx, insert, promo.PromoID, promo.UserID, promo.PromoReason, promo.PromoType)
		if insertErr != nil {
			requireNoError(insertErr)
		} else {
			log.Println("SUCCESS", ct)
		}
	}
}
