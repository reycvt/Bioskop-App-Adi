CREATE DATABASE bioskop_adi;

\c bioskop_adi;

CREATE TABLE bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    lokasi VARCHAR(100) NOT NULL,
    rating REAL
);
