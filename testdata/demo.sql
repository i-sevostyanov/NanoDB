CREATE DATABASE demo;

\use demo;

CREATE TABLE airports (
    id INTEGER PRIMARY KEY,
    code STRING NOT NULL,
    name STRING NOT NULL,
    city STRING NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    timezone STRING NOT NULL
);

INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('YKS', 'Yakutsk Airport', 'Yakutsk', 129.77099609375, 62.0932998657226562, 'Asia/Yakutsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('MJZ', 'Mirny Airport', 'Mirnyj', 114.03900146484375, 62.534698486328125, 'Asia/Yakutsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('KHV', 'Khabarovsk-Novy Airport',  'Khabarovsk', 135.18800354004, 48.5279998779300001, 'Asia/Vladivostok');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('PKC', 'Yelizovo Airport', 'Petropavlovsk', 158.453994750976562, 53.1679000854492188, 'Asia/Kamchatka');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('UUS', 'Yuzhno-Sakhalinsk Airport', 'Yuzhno-Sakhalinsk', 142.718002319335938, 46.8886985778808594, 'Asia/Sakhalin');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('VVO', 'Vladivostok International Airport', 'Vladivostok', 132.147994995117188, 43.3989982604980469, 'Asia/Vladivostok');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('LED', 'Pulkovo Airport', 'St. Petersburg', 30.2625007629394531, 59.8003005981445312, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('KGD', 'Khrabrovo Airport', 'Kaliningrad', 20.5925998687744141, 54.8899993896484375, 'Europe/Kaliningrad');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('KEJ', 'Kemerovo Airport', 'Kemorovo', 86.1072006225585938, 55.2700996398925781, 'Asia/Novokuznetsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('CEK', 'Chelyabinsk Balandino Airport', 'Chelyabinsk', 61.503300000000003, 55.3058010000000024, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('MQF', 'Magnitogorsk International Airport', 'Magnetiogorsk', 58.7556991577148438, 53.3931007385253906, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('PEE', 'Bolshoye Savino Airport', 'Perm', 56.021198272705, 57.9145011901860016, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('SGC', 'Surgut Airport', 'Surgut', 73.4018020629882812, 61.3437004089355469, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('BZK', 'Bryansk Airport', 'Bryansk', 34.1763992309999978, 53.2141990661999955, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('MRV', 'Mineralnyye Vody Airport', 'Mineralnye Vody', 43.0819015502929688, 44.2251014709472656, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('STW', 'Stavropol Shpakovskoye Airport', 'Stavropol', 42.1128005981445312, 45.1091995239257812, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('ASF', 'Astrakhan Airport', 'Astrakhan', 48.0063018799000005,46.2832984924000002, 'Europe/Samara');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('NJC', 'Nizhnevartovsk Airport', 'Nizhnevartovsk', 76.4835968017578125,60.9492988586425781, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('SVX', 'Koltsovo Airport', 'Yekaterinburg', 60.8027000427250002,56.7430992126460012, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('SVO', 'Sheremetyevo International Airport', 'Moscow', 37.4146000000000001,55.9725990000000024, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('VOZ', 'Voronezh International Airport', 'Voronezh', 39.2295989990234375,51.8142013549804688, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('VKO', 'Vnukovo International Airport', 'Moscow', 37.2615013122999983,55.5914993286000012, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('SCW', 'Syktyvkar Airport', 'Syktyvkar', 50.8451004028320312,61.6469993591308594, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('KUF', 'Kurumoch International Airport', 'Samara', 50.1642990112299998,53.5049018859860013, 'Europe/Samara');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('DME', 'Domodedovo International Airport', 'Moscow', 37.9062995910644531,55.4087982177734375, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('TJM', 'Roshchino International Airport', 'Tyumen', 65.3243026732999965,57.1896018981999958, 'Asia/Yekaterinburg');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('GOJ', 'Nizhny Novgorod Strigino International', 'Nizhniy Novgorod', 43.7840003967289988, 56.2300987243649999, 'Europe/Moscow');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('TOF', 'Bogashevo Airport', 'Tomsk', 85.2082977294920028, 56.3802986145020029, 'Asia/Krasnoyarsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('UIK', 'Ust-Ilimsk Airport', 'Ust Ilimsk', 102.56500244140625, 58.1361007690429688, 'Asia/Irkutsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('NSK', 'Norilsk-Alykel Airport', 'Norilsk', 87.3321990966796875, 69.31109619140625, 'Asia/Krasnoyarsk');
INSERT INTO airports (code, name, city, lat, lon, timezone) VALUES ('ARH', 'Talagi Airport', 'Arkhangelsk', 40.7167015075683594, 64.6003036499023438, 'Europe/Moscow');

CREATE TABLE aircrafts (
    id INTEGER PRIMARY KEY,
    code STRING NOT NULL,
    model STRING NOT NULL,
    range INTEGER NOT NULL
);

INSERT INTO aircrafts (code, model, range) VALUES ('773', 'Boeing 777-300', 11100);
INSERT INTO aircrafts (code, model, range) VALUES ('763', 'Boeing 767-300', 7900);
INSERT INTO aircrafts (code, model, range) VALUES ('SU9', 'Sukhoi Superjet-100', 3000);
INSERT INTO aircrafts (code, model, range) VALUES ('320', 'Airbus A320-200', 5700);
INSERT INTO aircrafts (code, model, range) VALUES ('321', 'Airbus A321-200', 5600);
INSERT INTO aircrafts (code, model, range) VALUES ('319', 'Airbus A319-100', 6700);
INSERT INTO aircrafts (code, model, range) VALUES ('733', 'Boeing 737-300', 4200);
INSERT INTO aircrafts (code, model, range) VALUES ('CN1', 'Cessna 208 Caravan', 1200);
INSERT INTO aircrafts (code, model, range) VALUES ('CR2', 'Bombardier CRJ-200', 2700);
