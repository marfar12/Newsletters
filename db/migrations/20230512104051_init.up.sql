CREATE TABLE IF NOT EXISTS editors
(
    editor_id uuid NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_editors_id PRIMARY KEY (editor_id),
    CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS newsletters
(
    newsletter_id uuid NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    "desc" text COLLATE pg_catalog."default",
    editor_id uuid,
    CONSTRAINT pk_newsletters_id PRIMARY KEY (newsletter_id),
    CONSTRAINT fk_newsletter_editor_id FOREIGN KEY (editor_id)
        REFERENCES editors (editor_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    newsletter_id uuid NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT subscriptions_pkey PRIMARY KEY (newsletter_id, email)
);