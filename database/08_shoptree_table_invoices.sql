USE `shoptree`;
-- --------------------------------------------------------

--
-- Table structure for table `invoices`
--

CREATE TABLE IF NOT EXISTS `invoices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `address_id` int(11) NOT NULL,
  `payment_evidence` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('Pending','VerifyPayment','AcceptOrder','Prepare','Sending','Done','Timeout') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `invoices_customer_id` (`customer_id`),
  KEY `invoices_address_id` (`address_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
