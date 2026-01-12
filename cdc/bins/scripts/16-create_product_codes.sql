SET search_path TO bins;

-- souuce: https://developers.tabapay.com/docs/mc-product-category-codes
CREATE TABLE bins.master_product_codes (
    licensed_product_id VARCHAR(10) NOT NULL,
    gcms_product_id VARCHAR(10) NOT NULL,
    card_program_identifier VARCHAR(10) NOT NULL,
    product_class VARCHAR(20),
    description VARCHAR(100),
    consumer_vs_commercial VARCHAR(20) NOT NULL
);

