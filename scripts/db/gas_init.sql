create table if not exists gas_price (
    id serial primary key,
    price numeric(10, 2) not null,
    timestamp timestamp not null
);