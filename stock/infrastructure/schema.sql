DROP DATABASE IF EXISTS equities;
CREATE DATABASE equities;
USE equities;
DROP TABLE IF EXISTS us_stock_meta_data;
CREATE TABLE us_stock_meta_data (
    stock_id varchar(50) not null,
    currency varchar(10) not null,
    description    varchar(200),
    display_symbol  varchar(10),
    figi           varchar(30) not null ,
    is_in          varchar(100),
    mic            varchar(10),
    share_class_figi varchar(30),
    symbol         varchar(10) not null,
    symbol2      varchar(10),
    equity_type     varchar(20),
    created_at timestamp,
    updated_at timestamp null,
    deleted_at timestamp null,
    primary key(stock_id),
    index(figi),
    index(symbol)
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;