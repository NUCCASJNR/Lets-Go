 -- sql script
CREATE DATABASE IF NOT EXISTS movies_db;
       CREATE USER IF NOT EXISTS 'm_user'@'127.0.0.1' IDENTIFIED BY 'M_pwd123@';
              GRANT ALL PRIVILEGES ON movies_db.* TO 'm_user'@'127.0.0.1';
                                      GRANT SELECT ON performance_schema.* TO 'm_user'@'127.0.0.1';
FLUSH PRIVILEGES;
