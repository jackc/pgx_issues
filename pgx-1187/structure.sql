create table test (hs hstore, hs2 hstore);
create table test2 (hs hstore, hs2 hstore);
insert into test2 (hs, hs2) values ('a=>1','a=>2');
insert into test2 (hs, hs2) values ('b=>1','b=>2');
