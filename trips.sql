BEGIN TRANSACTION;

    CREATE TABLE events(
    id            INTEGER PRIMARY KEY,  
    title         TEXT,
    location      TEXT, 
    description   TEXT, 
    date          TEXT,
    time          TEXT,
    cost          INTEGER,  
    max_people    INTEGER
);

insert into events(title, location, description, date, time, cost)     

values

( 'Humber River Bike Ride', 
    'Old Mill station',
    'A bike Ride an a picnic down Humber River trail followed by a fun picnic ',
    '2024-6-29',
    '16:00',
    0),


(
'Sidewalk Tunes'
, '879 St Clair aveune West' 
, 'Sing and make music together at the Shalom Catholic group.'
, '2024-7-6'
, '12:00'
, 0
),

(
    'Beach Day' 
    ,'Woodbine Beach 1675 Lake Shore Blvd East' 

    ,'A day of fun at the beach.' 

    ,'2024-7-27' 
    ,'12:00'
    ,0
),

(
    'Ice cream party' 
    ,'879 St Clair Avenue West' 
    ,'An ice cream party social' 
    ,'2024-8-17' 
    ,'12:00' 
    ,0

)
;

COMMIT;
