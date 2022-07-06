CREATE TABLE users(
   id SERIAL PRIMARY KEY     NOT NULL,
   name  VARCHAR(25)   NOT NULL,
   dial_code VARCHAR(5) NOT NULL,
   phone_number   VARCHAR(25) NOT NULL,
   email VARCHAR(25) DEFAULT NULL,
   verified_email boolean DEFAULT false,
   verified_phone boolean DEFAULT false,
   street VARCHAR(30) NOT NULL,
   city VARCHAR(30) NOT NULL,
   state VARCHAR(30) NOT NULL,
   created timestamp DEFAULT NULL,
   updated timestamp DEFAULT NULL
);
