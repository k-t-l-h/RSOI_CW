CREATE TABLE public.auth
(
    "UserUUID" uuid NOT NULL,
    "Login" text COLLATE pg_catalog."default" NOT NULL,
    "Password" text COLLATE pg_catalog."default" NOT NULL,
    "Role" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT auth_pkey PRIMARY KEY ("UserUUID")
);
