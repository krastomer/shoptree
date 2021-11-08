-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: database
-- Generation Time: Nov 08, 2021 at 06:10 AM
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `addresses_customer`
--

INSERT INTO `addresses_customer` (`id`, `customer_id`, `name`, `phone_number`, `address_line`, `country`, `state`, `city`, `district`, `postal_code`, `created_at`) VALUES
(1, 1, 'kasama thongsawang', '0828702739', '12345678', 'สยาม', 'บางกอก', 'เมือง', 'ตำบล', '12345', '2021-11-08 05:05:34'),
(2, 1, 'kasama thongsawang', '0828702739', '12345678', 'สยาม', 'บางกอก', 'เมือง', 'ตำบล', '12345', '2021-11-08 05:06:13'),
(3, 1, 'kasama thongsawang', '0828702739', '12345678', 'สยาม', 'บางกอก', 'เมือง', 'ตำบล', '12345', '2021-11-08 05:09:31');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE IF NOT EXISTS `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`) VALUES
(1, 'ไม้มงคล', '2021-11-08 05:31:42'),
(2, 'ไม้ยืนต้น', '2021-11-08 05:32:22'),
(3, 'ไม้ตกแต่งสวน', '2021-11-08 05:32:31'),
(4, 'บอนไซ', '2021-11-08 05:32:37'),
(5, 'ไม้เศรษฐกิจ', '2021-11-08 05:32:46');

-- --------------------------------------------------------

--
-- Table structure for table `categories_product`
--

CREATE TABLE IF NOT EXISTS `categories_product` (
  `category_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`category_id`,`product_id`),
  KEY `cp_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `categories_product`
--

INSERT INTO `categories_product` (`category_id`, `product_id`) VALUES
(1, 1);

-- --------------------------------------------------------

--
-- Stand-in structure for view `categories_product_name`
-- (See below for the actual view)
--
CREATE TABLE IF NOT EXISTS `categories_product_name` (
`id` int(11)
,`product_id` int(11)
,`name` varchar(100)
);

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE IF NOT EXISTS `customers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `email`, `password`, `phone_number`, `created_at`) VALUES
(1, 'Kasama Thongsawang', 'krastomer@gmail.com', '$2a$08$fYT.C5/D300FfHzLLF/PU.s14XtBytJwCwmhRU9n9oX/G3F3E0FUO', '0828702739', '2021-11-05 11:13:06');

-- --------------------------------------------------------

--
-- Table structure for table `employees`
--

CREATE TABLE IF NOT EXISTS `employees` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `level` enum('Undefined','Staff','Deliver','Admin') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `employees`
--

INSERT INTO `employees` (`id`, `name`, `email`, `password`, `phone_number`, `level`, `created_at`) VALUES
(1, 'Kasama Thongsawang', 'kasama.tsw@shoptree.com', '$2a$08$fYT.C5/D300FfHzLLF/PU.s14XtBytJwCwmhRU9n9oX/G3F3E0FUO', '0828702739', 'Admin', '2021-11-05 11:14:43');

-- --------------------------------------------------------

--
-- Table structure for table `images_product`
--

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

CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `address_id` int(11) NOT NULL,
  `payment_id` int(11) DEFAULT NULL,
  `status` enum('Undefined','VerifyPayment','AcceptOrder','Prepare','Sending','Done','Failed') COLLATE utf8mb4_unicode_ci NOT NULL,
  `review` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL',
  `created_at` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `o_customer_id` (`customer_id`),
  KEY `o_address_id` (`address_id`),
  KEY `o_payment_id` (`payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

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

CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `scientific_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(2000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `scientific_name`, `description`, `price`, `created_at`) VALUES
(1, 'ยางอินเดีย', 'Ficus elastica, Moraceae', 'เป็นต้นไม้มงคล เพราะลักษณะใบกลมมนสื่อถึงความมั่งคั่ง ร่ำรวย และเจริญรุ่งเรืองในหลักฮวงจุ้ย ผู้คนจึงนิยมปลูกไว้ในบ้านเพื่อช่วยเรียกเงินทอง โชคดี และความสำเร็จ ทั้งยังเป็นต้นไม้ฟอกอากาศ ขจัดมลพิษ และลดความเครียดไปในตัวอีกด้วย โดยตำแหน่งที่เหมาะสมควรอยู่ในบริเวณทางเข้าหรือมุมมั่งคั่งของบ้านลักษณะเป็นไม้ยืนต้น ไม่ผลัดใบ มีน้ำยางสีขาวและรากอากาศ ใบเรียงสลับ ทรงไข่ ขนาดใหญ่ ปลายเรียว โคนสอบ แผ่นหนาแข็ง มีหลายสีตามสายพันธุ์ เช่น เขียว เหลือง แดง ดอกออกเป็นช่อ ขนาดเล็ก ผลเป็นทรงกลมรี การปลูกและดูแลไม่ยาก นิยมขยายพันธุ์ด้วยการปักชำและตอนกิ่ง โตได้ในดินทุกประเภท แต่ชอบดินที่ระบายน้ำดีเป็นพิเศษ ทนแล้ง-ทนแดดจัดได้ แต่ควรวางไว้ในบริเวณที่ไม่โดนแดดโดยตรงจะดีที่สุด', '5259.00', '2021-11-08 05:30:48');

-- --------------------------------------------------------

--
-- Table structure for table `products_order`
--

CREATE TABLE IF NOT EXISTS `products_order` (
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `op_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Structure for view `categories_product_name`
--
DROP TABLE IF EXISTS `categories_product_name`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `categories_product_name`  AS SELECT `categories`.`id` AS `id`, `categories_product`.`product_id` AS `product_id`, `categories`.`name` AS `name` FROM (`categories_product` join `categories` on(`categories_product`.`category_id` = `categories`.`id`)) ;

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
