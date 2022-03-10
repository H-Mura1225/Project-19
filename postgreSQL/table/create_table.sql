create table masterkey (
  id	serial primary key,
  email	varchar,
  key	varchar,
  start	timestamp without time zone,
  "end"	timestamp without time zone
);