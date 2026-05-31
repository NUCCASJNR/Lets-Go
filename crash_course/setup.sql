 -- sql script
CREATE DATABASE IF NOT EXISTS course_db;
       CREATE USER IF NOT EXISTS 'course_user'@'localhost' IDENTIFIED BY 'Course_pwd123@';
              GRANT ALL PRIVILEGES ON course_db.* TO 'course_user'@'localhost';
                                      GRANT SELECT ON performance_schema.* TO 'course_user'@'localhost';
FLUSH PRIVILEGES;