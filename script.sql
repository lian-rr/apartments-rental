create table buildings
(
  id          int auto_increment
    primary key,
  shortName   varchar(50)            null,
  fullName    varchar(50)            null,
  address     varchar(100)           null,
  manager     varchar(50)            null,
  phone       varchar(50)            null,
  description varchar(255)           null,
  active      tinyint(1) default '0' not null
);

create table apartments
(
  id        int auto_increment
    primary key,
  number    varchar(10)  null,
  bathrooms float        null,
  bedrooms  int          null,
  rooms     int          null,
  details   varchar(255) null,
  building  int          null,
  active    tinyint(1)   null,
  constraint apartments_buildings_fk
  foreign key (building) references buildings (id)
);

create table guest
(
  id        int auto_increment
    primary key,
  firstName varchar(50) charset utf8  null,
  lastName  varchar(50) charset utf8  null,
  birthDay  date                      null,
  details   varchar(500) charset utf8 null,
  gender    varchar(50)               null,
  active    tinyint(1)                null
);

create table bookings
(
  id        int auto_increment
    primary key,
  status    varchar(100) null,
  startDate date         null,
  endDate   date         null,
  details   varchar(255) null,
  apartment int          null,
  guest     int          null,
  active    tinyint(1)   null,
  constraint bookings_apartments_id_fk
  foreign key (apartment) references apartments (id),
  constraint bookings_guest_id_fk
  foreign key (guest) references guest (id)
);


