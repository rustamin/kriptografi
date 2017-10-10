-- phpMyAdmin SQL Dump
-- version 4.7.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: 03 Okt 2017 pada 05.02
-- Versi Server: 10.1.22-MariaDB
-- PHP Version: 7.1.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang-example`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `coba`
--

CREATE TABLE `coba` (
  `nama` varchar(100) NOT NULL,
  `alamat` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `coba`
--

INSERT INTO `coba` (`nama`, `alamat`) VALUES
('rustam', 'alamat1'),
('rustam', 'alamat2'),
('rustam', 'alamat3'),
('rustam', 'alamat4'),
('amel', 'alamat3'),
('ameel', 'alamat4');

-- --------------------------------------------------------

--
-- Struktur dari tabel `code_verification`
--

CREATE TABLE `code_verification` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `code` varchar(120) NOT NULL,
  `active` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `cookies`
--

CREATE TABLE `cookies` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `code` varchar(140) NOT NULL,
  `cookie_value` varchar(150) DEFAULT NULL,
  `active` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1 active, 0 tdk'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `cookies`
--

INSERT INTO `cookies` (`id`, `user_id`, `code`, `cookie_value`, `active`) VALUES
(22, 1, '036738', NULL, 0),
(23, 1, '828725', NULL, 0),
(24, 1, '281239', NULL, 0),
(25, 1, '987731', NULL, 0),
(26, 1, '628715', NULL, 0),
(27, 1, '875898', NULL, 0),
(28, 1, '419928', NULL, 0),
(29, 1, '632341', NULL, 0),
(30, 0, '', 'bFiYPVPVFunvoKKiIZzEIkMFp', 1),
(31, 1, '951382', 'wnKWkVwQzPZCKXHLjgkyvZnWv', 1),
(32, 1, '175524', 'eEUwcVzSFDUNUYcmkRXYegQHq', 1),
(33, 1, '135132', 'YLktHZvBenHtmMbHfoajxTify', 1),
(34, 1, '461224', 'NguTwrCnooIJnLsFexAunPGWC', 1),
(35, 1, '356839', 'UDOBxlhKDFUjfgWjqpIIRrYgf', 1),
(36, 1, '483114', 'pAtMAKEGRxSATIxBrUsxmoumH', 1),
(37, 1, '157917', 'orQuNtapaVhyhapETAxmAHSyS', 1),
(38, 1, '845224', 'LqObcWNUvLjeOeMSKvnooTyfG', 1),
(39, 1, '587149', 'iKwDsNyCbwZLmUrLHpRIMOtnz', 1),
(40, 1, '617774', 'HtlgdhqEhtdriwjFvJGNEIiKu', 1),
(41, 1, '963211', 'iDaPnAImpcbwdMTpcrREGUiow', 1),
(42, 1, '376484', 'fraNCgLcerypjvNqxfmAGrReo', 1),
(43, 1, '181486', 'QEyVnmeuTvHlciyDkxXAUDiQK', 1),
(44, 1, '259392', 'auZiNvSXePmNqdhuMrPixAujI', 1),
(45, 1, '298468', NULL, 0),
(46, 1, '821356', 'cMllcPGPlEGjsMppZudxlKsiz', 1),
(47, 1, '847545', 'DlWEBdthpcBVLvmbgmzurCAfe', 1),
(48, 1, '965137', 'gAhUlnafHBQXNBysvUbQBKnwD', 1),
(49, 1, '424935', 'renhuDSoTlemtzjocllBXCWfx', 1),
(50, 1, '889622', 'vaXVJMOxRJmcxqoJQpgBYedAU', 1),
(51, 1, '319358', 'hjywFNEghzAXXfjsTMjIlXhsQ', 1),
(52, 1, '365819', 'XhYatZBzPKBimDzqGmJtpWRUk', 1),
(53, 1, '387679', 'IlkZXsndPKcYUtkYijkQClkki', 1),
(54, 1, '794234', 'ZjbfFHNDUbqyplckJLjgSCpnT', 1),
(55, 1, '528896', 'JvdDzjlfdgSTPceZCRvrXtDWw', 1),
(56, 1, '823224', 'CjufjepsjMFDGzNINuXaxUmMb', 1),
(57, 1, '538819', 'OpTYLNDNRdEcOSRYkvXWfpuCa', 1),
(58, 1, '619721', 'lotZHnINUoyonGZxgAtQnuyhL', 1),
(59, 1, '732947', 'RBSFLUBhpvqANDYEvpGVtCrQq', 1),
(60, 1, '933129', 'SojTgVGrZztByvErfngnoWLyC', 1),
(61, 1, '647674', 'okKbywysuPlrBDcjrPzxOYgoT', 1),
(62, 1, '318353', 'atDmzNFIZBuGuyjysJmLTBtdp', 1),
(63, 1, '417184', 'nJKjmrmNYYMOiHmOCIWozxqSY', 1),
(64, 1, '177744', 'kYlSmMJayNWCOiXFZRbTkmVmG', 1),
(65, 6, '433322', NULL, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `email` varchar(120) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(120) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `email`, `username`, `password`) VALUES
(1, 'rustamjr.11@gmail.com', 'coba', '$2a$10$si1oL3133XgLpUUE6tl4CeGFIamFhT3/XzBuDQ5cER2RVPI/jM7GC'),
(2, 'mail@mail.com', 'mail', ' ??p|x}A@'),
(6, 'rustamjr.2@gmail.com', 'rustam', 'xuruopfukuqj');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `code_verification`
--
ALTER TABLE `code_verification`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `cookies`
--
ALTER TABLE `cookies`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `cookie_value_unique` (`cookie_value`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `code_verification`
--
ALTER TABLE `code_verification`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `cookies`
--
ALTER TABLE `cookies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=66;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
