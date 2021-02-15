CREATE USER 'pedals'@'%' IDENTIFIED BY 'DONT_USE_ME';
GRANT INSERT, SELECT, UPDATE ON pedal_data.* TO 'pedals';

CREATE USER 'pedalboards'@'%' IDENTIFIED BY 'DONT_USE_ME';
GRANT INSERT, SELECT, UPDATE ON pedalboard_data.* TO 'pedalboards';
