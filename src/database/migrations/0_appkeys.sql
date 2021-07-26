CREATE TABLE app_keys (
	id SERIAL NOT NULL UNIQUE,
	"user" VARCHAR(255) NOT NULL,
	date_created TIMESTAMP NOT NULL DEFAULT NOW(),
	key VARCHAR(128) NOT NULL,
	PRIMARY KEY (id)
);