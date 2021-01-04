package db

const (
	CreateUsersAccount = `create table if not exists Users
(
    Id integer primary key autoincrement,
    Name text not null,
    Surname text not null,
    Age integer not null,
    Gender text not null,
	Status text not null,
    Login text not null,
    Password text not null,
    Remove boolean not null default false
);
`
	CreatNewsTable = `create table if not exists News
(
    Id integer primary key autoincrement,
    Name text not null,
    Data text not null,
    Textarea text not null

);
`
	CreatVacansyTable = `create table if not exists Vacancy
(
    Id integer primary key autoincrement,
    Name text not null,
	Salary text not null,
    Description text not null,
    DataAdd text not null
);
`
)
