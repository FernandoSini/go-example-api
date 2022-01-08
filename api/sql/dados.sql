INSERT INTO usuarios (name, nick, email, password)
VALUES
("teste","testeteste", "teste@gmail.com","$2a$10$RUVpUGJJz24jjXQnIlMHVO.xtakbuGZrt/vwAQFhU3a69FvBVkzci"),
("teste2","teste2teste2", "teste2@gmail.com","$2a$10$RUVpUGJJz24jjXQnIlMHVO.xtakbuGZrt/vwAQFhU3a69FvBVkzci"),
("teste3","teste3teste3", "teste3@gmail.com","$2a$10$RUVpUGJJz24jjXQnIlMHVO.xtakbuGZrt/vwAQFhU3a69FvBVkzci");


INSERT INTO followers (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts (title, content, author_id) VALUES 
("p1","p2p2",1),
("p2","p3p3",3),
("p3","p4p4",1);