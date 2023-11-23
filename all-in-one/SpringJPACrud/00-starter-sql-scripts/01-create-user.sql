-- Drop user first if they exist
DROP USER if exists 'keshav'@'%' ;

-- Now create user with prop privileges
CREATE USER 'keshav'@'%' IDENTIFIED BY 'keshav';

GRANT ALL PRIVILEGES ON * . * TO 'keshav'@'%';