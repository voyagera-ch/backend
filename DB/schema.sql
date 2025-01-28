CREATE EXTENSION pgcrypto;


CREATE TABLE USERS (
  id   BIGSERIAL PRIMARY KEY,
  password text NOT NULL,
  phonenumber  text,
  email text,
  is_deleted boolean DEFAULT FALSE
);

CREATE TABLE TRIPS (
  id   BIGSERIAL PRIMARY KEY,
  start_at timestamp,
  end_at timestamp,
  name text NOT NULL,
  bio  text,
  is_deleted boolean DEFAULT FALSE
);

CREATE TABLE USER_TRIP (
  id   BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL references USERS(id),
  trip_id BIGSERIAL references TRIPS(id),
  is_admin  boolean,
  is_deleted boolean DEFAULT FALSE
);


CREATE TABLE ACTIVITIES (
  id   BIGSERIAL PRIMARY KEY,
  trip_id BIGSERIAL references TRIPS(id),
  address text      NOT NULL,
  name text,
  bio text,
  start_at timestamp,
  end_at timestamp,
  latitude float,
  longitude float,
  is_hotel boolean,
  is_deleted boolean  DEFAULT FALSE
);


CREATE TABLE IMAGES (
  id   BIGSERIAL PRIMARY KEY,
  activities_id BIGSERIAL references ACTIVITIES(id),
  url text NOT NULL,
  latitude float,
  longitude float,
  is_activity boolean,
  is_deleted boolean DEFAULT FALSE
);