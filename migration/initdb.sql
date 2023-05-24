-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 24 Bulan Mei 2023 pada 02.17
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!40101 SET NAMES utf8mb4 */
;
--
-- Database: `steradian`
--

-- --------------------------------------------------------
--
-- Struktur dari tabel `admin`
--
CREATE DATABASE `steradian`;
USE `steradian`;
CREATE TABLE `admin` (
  `id` int(11) UNSIGNED NOT NULL,
  `email` varchar(250) NOT NULL,
  `password` varchar(250) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
--
-- Dumping data untuk tabel `admin`
--

INSERT INTO `admin` (`id`, `email`, `password`)
VALUES (1, 'audeta35@mail.com', 'asdasdasd'),
  (3, 'audeta@mail.co', 'asdasdasd'),
  (4, 'audeta@mail.co', 'asdasdasd'),
  (5, '', ''),
  (6, 'audeta@email.co', 'asdasdasd');
-- --------------------------------------------------------
--
-- Struktur dari tabel `article`
--

CREATE TABLE `article` (
  `id` int(11) UNSIGNED NOT NULL,
  `title` varchar(250) NOT NULL DEFAULT '',
  `body` longtext NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
-- --------------------------------------------------------
--
-- Struktur dari tabel `cars`
--

CREATE TABLE `cars` (
  `carId` int(11) UNSIGNED NOT NULL,
  `name` varchar(250) NOT NULL,
  `carType` varchar(250) DEFAULT NULL,
  `rating` varchar(11) DEFAULT NULL,
  `fuel` int(11) DEFAULT NULL,
  `image` text DEFAULT NULL,
  `hourRate` varchar(250) DEFAULT NULL,
  `dayRate` varchar(250) DEFAULT NULL,
  `monthRate` varchar(250) DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
--
-- Dumping data untuk tabel `cars`
--

INSERT INTO `cars` (
    `carId`,
    `name`,
    `carType`,
    `rating`,
    `fuel`,
    `image`,
    `hourRate`,
    `dayRate`,
    `monthRate`
  )
VALUES (
    1,
    'Rush',
    'Minibus',
    '4',
    60,
    'jeep.png',
    '12',
    '12',
    '9'
  ),
  (
    2,
    'avanza',
    'city car',
    '3',
    50,
    'suv.png',
    '12',
    '6',
    '1'
  );
-- --------------------------------------------------------
--
-- Struktur dari tabel `orders`
--

CREATE TABLE `orders` (
  `orderId` int(11) UNSIGNED NOT NULL,
  `pickUpLoc` varchar(250) NOT NULL,
  `dropOffLoc` varchar(250) DEFAULT NULL,
  `pickupDate` text DEFAULT NULL,
  `dropOffDate` text DEFAULT NULL,
  `pickUpTime` text DEFAULT NULL,
  `cardId` int(11) DEFAULT NULL,
  `userId` int(11) DEFAULT NULL,
  `adminId` int(11) DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
--
-- Dumping data untuk tabel `orders`
--

INSERT INTO `orders` (
    `orderId`,
    `pickUpLoc`,
    `dropOffLoc`,
    `pickupDate`,
    `dropOffDate`,
    `pickUpTime`,
    `cardId`,
    `userId`,
    `adminId`
  )
VALUES (2, 'asd', 'asd', 'asd', 'asd', 'asd', 1, 1, 1),
  (
    3,
    'Tanggerang',
    'Banten',
    '2023-05-11',
    '2023-05-11',
    'asdasd',
    1,
    1,
    1
  ),
  (
    4,
    'Banten',
    'Depok',
    '2023-05-10',
    '2023-05-17',
    '12:00',
    2,
    1,
    4
  ),
  (
    5,
    'test',
    'test',
    '2023-05-18',
    '2023-05-11',
    '12',
    1,
    1,
    1
  );
-- --------------------------------------------------------
--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `userId` int(11) UNSIGNED NOT NULL,
  `email` varchar(250) NOT NULL,
  `phoneNumber` varchar(250) DEFAULT NULL,
  `city` varchar(250) DEFAULT NULL,
  `zip` varchar(250) DEFAULT NULL,
  `message` varchar(250) DEFAULT NULL,
  `password` varchar(250) NOT NULL,
  `username` varchar(250) NOT NULL,
  `address` varchar(250) DEFAULT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (
    `userId`,
    `email`,
    `phoneNumber`,
    `city`,
    `zip`,
    `message`,
    `password`,
    `username`,
    `address`
  )
VALUES (
    1,
    'audeta@mail.co',
    '081212121212',
    'Bogor',
    '16820',
    '',
    'asdasdasd',
    'audeta',
    'Jln Apel'
  ),
  (
    2,
    'audeta@mail.co',
    '081212121212',
    'Bogor',
    '16820',
    '',
    'asdasdasd',
    'audeta',
    'Jln Apel'
  );
--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `admin`
--
ALTER TABLE `admin`
ADD PRIMARY KEY (`id`);
--
-- Indeks untuk tabel `article`
--
ALTER TABLE `article`
ADD PRIMARY KEY (`id`);
--
-- Indeks untuk tabel `cars`
--
ALTER TABLE `cars`
ADD PRIMARY KEY (`carId`);
--
-- Indeks untuk tabel `orders`
--
ALTER TABLE `orders`
ADD PRIMARY KEY (`orderId`);
--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
ADD PRIMARY KEY (`userId`);
--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `admin`
--
ALTER TABLE `admin`
MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 7;
--
-- AUTO_INCREMENT untuk tabel `article`
--
ALTER TABLE `article`
MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT untuk tabel `cars`
--
ALTER TABLE `cars`
MODIFY `carId` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 3;
--
-- AUTO_INCREMENT untuk tabel `orders`
--
ALTER TABLE `orders`
MODIFY `orderId` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 6;
--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
MODIFY `userId` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;
COMMIT;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;