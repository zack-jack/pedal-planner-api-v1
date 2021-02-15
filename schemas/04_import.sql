-- pedals.sql
USE pedal_data;

INSERT INTO `pedals` (`id`, `brand`, `name`, `width`, `height`, `image`)
VALUES
    (3, 'Strymon', 'El Capistan', 4.0, 4.5, "strymon-elcap.png"),
    (4, 'Strymon', 'BigSky', 6.75, 5.1, "strymon-bigsky.png");

-- pedalboards.sql
USE pedalboard_data;

INSERT INTO `pedalboards` (`id`, `brand`, `name`, `width`, `height`, `image`)
VALUES
    (5, 'Creation Music Co', 'Elevation 24x12.5', 24, 12.5, "creation-elevation24125.png"),
    (6, 'Pedaltrain', 'Classic 1', 22, 12.5, "pedaltrain-classic1.png");
