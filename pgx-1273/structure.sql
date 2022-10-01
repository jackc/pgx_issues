create table example
(
   data jsonb,
   id   serial
       constraint example_pk
           primary key
);

create unique index example_id_uindex
   on example (id);
