--- TABLES
CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
	price INTEGER NOT NULL,
	quantity INTEGER NOT NULL DEFAULT 0,
	image_url VARCHAR
);

--- INDEX
CREATE INDEX idx_product_id ON product (id);

--- SEED
DO $$
BEGIN
    INSERT INTO product (name, price, quantity, image_url) VALUES
    ('iPhone 13 Pro', 99900, 50, 'http://localhost:3131/image-iphone'),
    ('Samsung Galaxy S22 Ultra', 109900, 40, 'http://localhost:3131/image-galaxy'),
    ('Sony PlayStation 5', 49900, 30, 'http://localhost:3131/image-ps5'),
    ('Xbox Series X', 49900, 35, 'http://localhost:3131/image-xbox'),
    ('Canon EOS R5', 389900, 20, 'http://localhost:3131/image-canon'),
    ('Nikon Z9', 549900, 15, 'http://localhost:3131/image-nikon'),
    ('LG OLED C1 65" TV', 229900, 10, 'http://localhost:3131/image-lg-tv'),
    ('Dell XPS 15 Laptop', 179900, 25, 'http://localhost:3131/image-dell-laptop'),
    ('Amazon Echo Dot', 4999, 60, 'http://localhost:3131/image-echo'),
    ('Google Nest Hub', 8999, 45, 'http://localhost:3131/image-nest-hub'),
    ('Nintendo Switch', 29900, 55, 'http://localhost:3131/image-switch'),
    ('Fitbit Charge 5', 14900, 70, 'http://localhost:3131/image-fitbit'),
    ('Bose QuietComfort 45', 32900, 30, 'http://localhost:3131/image-bose'),
    ('Microsoft Surface Pro 8', 89900, 20, 'http://localhost:3131/image-surface'),
    ('GoPro Hero 10', 44900, 40, 'http://localhost:3131/image-gopro'),
    ('Apple Watch Series 7', 34900, 50, 'http://localhost:3131/image-watch'),
    ('Samsung Odyssey G9 Monitor', 199900, 15, 'http://localhost:3131/image-odyssey'),
    ('Logitech MX Master 3 Mouse', 9999, 65, 'http://localhost:3131/image-mouse'),
    ('Razer BlackWidow V3 Keyboard', 13999, 40, 'http://localhost:3131/image-keyboard'),
    ('Anker PowerCore 26800 Power Bank', 5999, 80, 'http://localhost:3131/image-powerbank'),
    ('Apple AirPods Pro', 24900, 50, 'http://localhost:3131/image-airpods'),
    ('Sony WH-1000XM4 Headphones', 34900, 30, 'http://localhost:3131/image-sony-headphones'),
    ('Garmin Fenix 7', 79900, 25, 'http://localhost:3131/image-garmin'),
    ('DJI Mavic 3 Drone', 249900, 10, 'http://localhost:3131/image-mavic'),
    ('Corsair Vengeance RGB Pro RAM', 6999, 55, 'http://localhost:3131/image-ram'),
    ('AMD Ryzen 9 5950X CPU', 74900, 20, 'http://localhost:3131/image-amd-cpu'),
    ('NVIDIA GeForce RTX 3080 GPU', 69900, 15, 'http://localhost:3131/image-nvidia-gpu'),
    ('Segway Ninebot ES4 Electric Scooter', 49900, 30, 'http://localhost:3131/image-scooter'),
    ('TCL 6-Series 55" TV', 64900, 10, 'http://localhost:3131/image-tcl-tv'),
    ('Roku Ultra Streaming Device', 9999, 40, 'http://localhost:3131/image-roku'),
    ('Fujifilm X-T4 Camera', 169900, 20, 'http://localhost:3131/image-fujifilm'),
    ('Lenovo Legion 5 Gaming Laptop', 119900, 15, 'http://localhost:3131/image-lenovo-laptop'),
    ('JBL Charge 5 Bluetooth Speaker', 17900, 50, 'http://localhost:3131/image-jbl-speaker'),
    ('Sony A8H OLED TV', 249900, 10, 'http://localhost:3131/image-sony-tv'),
    ('Microsoft Xbox Elite Series 2 Controller', 14900, 35, 'http://localhost:3131/image-xbox-controller'),
    ('Sennheiser HD 800 S Headphones', 149900, 10, 'http://localhost:3131/image-sennheiser-headphones'),
    ('Samsung Odyssey G7 Monitor', 79900, 20, 'http://localhost:3131/image-odyssey-monitor'),
    ('Corsair K70 RGB MK.2 Keyboard', 13900, 30, 'http://localhost:3131/image-corsair-keyboard'),
    ('Logitech G Pro X Superlight Mouse', 14900, 45, 'http://localhost:3131/image-logitech-mouse'),
    ('Apple iPad Air 5th Gen', 59900, 25, 'http://localhost:3131/image-ipad'),
    ('Google Pixel 6 Pro', 89900, 30, 'http://localhost:3131/image-pixel'),
    ('Razer Blade 15 Gaming Laptop', 179900, 15, 'http://localhost:3131/image-razer-laptop'),
    ('Jabra Elite 85t Earbuds', 19900, 40, 'http://localhost:3131/image-jabra-earbuds'),
    ('Huawei MateBook X Pro', 149900, 10, 'http://localhost:3131/image-huawei-laptop'),
    ('NVIDIA GeForce RTX 3090 GPU', 149900, 10, 'http://localhost:3131/image-nvidia-rtx3090'),
    ('SteelSeries Arctis 7 Headset', 14900, 25, 'http://localhost:3131/image-steelseries-headset'),
    ('Samsung 980 PRO SSD', 14900, 35, 'http://localhost:3131/image-samsung-ssd'),
    ('Bose SoundLink Revolve+ Speaker', 29900, 20, 'http://localhost:3131/image-bose-speaker'),
    ('Apple Mac Mini M1', 69900, 20, 'http://localhost:3131/image-mac-mini'),
    ('Canon EOS R6 Camera', 239900, 15, 'http://localhost:3131/image-canon-r6'),
    ('Logitech G Pro X Keyboard', 12900, 30, 'http://localhost:3131/image-logitech-keyboard'),
    ('AMD Ryzen 7 5800X CPU', 44900, 25, 'http://localhost:3131/image-amd-ryzen7'),
    ('Nintendo Switch OLED Model', 34900, 15, 'http://localhost:3131/image-switch-oled'),
    ('Roku Streaming Stick+', 4999, 50, 'http://localhost:3131/image-roku-stick'),
    ('Sony WH-CH710N Headphones', 9999, 40, 'http://localhost:3131/image-sony-whch710n'),
    ('Dell UltraSharp U2720Q Monitor', 54900, 20, 'http://localhost:3131/image-dell-monitor'),
    ('Samsung T7 Portable SSD', 9999, 30, 'http://localhost:3131/image-samsung-t7'),
    ('HP Pavilion 15 Gaming Laptop', 99900, 10, 'http://localhost:3131/image-hp-laptop'),
    ('Anker Soundcore 2 Speaker', 3999, 55, 'http://localhost:3131/image-anker-speaker'),
    ('Garmin Venu 2 Smartwatch', 39900, 30, 'http://localhost:3131/image-garmin-venu'),
    ('Bowers & Wilkins PX7 Headphones', 39900, 15, 'http://localhost:3131/image-bowers-wilkins'),
    ('LG UltraGear 27GN950-B Monitor', 69900, 10, 'http://localhost:3131/image-lg-monitor'),
    ('Canon RF 24-70mm f/2.8L Lens', 229900, 20, 'http://localhost:3131/image-canon-lens'),
    ('ASUS ROG Strix Scar 15 Gaming Laptop', 199900, 10, 'http://localhost:3131/image-asus-laptop'),
    ('Bose Frames Audio Sunglasses', 19900, 30, 'http://localhost:3131/image-bose-frames'),
    ('Sony WF-1000XM4 Earbuds', 24900, 25, 'http://localhost:3131/image-sony-wf1000xm4'),
    ('DJI OM 5 Gimbal', 15900, 35, 'http://localhost:3131/image-dji-om5'),
    ('Intel Core i9-12900K CPU', 59900, 15, 'http://localhost:3131/image-intel-cpu'),
    ('NVIDIA GeForce RTX 3070 GPU', 49900, 20, 'http://localhost:3131/image-nvidia-rtx3070'),
    ('Logitech MX Keys Keyboard', 9999, 40, 'http://localhost:3131/image-logitech-mxkeys'),
    ('Samsung Odyssey G5 Monitor', 34900, 15, 'http://localhost:3131/image-odyssey-g5'),
    ('Google Pixel Buds A-Series', 9999, 30, 'http://localhost:3131/image-pixel-buds'),
    ('AMD Ryzen 5 5600X CPU', 29900, 25, 'http://localhost:3131/image-amd-ryzen5'),
    ('Microsoft Xbox Wireless Controller', 5999, 45, 'http://localhost:3131/image-xbox-controller'),
    ('Corsair HS70 Bluetooth Headset', 7999, 20, 'http://localhost:3131/image-corsair-hs70'),
    ('Razer DeathAdder V2 Mouse', 5999, 35, 'http://localhost:3131/image-razer-mouse'),
    ('Apple TV 4K', 17900, 25, 'http://localhost:3131/image-apple-tv'),
    ('Sony Alpha a7 IV Camera', 249900, 10, 'http://localhost:3131/image-sony-a7iv'),
    ('Garmin Forerunner 945 Smartwatch', 59900, 20, 'http://localhost:3131/image-garmin-forerunner'),
    ('Samsung 34-inch Odyssey G3 Ultrawide Monitor', 34900, 15, 'http://localhost:3131/image-odyssey-g3'),
    ('JBL Flip 5 Bluetooth Speaker', 11900, 35, 'http://localhost:3131/image-jbl-flip5'),
    ('Sony SRS-XB33 Portable Speaker', 14900, 30, 'http://localhost:3131/image-sony-srsxb33'),
    ('Apple Magic Mouse', 7499, 45, 'http://localhost:3131/image-apple-magicmouse'),
    ('ASUS ROG Swift PG279Q Monitor', 49900, 20, 'http://localhost:3131/image-asus-monitor'),
    ('GoPro Hero 9', 49900, 20, 'http://localhost:3131/image-gopro-hero9'),
    ('Amazon Kindle Paperwhite', 12900, 50, 'http://localhost:3131/image-kindle'),
    ('Fitbit Versa 3', 19900, 25, 'http://localhost:3131/image-fitbit-versa'),
    ('NVIDIA GeForce RTX 3060 Ti GPU', 39900, 15, 'http://localhost:3131/image-nvidia-rtx3060ti'),
    ('Samsung Galaxy Tab S7', 64900, 10, 'http://localhost:3131/image-galaxy-tab'),
    ('LG CX 48" OLED TV', 119900, 10, 'http://localhost:3131/image-lg-cx'),
    ('SteelSeries Apex Pro TKL Keyboard', 15900, 30, 'http://localhost:3131/image-steelseries-apex'),
    ('Apple Pencil (2nd Generation)', 9999, 40, 'http://localhost:3131/image-apple-pencil'),
    ('Logitech G Pro X Superlight Wireless Mouse', 9999, 40, 'http://localhost:3131/image-logitech-gprox-superlight'),
    ('Sony Alpha a6600 Camera', 109900, 15, 'http://localhost:3131/image-sony-alpha-a6600'),
    ('Samsung Galaxy Buds Pro', 19900, 30, 'http://localhost:3131/image-galaxy-buds-pro');
END;
$$;
