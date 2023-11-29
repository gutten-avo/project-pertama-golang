CREATE TABLE "public"."tbl_user" ( 
  "id" SERIAL,
  "first_name" VARCHAR(250) NOT NULL,
  "last_name" VARCHAR(250) NOT NULL,
  "phone_number" VARCHAR(15) NULL,
  CONSTRAINT "tbl_user_pkey" PRIMARY KEY ("id")
);
