--- TABLES
CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
	price INTEGER NOT NULL,
	quantity INTEGER NOT NULL DEFAULT 0,
	image_url VARCHAR
);

--- INDEX
CREATE INDEX idx_product_id ON product (id);

--- SEED
DO $$
BEGIN
    INSERT INTO product (name, description, price, quantity, image_url) VALUES
    ('iPhone 14 Pro', 'The latest flagship smartphone from Apple.', 99900, 50, 'https://d1p6nzzdute2g.cloudfront.net/lojas/loja-1517/079fb574-e319-493d-a49f-e8fed3994201'),
    ('Samsung Galaxy S22 Ultra', 'Top-of-the-line smartphone with cutting-edge features.', 109900, 40, 'https://images.samsung.com/is/image/samsung/p6pim/pt/2202/gallery/pt-galaxy-s22-ultra-s908-sm-s908bzggeub-530885352?$650_519_PNG$'),
    ('Sony PlayStation 5', 'Next-generation gaming console with stunning graphics and performance.', 49900, 30, 'https://files.tecnoblog.net/wp-content/uploads/2020/11/playstation_5_produto-700x700.png'),
    ('Xbox Series X', 'Powerful gaming console for immersive gaming experiences.', 49900, 35, 'https://assets.xboxservices.com/assets/fb/d2/fbd2cb56-5c25-414d-9f46-e6a164cdf5be.png?n=XBX_A-BuyBoxBGImage01-D.png'),
    ('Canon EOS R5', 'Professional-grade mirrorless camera with outstanding image quality.', 389900, 20, 'https://i1.adis.ws/i/canon/eos-r5_front_rf24-105mmf4lisusm_square_32c26ad194234d42b3cd9e582a21c99b'),
    ('Nikon Z9', 'High-performance flagship camera for professional photographers.', 549900, 15, 'https://cdn-4.nikon-cdn.com/e/Q5NM96RZZo-RRZZFeeMiveET0gVQ--AxJI7g-xcIXdWk_L9xvLGcj82WzBlHRJ_B/Views/z-9-24-70-front.png'),
    ('LG OLED C1 65" TV', 'Immersive OLED TV with stunning picture quality.', 229900, 10, 'https://www.lg.com/content/dam/channel/wcms/pt/images/oled-tv/2021/oled-buying-guide/TV-OLED-Buying-Guide-06-Specs-Desktop-A-4.png'),
    ('Dell XPS 15 Laptop', 'Powerful laptop with high-resolution display and premium build quality.', 179900, 25, 'https://upload.wikimedia.org/wikipedia/commons/1/18/Dell_XPS_15_%282015%29.png'),
    ('Amazon Echo Dot', 'Smart speaker with Alexa voice control for hands-free convenience.', 4999, 60, 'https://crdms.images.consumerreports.org/prod/products/cr/models/402087-smart-speakers-amazon-echo-dot-4th-gen-10015594.png'),
    ('Nintendo Switch', 'Versatile gaming console for gaming at home or on the go.', 29900, 55, 'https://assets.nintendo.com/image/upload/f_auto/q_auto/dpr_2.0/c_scale,w_300/ncom/en_US/switch/refresh/switchindock'),
    ('Nike Air Max 270', 'Iconic Nike sneakers with Air Max cushioning for comfort.', 79900, 20, 'https://acdn.mitiendanube.com/stores/850/013/products/nike-270-preto-2-e1ff6e8324ee22812c16314691620531-640-0.png'),
    ('Instant Pot Duo 6-Quart', 'Multi-functional pressure cooker for easy and quick cooking.', 7999, 30, 'https://instanthome.in/images/Instant-Pot-Duo-6QT.png'),
    ('Samsung 55" 4K Smart TV', 'High-resolution smart TV with immersive viewing experience.', 69900, 15, 'https://images.samsung.com/is/image/samsung/p6pim/br/un60au7700gxzd/gallery/br-uhd-au7000-un60au7700gxzd-537668738?$650_519_PNG$'),
    ('Apple AirPods Pro', 'Wireless earbuds with active noise cancellation for immersive sound.', 19900, 40, 'https://png.pngtree.com/png-clipart/20230504/ourmid/pngtree-airpods-png-image_7081756.png'),
    ('KitchenAid Stand Mixer', 'Powerful stand mixer for effortless baking and cooking.', 34900, 20, 'https://www.mitchellcooper.co.uk/media/catalog/product/1/2/12847-05-1.png'),
    ('Adidas Ultraboost Running Shoes', 'High-performance running shoes with responsive cushioning.', 99900, 25, 'https://acdn.mitiendanube.com/stores/001/038/770/products/hq6343-png-8efc6199de615b391816964239043313-640-0.png'),
    ('Cuisinart 14-Cup Food Processor', 'Versatile food processor for easy meal preparation.', 15900, 20, 'https://res.cloudinary.com/hksqkdlah/image/upload/SIL_Hamilton-Beach_14-Cup-Professional-Food-Processor_7941_nqggbn.png'),
    ('Fujifilm Instax Mini 11', 'Compact instant camera for capturing memories on-the-go.', 6999, 35, 'https://mediamarkt.pt/cdn/shop/files/pic_mini12_GREEN_01.webp?v=1685462320'),
    ('Le Creuset Dutch Oven', 'Premium enameled cast iron Dutch oven for cooking delicious meals.', 349900, 10, 'https://res.cloudinary.com/hksqkdlah/image/upload/41956-sil-dutch-ovens-le-creuset-7-14-quart-round-dutch-oven-ls2501-28-1.png'),
    ('Vitamix Blender', 'High-performance blender for smoothies, soups, and more.', 39900, 15, 'https://www.vitamix.com/content/dam/vitamix/migration/media/other/images/a/A3500i_BrushedStainless_Front_64oz_Interlock.png'),
    ('Bose QuietComfort 35 II Headphones', 'Noise-cancelling headphones for immersive audio experience.', 29900, 25, 'https://avitlifestyle.com/cdn/shop/products/cq5dam.web.600.600_e45b65a5-72ce-42a9-8fa8-cb7f563f0e82_1024x1024@2x.png?v=1557482752'),
    ('Lego Millennium Falcon', 'Iconic Star Wars Lego set for fans and collectors.', 79900, 10, 'https://cdn.rebrickable.com/media/thumbs/mocs/moc-116416.jpg/1000x800.jpg'),
    ('Crock-Pot Slow Cooker', 'Convenient slow cooker for delicious, fuss-free meals.', 2999, 30, 'https://crdms.images.consumerreports.org/prod/products/cr/models/397829-slow-cookers-crock-pot-6-quart-cook-carry-manual-slow-cooker-sccpvl600-s-10003430.png'),
    ('Fitbit Versa 3', 'Advanced fitness tracker with built-in GPS and heart rate monitoring.', 22900, 35, 'https://t.ctcdn.com.br/JltYZ9dcG7rKic3ESjfRth3DCoI=/fit-in/600x600/filters:fill(transparent):watermark(wm/prd.png,-32p,center,1,none,15)/i599322.png'),
    ('Roku Streaming Stick+', 'Compact streaming device for high-definition entertainment.', 4999, 40, 'https://cigars.roku.com/v1/http%3A%2F%2Fimage.roku.com%2Fw%2Frapid%2Fimages%2Fundefined%2F0722a283-fae2-4a53-93cf-046406281176.png'),
    ('Hamilton Beach Coffee Maker', 'Programmable coffee maker for brewing fresh coffee.', 4999, 30, 'https://crdms.images.consumerreports.org/prod/products/cr/models/400768-drip-coffee-makers-with-carafe-hamilton-beach-programmable-alexa-smart-49350-10011102.png'),
    ('Coleman Sundome Tent', 'Spacious camping tent for outdoor adventures.', 4999, 20, 'https://assets.trailspace.com/assets/7/1/a/11413274/image.png'),
    ('Razer DeathAdder Elite Gaming Mouse', 'High-precision gaming mouse for competitive gaming.', 5999, 30, 'https://assets.razerzone.com/eeimages/support/products/724/724_deathadderelite_500x500.png'),
    ('NutriBullet Blender Combo', 'Versatile blender for making smoothies, soups, and more.', 13900, 25, 'https://mopani.co.za/cdn/shop/files/278225.webp?v=1705394160'),
    ('JBL Flip 5 Portable Bluetooth Speaker', 'Portable speaker with powerful sound for on-the-go entertainment.', 8999, 40, 'https://www.jbl.com.br/dw/image/v2/BFND_PRD/on/demandware.static/-/Sites-masterCatalog_Harman/default/dwd2aae112/JBL_Flip5_Acc_Black_006-1605x1605px_Front.png?sw=535&sh=535'),
    ('Puma RS-X Sneakers', 'Stylish sneakers with retro-inspired design.', 6999, 25, 'https://cdn.vnda.com.br/pardalsneakers/2023/08/21/22_11_48_284_tenis_puma_rs_x_core_white_black_35_1_579badd4e417bf41f2dce5dd6322d116_20210812081920.png?v=1692666708'),
    ('Black+Decker 20V MAX Drill', 'Comprehensive tool kit for DIY projects and home repairs.', 6999, 20, 'https://media-www.canadiantire.ca/product/fixing/tools/portable-power-tools/0547286/black-decker-20v-li-ion-3-8-drill-kit-09c0c875-f7d8-4f65-87d5-4e5b289e5c11.png?imdensity=1&imwidth=640&impolicy=mZoom'),
    ('Sony WH-1000XM4 Wireless Noise-Cancelling Headphones', 'Premium headphones with industry-leading noise cancellation technology.', 34900, 30, 'https://store.sony.co.nz/dw/image/v2/ABBC_PRD/on/demandware.static/-/Sites-sony-master-catalog/default/dw4f020a43/images/WH1000XM4B/WH1000XM4B.png?sw=442&sh=442&sm=fit'),
    ('Ring Video Doorbell', 'Smart doorbell with HD video and two-way talk.', 9999, 35, 'https://png.pngtree.com/png-clipart/20231105/original/pngtree-ring-video-doorbell-png-image_13508243.png'),
    ('L.L.Bean Bean Boots', 'Iconic outdoor boots for all-weather adventures.', 12900, 25, 'https://images.downeast.com/wp-content/uploads/2020/09/2010_GTFM_LLBean2-web.png'),
    ('Google Nest Thermostat', 'Smart thermostat for energy-saving home temperature control.', 12900, 30, 'https://cdn11.bigcommerce.com/s-5qnr8jdmkn/product_images/uploaded_images/image-enus-nlt3-f-ssteel-heat-rgb-simp-3000x3000.png'),
    ('Cuisinart Griddler', 'Versatile countertop grill for indoor grilling and cooking.', 9999, 25, 'https://cuisinart.com.au/sites/cuisinart/media/gr4-na/gr-4na.png'),
    ('Play-Doh Fun Tub', 'Creative playset for molding and shaping with Play-Doh.', 1499, 40, 'https://assets.target.com.au/transform/62a84faa-ce06-4a9f-abfa-2df2d0d84be8/10169544-IMG-007?io=transform%3Afit%2Cwidth%3A1400%2Cheight%3A1600&quality=90&output=webp'),
    ('Gaiam Yoga Mat', 'Non-slip yoga mat for comfortable and stable yoga practice.', 2499, 45, 'https://media-www.canadiantire.ca/product/playing/exercise/exercise-accessories/1840599/gaiam-4mm-spring-fern-yoga-mat-7a09235f-a9e9-4731-9114-d1f3765fbd8e.png?imdensity=1&imwidth=1244&impolicy=mZoom'),
    ('Anker PowerCore Portable Charger', 'High-capacity portable charger for charging devices on-the-go.', 2999, 50, 'https://cdn.shopify.com/s/files/1/0493/9834/9974/files/A16171.png?v=1686209148'),
    ('Dyson V11 Torque Drive Cordless Vacuum Cleaner', 'Powerful cordless vacuum cleaner with advanced filtration.', 59900, 15, 'https://crdms.images.consumerreports.org/f_auto,w_600/prod/products/cr/models/398145-stick-vacuums-greater-than-6-lbs-dyson-v11-torque-drive-10005020.png'),
    ('Calphalon Nonstick Cookware Set', 'Premium nonstick cookware set for versatile cooking.', 24900, 20, 'https://curated-upload.imgix.net/AgAAAB0AhAK5UIMKO6urhf8S-AEKzg.png'),
    ('Jabra Elite 75t True Wireless Earbuds', 'Wireless earbuds with customizable sound and secure fit.', 17999, 30, 'https://www.jabra.br.com/-/media/Images/Products/Jabra-Elite-Active-75t/Layout/Extra_Hotspots/elite75t.png?w=800&hash=24271FCE308B225C54E3059E27E3C8FE'),
    ('Coleman Portable Camping Chair', 'Foldable camping chair for outdoor relaxation.', 1999, 35, 'https://media-www.canadiantire.ca/product/playing/camping/camping-furniture/0762183/coleman-steel-sling-chair-722941b7-d49c-4f20-b324-9988d6d1253c.png');
END;
$$;
