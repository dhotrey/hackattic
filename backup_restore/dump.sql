--
-- PostgreSQL database dump
--

-- Dumped from database version 10.19 (Debian 10.19-1.pgdg90+1)
-- Dumped by pg_dump version 10.19 (Debian 10.19-1.pgdg90+1)

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

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = true;

--
-- Name: criminal_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.criminal_records (
    id integer NOT NULL,
    name character varying(120) NOT NULL,
    felony character varying(30) NOT NULL,
    ssn character varying(11) NOT NULL,
    home_address character varying(100) NOT NULL,
    entry timestamp without time zone NOT NULL,
    city character varying(100) NOT NULL,
    status character varying(16) NOT NULL
);


ALTER TABLE public.criminal_records OWNER TO postgres;

--
-- Name: criminal_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.criminal_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.criminal_records_id_seq OWNER TO postgres;

--
-- Name: criminal_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.criminal_records_id_seq OWNED BY public.criminal_records.id;


--
-- Name: criminal_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.criminal_records ALTER COLUMN id SET DEFAULT nextval('public.criminal_records_id_seq'::regclass);


--
-- Data for Name: criminal_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.criminal_records (id, name, felony, ssn, home_address, entry, city, status) FROM stdin;
1	Christopher Olson	Perjury	744-43-8352	91587 Vazquez Ports	1978-10-17 00:00:00	Georgeshire, AL 21973-6031	missing
2	Jenny Crawford	Vehicular homicide	551-40-7071	560 Johnson Gardens	1993-10-03 00:00:00	Lake Keithmouth, TN 18079-7577	alive
3	Cassandra Parks	Obstruction of justice	398-83-6338	5848 Jessica Junctions Suite 722	1979-03-26 00:00:00	Lake Darius, VA 55226	terminated
4	Eric Miller	Vehicular homicide	887-86-5877	Unit 3752 Box 0293	1978-11-15 00:00:00	DPO AA 72246	terminated
5	Heather Sanders	Arson	634-77-0467	USNS Frazier	1975-05-09 00:00:00	FPO AE 05346	missing
6	James Martin	Obstruction of justice	429-55-2039	75075 Powell Station	2003-05-27 00:00:00	Morganstad, KY 06862-7427	alive
7	Richard Mann	Manslaughter	807-09-7613	63684 Zamora Plaza Apt. 500	2001-09-28 00:00:00	Jayfurt, WV 02662-6260	missing
8	Alexis Oliver	Perjury	033-80-7541	1980 Gibson Points Suite 344	1982-04-12 00:00:00	New Loriside, WA 82706	missing
9	Joanna Lamb	Larceny	132-55-2153	0511 Valdez Expressway Suite 935	2013-02-01 00:00:00	West Joshua, NY 90330-8367	alive
10	Misty Shelton	Vehicular homicide	800-85-2317	80920 Jessica Loaf	2014-10-30 00:00:00	Lawsonville, UT 56672	alive
11	Zachary Douglas	Vehicular homicide	859-31-5424	533 Kathryn Lights Suite 640	1996-12-18 00:00:00	Millerburgh, NV 26104-3159	alive
12	Bryce Powell	Obstruction of justice	719-18-0762	721 Henry Village Apt. 225	1988-04-27 00:00:00	Ericside, AS 36465	missing
13	Nicholas Walters	Obstruction of justice	656-98-9896	184 Hodge Meadows	1982-11-02 00:00:00	East Edward, NC 86242	alive
14	Bradley Brown	Check fraud	509-76-9168	3831 Reginald Fork	2005-06-03 00:00:00	South Joshua, NE 09477	terminated
15	Gerald Fletcher	Burglary	018-38-8054	820 Christopher Row Apt. 586	2019-02-19 00:00:00	New Heather, MT 50862-4781	alive
16	Jennifer Young	Perjury	312-21-0721	42214 Davis Extensions	2014-10-22 00:00:00	Smithtown, WA 38233-0582	alive
17	Kristen Price	Arson	416-88-6787	1126 Smith Stream Apt. 897	1988-02-09 00:00:00	New Lawrenceborough, ME 01850-3734	missing
18	Michelle Wallace	Burglary	334-12-1401	513 Colon Trail	2000-10-19 00:00:00	New Bryanborough, NC 40335-8405	missing
19	Hector Scott	Burglary	132-49-1933	01772 Wallace Ridges	2001-01-20 00:00:00	Monicaberg, CO 39043-3439	terminated
20	Joseph Clark	Manslaughter	033-31-9020	37931 Nicole Shores Apt. 065	1997-12-26 00:00:00	Jenniferview, NJ 65875	missing
21	Theresa Anderson	Obstruction of justice	110-39-2300	054 Huff Center Apt. 705	1975-07-13 00:00:00	East Eugene, HI 33768	terminated
22	Holly Lyons	Obstruction of justice	217-73-9039	5617 Claudia Course Apt. 128	2005-12-23 00:00:00	Fordmouth, AR 34702	alive
23	Christie Jackson	Vehicular homicide	754-69-7409	44902 Brooke Stravenue	1997-10-19 00:00:00	Jeffreyberg, NY 03021-4039	terminated
24	Carrie Wilson	Arson	250-98-1274	064 Andrew Park	1988-01-17 00:00:00	Port Robert, FM 55564	terminated
25	Clarence Dawson	Burglary	438-79-2517	1587 Watson Corner	1981-06-16 00:00:00	West Bryanstad, MO 17667	missing
26	Sandra Yates	Vehicular homicide	159-65-0902	65347 Pham Mountains Apt. 913	2003-11-19 00:00:00	Jonathonshire, NH 75811-9657	terminated
27	William Young	Manslaughter	368-31-7757	0421 Vanessa Corners Apt. 197	2010-05-12 00:00:00	Henryview, WA 98791	terminated
28	Mr. Thomas Johnson MD	Animal cruelty	469-26-0495	773 Richardson Fork	1974-06-02 00:00:00	West Brianton, NY 50587	alive
29	David Rice	Vehicular homicide	798-54-1327	6218 Tabitha Union Apt. 308	1994-03-16 00:00:00	Michaelshire, OH 57532-2509	missing
30	Gabriella Walters	Larceny	086-76-2432	359 Rebecca Center	1971-04-18 00:00:00	North Susan, DC 93900-6661	alive
31	Christine Cook	Animal cruelty	338-98-2450	3466 Caleb Extension	1995-07-20 00:00:00	Destinymouth, IN 23031-9685	missing
32	David Smith	Tax evasion	019-80-2037	5312 Castillo Mission	2003-05-04 00:00:00	Hayleyhaven, OR 40514-0119	alive
33	Lori Anderson	Tax evasion	586-99-8018	46774 Burton Falls	2015-07-28 00:00:00	West Danaport, OK 04031	terminated
34	Rebecca Velez	Burglary	183-63-4348	USS Lee	1990-03-12 00:00:00	FPO AA 19780-2384	alive
35	Thomas Andrews	Perjury	022-64-7400	747 Andrew Bypass	1991-03-01 00:00:00	Toddland, OR 35188	alive
36	Aaron Jones	Burglary	210-79-5400	678 Hernandez Drive Suite 636	2002-05-05 00:00:00	Bakermouth, GU 70185-3194	alive
37	Taylor House	Manslaughter	322-42-2298	38874 William Mountains	1994-05-21 00:00:00	North Danielleton, GU 82331-6547	alive
38	Kara White	Vehicular homicide	630-91-4537	38051 John Drive	1989-04-16 00:00:00	Sarahview, ID 99172-7962	missing
39	Marie Perkins	Larceny	733-45-2124	1458 Natalie Divide Suite 155	2009-10-24 00:00:00	West Michelle, IL 89156-0366	terminated
40	Sheena Murphy	Obstruction of justice	495-48-6276	2084 Gregory Street Apt. 660	1991-06-25 00:00:00	Matthewshire, IA 53181	terminated
41	Michael Frank	Tax evasion	553-18-6638	53102 Conway Extension	1980-02-14 00:00:00	West Amyborough, NY 83879	terminated
42	Lisa Dodson	Burglary	342-48-2778	4112 Brandy Port	2013-11-17 00:00:00	Faulknerbury, OH 00070	alive
43	Cindy Cox	Burglary	584-27-0854	383 Chris Island	1983-04-21 00:00:00	Oliverville, NH 62103	missing
44	Mariah Stark	Perjury	798-10-2043	Unit 1568 Box 6556	1997-04-16 00:00:00	DPO AE 94335	alive
45	Catherine Higgins	Tax evasion	168-82-2521	399 James Flat	1996-02-13 00:00:00	Lake Angela, AK 63176	missing
46	Scott Garza	Burglary	730-58-5829	55139 Newton Point Apt. 973	1984-12-18 00:00:00	Wisemouth, ID 12547	missing
47	Richard Johnson	Obstruction of justice	166-56-1562	86234 Myers Rue	1972-09-01 00:00:00	East Travis, KY 20941	missing
48	Dennis Hernandez	Burglary	079-57-3940	5818 Johnson Light Suite 175	2021-08-02 00:00:00	Davidsonfurt, MH 90177-6234	terminated
49	Timothy Craig	Larceny	043-21-6674	72434 Evans Circles Apt. 588	1993-03-23 00:00:00	North Lorraine, WA 22482	missing
50	Ricky Gonzalez	Larceny	482-83-9288	898 Jackson Spring	2000-09-13 00:00:00	Elizabethbury, SC 69078-0712	alive
51	Jennifer Schneider	Larceny	527-43-7364	058 Cortez Parkways Apt. 285	2016-08-14 00:00:00	East Randy, MP 09876-2526	terminated
52	Margaret Sanders	Perjury	290-36-9206	850 Carson Union Suite 544	1977-07-23 00:00:00	New Margaretmouth, OH 51597-4891	alive
53	Tiffany Wright	Obstruction of justice	717-21-8524	72213 Brendan Manor Suite 487	2000-05-10 00:00:00	Youngburgh, PR 09395	missing
54	Victoria Hahn	Manslaughter	894-68-8477	15360 James Loaf Apt. 533	2002-03-25 00:00:00	Morrismouth, CO 73553-5373	terminated
55	Christina Williams	Perjury	088-58-8814	769 Huffman Vista	1974-03-05 00:00:00	Wernerchester, MP 71562-5349	terminated
56	Samuel Mccall	Tax evasion	132-45-3240	751 Kennedy Cape Apt. 820	1992-03-09 00:00:00	Caitlinborough, ID 55717-1397	alive
57	Ashley Camacho	Vehicular homicide	668-22-2794	04845 Joel Harbors Suite 394	1978-03-02 00:00:00	Walshchester, CT 90244	terminated
58	Deanna Ramos	Arson	414-15-2609	00756 Roy Forge Suite 452	1970-12-20 00:00:00	North Rebecca, AS 17879	missing
59	Ian Hansen	Larceny	806-37-8401	31254 Barbara Way	1986-07-20 00:00:00	Port Lucasside, NV 51775-7383	alive
60	Christopher Russo	Animal cruelty	757-24-9425	66791 Lloyd Bridge	2010-04-06 00:00:00	Thomasbury, SC 96842	missing
61	Brian Lester	Larceny	008-31-4704	71654 Chelsey Mountains Apt. 681	1972-05-04 00:00:00	North Jordanshire, TN 39629	alive
62	Briana Murray	Manslaughter	602-26-9479	291 Sandra Bridge	1970-11-11 00:00:00	Brendafurt, WI 95555	missing
63	Tracy Martin	Check fraud	392-96-7818	961 Kevin Mountains	2003-02-16 00:00:00	East Kristiemouth, NJ 74088-8117	terminated
64	Nicholas Sweeney	Check fraud	186-57-3607	6345 Albert Pass Apt. 590	2006-11-06 00:00:00	East Davidmouth, GA 19476-1232	alive
65	Robert Ellis	Perjury	772-10-9055	3607 Melissa Ports Apt. 452	1991-11-13 00:00:00	North Michelle, MD 85307	terminated
66	Kaitlyn Adams	Obstruction of justice	356-57-6391	0705 Micheal Junction	1972-03-23 00:00:00	New Bruceshire, NC 72343-1176	missing
67	Joseph Lopez	Perjury	082-01-2801	144 Jensen Stream	1992-03-03 00:00:00	New Ronaldstad, TX 79200	alive
68	Marc Rosario	Animal cruelty	757-26-0848	54869 Long Mills	1989-07-19 00:00:00	South Dawn, NM 88177-8272	alive
69	Ronald Thompson	Obstruction of justice	565-82-2191	2318 Denise Inlet Suite 637	2023-12-01 00:00:00	Feliciaton, AS 95196-9867	alive
70	Mitchell Lee	Arson	381-13-4044	7893 Mcintyre Wall	1977-08-11 00:00:00	Lake Johnland, DC 01662-1829	missing
71	John Smith	Tax evasion	752-35-1144	USS Clark	1974-08-10 00:00:00	FPO AE 94570-1347	missing
72	Michael Ruiz	Manslaughter	161-96-8271	874 Danielle Glens	1986-12-19 00:00:00	West Ericport, GU 35003-4930	missing
73	Mia Cook	Burglary	358-80-9665	1008 Amanda Center Suite 942	2009-09-07 00:00:00	West Tracey, NM 13889-3729	terminated
74	Theresa Thompson	Tax evasion	629-69-9304	29551 Tyrone Road	1993-12-27 00:00:00	West Erica, RI 89832	terminated
75	Susan Schroeder	Perjury	047-66-2627	409 Michael Rue Apt. 723	2012-01-09 00:00:00	East Michaelfort, MS 08279-6008	terminated
76	Benjamin Massey	Obstruction of justice	426-90-8494	123 Rhonda Haven Suite 039	1982-07-13 00:00:00	Andrewstad, SD 14185	missing
77	Pamela Castillo	Animal cruelty	235-14-6833	38506 Rebecca Viaduct Suite 753	1977-08-14 00:00:00	Robertland, WY 99209	alive
78	Mark Cline	Perjury	718-86-5392	279 Shields Mountains	2009-05-08 00:00:00	North Kimberly, NM 31930-0670	missing
79	Robert Cook	Obstruction of justice	756-70-0273	2759 Victoria Curve	2006-12-22 00:00:00	Adrianville, MN 88865-9337	alive
80	Patricia Poole	Manslaughter	483-23-6978	795 Daniel Mews Suite 135	2021-03-05 00:00:00	East Curtis, IA 49871	terminated
81	William Taylor	Obstruction of justice	846-12-2853	012 Hansen Prairie	1974-10-12 00:00:00	West Williambury, AK 50860-6163	alive
82	Eric Horn	Perjury	010-31-2222	USCGC Shaw	2005-08-16 00:00:00	FPO AE 34567-4486	alive
83	Jeffery Chen	Animal cruelty	344-89-4029	4675 Joshua Mews	1994-03-19 00:00:00	Jonesmouth, FM 90477	terminated
84	Brent Day	Burglary	363-17-2083	76142 Sullivan Estate	2004-12-30 00:00:00	Brownbury, TX 58446-9333	alive
85	Walter Serrano	Manslaughter	334-24-6199	45394 Hansen Divide Apt. 061	1990-11-15 00:00:00	New Thomasstad, AS 48921-6261	alive
\.


--
-- Name: criminal_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.criminal_records_id_seq', 85, true);


--
-- Name: criminal_records criminal_records_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.criminal_records
    ADD CONSTRAINT criminal_records_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

