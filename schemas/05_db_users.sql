CREATE USER 'pedals'@'%' IDENTIFIED BY 'DONT_USE_ME';
GRANT INSERT, SELECT, UPDATE ON pedal_db.* TO 'pedals';

CREATE USER 'pedalboards'@'%' IDENTIFIED BY 'DONT_USE_ME';
GRANT INSERT, SELECT, UPDATE ON pedalboard_db.* TO 'pedalboards';
