CREATE TABLE artist (
    id serial primary key ,
    artist_name varchar(255) not null unique ,
    password_hash varchar(255) not null
);

CREATE TABLE genre (
    id serial primary key,
    genre_type varchar(128) unique
);

CREATE TABLE album (
    id serial primary key ,
    title_album varchar(255) not null ,
    id_artist integer references artist(id) on delete cascade not null
);

CREATE TABLE songs (
    id serial primary key,
    title_song varchar(255) not null ,
    text_song text not null ,
    id_genre integer references genre(id) on delete cascade not null,
    id_album integer references album(id) on delete cascade not null,
    unique(id_album, id_genre)
);

CREATE TABLE playlist (
    id serial primary key ,
    title varchar(255) not null ,
    description text,
    id_artist integer references artist(id) on delete cascade not null
);

CREATE TABLE playlist_songs (
    id serial primary key ,
    id_song integer references songs(id) on delete cascade not null ,
    id_playlist integer references playlist(id)  on delete cascade not null
);