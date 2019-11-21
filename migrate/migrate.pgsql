--create database crudbd 
-- ENCODING "utf-8";
--ALTER DATABASE crudbd SET datestyle TO ISO, DMY;

create table cliente (
  clienteID integer,
  nome      varchar(40),
  endereco  varchar(30),
  datanasc  date,
  constraint pk_cliente PRIMARY KEY (clienteID)
);

-- curl --header "Content-Type: application/json" --request GET  --data '{"clienteID":"xyz","nome":"xyz", "endereco":"xyz","datanasc":"21/11/2019"}' http://localhost:8080/cliente/insere