CREATE TABLE shopclient(
   id SERIAL PRIMARY KEY     NOT NULL,
   `name`           VARCHAR(25)    NOT NULL,
   phone_number   VARCHAR(25) NOT NULL,
   email VARCHAR(25) DEFAULT NULL,
   shop_name VARCHAR(50) NOT NULL,
   service_type VARCHAR(20) NOT NULL,
   verified_email boolean DEFAULT false ,
   verified_phone boolean DEFAULT false,
   street VARCHAR(30) NOT NULL,
   city VARCHAR(30) NOT NULL,
   `state` VARCHAR(30) NOT NULL,
   pen_card_url VARCHAR(1000) DEFAULT '',
   aadhar_card_url VARCHAR(1000) DEFAULT '',
   created timestamp NOT NULL DEFAULT NOW(),
   updated timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE shopper_bank_details (
    shop_client_id int NOT NULL,
    account_number VARCHAR(50) NOT NULL,
    branch_name VARCHAR(50) NOT NULL,
    ifsc_code VARCHAR(10) NOT NULL,
    created timestamp NOT NULL DEFAULT NOW(),
    updated timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (shop_client_id),
    CONSTRAINT fk_shop_client_id FOREIGN KEY (shop_client_id) REFERENCES shopclient (id)
);

CREATE TABLE beauty_parlour(
    shop_client_parlour_id int NOT NULL,
    catlog_name VARCHAR(30) NOT NULL,
    price INT NOT NULL,
    `description` VARCHAR(200) DEFAULT NULL,
    created timestamp NOT NULL DEFAULT NOW(),
    updated timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (shop_client_id),
    CONSTRAINT fk_shop_client_parlour_id FOREIGN KEY (shop_client_parlour_id) REFERENCES shopclient (id)
)


