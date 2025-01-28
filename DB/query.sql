-- name: GetUser :one
SELECT * FROM USERS
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM USERS;

-- name: CreateUser :one
INSERT INTO USERS (
  password, phonenumber, email
) VALUES (
  crypt($1, gen_salt('testSalt12341808')), $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE USERS SET
  password = $2, 
  phonenumber = $3, 
  email = $4
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE USERS SET
  password = crypt($3, gen_salt('testSalt12341808'))
WHERE id = $1 
AND password = crypt($2, password);

-- name: DeleteUser :exec
UPDATE USERS SET
  is_deleted = TRUE
WHERE id = $1;




-- name: GetUserTrip :one
SELECT * FROM USER_TRIP
WHERE id = $1 LIMIT 1;

-- name: ListUserTrips :many
SELECT * FROM USER_TRIP;

-- name: CreateUserTrip :one
INSERT INTO USER_TRIP (
  user_id, trip_id, is_admin
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAdminUserTrip :exec
UPDATE USER_TRIP SET
  is_admin = $2
WHERE id = $1;

-- name: DeleteUserTrip :exec
UPDATE USER_TRIP SET
  is_deleted = TRUE
WHERE id = $1;




-- name: GetTrip :one
SELECT * FROM TRIPS
WHERE id = $1 LIMIT 1;

-- name: ListTrips :many
SELECT * FROM TRIPS;

-- name: CreateTrip :one
INSERT INTO TRIPS (
  start_at, end_at, name, bio
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTrip :exec
UPDATE TRIPS SET
  start_at = $2, 
  end_at = $3, 
  name = $4,
  bio = $5
WHERE id = $1;

-- name: DeleteTrip :exec
UPDATE TRIPS SET
  is_deleted = TRUE
WHERE id = $1;



-- name: GetActivity :one
SELECT * FROM ACTIVITIES
WHERE id = $1 LIMIT 1;

-- name: ListActivities :many
SELECT * FROM ACTIVITIES;

-- name: CreateActivity :one
INSERT INTO ACTIVITIES (
  trip_id, address, name, bio, start_at, end_at, latitude, longitude, is_hotel
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateActivity :exec
UPDATE ACTIVITIES SET
  trip_id = $2, 
  address = $3, 
  name = $4, 
  bio = $5, 
  start_at = $6, 
  end_at = $7, 
  latitude = $8, 
  longitude = $9, 
  is_hotel = $10
WHERE id = $1;

-- name: DeleteActivity :exec
UPDATE ACTIVITIES SET
  is_deleted = TRUE
WHERE id = $1;





-- name: GetImage :one
SELECT * FROM IMAGES
WHERE id = $1 LIMIT 1;

-- name: ListImages :many
SELECT * FROM IMAGES;

-- name: CreateImage :one
INSERT INTO IMAGES (
  activities_id, url, latitude, longitude, is_activity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateImage :exec
UPDATE IMAGES SET
  activities_id = $2, 
  url = $3, 
  latitude = $4,
  longitude = $5,
  is_activity = $6
WHERE id = $1;

-- name: DeleteImage :exec
UPDATE IMAGES SET
  is_deleted = TRUE
WHERE id = $1;
