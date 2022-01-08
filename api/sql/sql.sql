Create DATABASE IF NOT EXISTS schema_go;
USE schema_go;

DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(200) not null,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE followers(
    usuario_id int not null,
    FOREIGN KEY(usuario_id)
    REFERENCES  usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
)ENGINE=INNODB;

CREATE TABLE posts(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    FOREIGN KEY(author_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp() 
) ENGINE=INNODB;
