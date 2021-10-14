USE `shoptree`;
-- --------------------------------------------------------

--
-- Table structure for table `invoice_products`
--

CREATE TABLE IF NOT EXISTS `invoice_products` (
  `invoice_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  PRIMARY KEY (`invoice_id`,`product_id`),
  KEY `invoice_products_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
