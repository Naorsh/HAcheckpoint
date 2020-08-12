
CREATE TABLE public."Applications"
(
    id character varying(36) COLLATE pg_catalog."default" NOT NULL,
    name character varying(15) COLLATE pg_catalog."default" NOT NULL,
    key character varying(100) COLLATE pg_catalog."default" NOT NULL,
    creation_time character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Applications_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public."Applications"
    OWNER to postgres;