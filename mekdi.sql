--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-1.pgdg22.04+1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: menu; Type: TABLE; Schema: public; Owner: efrayanglain
--

CREATE TABLE public.menu (
    name text NOT NULL,
    serving_size text NOT NULL,
    ingredients text NOT NULL,
    tag text NOT NULL,
    allergy text NOT NULL,
    energy real DEFAULT 0 NOT NULL,
    protein real DEFAULT 0 NOT NULL,
    total_fat real DEFAULT 0 NOT NULL,
    sat_fat real DEFAULT 0 NOT NULL,
    trans_fat real DEFAULT 0 NOT NULL,
    chol real DEFAULT 0 NOT NULL,
    carbs real DEFAULT 0 NOT NULL,
    total_sugar real DEFAULT 0 NOT NULL,
    added_sugar real DEFAULT 0 NOT NULL,
    sodium real DEFAULT 0 NOT NULL,
    description text,
    id uuid DEFAULT gen_random_uuid() NOT NULL
);


ALTER TABLE public.menu OWNER TO efrayanglain;

--
-- Data for Name: menu; Type: TABLE DATA; Schema: public; Owner: efrayanglain
--

COPY public.menu (name, serving_size, ingredients, tag, allergy, energy, protein, total_fat, sat_fat, trans_fat, chol, carbs, total_sugar, added_sugar, sodium, description, id) FROM stdin;
Chicken McNuggets®	 96g 	Chicken bites	Each bite is better than the last.	Cereal containing gluten	254.52	15.04	14.3	6.68	0.1	36.99	15.74	0.48	0	469.87	Bite-sized pieces of breaded, boneless chicken formed in various shapes (Ball, Boot, Bell & Bone) fried and served hot with smoky Barbeque Sauce or Mustard Sauce.	ea4993cd-196c-46fe-afcc-1b48036fc048
McFlurry (Choco Crunch)®	 167.38g	Milk fat, Chocolate	Burgers aren't the only things we're good at.	Milk, Soya	703.17	10.99	19.68	13.8	0.21	11.06	117.78	66.93	54.51	295.89	Milk-based frozen dessert with chocolate crispies and chocolate dip.	d706339d-7974-46b3-bcdc-7fe8f48cdf7b
McFlurry (Oreo)®	 147.38g	Milk fat, Oreo	Burgers aren't the only things we're good at.	Cereal containing gluten, Milk, Soya, Sulphites	209.39	3.58	6.81	4.07	0.12	8	33.42	25.35	19.23	150.9	Milk-based frozen dessert with oreo cookies.	fdeea365-3332-4747-a9cc-c258b677636c
Fillet-O-Fish®	136g	Steamed regular bun, Chunky tartar mayo, Fish fillet patty, Cheese	This one's a great catch.	Cereal containing gluten, Milk, Egg, Fish, Soya	348.11	15.44	14.16	5.79	0.21	32.83	38.85	5.58	3.54	530.54	Made with Alaskan Pollock sourced from sustainable fisheries; topped with a melty American cheese slice and creamy tartar sauce; and served on a soft, steamed bun.	f105a7a8-104f-453d-853a-4ff627d94e38
McChicken®	173g	Quarter bun crown, Veg mayonnaise, Shredded lettuce, McChicken patty, Quarter bun heel	Familiarity breeds confidence.	Cereal containing gluten, Milk, Soya	400.8	15.66	15.7	5.47	0.16	31.17	47.98	5.53	4.49	766.33	Batter & breaded chicken patty containing green peas, carrots, green beans, onion, potatoes, rice and spices, served in a bun with eggless mayonnaise and lettuce.	75b44ffb-1635-4412-968b-38598758e32b
Chicken McNuggets®	 144g 	Chicken bites	Each bite is better than the last.	Cereal containing gluten	381.77	22.56	21.46	10.02	0.14	55.48	23.62	0.72	0	704.81	Bite-sized pieces of breaded, boneless chicken formed in various shapes (Ball, Boot, Bell & Bone) fried and served hot with smoky Barbeque Sauce or Mustard Sauce.	33f7464f-7180-4eba-9145-8c2dcb1b9a89
Cold Coffee Mcfloat®	270ml	i forgor 💀		Milk	270.05	5.91	7.18	4.89	0	15.81	45.44	36.18	19.98	173.59	Rich smooth creamy cold coffee along with sweet vanilla flavoured soft serve mix.\n    \n    	9ca4b2ea-ab94-4c80-b77d-8bb09985c536
Masala Chai®	 150ml			Milk	94.23	1.37	1.46	0.87	0.04	0	18.9	15.06	13.68	7.08	Masala chai is a smooth and spicy blend of natural tea extract, milk solids, black pepper, fennel, clove and cinnamon.	772350ad-63de-47e2-a9d8-ca352915c9f7
McAloo Tikki Burger®	146g	Regular bun crown, Tom-Mayo sauce, Sliced tomatoes, Shredded onion, Aloo tikki patty, Regular bun heel	The one that never goes out of f(l)avour.	Cereal containing gluten, Milk, Soya	339.52	8.5	11.31	4.27	0.2	1.47	50.27	7.05	4.07	545.34	A golden fried vegetarian patty prepared with peas, potato and infused with aromatic spices. Clubbed with sliced tomatoes, shredded red onion, and tangy tomato mayonnaise. Served in a warm toasted bun.	8920f039-3ad4-40ce-8d17-102bd44f7c7d
Cold Coffee®	250ml			Milk	301.1	9.75	11.15	7.45	0	27.4	40.2	37.5	21.25	175	A rich smooth creamy cold coffee made with coffee powder and milk.	b08952f5-297c-48ed-8b6f-91e837aa789d
McEgg®	115g	Regular bun crown, Mayonnaise, Crisp onion, Egg, Curry spices, Regular bun heel	Because eggs are great any time of the day.	Cereal containing gluten, Milk, Egg, Soya	265	12	10	0.8	0.1	76.88	31	5	1.6	675	Made with the freshest, warm, off-the-farm egg; steamed to perfection in our specialised steamer; and made tasty with a sprinkling of magic masala. Sandwiched between freshly toasted buns, topped off with creamy mayo and some crunchy onion.	8d2a816a-a237-42f0-9822-592f8928aabd
Butter Chicken Grilled Burger	153g	Premium grilled chicken patty, Shredded onion, Sesame seeded buns	Dil se makhani burger.	Cereal containing gluten, Milk, Soya	357.01	17.06	14.41	4.65	0	31.93	39.76	6.55	4.71	919.59	Premium grilled chicken patty topped with makhani sauce and shredded onions placed between freshly toasted sesame seeded buns.	792c3cca-8fad-44cf-b93e-13f1de03c429
Iced Tea®	400ml			No Allergens	242.52	1.08	0.12	0	0	0	59.28	58.08	55.16	16.68	A blend of aromatic tea and the fruity flavour of lemon.	8ff3f47e-2b53-4fae-bd94-e52cbe79274b
McFlurry (Oreo)®	86.79g 	Milk fat, Oreo	Burgers aren't the only things we're good at.	Cereal containing gluten, Milk, Soya, Sulphites	116.36	2.05	3.7	2.25	0.07	4.8	18.69	14.49	10.8	80.73	Milk-based frozen dessert with oreo cookies.	cec37715-b99e-4c38-87bd-1b6c71714c46
Black Coffee®	200ml			No Allergens	6.8	0	0	0	0	0	1.7	0	0	0	Perfectly brewed for any time of the day.	5556102f-d064-43b1-8219-b5585d6a41e0
Chicken Maharaja Mac®	296g	Maharaja bun crown, Haberno sauce, Shredded lettuce, Shredded onion, Jalapenos, Flame-grilled chicken patty, Sliced cheese, Maharaja bun heel	A royal treat.	Cereal containing gluten, Milk, Soya	689.12	34	36.69	10.33	0.25	81.49	55.39	8.92	6.14	1854.71	A double-decker toasted Maharaja bun sandwiched with one layer of flame-grilled chicken patty; crunchy iceberg lettuce; shredded onion; and a slice of cheese. Topped with another layer of flame-grilled chicken patty; tomato slices; and crunchy iceberg lettuce infused with harberno sauce.	60cef260-2be8-4feb-932d-9fa2b65e4fc4
Cheesy Fries	150g	Salted Fries, Smoky Chipotle Sauce	Hi, it’s cheese on this side!	Milk	453.92	7.19	21.1	10.64	0.59	3.89	41.94	0.95	0.4	430.79	The all-time favourite fries with a generous dollop of cheesy yet smoky chipotle sauce.	50931828-d842-438a-a8cd-1af25bcc6761
Dosa Masala Burger®	138g	Whole wheat bun, Rasam mayo, Dosa masala aloo patty	A new taste of India.	Cereal containing gluten, Milk, Soya	340.23	5.66	12.39	4.22	0	0	51.52	12.68	6.4	710.54	Turmeric-spiced mashed potato filling topped with fresh peas. Grilled and placed inside a soft whole wheat bun with a spicy molagapodi chutney mayo.	3ebcc085-833b-42c2-8f83-b7a955d0e33d
McSpicy Paneer®	199g	Quarter pounder bun crown, Shredded lettuce, Tandoori mayo, Spicy paneer patty, Quarter pounder bun heel	Let paneer surprise you.	Cereal containing gluten, Milk, Soya	652.76	20.29	39.45	17.12	0.18	21.85	52.33	8.35	5.27	1074.58	Crispy and spicy paneer patty with creamy tandoori sauce and crispy lettuce topping.	a7a2e45c-ad65-4a7e-8077-d6210406b23c
McSwirl Chocolate®	93.29g	Soft serve mix (100% diary product), Chocolate dip	Melts in your mouth. Goes to your heart.	Cereal containing gluten, Milk, Soya	160.14	2.71	7.14	5.25	0.07	5.71	20.92	15.39	11.31	51.31	Delightful soft-serve with a delectable chocolate topping.	ed414dd3-0d1f-4fb0-8b79-14eea7fe0d94
McVeggie®	168g	Quarter bun crown, Veg mayonnaise, Shredded lettuce, Vegetable patty, Quarter bun heel	Pure taste. Pure veg.	Cereal containing gluten, Milk, Soya	402.05	10.24	13.83	5.34	0.16	2.49	56.54	7.9	4.49	706.13	A delectable patty made of green goodness, potatoes, peas, carrots and a selection of Indian spices. Topped with crispy lettuce, mayonnaise, and packed into sesame toasted buns.	d7da3c93-5112-42e7-bafe-dee5302911ba
Our World Famous Fries®	 109g 	Potato and salt	The legend among legends.	No Allergens	317.92	4.79	14.7	7.04	0.11	1.09	38.34	0.55	0	216.79	The crisp, craveable, fan favourite: our World Famous Fries®. These epic fries are crispy and golden on the outside and fluffy on the inside.	6b6d4491-07a0-4cab-873f-0b2f2f0f8f0d
Pizza McPuff®	87g	Assorted vegetables, Refined wheat flour, Pizza seasoning	Something different. Something delicious.	Cereal containing gluten, Milk, Soya	228.21	5.45	11.44	5.72	0.09	5.17	24.79	2.73	0.35	390.74	A blend of assorted vegetables (carrot, beans,capsicum, onion & green peas); mozzarella cheese mixed with tomato sauce; and exotic spices stuffed in rectangle shaped savoury dough. Quick frozen.	fc4ca6d1-d816-4db9-abda-d5b282e7b6f4
Soft Serve Cone®	81.29g	Soft serve mix (100% diary product)	Delightfully basic.	Cereal containing gluten, Milk, Soya	85.73	1.99	1.82	1.31	0.05	4.75	15.23	10.68	6.99	40.78	Creamy vanilla soft-serve on a cone.	4f496912-9002-4b1e-8104-3db354d1273d
Spicy Paneer Wrap®	250g	Whole wheat flat bread, Spicy paneer patty, Veg mayo, Lettuce, Onion, Tomato, Mustard sauce, Cheese	Unwrap deliciousness.	Cereal containing gluten, Milk, Soya	674.68	20.96	39.1	19.73	0.26	40.93	59.278	3.5	1.08	1087.46	Tender paneer patty with a fiery, crunchy batter coating; dressed with fresh veggies and seasonings; topped with creamy sauce; and a dash of mustard and melted cheese.	de073a88-ac28-4f15-9a4c-f1049e78391f
Sundae (Strawberry)®	91.79g 	Soft serve mix (100% diary), Strawberry topping	A little zing goes a long way.	Milk	100.99	1.54	1.77	1.3	0.06	4.85	19.78	17.66	12.49	34.51	Creamy vanilla soft-serve with strawberry topping.	945513be-cea6-48c6-ac96-3cc8b995ee24
Sundae (Chocolate)®	91.79g 	Soft serve mix (100% diary), Hot fudge topping	Much more than ice-cream.	Milk	121.64	2.25	4.02	3.01	0.08	5.85	19.11	17.07	10.78	65.56	Creamy vanilla soft-serve topped with thick and rich hot fudge.	4b89d4a5-7b87-4f7c-8d57-1e23cf899b47
McSpicy Chicken®	186g	Quarter pounder bun crown, Veg sauce, Shredded lettuce, McSpicy chicken patty, Quarter pounder bun heel	Hot. In more ways than one.	Cereal containing gluten, Milk, Egg, Soya	451.92	21.46	19.36	7.63	0.18	66.04	46.08	5.88	4.49	928.52	Zesty and redolent whole muscle leg meat patty: Fried to perfect golden tan; quenched with creamy veg mayo and garden-fresh shredded iceberg lettuce. The sandwich is served in fresh, sesame-studded quarter pounder bun.	9a2e88f0-969b-476f-b185-58790505a7ca
Sundae (Chocolate Brownie)®	110.79g 	Soft serve mix (100% diary), Hazelnut brownie, Hot fudge topping	A global favourite.	Cereal containing gluten, Milk, Nuts	205.26	3.2	5.45	3.65	0.1	6.04	35.26	20.75	14.39	100.89	An iconic premium dessert option. Can be bought as an add-on to make it a "full meal" or simply as an indulgence.	009d025f-0dc8-44e2-8053-7f154124eea4
Butter Paneer Grilled Burger	142g	Mildly spiced grilled paneer patty, Shredded onion, Sesame seeded buns	Dil se makhani burger.	Cereal containing gluten, Milk, Soya	382.26	12.85	17.15	8.29	0	6.62	44.12	8.78	5.08	900.37	Mildly spiced grilled paneer patty topped with makhani sauce and shredded onions placed between freshly toasted sesame seeded buns.	5f40fe76-6b10-4474-8b3f-96a56f5d0832
Veg Surprise Burger	132g	Regular Bun, Italian mayo, Shredded onion, Herb Chilli Potato patty	A surprise that will leave you wide-eyed.	Cereal containing gluten, Milk, Soya	313.44	5.71	14.95	3.73	0.14	0	39.84	5.66	1.64	504.19	A scrumptious potato patty topped with a delectable Italian herb sauce and shredded onions placed between perfectly toasted buns.	d23ef841-a518-47e1-a760-9cb6ef8951ff
Spicy Chicken Wrap®	257g	Whole wheat flat bread, Spicy chicken patty, Veg mayo, Lettuce, Onion, Tomato, Cheese	Familiar, yet different.	Cereal containing gluten, Milk, Egg, Soya	567.19	23.74	26.89	12.54	0.27	87.63	57.06	2.52	1.08	1152.38	Juicy chicken coated with hot and crispy batter; dressed with a fresh salad of lettuce, onions, tomatoes and seasonings. Served with creamy sauce and supple cheese slices.\n    	a3ec2a1b-152e-4c8f-b59a-42623ec4de66
Veg Maharaja Mac®	306g	Maharaja bun crown, Cocktail sauce, Shredded lettuce, Shredded onion, Jalepenos, Corn & cheese patty, club, Sliced cheese, Maharaja bun heel	A feast fit for kings (and queens).	Cereal containing gluten, Milk, Soya	832.67	24.17	37.94	16.83	0.28	36.19	93.84	11.52	6.92	1529.22	A double-decker toasted Maharaja bun sandwiched with one layer of corn & cheese patty; crunchy iceberg lettuce; shredded onion; and a slice of cheese. Topped with another layer of corn & cheese patty.	60cb1b8a-8abe-4627-8c48-a65b745b5cba
Abimanyu Burger	250g	patties and meat and stuff	delicious	no allergies	0	0	0	0	0	0	0	0	0	0	Why do I like burgers so much there are many more kinds of items out there you know	eda8a54c-0fa5-4082-8dd4-67e57b05c453
Devano coffee	150g	200h	smk metland goodies.	no allergies	0	0	0	0	0	0	0	0	0	0	it's black.    \n    	daea86d3-53df-458a-9b76-bec10dcb49d0
\.


--
-- Name: menu menu_pkey; Type: CONSTRAINT; Schema: public; Owner: efrayanglain
--

ALTER TABLE ONLY public.menu
    ADD CONSTRAINT menu_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

