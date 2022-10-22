module github.com/jackc/pgx_issues

go 1.18

replace github.com/jackc/pgx/v5 => ../pgx

// replace github.com/jackc/pgx/v4 => ../pgx

// replace github.com/jackc/pgtype => ../pgtype

require (
	github.com/google/uuid v1.3.0
	github.com/jackc/pgconn v1.13.0
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/jackc/pgtype v1.12.0
	github.com/jackc/pgx/v4 v4.17.2
	github.com/jackc/pgx/v5 v5.0.1
	github.com/lib/pq v1.10.7
	github.com/shopspring/decimal v1.3.1
	github.com/vgarvardt/pgx-google-uuid/v5 v5.0.0
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)
