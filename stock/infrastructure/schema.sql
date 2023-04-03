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
    unique index(figi)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS us_stock_daily_quote;
CREATE TABLE us_stock_daily_quote (
    ID bigint unsigned not null auto_increment,
    stock_id varchar(50) not null,
    symbol         varchar(10) not null,
    current float,
    last_close float,
    high float,
    low float,
    open float,
    chg float,
    percent float,
    high52w float,
    low52w float,
    volume bigint unsigned,
    amount float,
    market_capital bigint unsigned,
    total_shares bigint unsigned,
    dividend float,
    dividend_yield float,
    eps float,
    turnover_rate float,
    volume_ratio float,
    amplitude float,
    current_year_percent float,
    issue_date bigint unsigned,
    pe_ttm float,
    pe_lyr float,
    pe_forecast float,
    navps float,
    pb float,
    psr float,
    timestamp bigint unsigned,
    trading_day varchar(20),
    created_at timestamp,
    updated_at timestamp null,
    deleted_at timestamp null,
    primary key(ID),
    unique index(stock_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;