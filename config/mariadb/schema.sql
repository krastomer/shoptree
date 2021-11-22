-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Nov 22, 2021 at 08:26 AM
-- Server version: 10.5.12-MariaDB-1:10.5.12+maria~focal
-- PHP Version: 7.4.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shoptree`
--
CREATE DATABASE IF NOT EXISTS `shoptree` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `shoptree`;

-- --------------------------------------------------------

--
-- Table structure for table `addresses_customer`
--

DROP TABLE IF EXISTS `addresses_customer`;
CREATE TABLE IF NOT EXISTS `addresses_customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address_line` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `country` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `city` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `district` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `postal_code` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `ac_customer_id` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `categories_product`
--

DROP TABLE IF EXISTS `categories_product`;
CREATE TABLE IF NOT EXISTS `categories_product` (
  `category_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`category_id`,`product_id`),
  KEY `cp_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Stand-in structure for view `categories_product_name`
-- (See below for the actual view)
--
DROP VIEW IF EXISTS `categories_product_name`;
CREATE TABLE IF NOT EXISTS `categories_product_name` (
`id` int(11)
,`product_id` int(11)
,`name` varchar(100)
);

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
CREATE TABLE IF NOT EXISTS `customers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
CREATE TABLE IF NOT EXISTS `employees` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `level` enum('Undefined','Staff','Deliver','Admin') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `images_product`
--

DROP TABLE IF EXISTS `images_product`;
CREATE TABLE IF NOT EXISTS `images_product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) NOT NULL,
  `image_path` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `ip_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `address_id` int(11) DEFAULT NULL,
  `payment_id` int(11) DEFAULT NULL,
  `status` enum('Undefined','Pending','VerifyPayment','AcceptOrder','Prepare','Sending','Done','Failed') COLLATE utf8mb4_unicode_ci NOT NULL,
  `review` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `o_customer_id` (`customer_id`),
  KEY `o_payment_id` (`payment_id`),
  KEY `o_address_id` (`address_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
CREATE TABLE IF NOT EXISTS `payments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_path` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `scientific_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(2000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Stand-in structure for view `products_available`
-- (See below for the actual view)
--
DROP VIEW IF EXISTS `products_available`;
CREATE TABLE IF NOT EXISTS `products_available` (
`id` int(11)
,`name` varchar(100)
,`scientific_name` varchar(100)
,`description` varchar(2000)
,`price` decimal(10,2)
,`created_at` timestamp
);

-- --------------------------------------------------------

--
-- Table structure for table `products_order`
--

DROP TABLE IF EXISTS `products_order`;
CREATE TABLE IF NOT EXISTS `products_order` (
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `op_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Stand-in structure for view `products_pending`
-- (See below for the actual view)
--
DROP VIEW IF EXISTS `products_pending`;
CREATE TABLE IF NOT EXISTS `products_pending` (
`customer_id` int(11)
,`created_at` timestamp
,`product_id` int(11)
);

-- --------------------------------------------------------

--
-- Structure for view `categories_product_name`
--
DROP TABLE IF EXISTS `categories_product_name`;

DROP VIEW IF EXISTS `categories_product_name`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `categories_product_name`  AS SELECT `categories`.`id` AS `id`, `categories_product`.`product_id` AS `product_id`, `categories`.`name` AS `name` FROM (`categories_product` join `categories` on(`categories_product`.`category_id` = `categories`.`id`)) ;

-- --------------------------------------------------------

--
-- Structure for view `products_available`
--
DROP TABLE IF EXISTS `products_available`;

DROP VIEW IF EXISTS `products_available`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `products_available`  AS SELECT `products`.`id` AS `id`, `products`.`name` AS `name`, `products`.`scientific_name` AS `scientific_name`, `products`.`description` AS `description`, `products`.`price` AS `price`, `products`.`created_at` AS `created_at` FROM `products` WHERE !(`products`.`id` in (select `products`.`id` from (`products` join `products_order` on(`products_order`.`product_id` = `products`.`id`)))) ;

-- --------------------------------------------------------

--
-- Structure for view `products_pending`
--
DROP TABLE IF EXISTS `products_pending`;

DROP VIEW IF EXISTS `products_pending`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `products_pending`  AS SELECT `orders`.`customer_id` AS `customer_id`, `orders`.`created_at` AS `created_at`, `products`.`id` AS `product_id` FROM ((`orders` join `products_order` on(`orders`.`id` = `products_order`.`order_id`)) join `products` on(`products_order`.`product_id` = `products`.`id`)) WHERE `orders`.`status` = 'Undefined' ;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `addresses_customer`
--
ALTER TABLE `addresses_customer`
  ADD CONSTRAINT `ac_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `categories_product`
--
ALTER TABLE `categories_product`
  ADD CONSTRAINT `cp_category_id` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `cp_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `images_product`
--
ALTER TABLE `images_product`
  ADD CONSTRAINT `ip_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `o_address_id` FOREIGN KEY (`address_id`) REFERENCES `addresses_customer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `o_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `o_payment_id` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `products_order`
--
ALTER TABLE `products_order`
  ADD CONSTRAINT `op_order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `op_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
