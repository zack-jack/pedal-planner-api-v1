-- pedals.sql
USE pedal_db;

INSERT INTO `pedals` (`id`, `brand`, `name`, `width`, `height`, `image`)
VALUES
    (3, 'Strymon', 'El Capistan', 4.0, 4.5, 'image/pedal/strymon-elcap.png'),
    (4, 'Strymon', 'BigSky', 6.75, 5.1, 'image/pedal/strymon-bigsky.png');

-- pedalboards.sql
USE pedalboard_db;

INSERT INTO `pedalboards` (`id`, `brand`, `name`, `width`, `height`, `image`)
VALUES
    (5, 'Creation Music Co', 'Elevation 24x12.5', 24, 12.5, 'image/pedalboard/creation-elevation-24-125.png'),
    (6, 'Pedaltrain', 'Classic 1', 22, 12.5, 'image/pedalboard/pedaltrain-classic1.png');
