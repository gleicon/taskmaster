drop database if exists taskmaster;
create database taskmaster character set utf8 collate utf8_unicode_ci;
grant all privileges on taskmaster.* to 'taskmaster'@'localhost' identified by '4ssTruck3r';
use taskmaster;

create table User (
  Id integer not null auto_increment,
  Email varchar(50) unique not null,
  Passwd varchar(40) not null,
  FullName varchar(80) null,
  IsActive boolean not null,
  primary key(id)
);

create table Crontab (
  Id integer not null auto_increment,
  UserId integer not null,
  CrontabString varchar(100) not null,
  DateCreation timestamp not null default current_timestamp, 
  DateLastModified timestamp, 
  DateLastExecution timestamp,
  Errors integer default 0,
  Success integer default 0,
  IsActive boolean not null,
  SrcURL    varchar(1024),
  DstURL    varchar(1024),
  SrcMethod varchar(16),
  DstMethod varchar(16),
  SrcBody   varchar(512),
  Policy    varchar(16),
  primary key(id)
);

